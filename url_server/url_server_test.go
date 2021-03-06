package url_server

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	pb "golang_url_shortener/proto"
	"golang_url_shortener/repository"
	"log"
	"testing"
)

var s *ShortenerServer

func TestGenerateShortUrl(t *testing.T) {
	require.Equal(t, len(GenerateShortUrl(10)), 10)
	require.Equal(t, GenerateShortUrl(63), "0000000011")
}
func TestShortenerServer_Create(t *testing.T) {
	pool, err := repository.InitDb("localhost", "54320")
	if err != nil {
		log.Fatal(err)
	}
	repos := repository.NewRepository(pool)
	s = NewShortenerServer(repos)

	testInsert := `
	insert into urls (FullUrl,ShortUrl)
	values ('fulltest','shorttest3');`
	_, err = repos.Db.Exec(context.Background(), testInsert)
	if err != nil {
	}

	_, err = s.Create(context.Background(), &pb.FullUrl{Url: "testurl"})
	require.Nil(t, err)
}
func TestShortenerServer_Get(t *testing.T) {
	shortUrl, err := s.Get(context.Background(), &pb.ShortUrl{Link: "shorttest3"})
	require.Nil(t, err)
	require.NotNil(t, shortUrl)

	shortUrl, err = s.Get(context.Background(), &pb.ShortUrl{Link: "empty"})
	require.Equal(t, fmt.Sprintln(err), fmt.Sprint("no rows in result set\n"))
}
