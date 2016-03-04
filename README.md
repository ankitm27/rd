RD: A simple RPC Message Queue Server/Client with DiskQueue
==================

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/rd/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/rd?status.svg)](https://godoc.org/github.com/kkdai/rd)  [![Build Status](https://travis-ci.org/kkdai/rd.svg?branch=master)](https://travis-ci.org/kkdai/rd)
    

**RD** RD is a RPC client/server to implement Message Queue which support pubsub. Because of RPC communication protocol, we use `Consume` to get all message until now.

Install
---------------
Install package `go get github.com/kkdai/rd`

Install example server: `go get github.com/kkdai/rd/server`

Install example client: `go get github.com/kkdai/rd/client`


Usage
---------------

#### Server side example

Init a local UDP server on port "10001".

```go
package main

import (
	"log"

	rd "github.com/kkdai/rd"
)

func main() {
	log.Println("Server Starting...")

	//Init server in port 1234
	rd.NewServer(":1234")
}
```

#### Client side example

Send UDP socket to server "127.0.0.1:1234"

```go
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

```

###Use the binary directly

```
//Query a Queue Name "w1"
Evan:client Q w1
>> Q w1
>> 2016/03/04 09:04:37 WorkQueue:  w1

//Consume current data from w1
Evan:client c w1
>> c w1
>> 2016/03/04 09:04:41 Consume from  w1 : &{[]}

//Publish data "1234" to w1
Evan:client p w1 1234
>> p w1 1234
>> 2016/03/04 09:04:49 Publish  1234  to  w1  done.

//Publish data "23" to w1
Evan:client p w1 23
>> p w1 23
>> 2016/03/04 09:04:54 Publish  23  to  w1  done.

//Publish data "45" to w1
Evan:client p w1 45
>> p w1 45
>> 2016/03/04 09:04:58 Publish  45  to  w1  done.


//Consume data from w1 until now
Evan:client c w1
>> c w1
>> 2016/03/04 09:05:02 Consume from  w1 : &{[[49 50 51 52] [50 51] [52 53]]}
>> 2016/03/04 09:05:02  value: 1234
>> 2016/03/04 09:05:02  value: 23
>> 2016/03/04 09:05:02  value: 45
```

Inspired
---------------

- [RabbitMQ](https://www.rabbitmq.com/)


Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

This package is licensed under MIT license. See LICENSE for details.

