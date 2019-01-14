package main

import (
	"fmt"
	plugin "github.com/MQasimSarfraz/kubectl-release-plugin"
	"github.com/google/go-github/v21/github"
	"github.com/pkg/errors"
	"os"
)

func main() {

	// get the project information
	client := github.NewClient(nil)
	projectsInfo, err := plugin.GetProjectsInfo(client, plugin.Projects)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *github.RateLimitError:
			fmt.Println("Could not retrive information from github - Hitting rate limit")
			os.Exit(1)
		default:
			fmt.Println("Could not retrive information from github - error type Unknown")
			os.Exit(1)
		}
	}

	// print the table
	plugin.FormatAndPrintTable(os.Stdout, plugin.Headers, projectsInfo)
}
