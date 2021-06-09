docker_build(
    'minimal-app:dev',
    '.',
    only=['cmd', 'go.sum', 'go.mod'])
k8s_yaml(helm("helm/minimal-app", values="helm/dev-values.yaml"))
k8s_resource('chart-minimal-app', port_forwards=8000)
