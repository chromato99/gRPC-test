package main

import (
	"context"
	"flag"
	"fmt"
	"strconv"
	"time"

	pb "github.com/chromato99/gRPC-test/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr   = flag.String("addr", "localhost:50051", "the address to connect to")
	maxNum = flag.String("max-num", "10", "the max range of random number")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewRandomNumberClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	num, _ := strconv.ParseInt(*maxNum, 10, 64)
	r, err := c.GetRandomNumber(ctx, &pb.RandomNumberRequest{MaxNumber: num})
	if err != nil {
		fmt.Printf("could not greet: %v\n", err)
	}
	fmt.Printf("Greeting: %d\n", r.GetNumber())
}
