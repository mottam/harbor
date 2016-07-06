package docker

import (
	"fmt"
	"github.com/hashicorp/go-version"
	"os"
	"os/exec"
	"strings"
)

func GetDockerVersion(dockerVersion *string) {

	output, err := exec.Command("docker", "version", "--format='{{.Client.Version}}'").CombinedOutput()

	if err != nil {
		fmt.Println("Não consegui executar o comando 'docker'. Output: " + string(output))
		os.Exit(1)
	} else {
		*dockerVersion = strings.TrimSpace(string(output))
	}
}

// Retorna false se o docker for menor que a versão passada
// Returna true se o docker for maior ou igual que a versão passada
func CompareDockerVersion(dockerCmpVersion string) bool {
	var dockerVersion string
	GetDockerVersion(&dockerVersion)
	v1, _ := version.NewVersion(dockerVersion)
	v2, _ := version.NewVersion(dockerCmpVersion)

	if v1.LessThan(v2) {
		return false
	} else {
		return true
	}
}
