package kubectlreleaseplugin

import (
	"context"
	"fmt"
	"github.com/google/go-github/v21/github"
	"time"
)

type Project struct {
	owner string
	name  string
}

func GetProjectsInfo(c *github.Client, p *[]Project) ([][]string, error) {
	var projectsInfo [][]string
	for _, repo := range *p {
		gitRelease, _, err := c.Repositories.GetLatestRelease(context.Background(), repo.owner, repo.name)
		if err != nil {
			return nil, err
		}

		projectsInfo = append(projectsInfo, []string{repo.name, *gitRelease.TagName, getAge(gitRelease.CreatedAt.Time), *gitRelease.HTMLURL})

	}

	return projectsInfo, nil
}

func getAge(t time.Time) string {
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
