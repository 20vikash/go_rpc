package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:4000") // Connect to server at localhost:8080
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer client.Close()

	args := Args{A: 10}
	var result int

	err = client.Call("NextNumber.Next", args, &result)
	if err != nil {
		fmt.Println("Error calling NextNumber.Next:", err)
		return
	}

	fmt.Println("Result of nextNumber:", result)
}
