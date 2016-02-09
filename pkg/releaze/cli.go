package releaze

import (
	"encoding/json"
	"log"

	"github.com/codegangsta/cli"
)

func CliCommand() cli.Command {
	return cli.Command{
		Name:  "buildinfo",
		Usage: "Dump releaze embedded build info.",
		Action: func(c *cli.Context) {
			info := Get()
			bytes, err := json.MarshalIndent(info, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			c.App.Writer.Write(bytes)
			c.App.Writer.Write([]byte("\n"))
		},
	}
}
