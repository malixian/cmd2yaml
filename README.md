# cmd2yaml
This tool can easily convert docker commands into pod yaml of k8s

## Introduction
We are usually more familiar with docker commands than k8s yaml file.
so this tool can easily transform docker commands you offer into yaml file for k8s Pod API.

## Install

sh build.sh

## Example

use c2y like this
```
c2y -i 'docker run --network host -v /tmp:/tmp -p 12345:12345 --device /dev/ your-container bash -c "python test.py' -n "test"
```

then you can find test.yaml in current path

```
apiVersion: v1
kind: Pod
metadata:
  name: test
  labels:
    name: test
spec:
  hostNetwork: true
  containers:
  - name: your-container
    image: your-container
    command:
    - bash
    - -c
    args:
    - python test.py
    ports:
    - hostPort: 12345
      containerPort: 12345
    volumeMounts:
    - name: ssh
      mountPath: /tmp
  volumes:
  - name: ssh
    hostPath:
      path: /tmp
```

## Attention

Now support docker flag has:
- command
- args
- volumeMount
- network host

Now support pod options has:
- containers
- volumes
