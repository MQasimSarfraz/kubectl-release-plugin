package kubectlreleaseplugin

import (
	"context"
	"fmt"
	"github.com/google/go-github/v24/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"io"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

var titles = []string{"NAME", "VERSION", "AGE", "URL"}

var defaultProject = project{owner: "kubernetes", name: "kubernetes"}

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
	{
		owner: "etcd-io",
		name:  "etcd",
	},
	{
		owner: "containous",
		name:  "traefik",
	},
	{
		owner: "openfaas",
		name:  "faas",
	},
	{
		owner: "rancher",
		name:  "rancher",
	},
	{
		owner: "kubeflow",
		name:  "kubeflow",
	},
	{
		owner: "kubernetes-sigs",
		name:  "kubespray",
	},
	{
		owner: "kubernetes-sigs",
		name:  "cluster-api",
	},
	{
		owner: "rook",
		name:  "rook",
	},
	{
		owner: "kubeless",
		name:  "kubeless",
	},
	{
		owner: "coreos",
		name:  "flannel",
	},
	{
		owner: "cilium",
		name:  "cilium",
	},
	{
		owner: "argoproj",
		name:  "argo",
	},
	{
		owner: "zalando",
		name:  "patroni",
	},
	{
		owner: "kubernetes-incubator",
		name:  "external-dns",
	},
	{
		owner: "pulumi",
		name:  "pulumi",
	},
	{
		owner: "linkerd",
		name:  "linkerd2",
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

	ctx := context.Background()
	client := newGitClient(ctx)

	var projectsRelease [][]string
	if filterProject != "" {
		for _, project := range *projects {
			if strings.ToLower(filterProject) == project.name {
				release, err := release(client, ctx, project)
				if err != nil {
					return err
				}

				projectsRelease = append(projectsRelease, release)
				break
			}
		}
	} else {
		release, err := release(client, ctx, defaultProject)
		if err != nil {
			return err
		}

		projectsRelease = append(projectsRelease, release)
	}

	formatAndPrintTable(os.Stdout, titles, projectsRelease)

	return nil
}

func CheckError(err error) {
	if err != nil {
		switch errors.Cause(err).(type) {
		case *github.RateLimitError:
			fmt.Println("Hitting rate limit - Please set env GITHUB_TOKEN and retry")
			os.Exit(1)
		default:
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	}
}

func release(client *github.Client, ctx context.Context, project project) ([]string, error) {
	gitRelease, _, err := client.Repositories.GetLatestRelease(ctx, project.owner, project.name)
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

func newGitClient(ctx context.Context) *github.Client {
	gitToken := os.Getenv("GITHUB_TOKEN")
	if gitToken != "" {
		tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: gitToken})
		tokenClient := oauth2.NewClient(ctx, tokenSource)
		return github.NewClient(tokenClient)
	}
	return github.NewClient(nil)
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
