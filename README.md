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
# wait until all resources are up and running
❯ watch -n 0.1 kubectl get all --namespace dummy-com
```

Terminal 1

```bash
❯ kubectl exec -it -n dummy-com $(kubectl get pod -n dummy-com -l component=cli.peer0.org1.dummy.com -o jsonpath="{.items[0].metadata.name}") -- bash

❯ cd artifacts;
❯ peer channel create -c $CHANNEL_NAME -f ./channelall.tx -o orderer0-dummy-com:7050 --tls --cafile $ORDERER_CA;
❯ peer channel join -b ./channelall.block;
❯ cd -;
❯ peer lifecycle chaincode install ./chaincodes/chaincode-as-external-service/chaincode-as-external-service.tgz;
❯ peer lifecycle chaincode approveformyorg  -o orderer0-dummy-com:7050 --tls --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name chaincode-as-external-service --version 1.0 --init-required --package-id chaincode-as-external-service:33b295bb4ac3f8dead7bddb9e86315aa7b3729b76d6d53f9379ddba6db900f7f --sequence 1 --signature-policy "AND ('Org1MSP.peer','Org2MSP.peer')"
❯ peer lifecycle chaincode queryapproved -C $CHANNEL_NAME -n chaincode-as-external-service --sequence 1
❯ peer lifecycle chaincode checkcommitreadiness -o orderer0-dummy-com:7050 --channelID $CHANNEL_NAME --tls --cafile $ORDERER_CA --name chaincode-as-external-service --version 1.0 --init-required --sequence 1
```

Terminal 2

```bash
❯ kubectl exec -it -n dummy-com $(kubectl get pod -n dummy-com -l component=cli.peer0.org2.dummy.com -o jsonpath="{.items[0].metadata.name}") -- bash

❯ cd artifacts;
❯ peer channel join -b ./channelall.block
❯ cd -;
❯ peer lifecycle chaincode install ./chaincodes/chaincode-as-external-service/chaincode-as-external-service.tgz;
❯ peer lifecycle chaincode approveformyorg  -o orderer0-dummy-com:7050 --tls --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name chaincode-as-external-service --version 1.0 --init-required --package-id chaincode-as-external-service:33b295bb4ac3f8dead7bddb9e86315aa7b3729b76d6d53f9379ddba6db900f7f --sequence 1 --signature-policy "AND ('Org1MSP.peer','Org2MSP.peer')"
❯ peer lifecycle chaincode queryapproved -C $CHANNEL_NAME -n chaincode-as-external-service --sequence 1
❯ peer lifecycle chaincode checkcommitreadiness -o orderer0-dummy-com:7050 --channelID $CHANNEL_NAME --tls --cafile $ORDERER_CA --name chaincode-as-external-service --version 1.0 --init-required --sequence 1
```

## Cleanup

```bash
❯ kubectl delete pod,deployment,service,job,secrets,pvc --all --namespace dummy-com && kubectl delete pv local-volume
```
