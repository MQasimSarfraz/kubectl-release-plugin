package main

import (
	"fmt"
	plugin "github.com/MQasimSarfraz/kubectl-release-plugin"
	"github.com/google/go-github/v21/github"
	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
	"os"
)

func main() {

	var opts struct {
		Project string `short:"p" long:"project" description:"Latest release for the given project "`
		List    bool   `short:"l" long:"list" description:"List of the allowed projects "`
	}

	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	// print projects releases
	err = plugin.Execute(opts.Project, opts.List)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *github.RateLimitError:
			fmt.Println("Could not retrieve information from github - Hitting rate limit")
			os.Exit(1)
		default:
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	}
}
