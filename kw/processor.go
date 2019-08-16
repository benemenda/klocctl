package kw

import "github.com/spf13/afero"

var AppFs = afero.NewOsFs()

func processIssues(issuesJSON []byte) {
	afero.WriteFile(AppFs, "issues.json", issuesJSON, 0644)
}

func processBuilds(buildsJSON []byte) {
	afero.WriteFile(AppFs, "builds.json", buildsJSON, 0644)
}

func processProjects(projectsJSON []byte) {
	afero.WriteFile(AppFs, "projects.json", projectsJSON, 0644)
}
