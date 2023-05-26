package docker_build

import (
	"fmt"
	"github.com/jonas-be/papermcdl/pkg/paper_api"
	"log"
	"os"
	"os/exec"
	"strings"
)

const serverJarName = "server.jar"

type ImageBuilder struct {
	Api paper_api.PapermcAPI
}

func (i ImageBuilder) BuildAllProjects() error {
	projects, err := i.Api.GetProjects()
	if err != nil {
		fmt.Printf("Error getting projects: %v", err)
		return err
	}

	for _, project := range projects.Projects {
		versions, err := i.Api.GetVersions(project)
		if err != nil {
			fmt.Printf("Error getting versions of %s: %v", project, err)
			return err
		}
		err = i.BuildVersions(project, versions.Versions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i ImageBuilder) BuildAllVersions(project string) error {
	versions, err := i.Api.GetVersions(project)
	if err != nil {
		return err
	}
	err = i.BuildVersions(project, versions.Versions)
	if err != nil {
		return err
	}
	return nil
}

func (i ImageBuilder) BuildVersions(project string, versions []string) error {
	for _, version := range versions {
		builds, err := i.Api.GetBuilds(project, version)
		if err != nil {
			return err
		}
		latestBuild, err := builds.GetLatestBuild()
		if err != nil {
			return err
		}
		err = i.BuildDockerImage(project, version, latestBuild)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i ImageBuilder) BuildDockerImage(project string, version string, build string) error {
	err := i.deleteServerJAR()
	if err != nil {
		return err
	}
	err = i.downloadServerJAR(project, version, build)
	if err != nil {
		return err
	}
	dockerBuildCommand := fmt.Sprintf("docker build -t mc8s/%s:%s .", project, version)
	executeCommand(dockerBuildCommand)
	return nil
}

func executeCommand(command string) {
	fmt.Printf("[%s]: ", command)
	arr := strings.Split(command, " ")
	cmd := exec.Command(arr[0], arr[1:]...)

	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(output))
	fmt.Println("| Done")
}

func (i ImageBuilder) downloadServerJAR(project string, version string, build string) error {
	info, err := i.Api.GetBuildInfo(project, version, build)
	if err != nil {
		return err
	}
	filename := i.Api.GetFileName(info)
	err = i.Api.Download(project, version, build)
	if err != nil {
		return err
	}

	err = os.Rename(filename, serverJarName)
	if err != nil {
		return err
	}
	return nil
}

func (i ImageBuilder) deleteServerJAR() error {
	err := os.Remove(serverJarName)
	if err != nil {
		return err
	}
	return nil
}
