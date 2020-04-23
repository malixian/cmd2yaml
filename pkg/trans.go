package pkg

import (
	"strings"
	"fmt"

	"gopkg.in/yaml.v2"

	v1 "k8s.io/api/core/v1"

)

// docker run --network host -v /tmp:/tmp -p 12345:12345 --device /dev/ your-container bash -c "python test.py"

func Cmd2yaml(cmd string) (string, error) {
	strings.Trim(cmd, " ")
	strs := strings.Split(cmd, " ")
	// min command is docker run img bash
	if len(strs) < 4 {
		return "", fmt.Errorf("input cmd is invalid")
	}

	if strs[0] != "docker" {
		return "", fmt.Errorf("only support docker run")
	}

	if strs[1] != "run" {
		return "", fmt.Errorf("only support docker run")
	}

	for i:=2; i<len(strs); i++ {
		str := strs[i]
		if isFlag(str) {
			addFlag2Yaml(str)
		} else {
			addArg2Yaml(str)
		}
	}
}

func isFlag(input string) bool {
	for i:=0; i<len(input); i++ {
		if input[i] == '-' {
			return true
		}
	}
	return false
}

func initYaml() {
	podSpec := &v1.PodSpec{
		
	}
}

func addFlag2Yaml(flag string) {

}

func addArg2Yaml(arg string) {

}
