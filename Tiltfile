allow_k8s_contexts('minikube')

#docker_build('roukou/tilt-demo', '.')
k8s_yaml('k8s-app.yml')
k8s_resource('tilt-demo', port_forwards=42050)