package main

import (
	"context"
	"fmt"
	bc "github.com/chtison/baseconverter"
	"github.com/jackc/pgx/v4/pgxpool"
	pb "golang_url_shortener/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/url"
	"os"
	"strconv"
)

const (
	port = ":50051"
)

type LinkFromDb struct {
	id    int
	full  string
	short string
}

type ShortenerServer struct {
	db *pgxpool.Pool
	pb.UnimplementedShortenerServiceServer
}

func NewShortenerServer() *ShortenerServer {
	return &ShortenerServer{}
}

func GenerateShortUrl(id int) (surl string) {
	number := strconv.Itoa(id)
	var inBase string = "0123456789"
	var toBase string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	converted, _, _ := bc.BaseToBase(number, inBase, toBase)

	var nulled string = ""
	for i := 0; i < (10 - len(converted)); i++ {
		nulled = nulled + "0"
	}
	nulled = nulled + converted
	return nulled
}

func (s *ShortenerServer) Create(ctx context.Context, in *pb.FullUrl) (*pb.ShortUrl, error) {

	linkFromDb := &LinkFromDb{}
	err := s.db.QueryRow(ctx, "SELECT Id, FullUrl, ShortUrl FROM urls WHERE FullUrl=$1 LIMIT 1;", in.Url).Scan(&linkFromDb.id, &linkFromDb.full, &linkFromDb.short)
	if err != nil {
		fmt.Println(err)
	}
	var result string
	if linkFromDb.id != 0 {
		fmt.Println("найден в базе")
		result = linkFromDb.short
	} else {
		fmt.Println("не найден в базе")
		var id int
		err := s.db.QueryRow(ctx, "SELECT MAX(Id) FROM urls").Scan(&id)

		fmt.Println(id,"id из базы")
		if err != nil {
			fmt.Println(err)
		}
		result = GenerateShortUrl(id)
	}

	insertNewUrl := `
	insert into urls (FullUrl,ShortUrl)
	values ($1,$2);`
	_, err = s.db.Exec(context.Background(), insertNewUrl,in.Url,result)
	if err != nil {
	}
	return &pb.ShortUrl{Link: result}, nil
}

func (s *ShortenerServer) Get(ctx context.Context, in *pb.ShortUrl) (*pb.FullUrl, error) {
	//someshortlink := "abracadabra"
	//fmt.Println(reflect.TypeOf(in))
	linkFromDb := &LinkFromDb{}
	err := s.db.QueryRow(ctx, "SELECT Id, FullUrl, ShortUrl FROM urls WHERE ShortUrl=$1 LIMIT 1;", in.Link).Scan(&linkFromDb.id, &linkFromDb.full, &linkFromDb.short)
	if err != nil {
		fmt.Println(err)
	}

	return &pb.FullUrl{Url: linkFromDb.full}, nil
}

func main() {
	//database_url := "postgres://postgres:mysecretpassword@localhost:5432/postgres"
	//conn, err := pgx.Connect(context.Background(), database_url)
	//connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%d",
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
	poolConfig.MaxConns = 5

	//Получаем пул соединений, используя контекст и конфиг
	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connect to database failed: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Fprintf(os.Stderr, "good")
	}
	fmt.Println("Connection OK!")

	var server = NewShortenerServer()
	server.db = pool

	createSql := `
	create table if not exists urls (
		Id SERIAL PRIMARY KEY,
		FullUrl VARCHAR(2048),
		ShortUrl VARCHAR(2048)
	);`
	_, err = server.db.Exec(context.Background(), createSql)
	if err != nil {
	}

	testInsert := `
	insert into urls (FullUrl,ShortUrl)
	values ('fulltest','shorttest');`
	_, err = server.db.Exec(context.Background(), testInsert)
	if err != nil {
	}

	testInsert = `
	insert into urls (FullUrl,ShortUrl)
	values ('fulltest2','shorttest2');`
	_, err = server.db.Exec(context.Background(), testInsert)
	if err != nil {
	}

	testInsert = `
	insert into urls (FullUrl,ShortUrl)
	values ('fulltest3','shorttest3');`
	_, err = server.db.Exec(context.Background(), testInsert)
	if err != nil {
	}
	//Создаём grpc сервер и регистрируем его как сервер для ссервиса укорачивания
	lis, err := net.Listen("tcp", port)
	s := grpc.NewServer()
	pb.RegisterShortenerServiceServer(s, server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)

	}
	fmt.Println("Connection OK2!")
	select {}
}
