package main

import (
	"fmt"
	"net"
	"os"

	"github.com/kil-san/micro-serv/note-service/connection"
	"github.com/kil-san/micro-serv/note-service/pb"
	"github.com/kil-san/micro-serv/note-service/rpc"
	"github.com/kil-san/micro-serv/pkg/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

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

	db, err := connection.NewSqlDbConnection(model.DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
	})
	if err != nil {
		log.Fatal("could not open connection to db: %+v", err)
	}
	defer db.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterNoteRPCServer(s, rpc.New(db))

	log.Info("Serving gRPC on https://%s", addr)
	log.Fatal("Failed to serve: %+v", s.Serve(lis))
}
