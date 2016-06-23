package config

import (
	"github.com/elo7/harbor/commandline"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

var Options struct {
	Debug        bool
	DockerOpts   string
	NoDockerPush bool
	NoLatestTag  bool
}

type HarborFile struct {
	S3Path     string
	FileName   string
	Permission int
}

type HarborConfig struct {
	ImageTag      string
	CliConfigVars commandline.ConfigVarsMap
	Tags          []string
	S3            struct {
		Bucket   string
		BasePath string
		Region	 string
	}
	DownloadPath string `yaml:",omitempty"`
	ProjectPath  string `yaml:",omitempty"`
	Files        []HarborFile
	Commands     []string
}

func Load(cliConfigVars commandline.ConfigVarsMap, projectPath string) (HarborConfig, error) {
	harborConfig := HarborConfig{}

	configFile, err := LoadFile(projectPath)
	if err != nil {
		return harborConfig, err
	}

	configFile = SetEnv(cliConfigVars, configFile)

	err = yaml.Unmarshal(configFile, &harborConfig)

	harborConfig.ProjectPath = projectPath

	if err != nil {
		return harborConfig, err
	}

	return harborConfig, nil
}

func LoadFile(projectPath string) ([]byte, error) {
	return ioutil.ReadFile(projectPath + "/harbor.yml")
}
