package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/apfelkraepfla/tilt-my-def/protos/quotespb"
	"google.golang.org/grpc"
)

type server struct {
	pb.QuoteServiceServer
}

var (
	port = flag.Int("port", 3000, "The server port")
)

func (s *server) GetQuote(ctx context.Context, req *pb.QuoteRequest) (*pb.QuoteResponse, error) {
	// Return a hardcoded philosophical quote
	return &pb.QuoteResponse{
		Quote: "The only true wisdom is in knowing you know nothing. - Socrates",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen at %v: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterQuoteServiceServer(grpcServer, &server{})

	log.Println("gRPC server is running on port 3000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
