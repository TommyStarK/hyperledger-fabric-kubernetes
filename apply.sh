#!/usr/bin/env bash

kubectl apply -f namespaces/
kubectl apply -f secrets/org1
kubectl apply -f services/
kubectl apply -f services/org1
kubectl apply -f services/org2
kubectl apply -f services/org3
kubectl apply -f deployments/
kubectl apply -f deployments/org1
kubectl apply -f deployments/org2
kubectl apply -f deployments/org3

