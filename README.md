## Requirements

1. [minikube](https://minikube.sigs.k8s.io/docs/start/)
2. [docker & dockerhub](https://www.docker.com/get-started/)
3. [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)
4. [skaffold](https://skaffold.dev/docs/install/)

You need a minikube instance running, to start one run:

```bash
# starting minikube
minikube start

# stopping minikube
minikube stop
```

## Running tft-df via skaffold (ideally do this)

```bash
# skaffold dev or skaffold run
skaffold dev

# delete/wipe skaffold
skaffold delete
```


## Running tft-df in docker & kubernetes (manually / alternative methods)

You can run most of the listed actions below with this script from the root directory:

```bash
# linux/mac
./scripts/deploy.sh

# windows
./scripts/windows/deploy.ps1
```

Build the docker image in the root directory with the following command:

```bash
docker build -t tzdanows/df:latest -f ./data-fetcher/Dockerfile ./data-fetcher/
```

Run the docker container:

```bash
docker run --rm tzdanows/df
```

Deploying the container to k8s

```bash
# push it to docker hub (assuming you just rebuilt the image)
docker push tzdanows/df:latest

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

Note: You need to refresh the riot API key once every 24h unless you apply to be an official application endorsed by riot

### Resources

* [Skaffold Docs](https://skaffold.dev/docs/)
* [Docker/DockerHub Docs](https://docs.docker.com/docker-hub/)
