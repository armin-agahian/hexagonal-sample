package main

import (
	"context"
	"log"
	"net"

	"github.com/armin-agahian/hexagonal-sample/internal/adapters/infrastructure/redis_adapter"
	pb "github.com/armin-agahian/hexagonal-sample/internal/adapters/interface/grpc"
	"github.com/armin-agahian/hexagonal-sample/internal/core/application/ports"
	"github.com/armin-agahian/hexagonal-sample/internal/core/application/services/articlesrv"
	"github.com/armin-agahian/hexagonal-sample/internal/core/domain/entities"
	"google.golang.org/grpc"
)

const port = ":50051"

type server struct {
	pb.UnimplementedArticleServiceServer
	service ports.ArticleService
}

func NewServer(srv ports.ArticleService) pb.ArticleServiceServer {
	return &server{service: srv}
}

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	article, err := s.service.Create(ctx, in.Article.GetTitle(), in.Article.GetBody())
	if err != nil {
		return nil, err
	} else {
		return &pb.CreateResponse{Id: article.Id}, err
	}
}

func (s *server) Read(ctx context.Context, in *pb.ReadRequest) (*pb.ReadResponse, error) {
	article, err := s.service.Get(ctx, in.GetId())
	if err != nil {
		article = entities.Article{}
	}
	pb_article := &pb.Article{Id: article.Id, Title: article.Title, Body: article.Body}
	return &pb.ReadResponse{Article: pb_article}, err
}


func main() {
	repository := redis_adapter.NewArticleRepository()
	service := articlesrv.New(repository)
	grpc_server(service)
}

func grpc_server(service ports.ArticleService) {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
	} else {
		log.Printf("listening on port%v \n", port)
	}
	s := grpc.NewServer()
	server := NewServer(service)
	pb.RegisterArticleServiceServer(s, server)
	if err := s.Serve(listen); err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
