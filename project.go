package kubectlreleaseplugin

import (
	"context"
	"fmt"
	"github.com/google/go-github/v21/github"
	"io"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

// Titles for the table columns
var titles = []string{"NAME", "VERSION", "AGE", "URL"}

// names of the project to retrieve release information
var projects = &[]project{
	{
		owner: "kubernetes",
		name:  "kubernetes",
	},
	{
		owner: "kubernetes",
		name:  "kops",
	},
	{
		owner: "istio",
		name:  "istio",
	},
	{
		owner: "helm",
		name:  "helm",
	},
	{
		owner: "kubernetes",
		name:  "ingress-nginx",
	},
}

type project struct {
	owner string
	name  string
}

func Execute(filterProject string, showProjects bool) error {

	if showProjects {
		listProjects()
		return nil
	}

	client := github.NewClient(nil)

	var projectsRelease [][]string
	for _, project := range *projects {
		if filterProject != "" {
			if strings.ToLower(filterProject) == project.name {
				release, err := release(client, project)
				if err != nil {
					return err
				}

				projectsRelease = append(projectsRelease, release)
				break
			}
		} else {
			release, err := release(client, project)
			if err != nil {
				return err
			}
			projectsRelease = append(projectsRelease, release)
		}
	}

	formatAndPrintTable(os.Stdout, titles, projectsRelease)

	return nil
}

func release(client *github.Client, project project) ([]string, error) {
	gitRelease, _, err := client.Repositories.GetLatestRelease(context.Background(), project.owner, project.name)
	if err != nil {
		return nil, err
	}

	return []string{project.name, *gitRelease.TagName, age(gitRelease.CreatedAt.Time), *gitRelease.HTMLURL}, nil
}

func listProjects() {
	var list [][]string
	for _, project := range *projects {
		list = append(list, []string{project.name, project.owner})
	}

	formatAndPrintTable(os.Stdout, []string{"NAME", "OWNER"}, list)
}

func formatAndPrintTable(out io.Writer, titles []string, rows [][]string) error {
	w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
	fmt.Fprintf(w, strings.Join(titles, "\t"))
	fmt.Fprintln(w)
	for _, values := range rows {
		fmt.Fprintf(w, strings.Join(values, "\t"))
		fmt.Fprintln(w)
	}
	return w.Flush()
}

func age(t time.Time) string {
	age := time.Now().Sub(t)
	if age.Hours() != 0 {
		hours := int(age.Hours())
		if hours < 24 {
			return fmt.Sprintf("%d hours", hours)
		} else {
			return fmt.Sprintf("%d days", hours/24)
		}
	} else if age.Minutes() != 0 {
		return fmt.Sprintf("%d minutes", int(age.Minutes()))
	} else {
		return fmt.Sprintf("%d seconds", int(age.Seconds()))
	}
}
