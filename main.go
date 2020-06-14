package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	log "github.com/cihub/seelog"
	"github.com/reyahsolutions/orchestra/commands"
	"github.com/reyahsolutions/orchestra/config"
	"github.com/reyahsolutions/orchestra/services"
	"github.com/urfave/cli"
)

var app *cli.App

const defaultConfigFile = "orchestra.yml"

func main() {
	defer log.Flush()
	app = cli.NewApp()
	app.Name = "Orchestra"
	app.Usage = "Orchestrate Go Services"
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		commands.BuildCommand,
		commands.ExportCommand,
		commands.InstallCommand,
		commands.LogsCommand,
		commands.PsCommand,
		commands.RestartCommand,
		commands.StartCommand,
		commands.StopCommand,
		commands.TestCommand,
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "config, c",
			Value:   "orchestra.yml",
			Usage:   "Specify a different config file to use",
			EnvVars: []string{"ORCHESTRA_CONFIG"},
		},
	}
	// init checks for an existing orchestra.yml in the current working directory
	// and creates a new .orchestra directory (if doesn't exist)
	app.Before = func(c *cli.Context) error {
		confVal := c.String("config")
		if confVal == "" {
			confVal = defaultConfigFile
		}

		config.ConfigPath, _ = filepath.Abs(confVal)
		if _, err := os.Stat(config.ConfigPath); os.IsNotExist(err) {
			fmt.Printf("No %s found. Have you specified the right directory?\n", c.String("config"))
			os.Exit(1)
		}
		services.ProjectPath, _ = path.Split(config.ConfigPath)
		services.OrchestraServicePath = services.ProjectPath + ".orchestra"

		if err := os.Mkdir(services.OrchestraServicePath, 0766); err != nil && os.IsNotExist(err) {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		config.ParseGlobalConfig()
		services.Init()
		return nil
	}
	app.Version = "0.4.0"
	app.Run(os.Args)
	if commands.HasErrors() {
		os.Exit(1)
	}
}
