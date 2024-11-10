# kube-playground

## Installation
Tested on Fedora 41 workstation with 32GB RAM and 8CPU.
```
# Download and install minikube
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube && rm minikube-linux-amd64

# Configure minikube to use more resources and start
minikube config set rootless true
minikube start --cpus=6 --memory=24000 --disk-size=50g --driver=podman --container-runtime=containerd
minikube addons enable ingress

# Set up shortcuts
alias kubectl="minikube kubectl --"
source <(kubectl completion bash)
```