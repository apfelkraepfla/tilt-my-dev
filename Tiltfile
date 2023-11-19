# -*- mode: Python -*-

docker_build('tilt-my-dev', '.', dockerfile='deployments/Dockerfile')
k8s_yaml('deployments/kubernetes.yaml')
k8s_resource('example-go', port_forwards=3000)