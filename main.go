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
	pool, err := repository.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewRepository(pool)
	server := url_server.NewShortenerServer(repository)

	lis, err := net.Listen("tcp", port)
	s := grpc.NewServer()
	pb.RegisterShortenerServiceServer(s, server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)

	}
	//testInsert := `
	//insert into urls (FullUrl,ShortUrl)
	//values ('fulltest','shorttest');`
	//_, err = server.Db.Exec(context.Background(), testInsert)
	//if err != nil {
	//}
	//
	//testInsert = `
	//insert into urls (FullUrl,ShortUrl)
	//values ('fulltest2','shorttest2');`
	//_, err = server.Db.Exec(context.Background(), testInsert)
	//if err != nil {
	//}
	//
	//testInsert = `
	//insert into urls (FullUrl,ShortUrl)
	//values ('fulltest3','shorttest3');`
	//_, err = server.Db.Exec(context.Background(), testInsert)
	//if err != nil {
	//}
}
