
Build the docker image (in the data-fetcher directory) with the following command:

```bash
docker build -t <dockerhub-username>/df:latest -f ./data-fetcher/Dockerfile .
```

Run the docker container:

```bash
docker run --rm <dockerhub-username>/df
```

Deploying the container to k8s

```bash
# push it to docker hub (assuming you just rebuilt the image)
docker push <dockerhub-username>/df:latest

# apply the deployment to k8s
kubectl apply -f data-fetcher/deployment.yaml

# if you are updating a former k8s deployment, use this
kubectl rollout restart deployment/df

# verify the update
kubectl rollout status deployment/df
```

Run the container in k8s:

```bash
# get pod name
kubectl get pods

# run the code
kubectl exec [POD_NAME] -- ./df
```