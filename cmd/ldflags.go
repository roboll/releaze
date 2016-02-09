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
	"github.com/roboll/releaze/pkg/scm/git.commit": tryGitCommit,
	"github.com/roboll/releaze/pkg/scm/git.branch": tryGitBranch,
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

	for key, val := range flags {
		out.Write([]byte(fmt.Sprintf("-X %s=%s ", key, val)))
	}
}

func tryGitCommit() string {
	//TODO
	return "commit"
}

func tryGitBranch() string {
	//TODO
	return "branch"
}
