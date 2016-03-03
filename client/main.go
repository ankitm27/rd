package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"

	rd "github.com/kkdai/rd"
)

func QueryRPC(client *rpc.Client, name string) {
	var err error
	args := &rd.QArgs{QueueName: name}
	var reply int
	err = client.Call("WorkQueue.QueueDeclare", args, &reply)
	if err != nil {
		log.Fatal("rpc error:", err)
	}
	fmt.Println("WorkQueue: ", name)
}

func PublishRPC(client *rpc.Client, name string, value []byte) {
	var err error
	pReply := &rd.PubRet{}
	argsP := &rd.PArgs{QName: name, QValue: value}
	err = client.Call("WorkQueue.Publish", argsP, &pReply)
	if err != nil {
		log.Fatal("rpc error:", err)
	}
	fmt.Println("Publish ", string(value), " to ", name, " done.")
}

func ConsumeRPC(client *rpc.Client, name string) {
	var err error
	argsC := &rd.ConsumeArgs{QueueName: name}
	retC := &rd.ConsumeRet{}

	err = client.Call("WorkQueue.Consume", argsC, &retC)
	if err != nil {
		log.Fatal("rpc error:", err)
	}
	fmt.Println("Consume from ", name, ":", retC)

	for _, v := range retC.ReturnValue {
		fmt.Println(" value:", string(v))
	}
}

func main() {
	var cmd, param1, param2 string
	flag.Parse()
	cmd = flag.Arg(0)
	param1 = flag.Arg(1)
	param2 = flag.Arg(2)

	fmt.Println(cmd, param1, param2)
	client, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	switch cmd {
	case "Q", "q":
		QueryRPC(client, param1)
	case "P", "p":
		PublishRPC(client, param1, []byte(param2))
	case "C", "c":
		ConsumeRPC(client, param1)
	default:
		fmt.Println("Please check your parameter")
	}
}
