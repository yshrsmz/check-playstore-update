package main

import "github.com/urfave/cli"

// Flags for the app
var Flags = []cli.Flag{
	cli.StringFlag{
		Name:  "name, n",
		Value: "",
		Usage: "app package name to check.",
	},
	cli.StringFlag{
		Name:  "config, c",
		Value: "config.yml",
		Usage: "specify config file.",
	},
}
