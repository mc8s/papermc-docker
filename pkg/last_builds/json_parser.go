package last_builds

import (
	"encoding/json"
	"os"
)

type JSONParser struct {
	FilePath string
}

func (j JSONParser) EnsureExists() error {
	_, err := os.Stat(j.FilePath)
	if os.IsNotExist(err) {
		file, err := os.Create(j.FilePath)
		if err != nil {
			return err
		}
		defer file.Close()
		err = j.SaveLastBuilds(LastBuilds{})
	}
	if err != nil && err.Error() != "CreateFile last_builds.json: The system cannot find the file specified." {
		return err
	}
	return nil
}

func (j JSONParser) GetLastBuilds() (LastBuilds, error) {
	fileContent, err := os.ReadFile(j.FilePath)
	if err != nil {
		return LastBuilds{}, err
	}

	var lastBuilds LastBuilds
	err = json.Unmarshal(fileContent, &lastBuilds)
	if err != nil {
		return LastBuilds{}, err
	}

	return lastBuilds, nil
}

func (j JSONParser) SaveLastBuilds(lastBuilds LastBuilds) error {
	fileContent, err := json.MarshalIndent(lastBuilds, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(j.FilePath, fileContent, 0644)
	if err != nil {
		return err
	}

	return nil
}
