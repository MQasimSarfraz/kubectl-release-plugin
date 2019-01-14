package kubectlreleaseplugin

// column names for the table
var Headers = []string{"NAME", "VERSION", "AGE", "URL"}

// names of the project to retrieve release information
var Projects = &[]Project{
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
