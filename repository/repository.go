package repository

import (
	"context"
	"fmt"
	pb "golang_url_shortener/proto"

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
	InsertNewUrl(ctx context.Context, s1 string, s2 string) error
	SearchShortUrlInDb(ctx context.Context, link string) (string, error)
}

type Repository struct {
	Db *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{Db: pool}
}

func (r *Repository) FindMaxId(ctx context.Context) int {
	var id int
	err := r.Db.QueryRow(ctx, "SELECT MAX(Id) FROM urls").Scan(&id)
	//fmt.Println(id, "id из базы")
	if err != nil {
		fmt.Println(err)
	}
	return id
}

func (r *Repository) SearchFullUrlInDb(ctx context.Context, in *pb.FullUrl) *LinkFromDb {
	linkFromDb := &LinkFromDb{}
	err := r.Db.QueryRow(ctx, "SELECT Id, FullUrl, ShortUrl FROM urls WHERE FullUrl=$1 LIMIT 1;", in.Url).Scan(&linkFromDb.Id, &linkFromDb.Full, &linkFromDb.Short)
	if err != nil {
		fmt.Println(err)
	}
	return linkFromDb
}

func (r *Repository) InsertNewUrl(ctx context.Context, s1 string, s2 string) error {
	insertSql := `
	insert into urls (FullUrl,ShortUrl)
	values ($1,$2);`
	_, err := r.Db.Exec(ctx, insertSql, s1, s2)
	return err

}
func (r *Repository) SearchShortUrlInDb(ctx context.Context, link string) (string, error) {
	linkFromDb := &LinkFromDb{}
	err := r.Db.QueryRow(ctx, "SELECT Id, FullUrl, ShortUrl FROM urls WHERE ShortUrl=$1 LIMIT 1;", link).Scan(&linkFromDb.Id, &linkFromDb.Full, &linkFromDb.Short)
	if err != nil {
		fmt.Println(err)
	}
	return linkFromDb.Full, nil

}
