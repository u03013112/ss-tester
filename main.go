package main

import (
	// pb "github.com/u03013112/ss-pb/tester"
	"github.com/u03013112/ss-tester/tester"
)

const (
	port = ":50004"
)

// for ci
func main() {
	tester.InitDB()
	tester.ScheduleInit()
	// lis, err := net.Listen("tcp", port)
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// log.Printf("listen %s", port)
	// s := grpc.NewServer()
	// pb.RegisterSSConfigServer(s, &tester.Srv{})
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
	tester.Test()
}
