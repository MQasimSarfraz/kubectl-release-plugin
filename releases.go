package main

import (
	"github.com/google/go-github/v21/github"
	"context"
	"fmt"
	"io"
	"text/tabwriter"
	"strings"
	"os"
	"github.com/pkg/errors"
)

func main() {

	// column names for the table
	headers := []string{"NAME", "VERSION", "AGE", "URL"}

	// names of the project to retrieve release information
	k8s := project{"kubernetes", "kubernetes"}
	kops := project{"kubernetes", "kops"}
	projects := &[]project{k8s, kops}

	// get the project information
	client := github.NewClient(nil)
	projectsInfo, err := getProjectsInfo(client, projects)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *github.RateLimitError:
			fmt.Println("Could not retrive information for github - Hitting rate limit")
			os.Exit(1)
		default:
			fmt.Println("Could not retrive information for github - error type Unknown")
			os.Exit(1)
		}
	}

	// print the table
	formatAndPrintTable(os.Stdout, headers, projectsInfo)
}

type project struct {
	owner string
	name  string
}

func getProjectsInfo(c *github.Client, p *[]project) ([][]string, error) {
	var projectsInfo [][]string
	for _, repo := range *p {
		gitRelease, _, err := c.Repositories.GetLatestRelease(context.Background(), repo.owner, repo.name)
		if err != nil {
			return nil, err
		}
		projectsInfo = append(projectsInfo, []string{repo.name, *gitRelease.TagName, (*gitRelease.CreatedAt).String(), *gitRelease.HTMLURL})

	}

	return projectsInfo, nil
}

func formatAndPrintTable(out io.Writer, headers []string, rows [][]string) error {
	w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
	fmt.Fprintf(w, strings.Join(headers, "\t"))
	fmt.Fprintln(w)
	for _, values := range rows {
		fmt.Fprintf(w, strings.Join(values, "\t"))
		fmt.Fprintln(w)
	}
	return w.Flush()
}
