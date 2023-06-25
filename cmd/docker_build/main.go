package main

import (
	"flag"
	"fmt"
	"github.com/jonas-be/papermcdl/pkg/paper_api"
	"os"
	"papermc-docker/pkg/docker_build"
	"papermc-docker/pkg/last_builds"
)

func main() {
	project, ok := getConfiguredProject()
	if !ok {
		fmt.Println("No project configured")
		return
	}

	api := paper_api.PapermcAPI{URL: "https://papermc.io"}

	if projects, err := api.GetProjects(); err != nil {
		fmt.Println("Error getting projects: ", err)
		return
	} else {
		if !sliceContains(projects.Projects, project) {
			fmt.Printf("Project %v not found\n", project)
			return
		}
	}

	parser := last_builds.JSONParser{FilePath: fmt.Sprintf("last-builds-%v.json", project)}
	err := parser.EnsureExists()
	if err != nil {
		fmt.Println("Error ensuring file exists: ", err)
		return
	}
	lastBuilds, err := parser.GetLastBuilds()
	if err != nil {
		fmt.Println("Error getting last builds: ", err)
		return
	}
	builder := docker_build.ImageBuilder{Api: api, LastBuilds: &lastBuilds}

	err = builder.BuildAllVersions(project)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = parser.SaveLastBuilds(*builder.LastBuilds)
	if err != nil {
		fmt.Println("Error saving last builds: ", err)
		return
	}
}

func getConfiguredProject() (string, bool) {
	project := os.Getenv("PROJECT")
	if project != "" {
		return project, true
	}
	var projectFlag = flag.String("project", "", "specify the project to build")
	flag.Parse()
	if projectFlag != nil && *projectFlag != "" {
		return *projectFlag, true
	}
	return "", false
}

func sliceContains(slice []string, element string) bool {
	for _, sliceElement := range slice {
		if sliceElement == element {
			return true
		}
	}
	return false
}
