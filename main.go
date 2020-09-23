package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/kaidev1024/protobuf/protos/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	protos "github.com/kaidev1024/protobuf/protos/currency"
)

func main() {
	logger := log.New(os.Stdout, "proto-currency:", http.StatusBadRequest)

	gs := grpc.NewServer()
	cs := server.NewCurrency(logger)

	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		logger.Println("Unable to listen: ", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
