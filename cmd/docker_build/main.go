package main

import (
	"fmt"
	"github.com/jonas-be/papermcdl/pkg/paper_api"
	"papermc-docker/pkg/docker_build"
	"papermc-docker/pkg/docker_hub"
)

func main() {
	api := paper_api.PapermcAPI{URL: "https://papermc.io"}
	builder := docker_build.ImageBuilder{Api: api}

	printTags()

	err := builder.BuildAllVersions("paper")
	if err != nil {
		return
	}
}

func printTags() {
	allTags := docker_hub.GetTags("jonasbe25/my-portfolio")
	for _, image := range allTags {
		fmt.Printf("Name: %s\n", image.Name)
		fmt.Printf("Last Updated: %s\n\n", image.LastUpdated)
	}
	fmt.Println(allTags)
	fmt.Println(len(allTags))
}
