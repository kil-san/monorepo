load('ext://restart_process', 'docker_build_with_restart')
load('ext://helm_remote', 'helm_remote')

default_registry('localhost:56406')

k8s_yaml('k8s/mysql-deployment.yaml')
helm_remote('phpmyadmin',
  repo_name="bitnami",
  repo_url="https://charts.bitnami.com/bitnami",
  release_name="dev-release",
  namespace="kube-system",
  set=["db.host=db", "db.port=3306"],
)

k8s_resource('db', new_name='run svc: db', port_forwards='30000:3306')
k8s_resource('dev-release-phpmyadmin', new_name='run svc: phpMyAdmin', port_forwards='3308:8080')

for service in [
  'gateway-service',
  'note-service'
]:
  k8s_yaml('k8s/service/' + service + '-deployment.yaml')
  docker_build_with_restart(
    ref=service + ':latest',
    entrypoint='/go/bin/' + service,
    build_args={'ServiceRelativePath': './services'},
    context='.',
    dockerfile='services/' + service + '/Dockerfile',
    live_update=
      [
        sync('services/' + service , '/dev/null')
      ]
  )

for service, port in [
  ['gateway-service', '50001:8000'],
  ['note-service', '50002:8008']
]:
  k8s_resource(
    service,
    new_name=service,
    port_forwards=port,
  )
