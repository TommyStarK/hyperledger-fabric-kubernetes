# hyperledger-fabric-kubernetes

This repository aims to demonstrate how to deploy an Hyperledger Fabric [v2.3](https://hyperledger-fabric.readthedocs.io/en/release-2.3/) network on Kubernetes, and use chaincodes
as external services.

Legacy way of building and runing chaincodes required from the peer, a binding to the Docker socket for being able to
talk with the Docker daemon.

Using chaincode as an external service, the chaincode endpoint is deployed to the peer. The chaincode can be built and lauched separated from the peer. Therefore, there is no more dependency on the Kubernetes CRI implementation.

The source code herein is not production ready.

## Usage

- Assuming you have a running cluster or minikube:

```bash
❯ ./apply.sh
```

- Delete the network

```bash
❯ kubectl delete pod,deployment,service,job,secrets,pvc --all --namespace dummy-com && kubectl delete pv local-volume
```
