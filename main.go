package main

import (
	pb "golang_url_shortener/proto"
	repository "golang_url_shortener/repository"
	"golang_url_shortener/url_server"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	pool, err := repository.InitDb("pgdb", "5432")
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(pool)
	server := url_server.NewShortenerServer(repos)

	lis, err := net.Listen("tcp", port)
	s := grpc.NewServer()
	pb.RegisterShortenerServiceServer(s, server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
