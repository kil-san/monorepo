package main

import (
	"fmt"
	"net"

	"github.com/kil-san/micro-serv/note-service/connection"
	"github.com/kil-san/micro-serv/note-service/pb"
	"github.com/kil-san/micro-serv/note-service/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	log "unknwon.dev/clog/v2"
)

func main() {
	err := log.NewConsole()
	if err != nil {
		panic("unable to create new logger: " + err.Error())
	}

	addr := fmt.Sprintf("127.0.0.1:%d", 8008)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Failed to listen: %+v", err)
	}

	db, err := connection.NewSqliteConnection("sqlite.db")
	if err != nil {
		panic("could not open connection to db")
	}
	defer db.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterNoteRPCServer(s, rpc.New(db))

	log.Info("Serving gRPC on https://%s", addr)
	log.Fatal("Failed to serve: %+v", s.Serve(lis))
}
