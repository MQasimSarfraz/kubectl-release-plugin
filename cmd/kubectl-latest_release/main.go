package main

import (
	plugin "github.com/MQasimSarfraz/kubectl-release-plugin"
	"github.com/jessevdk/go-flags"
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
	plugin.CheckError(err)
}
