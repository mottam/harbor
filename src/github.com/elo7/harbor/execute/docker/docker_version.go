package docker

import (
	"github.com/hashicorp/go-version"
	"os/exec"
	"strings"
)

func GetDockerVersion() (dockerVersion string, err error) {
	output, err := exec.Command("docker", "version", "--format='{{.Client.Version}}'").CombinedOutput()
	return strings.TrimSpace(string(output)), err
}

// Retorna false se o docker for menor que a versão passada
// Returna true se o docker for maior ou igual que a versão passada
func CompareDockerVersion(dockerCmpVersion string) bool {
	dockerVersion, _ := GetDockerVersion()
	v1, _ := version.NewVersion(dockerVersion)
	v2, _ := version.NewVersion(dockerCmpVersion)

	if v1.LessThan(v2) {
		return false
	} else {
		return true
	}
}
