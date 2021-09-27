package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	pb "golang_url_shortener/proto"
	"reflect"

	"google.golang.org/grpc"
	"log"
	"net"
	"net/url"
	"os"
)

const (
	port = ":50051"
)

type Link struct {
	id int
	full string
	short string
}

type ShortenerServer struct {
	db *pgxpool.Pool
	pb.UnimplementedShortenerServiceServer
}

func NewShortenerServer() *ShortenerServer {
	return &ShortenerServer{}
}

func GenerateShortUrl() (surl string){

	return
}

func (s *ShortenerServer) Create(ctx context.Context, in *pb.FullUrl) (*pb.ShortUrl, error) {

	var sr string = "33222"
	var sr2 string = "33222333"

	tx, err := s.db.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec(context.Background(), "insert into urls(full,short) values($1,$2)", sr, sr2)
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit(context.Background())

	return &pb.ShortUrl{Link: sr}, nil
}

func (s *ShortenerServer) Get(ctx context.Context, in *pb.ShortUrl) (*pb.FullUrl, error) {
	someshortlink := "abracadabra"
	fmt.Println(reflect.TypeOf(in))
	//link := &Link{}
	//err := s.db.QueryRow(ctx, "SELECT id, full, short FROM urls WHERE short=$1 LIMIT 1;", in ).Scan(&link.id, &link.full, &link.short)
	//if err != nil {
	//	fmt.Println(err)}

//	sql_query := fmt.Sprintf(`
//	INSERT INTO urls (full,short)
//SELECT * FROM (SELECT %s AS full, %s AS short) AS temp
//WHERE NOT EXISTS (
//    SELECT full FROM urls WHERE full = %s
//) LIMIT 1;;
//	`, someshortlink, someshortlink, someshortlink)
//
//	_, err := s.db.Exec(ctx, sql_query)

	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Ping failed: %v\n", err)
	//	os.Exit(1)
	//}
	return &pb.FullUrl{Url: someshortlink}, nil
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
	poolConfig.MaxConns = 1

	//Получаем пул соединений, используя контекст и конфиг
	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connect to database failed: %v\n", err)
		os.Exit(1)
	}else{fmt.Fprintf(os.Stderr, "good")}
	fmt.Println("Connection OK!")

	var server = NewShortenerServer()
	server.db = pool

	createSql := `
	create table if not exists urls (
		Id int,
		FullUrl VARCHAR(2048),
		ShortUrl VARCHAR(2048)
	);`
	_, err = server.db.Exec(context.Background(), createSql)
	if err != nil {}


	//Создаём grpc сервер и регистрируем его как сервер для ссервиса укорачивания
	lis, err := net.Listen("tcp", port)
	s := grpc.NewServer()
	pb.RegisterShortenerServiceServer(s, server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)

	}

	select {}
}
