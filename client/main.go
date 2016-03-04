package main

import (
	"flag"
	"fmt"

	rd "github.com/kkdai/rd"
)

func main() {
	var cmd, param1, param2 string
	flag.Parse()
	//Get command from console
	cmd = flag.Arg(0)
	//Get parameter 1
	param1 = flag.Arg(1)
	//Get parameter 2
	param2 = flag.Arg(2)

	fmt.Println(cmd, param1, param2)
	client := rd.NewClientRPC("127.0.0.1:1234")

	switch cmd {
	case "Q", "q":
		client.QueryRPC(param1)
	case "P", "p":
		client.PublishRPC(param1, []byte(param2))
	case "C", "c":
		client.ConsumeRPC(param1)
	default:
		fmt.Println("Please check your parameter")
	}
}
