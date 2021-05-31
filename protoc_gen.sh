#!/usr/bin/env bash

services=(
    'note-service'
)

for service in ${services[@]}
do
  protoc -I=services/pkg/proto/ \
    --go_out=services/"$service"/pb \
    --go_opt=paths=source_relative \
    --go-grpc_out=services/"$service"/pb \
    --go-grpc_opt=paths=source_relative \
    services/pkg/proto/"${service//-/_}".proto
done
