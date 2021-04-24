# hyperledger-fabric-kubernetes

This repository aims to demonstrate how to deploy an Hyperledger Fabric [v2.3](https://hyperledger-fabric.readthedocs.io/en/release-2.3/) network on Kubernetes, and use chaincodes
as external services.

Legacy way of building and runing chaincodes required from the peer, a binding to the Docker socket for being able to
talk with the Docker daemon.

Using chaincode as an external service, the chaincode endpoint is deployed to the peer. The chaincode can be built and lauched separated from the peer. Therefore, there is no more dependency on the Kubernetes CRI implementation.

The source code herein is not production ready.

## Usage

- Assuming you have minikube running:

```bash
❯ ./deploy.sh
# wait until all resources are up and running
❯ watch -n 1 kubectl get all --namespace dummy-com
```

Now, we are going to operate the peer of Org1 and the peer of Org2 in order to create and join a channel as well as install the chaincode.

You'll need two different terminals and run the following commands:

- Terminal 1: CLI configured for the peer of Org1

```bash
# connect to the pod running the CLI for peer0.org1
❯ kubectl exec -it -n dummy-com $(kubectl get pod -n dummy-com -l component=cli.peer0.org1.dummy.com -o jsonpath="{.items[0].metadata.name}") -- bash

❯ cd artifacts;
❯ peer channel create -c $CHANNEL_NAME -f ./channelall.tx -o orderer0-dummy-com:7050 --tls --cafile $ORDERER_CA;
❯ peer channel join -b ./channelall.block;
❯ cd -;
```

- Terminal 2: CLI configured for the peer of Org2

```bash
# connect to the pod running the CLI for peer0.org2
❯ kubectl exec -it -n dummy-com $(kubectl get pod -n dummy-com -l component=cli.peer0.org2.dummy.com -o jsonpath="{.items[0].metadata.name}") -- bash

❯ cd artifacts;
❯ peer channel join -b ./channelall.block
❯ cd -;
```

At this point the channel has been created, and both peers of Org1 and Org2 have joined the channel.

We can now proceed to install the chaincode.

- Terminal 1: CLI configured for the peer of Org1

```bash
# install the chaincode
❯ peer lifecycle chaincode install ./chaincodes/chaincode-as-external-service/chaincode-as-external-service.tgz;
# approve the chaincode for Org1
❯ peer lifecycle chaincode approveformyorg  -o orderer0-dummy-com:7050 --tls --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name chaincode-as-external-service --version 1.0 --init-required --package-id chaincode-as-external-service:33b295bb4ac3f8dead7bddb9e86315aa7b3729b76d6d53f9379ddba6db900f7f --sequence 1
```

- Terminal 2: CLI configured for the peer of Org2

```bash
# install the chaincode
❯ peer lifecycle chaincode install ./chaincodes/chaincode-as-external-service/chaincode-as-external-service.tgz;
# approve the chaincode for Org2
❯ peer lifecycle chaincode approveformyorg  -o orderer0-dummy-com:7050 --tls --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name chaincode-as-external-service --version 1.0 --init-required --package-id chaincode-as-external-service:33b295bb4ac3f8dead7bddb9e86315aa7b3729b76d6d53f9379ddba6db900f7f --sequence 1
```

We are almost done, chaincode has been approved by both Org1 and Org2, we can now commit the chaincode before invoking it.

- Terminal 1: CLI configured for the peer of Org1

```bash
# You can check the commit readiness of the chaincode by running:
❯ peer lifecycle chaincode checkcommitreadiness -o orderer0-dummy-com:7050 --channelID $CHANNEL_NAME --tls --cafile $ORDERER_CA --name chaincode-as-external-service --version 1.0 --init-required --sequence 1

# It should output something like:
# Chaincode definition for chaincode 'chaincode-as-external-service', version '1.0', sequence '1' on channel 'channelall' approval status by org:
# Org1MSP: true
# Org2MSP: true
# Org3MSP: false


# commit the chaincode
❯ peer lifecycle chaincode commit -o orderer0-dummy-com:7050 --channelID $CHANNEL_NAME --name chaincode-as-external-service --version 1.0 --sequence 1 --init-required --tls --cafile $ORDERER_CA --peerAddresses peer0-org1-dummy-com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0-org2-dummy-com:7051 --tlsRootCertFiles /etc/hyperledger/fabric/crypto/peerOrganizations/org2.dummy.com/peers/peer0.org2.dummy.com/tls/ca.crt
```

That's it ! The chaincode is ready to be invoked :smile: .

- Terminal 1: CLI configured for the peer of Org1

```bash
# init and invoke the chaincode
❯ peer chaincode invoke -o orderer0-dummy-com:7050 --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n chaincode-as-external-service  --peerAddresses peer0-org1-dummy-com:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE  --peerAddresses peer0-org2-dummy-com:7051 --tlsRootCertFiles /etc/hyperledger/fabric/crypto/peerOrganizations/org2.dummy.com/peers/peer0.org2.dummy.com/tls/ca.crt --isInit -c '{"function":"Init","Args":[]}'
```

- Terminal 2: CLI configured for the peer of Org2

```bash
# query the chaincode
❯ peer chaincode query -C $CHANNEL_NAME -n chaincode-as-external-service -c '{"Args":["Query", "default-asset"]}'
```

## Cleanup

```bash
❯ kubectl delete pod,deployment,service,job,secrets,pvc --all --namespace dummy-com && kubectl delete pv local-volume
```
