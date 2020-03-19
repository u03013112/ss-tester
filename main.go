package main

import (
	"log"
	"net"

	pb "github.com/u03013112/ss-pb/tester"
	"google.golang.org/grpc"

	"github.com/u03013112/ss-tester/spider"
	"github.com/u03013112/ss-tester/tester"
)

const (
	port = ":50004"
)

// for ci
func main() {
	tester.InitDB()
	tester.ScheduleInit()
	spider.ScheduleInit()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen %s", port)
	s := grpc.NewServer()
	pb.RegisterSSTesterServer(s, &tester.Srv{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
