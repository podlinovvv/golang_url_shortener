package main

import (
	"context"
	pb "golang_url_shortener/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {

	//дрес, незащищённое соединение, блок означает что функция не вернёт значение пока соединение активно
	conn, err := grpc.Dial(address,grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	c := pb.NewShortenerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_urls  = [3]string{"32243", "4343", "435435"}
	for _, val := range new_urls {
		response, err := c.Get(ctx, &pb.ShortUrl{Link: val})
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println(response)
		}
	}


}