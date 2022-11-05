package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Action = func(c *cli.Context) {
		fmt.Println(c.FlagNames())
	}
	app.Name = "cli test"
	app.Usage = "cli test"
	app.Version = "1.0.0"
	app.HelpName = "help"
	app.Run(os.Args)
}
