package main

import (
	"context"
	"go-micro.dev/v5"
	"log"

	pb "github.com/wandhie08/tinyurl/proto"
)

type URLShortener struct{}

func (s *URLShortener) Create(ctx context.Context, req *pb.CreateRequest, res *pb.CreateResponse) error {
	log.Printf("Received URL: %s", req.OriginalUrl)
	res.ShortenedUrl = "http://short.url/" + req.OriginalUrl
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("shortener"),
		micro.Address(":8081"),
	)

	service.Init()

	if err := pb.RegisterURLShortenerHandler(service.Server(), new(URLShortener)); err != nil {
		log.Fatal(err)
	}

	log.Println("Service running on :8081")
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
