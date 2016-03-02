package main

import (
	"fmt"
	"log"
	"net/rpc"

	rd "github.com/kkdai/rd"
)

func main() {

	client, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args1 := &rd.QArgs{QueueName: "Work1"}
	var reply int
	err = client.Call("WorkQueue.QueueDeclare", args1, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println("WorkQueue: add1 ")

	args2 := &rd.QArgs{QueueName: "Work2"}
	err = client.Call("WorkQueue.QueueDeclare", args2, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println("WorkQueue: add2 ")

}
