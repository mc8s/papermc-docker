package main

import (
	"fmt"
	"github.com/jonas-be/papermcdl/pkg/paper_api"
	"papermc-docker/pkg/docker_build"
	"papermc-docker/pkg/last_builds"
)

func main() {
	api := paper_api.PapermcAPI{URL: "https://papermc.io"}
	parser := last_builds.JSONParser{FilePath: "last-builds.json"}
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

	err = builder.BuildAllVersions("paper")
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
