package cmd

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
)

var v string

var version = cli.Command{
	Name:  "version",
	Usage: "Prints ldflags to embed version at compile time.",
	Action: func(c *cli.Context) {
		if len(v) < 1 {
			log.Fatal("Version is required!")
		}
		c.App.Writer.Write([]byte(fmt.Sprintf("-X %s=%s ", "github.com/roboll/releaze/pkg/releaze.version", v)))
	},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "version, v",
			Usage:       "Release version.",
			EnvVar:      "REL_VERSION",
			Destination: &v,
		},
	},
}
