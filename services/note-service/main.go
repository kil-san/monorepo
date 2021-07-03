package main

import (
	"fmt"
	"net"

	"github.com/kil-san/micro-serv/note-service/pb"
	"github.com/kil-san/micro-serv/note-service/rpc"
	"google.golang.org/grpc"

	log "unknwon.dev/clog/v2"
)

func main() {
	err := log.NewConsole()
	if err != nil {
		panic("unable to create new logger: " + err.Error())
	}

	addr := fmt.Sprintf("0.0.0.0:%d", 8008)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Failed to listen: %+v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNoteRPCServer(s, rpc.New())

	log.Info("Serving gRPC on https://%s", addr)
	log.Fatal("Failed to serve: %+v", s.Serve(lis))
}
