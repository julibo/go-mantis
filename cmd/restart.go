package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

var restart = cli.Command{
	Name: "restart",
	Usage: "程序重启命令",
	Category: "serve",
	Before: func(context *cli.Context) error {
		return nil
	},
	After: func(context *cli.Context) error {
		return nil
	},
	Action: func(c *cli.Context) error {
		fmt.Println("1 + 1 = ", 1 + 1)
		return nil
	},
	Flags: []cli.Flag{

	},
}

