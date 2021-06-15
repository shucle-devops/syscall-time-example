package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes"
	_ "github.com/lib/pq"
	"github.com/yeongjukang/syscall-time-example/proto"
	"github.com/yeongjukang/syscall-time-example/syscalltimeexample"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	proto.ExampleServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) PostgresExample(ctx context.Context, in *proto.Request) (*proto.Reply, error) {
	var pgNow, formattedPtypesTimestamp string
	for i := 0; i < int(in.GetCount()); i++ {
		syscalltimeexample.SetTimeWithSyscall(time.Now().Add(time.Minute * 1))
		formattedPtypesTimestamp = ptypes.TimestampNow().AsTime().Format("2006-01-02 15:04:05")
		pgNow = getPgNow()
		fmt.Printf("[SERVER] CURRENT PTYPES TIMESTAMP : %s\n ", formattedPtypesTimestamp)
		fmt.Printf("[SERVER] CURRENT POSTGRES NOW() : %s\n", getPgNow())
	}
	return &proto.Reply{
		LastPtypesTime: formattedPtypesTimestamp,
		LastPgTime:     pgNow,
	}, nil
}

var pg *sql.DB

func connectPg() {
	var connStr string
	var e error
	if os.Getenv("RUN_ENV") == "CONTAINER" {
		connStr = "host=postgres user=postgres password=postgres dbname=postgres sslmode=disable"
	} else {
		connStr = "host=localhost user=postgres password=postgres dbname=postgres sslmode=disable"
	}
	pg, e = sql.Open("postgres", connStr)
	if e != nil {
		panic(e)
	}
}

func getPgNow() string {
	var now string
	pg.QueryRow("SELECT NOW()").Scan(&now)
	return now
}

func main() {
	connectPg()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterExampleServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
