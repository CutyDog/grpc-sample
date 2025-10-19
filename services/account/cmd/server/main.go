package main

import (
	"log"
	"net"
	"os"

	accountv1 "github.com/CutyDog/grpc-sample/proto/gen/account/v1"
	"github.com/CutyDog/grpc-sample/services/account/internal/db"
	"github.com/CutyDog/grpc-sample/services/account/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db.ConnectDB()

	addr := os.Getenv("GRPC_ADDR")
	s := grpc.NewServer()
	accountv1.RegisterAccountServiceServer(s, server.NewAccountServer(db.DB))
	reflection.Register(s)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	log.Printf("account gRPC listening on %s", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
