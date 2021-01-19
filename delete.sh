#!/usr/bin/env bash

kubectl delete -f deployments/org3
kubectl delete -f deployments/org2
kubectl delete -f deployments/org1
kubectl delete -f deployments/
kubectl delete -f services/org3
kubectl delete -f services/org2
kubectl delete -f services/org1
kubectl delete -f services/
kubectl delete -f secrets/org1
kubectl delete -f namespaces/

