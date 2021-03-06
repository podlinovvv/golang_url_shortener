package repository

import (
	"context"
	"fmt"
	pb "golang_url_shortener/proto"
	"net/url"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type LinkFromDb struct {
	Id    int
	Full  string
	Short string
}

type IRepository interface {
	FindMaxId(ctx context.Context) int
	SearchFullUrlInDb(ctx context.Context, in *pb.FullUrl) *LinkFromDb
	InsertNewUrl(ctx context.Context, s1 string, s2 string)
	SearchShortUrlInDb(ctx context.Context, link string) (string, error)
}

type Repository struct {
	Db *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{Db: pool}
}

func InitDb(address string, port string) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/?sslmode=disable&connect_timeout=%d",
		"postgresql",
		url.QueryEscape("db_user"),
		url.QueryEscape("pwd123"),
		address,
		port,
		15)
	ctx := context.Background()

	poolConfig, _ := pgxpool.ParseConfig(connStr)
	poolConfig.MaxConns = 1

	//Получаем пул соединений, используя контекст и конфиг
	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connect to database failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connection OK!")

	//создание таблицы URLов
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

//поиск максимального ID в базе
func (r *Repository) FindMaxId(ctx context.Context) int {
	var id int
	err := r.Db.QueryRow(ctx, "SELECT MAX(Id) FROM urls").Scan(&id)
	//fmt.Println(id, "id из базы")
	if err != nil {
		fmt.Println(id, err)
	}
	return id
}

//поиск полного URL в базе
func (r *Repository) SearchFullUrlInDb(ctx context.Context, in *pb.FullUrl) *LinkFromDb {
	linkFromDb := &LinkFromDb{}
	err := r.Db.QueryRow(ctx, "SELECT Id, FullUrl, ShortUrl FROM urls WHERE FullUrl=$1 LIMIT 1;", in.Url).Scan(&linkFromDb.Id, &linkFromDb.Full, &linkFromDb.Short)
	if err != nil {
		fmt.Println(err)
	}
	return linkFromDb
}

//добавление новой пары из полного и сокращённого URL
func (r *Repository) InsertNewUrl(ctx context.Context, s1 string, s2 string) {
	insertSql := `
	insert into urls (FullUrl,ShortUrl)
	values ($1,$2);`
	_, _ = r.Db.Exec(ctx, insertSql, s1, s2)
	return
}

//поиск сокращённого URL в базе
func (r *Repository) SearchShortUrlInDb(ctx context.Context, link string) (string, error) {
	linkFromDb := &LinkFromDb{}
	err := r.Db.QueryRow(ctx, "SELECT Id, FullUrl, ShortUrl FROM urls WHERE ShortUrl=$1 LIMIT 1;", link).Scan(&linkFromDb.Id, &linkFromDb.Full, &linkFromDb.Short)
	if err != nil {
		return "", err
	}
	return linkFromDb.Full, nil

}
