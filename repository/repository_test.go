package repository

import (
	"context"
	"github.com/stretchr/testify/require"

	"log"
	"testing"
)

func TestRepository_FindMaxId(t *testing.T) {
	pool, err := InitDb("localhost", "54320")
	if err != nil {
		log.Fatal(err)
	}
	repos := NewRepository(pool)

	testInsert := `
	insert into urls (FullUrl,ShortUrl)
	values ('fulltest','shorttest');`
	_, err = repos.Db.Exec(context.Background(), testInsert)
	if err != nil {
	}

	require.GreaterOrEqual(t, repos.FindMaxId(context.Background()), 1)
}

func TestRepository_InsertNewUrl(t *testing.T) {

}

func TestRepository_SearchFullUrlInDb(t *testing.T) {

}

func TestRepository_SearchShortUrlInDb(t *testing.T) {

}
