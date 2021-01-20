# hyperledger-fabric-kubernetes

The source code herein is not production ready. It is a demonstration of how to deploy [this](https://github.com/TommyStarK/hyperledger-fabric-network/tree/kubernetes) Hyperledger Fabric network on Kubernetes.

## Usage

- Assuming you have a running cluster or minikube:

```bash
❯ ./apply.sh
```

- Delete the network

```bash
❯ kubectl delete all --all --namespace dummy-com
```
