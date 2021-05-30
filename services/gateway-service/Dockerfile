FROM golang:1.16-alpine3.12

ARG ServiceRelativePath=.
WORKDIR /go/src/github.com/kil-san/micro-serv
COPY $ServiceRelativePath/go.mod $ServiceRelativePath/go.sum ./
RUN go mod download
COPY $ServiceRelativePath/ .
RUN go install github.com/kil-san/micro-serv/gateway-service

CMD ["gateway-service"]
