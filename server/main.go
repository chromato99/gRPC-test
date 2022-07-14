package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"time"

	pb "github.com/chromato99/gRPC-test/api"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement api.UnimplementedRandomNumberServer
type server struct {
	pb.UnimplementedRandomNumberServer
}

// GetRandomNumber implements api.RandomNumberServer
func (s *server) GetRandomNumber(ctx context.Context, in *pb.RandomNumberRequest) (*pb.RandomNumberReply, error) {
	fmt.Printf("Received: %v\n", in.GetMaxNumber())
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return &pb.RandomNumberReply{Number: r1.Int63n(in.GetMaxNumber())}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterRandomNumberServer(s, &server{})
	fmt.Printf("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
