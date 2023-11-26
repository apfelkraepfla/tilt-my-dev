# Basic Tilt example on Windows

## Requirements
* install Docker Desktop (make sure that WSL2 feature installation is checked in the installer) (s. [here](https://docs.docker.com/desktop/install/windows-install/))
* enable Kubernetes in Docker Desktop GUI (`Settings`>`Kubernetes`>:white_check_mark: Enable Kubernetes)
* install `kubectl` s. [here](https://kubernetes.io/docs/tasks/tools/install-kubectl-windows/#install-kubectl-binary-with-curl-on-windows)
* install `Minikube` s. [here](https://minikube.sigs.k8s.io/docs/start/)
* install `Tilt`, and create rudimentary project, s. https://github.com/tilt-dev/tilt-example-go/blob/master/0-base/Tiltfile
* for the devcontainer, I used the [devcontainer.json](https://github.com/grpc-ecosystem/grpc-gateway/blob/main/.devcontainer/devcontainer.json) from https://github.com/grpc-ecosystem/grpc-gateway, and the referenced [Dockerfile](https://github.com/grpc-ecosystem/grpc-gateway/blob/main/.github/Dockerfile) with some small changes

## Prerequisites for working with Tilt and cluster
* start `Minikube`: `$ minikube start`
* start Tilt from the root of the project with the Tilt file: `$ tilt up`; as the port-forward is already specified in the [Tiltfile](.Tiltfile), it does not have to be done manually via `kubectl`

### TODOs
* add SSH option in Dockerfile (https://medium.com/@quentin.mcgaw/ultimate-go-dev-container-for-visual-studio-code-448f5e031911)

### Creating the pb files
* the protocol compiler plugins for Go are already installed in the `Dockerfile` used by the devcontainer
* create proto files with your new specification: `protos/<NAME>.proto`
* add compilation task in `.vscode/tasks.json`, and run vai `F1` -> `Tasks: Run Task`


