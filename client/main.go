package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/yeongjukang/syscall-time-example/proto"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Printf("[CLIENT] %s\n", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second * 3)
		fmt.Printf("[CLIENT] %s\n", time.Now().Format("2006-01-02 15:04:05"))
		wg.Done()
	}()

	c := proto.NewExampleServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.PostgresExample(ctx, &proto.Request{Count: 5})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	wg.Wait()
	fmt.Printf("[CLIENT] RESULT : last ptypes - %s, last pg - %s\n", r.LastPtypesTime, r.LastPgTime)
}
