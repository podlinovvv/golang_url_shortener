package main

import (
	"context"
	"fmt"
	pb "golang_url_shortener/proto"
	repository "golang_url_shortener/repository"
	"golang_url_shortener/url_server"
	"log"
	"net"
	"net/url"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func InitDb() (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/?sslmode=disable&connect_timeout=%d",
		"postgresql",
		url.QueryEscape("db_user"),
		url.QueryEscape("pwd123"),
		"pgdb",
		"5432",
		15)
	ctx := context.Background()

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

	createSql := `
	create table if not exists urls (
		Id SERIAL PRIMARY KEY,
		FullUrl VARCHAR(2048),
		ShortUrl VARCHAR(2048)
	);`
	_, err = pool.Exec(context.Background(), createSql)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func main() {
	pool, err := InitDb()
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
	//Создаём grpc сервер и регистрируем его как сервер для ссервиса укорачивания
}
