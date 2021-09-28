package url_server

import (
	"context"
	"fmt"
	pb "golang_url_shortener/proto"
	"golang_url_shortener/repository"
	"strconv"

	bc "github.com/chtison/baseconverter"
	_ "google.golang.org/grpc"
)

type IShortener interface {
	Create(ctx context.Context, in *pb.FullUrl) (*pb.ShortUrl, error)
	Get(ctx context.Context, in *pb.ShortUrl) (*pb.FullUrl, error)
}

type ShortenerServer struct {
	r repository.IRepository
	pb.UnimplementedShortenerServiceServer
}

func NewShortenerServer(r *repository.Repository) *ShortenerServer {
	return &ShortenerServer{r: r}
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

func (s *ShortenerServer) Create(ctx context.Context, in *pb.FullUrl) (*pb.ShortUrl, error) {
	//поиск полного url в базе
	rowFromDb := s.r.SearchFullUrlInDb(ctx, in)

	var result string
	if rowFromDb.Id != 0 {
		result = rowFromDb.Short
	} else {
		//поиск максимального значения ID в базе
		var id int
		id = s.r.FindMaxId(ctx)
		//генерация нового short url
		result = GenerateShortUrl(id)
	}
	//добавление новых значений fullurl и shorturl в базу
	s.r.InsertNewUrl(context.Background(), in.Url, result)

	return &pb.ShortUrl{Link: result}, nil
}

func (s *ShortenerServer) Get(ctx context.Context, in *pb.ShortUrl) (*pb.FullUrl, error) {
	shorturl, err := s.r.SearchShortUrlInDb(ctx, in.Link)
	if err != nil {
		return nil, err
	}
	return &pb.FullUrl{Url: shorturl}, nil
}
