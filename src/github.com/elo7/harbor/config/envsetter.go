package config

import (
	"fmt"
	"github.com/elo7/harbor/commandline"
	"regexp"
	"strings"
)

func SetEnv(cliConfigVars commandline.ConfigVarsMap, configString []byte) []byte {
	str := string(configString)

	// FIXME: Parallelize replace
	for key, value := range cliConfigVars {
		str = strings.Replace(str, fmt.Sprintf("${%s}", key), value, -1)
	}

	return []byte(str)
}

func ReadEnv(configString []byte) []string {
	matcher := regexp.MustCompile(`\$\{[a-zA-Z0-9_\-]+\}`)

	return matcher.FindAllString(string(configString), -1)
}
