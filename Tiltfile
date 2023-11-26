# -*- mode: Python -*-

docker_build('ghcr.io/apfelkraepfla/quotes-service', 
    '../quotes-service')
k8s_yaml('../quotes-service/deployments/kubernetes.yaml')
k8s_resource('quotes-service', port_forwards=3000)