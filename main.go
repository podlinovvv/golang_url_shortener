package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	pb "golang_url_shortener/proto"
	"golang_url_shortener/url_server"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/url"
	"os"
)

const (
	port = ":50051"
)


func main() {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/?sslmode=disable&connect_timeout=%d",
		"postgresql",
		url.QueryEscape("db_user"),
		url.QueryEscape("pwd123"),
		"pgdb",
		"5432",
		15)
	ctx, _ := context.WithCancel(context.Background())

	//Сконфигурируем пул, задав для него максимальное количество соединений
	poolConfig, _ := pgxpool.ParseConfig(connStr)
	poolConfig.MaxConns = 1

	//Получаем пул соединений, используя контекст и конфиг
	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connect to database failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connection OK!")

	var server = url_server.NewShortenerServer()
	server.Db = pool

	createSql := `
	create table if not exists urls (
		Id SERIAL PRIMARY KEY,
		FullUrl VARCHAR(2048),
		ShortUrl VARCHAR(2048)
	);`
	_, err = server.Db.Exec(context.Background(), createSql)
	if err != nil {
	}

	testInsert := `
	insert into urls (FullUrl,ShortUrl)
	values ('fulltest','shorttest');`
	_, err = server.Db.Exec(context.Background(), testInsert)
	if err != nil {
	}

	testInsert = `
	insert into urls (FullUrl,ShortUrl)
	values ('fulltest2','shorttest2');`
	_, err = server.Db.Exec(context.Background(), testInsert)
	if err != nil {
	}

	testInsert = `
	insert into urls (FullUrl,ShortUrl)
	values ('fulltest3','shorttest3');`
	_, err = server.Db.Exec(context.Background(), testInsert)
	if err != nil {
	}
	//Создаём grpc сервер и регистрируем его как сервер для ссервиса укорачивания
	lis, err := net.Listen("tcp", port)
	s := grpc.NewServer()
	pb.RegisterShortenerServiceServer(s, server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)

	}
	select {}
}
