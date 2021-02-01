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
❯ ./deploy.sh
```

Terminal 1

```bash
❯ kubectl exec -it -n dummy-com $(kubectl get pod -n dummy-com -l component=cli.peer0.org1.dummy.com -o jsonpath="{.items[0].metadata.name}") -- bash

❯ cd artifacts;
❯ peer channel create -c $CHANNEL_NAME -f ./channelall.tx -o orderer0-dummy-com:7050 --tls --cafile $ORDERER_CA
❯ peer channel join -b ./channelall.block
❯ peer channel list
```



Terminal 2


```bash
❯ kubectl exec -it -n dummy-com $(kubectl get pod -n dummy-com -l component=cli.peer0.org2.dummy.com -o jsonpath="{.items[0].metadata.name}") -- bash

❯ cd artifacts;
❯ peer channel join -b ./channelall.block
❯ peer channel list
```

## Cleanup

```bash
❯ kubectl delete pod,deployment,service,job,secrets,pvc --all --namespace dummy-com && kubectl delete pv local-volume
```
