package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

// Define a struct. This struct will bind all the RPC methods
type NextNumber struct{}

// Arguments for the next number
type Args struct {
	A int
}

func (n *NextNumber) Next(args *Args, reply *int) error {
	if args == nil {
		return errors.New("tHE ARGUMENTS CANNOT BE EMPTY FOOL")
	}

	*reply += 1

	return nil
}

func main() {
	newNumber := new(NextNumber)
	rpc.Register(newNumber)

	listener, err := net.Listen("tcp", ":4000") // Listen on port 4000
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 4000...")
	rpc.Accept(listener) // Accept incoming RPC connections
}
