package commands

import (
	"fmt"

	"github.com/urfave/cli"
)

func run() {
	app := cli.NewApp()
	app.Name = "Sakara Module Builder"
	app.Usage = "Let's us build your module and extract the module API in seconds"

	ModuleBuilderFlags := []cli.Flag{
		cli.StringFlag{
			Name: "moduleName",
		},
		cli.StringFlag{
			Name: "data",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "sakara:module",
			Usage: "Create New module",
			Flags: ModuleBuilderFlags,
			Action: func(c *cli.Context) {
				fmt.Printf("S")
				ModuleBuilderInit(c)
			},
		}, {
			Name:  "sakara:model",
			Usage: "Create New model",
			Flags: ModuleBuilderFlags,
			Action: func(c *cli.Context) {
			},
		},
	}
}
