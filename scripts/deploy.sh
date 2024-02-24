#!/bin/bash

# Build the Docker image
docker build -t tzdanows/df:latest -f ./data-fetcher/Dockerfile ./data-fetcher/

# Apply the Kubernetes deployment configuration
kubectl apply -f data-fetcher/deployment.yaml

# Push the image to the Docker registry
docker push tzdanows/df:latest

# Pull the image from the Docker registry (to ensure it's available or updated)
docker pull tzdanows/df:latest

# Restart the Kubernetes deployment to pick up the new image
kubectl rollout restart deployment/df

# Monitor the status of the rollout to ensure it completes successfully
kubectl rollout status deployment/df
