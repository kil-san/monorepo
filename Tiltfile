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
  pwd = os.getcwd()
  deployment = local("bazel run --define service={service} --define pwd={pwd} :%s".format(
    service=service,
    pwd=pwd,
  ) % service)

  k8s_yaml(deployment)
  target = "//services/{}:image".format(service)
  dest = target.replace('//', 'bazel/')
  base_cmd = "bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64"
  cmd = "{cmd} {target} -- --norun && docker tag {dest} $EXPECTED_REF".format(
    cmd=base_cmd,
    target=target,
    dest=dest
  )
  custom_build(
    service + '-image',
    command=cmd,
    deps=['services/' + service]
  )


for service, port, deps in [
  ['gateway-service', '50001:8000', ['build: buf', 'build: graphql']],
  ['note-service', '50002:8008', ['build: buf', 'run svc: db']]
]:
  k8s_resource(
    service,
    new_name=service,
    port_forwards=port,
    resource_deps=deps
  )

local_resource(
  'build: buf',
  cmd='./protoc_gen.sh',
  deps=['services/**/*.proto'],
)

local_resource(
  'build: graphql',
  cmd='cd services/gateway-service && go run github.com/99designs/gqlgen',
  deps=['services/**/*.graphqls'],
)
