package pkg

import (
	"strings"
	"strconv"
	"fmt"
)

func addNetWork(yaml *Pod, i int, strs []string) {
	spec := yaml.Spec
	fmt.Println("network is ", strs[i+1])
	tmp := strings.Trim(strs[i+1], " ")
	if i+1 < len(strs) && tmp == "host" {
		spec.HostNetwork = true
		yaml.Spec = spec
	}
}

func addPort(yaml *Pod, i int, strs []string) error{
	strsize := len(strs)
	container := yaml.Spec.Containers[0]
	tmp := strings.Trim(strs[i+1], " ")
	fmt.Println("port is ", tmp)
	if i+1 < strsize && len(strings.Split(tmp, ":")) == 2 {
		ports := strings.Split(strs[i+1], ":")
		conPort, err1 := strconv.Atoi(ports[0])
		hostPort, err2 := strconv.Atoi(ports[1])
		if err1 != nil || err2 != nil {
			return fmt.Errorf("Parse port error is err1 is %v, err2 is %v", err1, err2)
		}
		container.Ports = append(container.Ports, ContainerPort{ContainerPort: int32(conPort), HostPort: int32(hostPort)})
		yaml.Spec.Containers[0] = container
	}
	return nil
}

func addVolume(yaml *Pod, i int, strs []string) {
	strsize := len(strs)
	spec := yaml.Spec
	container := spec.Containers[0]

	if i+1 < strsize && len(strings.Split(strs[i+1], ":")) == 2 {
		mounts := strings.Split(strs[i+1], ":")
		// todo change fixed mount name
		mountName := "ssh"
		hostPath := mounts[0]
		mountPath := mounts[1]
		container.VolumeMounts = append(container.VolumeMounts, VolumeMount{Name: mountName, MountPath: mountPath})
		spec.Volumes = append(spec.Volumes, Volume{Name: mountName, VolumeSource: VolumeSource{HostPath: &HostPathVolumeSource{Path: hostPath}}})
		yaml.Spec = spec
		yaml.Spec.Containers[0] = container
	}

}