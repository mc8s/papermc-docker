package main

import (
	"fmt"
	"github.com/jonas-be/papermcdl/pkg/paper_api"
	"os"
	"papermc-docker/internal/docker_hub"
)

const serverJarName = "server.jar"

func main() {
	printTags()
}

func printTags() {
	allTags := docker_hub.GetTags("library/alpine")
	for _, image := range allTags {
		fmt.Printf("Name: %s\n", image.Name)
		fmt.Printf("Last Updated: %s\n\n", image.LastUpdated)
	}
	fmt.Println(allTags)
	fmt.Println(len(allTags))
}

func buildAllVersions() {
	api := paper_api.PapermcAPI{URL: "https://papermc.io"}

	projects, err := api.GetProjects()
	if err != nil {
		fmt.Printf("error getting projects: %v", err)
		return
	}

	for _, project := range projects.Projects {
		versions, err := api.GetVersions(project)
		if err != nil {
			fmt.Printf("error getting versions of %s: %v", project, err)
			return
		}
		buildVersions(api, project, versions)
	}
}

func buildVersions(api paper_api.PapermcAPI, project string, versions paper_api.Versions) {
	for _, version := range versions.Versions {
		builds, err := api.GetBuilds(project, version)
		if err != nil {
			fmt.Printf("error getting builds of %s %s: %v", project, version, err)
			return
		}
		latestBuild, err := builds.GetLatestBuild()
		if err != nil {
			fmt.Printf("error getting latest build of %s %s: %v", project, version, err)
			return
		}
		buildDockerImage(api, project, version, latestBuild)
	}
}

func buildDockerImage(api paper_api.PapermcAPI, project string, version string, build string) {
	downloadServerJAR(api, project, version, build)
	fmt.Printf("docker build -t mc8s/%s:%s-%s .\n", project, version, build)
}

func downloadServerJAR(api paper_api.PapermcAPI, project string, version string, build string) {
	info, err := api.GetBuildInfo(project, version, build)
	if err != nil {
		fmt.Printf("error getting buildinfo of %s %s %s: %v", project, version, build, err)
		return
	}
	filename := api.GetFileName(info)
	err = api.Download(project, version, build)
	if err != nil {
		fmt.Printf("error downloading %s: %v", filename, err)
		return
	}

	err = os.Rename(filename, serverJarName)
	if err != nil {
		fmt.Printf("Failed to rename file: %s\n", err.Error())
		return
	}
}
