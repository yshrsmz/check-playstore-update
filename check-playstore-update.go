package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "check-playstore-update"
	app.Version = Version
	app.Usage = "Check if android app update is deployed to PlayStore"
	app.Author = "yshrsmz"
	app.Email = "the.phantom.bane@gmail.com"
	app.Flags = Flags
	app.Action = doAction

	app.Run(os.Args)
}
