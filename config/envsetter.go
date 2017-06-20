package config

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/elo7/harbor/commandline"
)

func setEnv(cliConfigVars commandline.ConfigVarsMap, configString []byte) []byte {
	str := string(configString)

	// FIXME: Parallelize replace
	for key, value := range cliConfigVars {
		str = strings.Replace(str, fmt.Sprintf("${%s}", key), value, -1)
	}

	return []byte(str)
}

//ReadEnv read envs from file
func ReadEnv(configString []byte) []string {
	return regexp.MustCompile(`\$\{[a-zA-Z0-9_\-]+\}`).FindAllString(string(configString), -1)
}
