# Basic Tilt example on Windows

## Requirements
* install Docker Desktop (make sure that WSL2 feature installation is checked in the installer) (s. [here](https://docs.docker.com/desktop/install/windows-install/))
* enable Kubernetes in Docker Desktop GUI (`Settings`>`Kubernetes`>:white_check_mark: Enable Kubernetes)
* install `kubectl` s. [here](https://kubernetes.io/docs/tasks/tools/install-kubectl-windows/#install-kubectl-binary-with-curl-on-windows)
* install `Minikube` s. [here](https://minikube.sigs.k8s.io/docs/start/)
* install `Tilt`, and create rudimentary project, s. https://github.com/tilt-dev/tilt-example-go/blob/master/0-base/Tiltfile

## Start working
* start `Minikube`: `$ minikube start`
* start Tilt from the root of the project with the Tilt file: `$ tilt up`; as the port-forward is already specified in the [Tiltfile](.Tiltfile), it does not have to be done manually via `kubectl`


