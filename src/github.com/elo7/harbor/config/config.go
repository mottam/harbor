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
		Region   string
	}
	DownloadPath string `yaml:",omitempty"`
	ProjectPath  string `yaml:",omitempty"`
	Files        []HarborFile
	Commands     []string
	BuildArgs    map[string]string
}

func Load(cliConfigVars commandline.ConfigVarsMap, projectPath string, configFile string) (HarborConfig, error) {
	harborConfig := HarborConfig{}

	// Loads config file contents
	config, err := LoadFile(configFile)
	if err != nil {
		return harborConfig, err
	}

	// Replaces variables (${KEY} format) from -e parameter
	config = SetEnv(cliConfigVars, config)

	// Parses file content (YAML expected)
	err = yaml.Unmarshal(config, &harborConfig)

	harborConfig.ProjectPath = projectPath

	if err != nil {
		return harborConfig, err
	}

	return harborConfig, nil
}

func LoadFile(configFile string) ([]byte, error) {
	return ioutil.ReadFile(configFile)
}
