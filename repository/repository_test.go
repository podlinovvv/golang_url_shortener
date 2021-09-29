package repository

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	pb "golang_url_shortener/proto"
	"log"
	"testing"
)

var repos *Repository

func TestRepository_FindMaxId(t *testing.T) {
	pool, err := InitDb("localhost", "54320")
	if err != nil {
		log.Fatal(err)
	}
	repos = NewRepository(pool)

	testInsert := `
	insert into urls (FullUrl,ShortUrl)
	values ('fulltest','shorttest1');`
	_, err = repos.Db.Exec(context.Background(), testInsert)
	if err != nil {
	}

	require.GreaterOrEqual(t, repos.FindMaxId(context.Background()), 1)
}

func TestRepository_SearchFullUrlInDb(t *testing.T) {

	l := repos.SearchFullUrlInDb(context.Background(), &pb.FullUrl{Url: "fulltest"})
	require.GreaterOrEqual(t, l.Id, 0)
	fmt.Println(l.Short)
	require.Equal(t, 10, len(l.Short))
}

func TestRepository_SearchShortUrlInDb(t *testing.T) {
	_, err := repos.SearchShortUrlInDb(context.Background(), "shorttest1")
	require.Nil(t, err)
}
