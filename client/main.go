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
		log.Fatal("rpc error1:", err)
	}
	fmt.Println("WorkQueue: add1 ")

	args2 := &rd.QArgs{QueueName: "Work2"}
	err = client.Call("WorkQueue.QueueDeclare", args2, &reply)
	if err != nil {
		log.Fatal("rpc error2:", err)
	}
	fmt.Println("WorkQueue: add2 ")

	//Try to publish data
	pReply := &rd.PubRet{}
	argsP := &rd.PArgs{QName: "Work1", QValue: []byte("test1")}
	err = client.Call("WorkQueue.Publish", argsP, &pReply)
	if err != nil {
		log.Fatal("rpc error3:", err)
	}
	fmt.Println("Publish to Work1")

	//Try to push again
	pReply = &rd.PubRet{}
	argsP = &rd.PArgs{QName: "Work1", QValue: []byte("test2")}
	err = client.Call("WorkQueue.Publish", argsP, &pReply)
	if err != nil {
		log.Fatal("rpc error3:", err)
	}
	fmt.Println("Publish to Work1")

	argsP2 := &rd.PArgs{QName: "Work2", QValue: []byte("test2")}
	err = client.Call("WorkQueue.Publish", argsP2, &pReply)
	if err != nil {
		log.Fatal("rpc error4:", err)
	}
	fmt.Println("Publish to Work2")

	//Get some consume
	argsC := &rd.ConsumeArgs{QueueName: "Work1"}
	retC := &rd.ConsumeRet{}

	err = client.Call("WorkQueue.Consume", argsC, &retC)
	if err != nil {
		log.Fatal("rpc error5:", err)
	}
	fmt.Println("Consume to Work1:", retC)

	for _, v := range retC.ReturnValue {
		fmt.Println(" value:", string(v))
	}
}
