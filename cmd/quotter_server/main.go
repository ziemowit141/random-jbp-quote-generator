package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/ziemowit141/random-jbp-quote-generator/src/quotes"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedQuotterServer
}

func (s *server) GetQuote(ctx context.Context, in *pb.Empty) (*pb.Quote, error) {
	log.Printf("Received quote request")
	inner_wisdom := make(chan string)
	go pb.GetQuote(inner_wisdom)
	return &pb.Quote{Quote: <-inner_wisdom}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterQuotterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}