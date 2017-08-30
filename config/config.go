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
	Repository string
}

//HarborConfig harbor config structure
type HarborConfig struct {
	ImageTag       string
	CliConfigVars  commandline.ConfigVarsMap
	Tags           []string
	S3             S3HarborConfig //retrocompatibility
	S3Repositories S3HarborRepositoriesConfig
	DownloadPath   string `yaml:",omitempty"`
	ProjectPath    string `yaml:",omitempty"`
	Files          []HarborFile
	Commands       []string
	BuildArgs      map[string]string
}

//S3HarborConfig harbor s3 config structure
type S3HarborConfig struct {
	Bucket   string
	BasePath string
	Region   string
}

//S3HarborRepositoriesConfig map with harbor s3 config structure
type S3HarborRepositoriesConfig map[string]S3HarborConfig

//Get get harbor s3 config, return default if not found
func (c S3HarborRepositoriesConfig) Get(key string) S3HarborConfig {
	if result, exists := c[key]; exists {
		return result
	}
	return c["default"]
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
	if harborConfig.S3Repositories == nil { //S3 -> S3Repositories retrocompatibility
		harborConfig.S3Repositories = make(S3HarborRepositoriesConfig)
		harborConfig.S3Repositories["default"] = harborConfig.S3
	}

	return *harborConfig, nil
}
