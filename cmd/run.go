package cmd

import "github.com/codegangsta/cli"

var cmds = []cli.Command{ldflags, version}

func RunCli(args []string) {
	app := cli.NewApp()

	app.Name = "releaze"
	app.Usage = `embed release info in build artifacts.
   https://github.com/roboll/releaze`
	app.Version = ""

	app.Commands = cmds
	app.Run(args)
}
