package main

import (
	"github.com/sirupsen/logrus"
	"go-mantis/cmd"
)

func main() {
	logrus.Debug("我日")
	cmd.Run()
}
