package pkg

import (
	"fmt"
	"strings"

	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

// docker run --network host -v /tmp:/tmp -p 12345:12345 --device /dev/ your-container bash -c "python test.py"

func Show(cmd string, name string) {
	if data, err := cmd2yaml(cmd, name); err != nil {
		fmt.Println("========= trans yaml error =======")
	} else {
		writeYamlToLocal(name+".yaml", data)
		fmt.Println(data)
	}
}

func writeYamlToLocal(filename string, data string) error{
	var d = []byte(data)
	if err := ioutil.WriteFile(filename, d, 0666); err != nil {
		return  fmt.Errorf("write data error is %v", err)
	} else {
		return nil
	}
}

func cmd2yaml(cmd string, name string) (string, error) {
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

	image, cmds, args := getImage(strs)

	genYaml := initYaml(name, image, cmds, args)

	for i := 2; i < len(strs); i++ {
		str := strs[i]
		if isFlag(str) {
			addFlag2Yaml(genYaml, i, strs)
		} else if str == image {
			break
		}
	}

	data, err := yaml.Marshal(genYaml)
	if err != nil {
		return "", fmt.Errorf("marshal yaml failed, error is %v", err)
	}
	return string(data), nil

}

func isFlag(input string) bool {
	if len(input) >= 2 {
		return input[0] == '-'
	}
	return false
}

// genrate a yaml template，which You can add elements later
func initYaml(name string, image string, cmds []string, args string) *Pod {
	podSpec := &Pod{
		TypeMeta: TypeMeta{
			APIVersion: APIVERSION,
			Kind:       KIND,
		},
		ObjectMeta: ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"name": name,
			},
		},
		Spec: PodSpec{
				Containers: []Container{
					Container{
						Name:    image,
						Image:   image,
						Command: cmds,
						Args:    []string{args},
					},
				},
			},
	}
	return podSpec
}

func addFlag2Yaml(yaml *Pod, i int, strs []string) error {
	flag := strs[i]
	flag = strings.ReplaceAll(flag, "-", "")
	if len(yaml.Spec.Containers) == 0 {
		return fmt.Errorf("template no one container")
	}
	switch flag {
	case "network":
		addNetWork(yaml, i, strs)
	case "p":
		if err := addPort(yaml, i, strs); err != nil {
			return err
		}
	case "v":
		addVolume(yaml, i, strs)

	}
	return nil
}




func getImage(strs []string) (image string, cmd []string, args string) {

	i := 0
	for i = 2; i < len(strs); i++ {
		if i-1 >=0 && i+1 < len(strs) && !isFlag(strs[i-1]) && !isFlag(strs[i]) && !isFlag(strs[i+1]) {
			image = strs[i]
			break
		}
	}

	// get cmd list
	i++
	for ; i < len(strs); i++ {
		if !strings.Contains(strs[i], `"`) {
			cmd = append(cmd, strs[i])
		} else {
			break
		}
	}

	// get args
	for ; i < len(strs); i++ {
		tmp := strings.ReplaceAll(strs[i], `"`, "")
		args += " " + tmp
	}
	args = strings.Trim(args, " ")

	return image, cmd, args
}
