package cmd

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func Run()  {
	app := cli.NewApp()
	app.Name = "Go-Mantis"
	app.Usage = "基于gin的API服务demo"
	app.Version = "0.0.1"

	app.Commands = []cli.Command {
		start,
		stop,
		reload,
		restart,
	}

	app.Action = func(c *cli.Context) error {
		_ = cli.ShowAppHelp(c)
		return nil
	}

	app.Before = func(c *cli.Context) error {
		return nil
	}
	app.After = func(c *cli.Context) error {
		return nil
	}

	cli.HelpFlag = cli.BoolFlag {
		Name: "help, h",
		Usage: "show help",
	}

	cli.VersionFlag = cli.BoolFlag {
		Name: "version, v",
		Usage: "print version",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
