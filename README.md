# monorepo

This project makes use of [Tilt](https://tilt.dev/) and the lightweight [k3s](https://github.com/rancher/k3s) wrapper [k3d](https://k3d.io/).


## Dev Setup

This guide assumes that the machine is already setup with golang, node and yarn

- Install [docker](https://docs.docker.com/get-docker) on your machine if you don't already have it
- Install [k3d](https://k3d.io/#installation) on your machine depending on your operating system
- Create the k3d cluser with registry and also mount the project in the cluster. The command below assumes we have the project in `workspace` directory.  
  You may need to switch k8s context after this
  ```
  k3d cluster create dev-cluster --volume $HOME/workspace:$HOME/workspace --registry-create
  ```
- Install [tilt](https://docs.tilt.dev/install.html)
- Install [bazel](https://docs.bazel.build/versions/main/install.html)
- Install [protoc-gen-go](https://grpc.io/docs/languages/go/quickstart)


## Start Project

- cd into the project directory and run `tilt up`
- navigate to `localhost:10350` on your browser to view the tilt dashboard
- you can reach the app at `localhost:3000`


