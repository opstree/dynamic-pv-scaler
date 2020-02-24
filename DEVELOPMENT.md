### Development

#### Pre-requisites

First of all, you need access on a Kubernetes cluster. The other dependencies are:-

- **[Kubernetes](https://kubernetes.io/)**
- **[Prometheus](https://prometheus.io/)**
- **[Golang](https://golang.org/)**

In storageclass, *allowVolumeExpansion* flag should be `true`

Example:-

```yaml
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: standard
provisioner: kubernetes.io/aws-ebs
parameters:
  type: gp2
reclaimPolicy: Retain
allowVolumeExpansion: true # This should be true for every storage class
mountOptions:
  - debug
volumeBindingMode: Immediate
```

#### Code Directories Overview

```s
dynamic-pv-scaler
├── api                  ---> API machinery for k8s client and prometheus client
├── deploy               ---> Deployment related files, for deployment on k8s
│   └── helm             ---> Fully functional helm chart
│       └── templates    ---> Helm templates
├── logger               ---> JSON logger interface for logging
├── pkg                  ---> K8s related operations like pv resizing, pod deletion
├── static               ---> Static images for documentation
└── utils                ---> Utilities which are being used in codebase like yaml to json conversion
```

#### Building Code

For building the code, execute `make` steps:-

```shell
# For building the code
make build-code

#  For building the image
make build-image
```

#### Checking Code Formatting

```shell
make check-fmt
make lint
```

#### Testing the Code

```shell
make test
```