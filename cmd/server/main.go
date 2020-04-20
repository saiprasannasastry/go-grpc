package main

import (
	"fmt"
	albumpb "github.com/crud-grpc/pkg/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"github.com/crud-grpc/pkg/svc"
)

type AlbumServiceServer struct{}

func main() {

	listener, err := net.Listen("tcp", "10.196.105.125:50051")
	if err != nil {
		log.Fatalf("Unable to listen on port :50051: %v", err)
	}
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	srv := svc.NewAlbumServer()
	albumpb.RegisterAlbumServiceServer(s, srv)

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("\nStopping the server...")
	s.Stop()
	listener.Close()
}
