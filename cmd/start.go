package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

var host string
var port int
var env string

var start = cli.Command{
	Name: "start",
	Usage: "程序启动命令",
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
		cli.StringFlag{
			Name:        "host",           //参数名称
			Value:       "0.0.0.0",        //参数默认值
			Usage:       "Server Address", //参数功能描述
			Destination: &host,            //接收值的变量
		},
		cli.IntFlag{
			Name:  "port",
			Value: 9111,
			Usage: "Listening Port",
			Destination: &port,
		},
		cli.StringFlag{
			Name:        "env, e",            	//参数名称
			Value:       "dev",      	   		//参数默认值
			Usage:       "Runtime Environment",	//参数功能描述
			Destination: &env,            		//接收值的变量
		},
		cli.BoolFlag{
			Name:        "daemonize, d",
			Usage:       "Running Mode",
		},
	},
}



