package url_server

import (
	"context"
	pb "golang_url_shortener/proto"
	"golang_url_shortener/repository"
	"golang_url_shortener/repository/mocks"
	"reflect"
	"testing"
)

func TestGenerateShortUrl(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "GenerateShortUrl",
			args: args{1},
			want: "0000000001",
		},

		{
			name: "GenerateShortUrl",
			args: args{63},
			want: "0000000010",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateShortUrl(tt.args.id); got != tt.want {
				t.Errorf("GenerateShortUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestShortenerServer_Create(t *testing.T) {
	type fields struct {
		r                                   repository.IRepository
		UnimplementedShortenerServiceServer pb.UnimplementedShortenerServiceServer
	}
	type args struct {
		ctx context.Context
		in  *pb.FullUrl
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ShortUrl
		wantErr bool
	}{
		{
			name:   "test1",
			fields: fields{r: new(mocks.IRepository), UnimplementedShortenerServiceServer: pb.UnimplementedShortenerServiceServer{}},
			args: args{ctx:context.Background(), in:&pb.FullUrl{Url:""}},
			want: 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ShortenerServer{
				r:                                   tt.fields.r,
				UnimplementedShortenerServiceServer: tt.fields.UnimplementedShortenerServiceServer,
			}
			got, err := s.Create(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortenerServer_Get(t *testing.T) {
	type fields struct {
		r                                   repository.IRepository
		UnimplementedShortenerServiceServer pb.UnimplementedShortenerServiceServer
	}
	type args struct {
		ctx context.Context
		in  *pb.ShortUrl
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.FullUrl
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ShortenerServer{
				r:                                   tt.fields.r,
				UnimplementedShortenerServiceServer: tt.fields.UnimplementedShortenerServiceServer,
			}
			got, err := s.Get(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}