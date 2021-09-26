package main

import (
	//"google.golang.org/grpc"

	"context"
	pb "golang_url_shortener/proto"
	"log"
	"net"
)


const (
	port=":50051"
)

type ShortenerServer struct {
	pb.UnimplementedShortenerServiceServer
}

func (s *ShortenerServer) Create (ctx context.Context, in *pb.FullUrl) (*pb.ShortUrl, error){
	var sr string = "33222"
	return &pb.ShortUrl{ Link : sr},nil
}


func main(){
	_,err :=net.Listen("tcp",port)
	if err != nil {
		log.Fatal(err)
	}

	//s:=grpc.NewServer(lis)
}
