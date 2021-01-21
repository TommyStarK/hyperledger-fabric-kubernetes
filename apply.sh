#!/usr/bin/env bash

kubectl apply -f k8s/namespaces/
kubectl apply -f k8s/secrets/org1.dummy.com/
kubectl apply -f k8s/storage/pv
kubectl apply -f k8s/storage/pvc
kubectl apply -f k8s/job
kubectl wait --for=condition=complete --namespace dummy-com job.batch/setup
kubectl apply -f k8s/services/
kubectl apply -f k8s/services/org1.dummy.com/
kubectl apply -f k8s/services/org2.dummy.com/
kubectl apply -f k8s/services/org3.dummy.com/
kubectl apply -f k8s/deployments/
kubectl apply -f k8s/deployments/org1.dummy.com/
kubectl apply -f k8s/deployments/org2.dummy.com/
kubectl apply -f k8s/deployments/org3.dummy.com/
