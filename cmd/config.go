package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"errors"
	"os"
)

type (
	Configuration struct {
		GitRepo      string
		FilePath     string
		ManagedVersions []ManagedVersion
		GithubCreds GithubCreds
	}


	ManagedVersion struct {
		Name            string
		RegexCoordinate string
		DepResource     string
		VersionPattern  string `"major.minor.build"`
	}

	GithubCreds struct {
		Token string
	}

)

var configuration Configuration

func LoadConfigFile(configFilePath string) (Configuration, error) {
	fmt.Printf("// reading file %s\n", configFilePath)
	configFile, err1 := ioutil.ReadFile(configFilePath)

	if err1 != nil {
		fmt.Printf("// Found error reading file. ")
		return configuration, errors.New("Could not open file.")
	}

	err2 := json.Unmarshal(configFile, &configuration)


	if err2 != nil {
		return configuration, errors.New("Could not wrangle json into a Configuration.")
	}

	githubCreds := GithubCreds{}
	githubCreds.Token = os.Getenv("GITHUB_TOKEN")
	configuration.GithubCreds=githubCreds

	return configuration, nil
}
