#!/usr/bin/env bash

kubectl apply -f namespaces/
kubectl apply -f secrets/org1.dummy.com/
kubectl apply -f services/
kubectl apply -f services/org1.dummy.com/
kubectl apply -f services/org2.dummy.com/
kubectl apply -f services/org3.dummy.com/
kubectl apply -f deployments/
kubectl apply -f deployments/org1.dummy.com/
kubectl apply -f deployments/org2.dummy.com/
kubectl apply -f deployments/org3.dummy.com/
