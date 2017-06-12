package config

import (
	"io/ioutil"

	"github.com/elo7/harbor/commandline"
	yaml "gopkg.in/yaml.v2"
)

//Options harbor options
var Options struct {
	Debug        bool
	DockerOpts   string
	NoDockerPush bool
	NoLatestTag  bool
}

//HarborFile harbor file config structure
type HarborFile struct {
	S3Path     string
	FileName   string
	Permission int
}

//HarborConfig harbor config structure
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

//Load load configs from file
func Load(cliConfigVars commandline.ConfigVarsMap, projectPath string, configFile string) (HarborConfig, error) {
	// Loads config file contents
	config, err := ioutil.ReadFile(configFile)
	if err != nil {
		return HarborConfig{}, err
	}

	harborConfig := &HarborConfig{}

	// Replaces variables (${KEY} format) from -e parameter
	config = setEnv(cliConfigVars, config)

	// Parses file content (YAML expected)
	if err = yaml.Unmarshal(config, harborConfig); err != nil {
		return HarborConfig{}, err
	}
	harborConfig.ProjectPath = projectPath
	return *harborConfig, nil
}
