package docker_hub

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type DockerImage struct {
	LastUpdated string `json:"last_updated"`
	Name        string `json:"name"`
}

type APIResponse struct {
	Count    int           `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []DockerImage `json:"results"`
}

func GetTags(imageName string) (allTags []DockerImage) {
	baseURL := fmt.Sprintf("https://hub.docker.com/v2/repositories/%s/tags/", imageName)
	for baseURL != "" {
		resp, err := http.Get(baseURL)
		if err != nil {
			log.Fatal(err)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(resp.Body)

		var apiResponse APIResponse
		err = json.NewDecoder(resp.Body).Decode(&apiResponse)
		if err != nil {
			log.Fatal(err)
		}

		allTags = append(allTags, apiResponse.Results...)
		baseURL = apiResponse.Next
	}

	return allTags
}
