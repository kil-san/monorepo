load('ext://restart_process', 'docker_build_with_restart')
load('ext://helm_remote', 'helm_remote')

default_registry('localhost:56406')

k8s_yaml('k8s/config-map/env.yaml')
k8s_yaml('k8s/config-map/dev.yaml')

# k8s_yaml('k8s/mysql-deployment.yaml')
# helm_remote('phpmyadmin',
#   repo_name="bitnami",
#   repo_url="https://charts.bitnami.com/bitnami",
#   release_name="dev-release",
#   namespace="kube-system",
#   set=["db.host=db", "db.port=3306"],
# )
helm_remote('mongodb',
  repo_name="mongodb",
  repo_url="https://charts.bitnami.com/bitnami",
  release_name="dev-release",
  namespace="kube-system",
  set=[
    "auth.enabled=false"
  ],
)

# k8s_resource('db', new_name='run side: db', port_forwards='30000:3306')
# k8s_resource(
#   'dev-release-phpmyadmin',
#   new_name='run side: phpMyAdmin',
#   port_forwards='3308:8080',
#   resource_deps=['run side: db']
# )
k8s_resource(
  'dev-release-mongodb',
  new_name='run side: mongoDB',
  port_forwards='27000:27017',
)

# Firebase emulator
compiledFirebaseEmulatorYaml = local('./formatDeploymentYaml.sh k8s/firebase-emulator.yaml')
k8s_yaml(compiledFirebaseEmulatorYaml)
docker_build('firebase-emulator', './firebase-emulator')
k8s_resource('firebase-emulator',
  new_name='run svc: firebase-emulator',
  port_forwards=['4000:4000', '8080:8080', '9099:9099'],
)

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
  ['note-service', '50002:443', ['build: buf', 'run side: mongoDB']]
]:
  k8s_resource(
    service,
    new_name="run svc: " + service,
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

local_resource(
  'build: export firebase',
  serve_cmd='./hack/firebase_export.sh',
  allow_parallel=True,
  deps=[
    'hack/firebase_export.sh',
  ],
  resource_deps=[
    'run svc: firebase-emulator',
  ],
)

