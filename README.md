# kubectl-release-plugin
This project contains a simple kubectl plugin to retrieve the latest release of your favourite projects around kubernetes, usage as follow:
```
[qasim.sarfraz@~ ]$  kubectl latest-release -h
Usage:
  kubectl-latest_release [OPTIONS]

Application Options:
  -p, --project= Latest release for the given project
  -l, --list     List of the allowed projects

Help Options:
  -h, --help     Show this help message

```

## Example:
List all the valid projects using:
```
[qasim.sarfraz@~ ]$  kubectl latest-release -l
NAME           OWNER
kubernetes     kubernetes
kops           kubernetes
istio          istio
helm           helm
ingress-nginx  kubernetes
etcd           etcd-io
traefik        containous
faas           openfaas
rancher        rancher
kubeflow       kubeflow
kubespray      kubernetes-sigs
cluster-api    kubernetes-sigs
rook           rook
kubeless       kubeless
flannel        coreos
cilium         cilium
argo           argoproj
patroni        zalando
external-dns   kubernetes-incubator
pulumi         pulumi
linkerd2       linkerd
```
Check the release for a given project:
```
[qasim.sarfraz@~ ]$  kubectl latest-release -p helm
NAME  VERSION  AGE     URL
helm  v2.12.2  4 days  https://github.com/helm/helm/releases/tag/v2.12.2
```
## Installation:
You can install the plugin using:
```
go get -u github.com/mqasimsarfraz/kubectl-release-plugin/cmd/kubectl-latest_release
```

## Requirements:

- `kubectl` > `1.12.0`
