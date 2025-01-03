package main

import (
	"context"
	"github.com/tinyurl/proto/shortener"
	"go-micro.dev/v4"
	"log"
)

type URLShortener struct{}

func (s *URLShortener) Create(ctx context.Context, req *shortener.URLRequest, res *shortener.URLResponse) error {
	log.Printf("Received URL: %s", req.LongUrl)
	res.ShortUrl = "http://short.url/" + req.LongUrl
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("shortener"),
		micro.Address(":8081"),
	)

	service.Init()

	// Register handler using the generated code
	if err := shortener.RegisterURLShortenerHandler(service.Server(), new(URLShortener)); err != nil {
		log.Fatal(err)
	}
	log.Println("Service running on :8081")
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
