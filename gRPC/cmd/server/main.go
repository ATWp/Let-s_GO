package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

/*go buil  -v ./cmd/server
./server

evans api/proto/adder.proto -p 8080
call Add
3
4
*/
func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Server(l); err != nil {
		log.Fatal(err)
	}

}
