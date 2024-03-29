package main

import (
	"os"
	"sorcia/cmd"

	"github.com/urfave/cli"
)

// Version ...
// Write sorcia version here
const Version = "0.1.0"

func main() {
	app := cli.NewApp()
	app.Name = "sorcia"
	app.Usage = "Self-hosted and modular services for your Git projects"
	app.Version = Version
	app.Commands = []cli.Command{
		cmd.Core,
		cmd.SSHServe,
	}
	app.Run(os.Args)
}
