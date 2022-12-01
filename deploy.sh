#!/usr/bin/env bash

kubectl apply -f k8s/namespaces
kubectl apply -f k8s/secrets/org1.dummy.com
kubectl apply -f k8s/secrets/org2.dummy.com
kubectl apply -f k8s/secrets/org3.dummy.com
kubectl apply -f k8s/storage/pv
kubectl apply -f k8s/storage/pvc
kubectl apply -f k8s/job
kubectl wait --timeout=120s --for=condition=complete --namespace dummy-com job.batch/setup
kubectl delete -f k8s/job
kubectl apply -k k8s
kubectl wait --timeout=180s --for=condition=Ready --namespace dummy-com pods --all
