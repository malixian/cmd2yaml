package pkg

const (
	APIVERSION = "v1"
	KIND = "Pod"
)

type TypeMeta struct {
	APIVersion string  `yaml:"apiVersion,omitempty"`
	Kind string `yaml:"kind,omitempty"`
}

type ObjectMeta struct {
	Name string `yaml:"name,omitempty"`
	Labels map[string]string `yaml:"labels,omitempty"`
	Namespace string `yaml:"namespace,omitempty"`
}


type PodSpec struct {
	HostNetwork bool `yaml:"hostNetwork,omitempty"`
	Containers []Container `yaml:"containers" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=containers"`
	Volumes []Volume `yaml:"volumes,omitempty"`
}


type Pod struct{
	TypeMeta `yaml:",inline"`
	ObjectMeta `yaml:"metadata,omitempty"`
	Spec PodSpec `yaml:"spec,omitempty"`
}

type HostPathVolumeSource struct {
	Path string `yaml:"path" protobuf:"bytes,1,opt,name=path"`
}

type VolumeSource struct {
	HostPath *HostPathVolumeSource `yaml:"hostPath,omitempty" protobuf:"bytes,1,opt,name=hostPath"`
}

type Volume struct {
	Name string `yaml:"name" protobuf:"bytes,1,opt,name=name"`
	VolumeSource `yaml:",inline" protobuf:"bytes,2,opt,name=volumeSource"`
}

type ContainerPort struct {
	Name string `yaml:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	HostPort int32 `yaml:"hostPort,omitempty" protobuf:"varint,2,opt,name=hostPort"`
	ContainerPort int32 `yaml:"containerPort" protobuf:"varint,3,opt,name=containerPort"`
}

type VolumeMount struct {
	Name string `yaml:"name" protobuf:"bytes,1,opt,name=name"`
	MountPath string `yaml:"mountPath" protobuf:"bytes,3,opt,name=mountPath"`
}

type Container struct {
	Name string `yaml:"name" protobuf:"bytes,1,opt,name=name"`
	Image string `yaml:"image,omitempty" protobuf:"bytes,2,opt,name=image"`
	Command []string `yaml:"command,omitempty" protobuf:"bytes,3,rep,name=command"`
	Args []string `yaml:"args,omitempty" protobuf:"bytes,4,rep,name=args"`
	Ports []ContainerPort `yaml:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"containerPort" protobuf:"bytes,6,rep,name=ports"`
	VolumeMounts []VolumeMount `yaml:"volumeMounts,omitempty" patchStrategy:"merge" patchMergeKey:"mountPath" protobuf:"bytes,9,rep,name=volumeMounts"`
}


