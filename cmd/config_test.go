package cmd

import (
	"testing"
)

func TestLoadConfigFileNoFile(t *testing.T) {
	_, err := LoadConfigFile("")
	expectedError := "Could not open file."
	if err.Error() != expectedError {
		t.Error("Expected error")
	}
}

func TestConfigurationStructFilePath(t *testing.T) {
	config, _ := LoadConfigFile("../config/conf_test.json")
	if config.FilePath != "./Dockerfile" {
		t.Error("Configuration file path was not as expected")
	}

}

func TestConfigurationStructFileGitRepo(t *testing.T) {
	config, _ := LoadConfigFile("../config/conf_test.json")
	if config.GitRepo != "git@github.com:samsung-cnct/k2-tools.git" {
		t.Error("Configuration file path was not as expected")
	}

}

func TestConfigurationStructFileGitHubCreds(t *testing.T) {
	config, _ := LoadConfigFile("../config/conf_test.json")
	if config.GithubCreds.Token == "" {
		t.Error("Configuration file path was not as expected")
	}

}

func TestConfigurationStuctManagedVersionsLength(t *testing.T) {
	config, _ := LoadConfigFile("../config/conf_test.json")
	if len(config.ManagedVersions) != 7 {
		t.Error("The number of managed versions was not as expected")
	}
}
