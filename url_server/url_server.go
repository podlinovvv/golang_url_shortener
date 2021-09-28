package url_server

import (
	"context"
	"fmt"
	bc "github.com/chtison/baseconverter"
	"github.com/jackc/pgx/v4/pgxpool"
	pb "golang_url_shortener/proto"
	_ "google.golang.org/grpc"
	"strconv"
)

type LinkFromDb struct {
	id    int
	full  string
	short string
}

type ShortenerServer struct {
	Db *pgxpool.Pool
	pb.UnimplementedShortenerServiceServer
}

func NewShortenerServer() *ShortenerServer {
	return &ShortenerServer{}
}

func GenerateShortUrl(id int) string {
	number := strconv.Itoa(id)
	var inBase string = "0123456789"
	var toBase string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	converted, _, _ := bc.BaseToBase(number, inBase, toBase)
	var nulled string = ""
	for i := 0; i < (10 - len(converted)); i++ {
		nulled = nulled + "0"
	}
	nulled = nulled + converted
	fmt.Println(nulled)
	return nulled
}
func findMaxId(ctx context.Context, s *ShortenerServer) int {
	var id int
	err := s.Db.QueryRow(ctx, "SELECT MAX(Id) FROM urls").Scan(&id)
	//fmt.Println(id, "id из базы")
	if err != nil {
		fmt.Println(err)
	}
	return id
}

func SearchFullUrlInDb(ctx context.Context, s *ShortenerServer, in *pb.FullUrl) *LinkFromDb {
	linkFromDb := &LinkFromDb{}
	err := s.Db.QueryRow(ctx, "SELECT Id, FullUrl, ShortUrl FROM urls WHERE FullUrl=$1 LIMIT 1;", in.Url).Scan(&linkFromDb.id, &linkFromDb.full, &linkFromDb.short)
	if err != nil {
		fmt.Println(err)
	}
	return linkFromDb
}

func insertNewUrl(ctx context.Context, s *ShortenerServer, s1 string, s2 string) {
	insertSql := `
	insert into urls (FullUrl,ShortUrl)
	values ($1,$2);`
	_, err := s.Db.Exec(ctx, insertSql, s1, s2)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *ShortenerServer) Create(ctx context.Context, in *pb.FullUrl) (*pb.ShortUrl, error) {
	//поиск полного url в базе
	rowFromDb := SearchFullUrlInDb(ctx, s, in)

	var result string
	if rowFromDb.id != 0 {
		result = rowFromDb.short
	} else {
		//поиск максимального значения ID в базе
		var id int
		id = findMaxId(ctx, s)
		//генерация нового short url
		result = GenerateShortUrl(id)
	}
	//добавление новых значений fullurl и shorturl в базу
	insertNewUrl(context.Background(), s, in.Url, result)

	return &pb.ShortUrl{Link: result}, nil
}

func (s *ShortenerServer) Get(ctx context.Context, in *pb.ShortUrl) (*pb.FullUrl, error) {
	linkFromDb := &LinkFromDb{}
	err := s.Db.QueryRow(ctx, "SELECT Id, FullUrl, ShortUrl FROM urls WHERE ShortUrl=$1 LIMIT 1;", in.Link).Scan(&linkFromDb.id, &linkFromDb.full, &linkFromDb.short)
	if err != nil {
		fmt.Println(err)
	}
	return &pb.FullUrl{Url: linkFromDb.full}, nil
}
