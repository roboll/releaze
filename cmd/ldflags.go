package cmd

import (
	"fmt"
	"io"
	"sync"

	"github.com/codegangsta/cli"
)

var ldflags = cli.Command{
	Name:  "ldflags",
	Usage: "Prints ldflags to embed build info at compile time.",
	Action: func(c *cli.Context) {
		printLdflags(c.App.Writer, defaultCollectors)
	},
}

var defaultCollectors = map[string]func() string{
	"github.com/roboll/releaze/pkg/releaze.version": func() string { return "version" },
	"github.com/roboll/releaze/pkg/scm/git.commit":  func() string { return "commit" },
	"github.com/roboll/releaze/pkg/scm/git.branch":  func() string { return "branch" },
}

func printLdflags(out io.Writer, collectors map[string]func() string) {
	var wg sync.WaitGroup

	flags := make(map[string]string, len(collectors))
	for key, collector := range collectors {
		wg.Add(1)
		go func(key string, collector func() string) {
			defer wg.Done()
			flags[key] = collector()
		}(key, collector)
	}

	wg.Wait()

	//out.Write([]byte(`"`))
	for key, val := range flags {
		out.Write([]byte(fmt.Sprintf("-X %s=%s ", key, val)))
	}
	//out.Write([]byte(`" `))
}
