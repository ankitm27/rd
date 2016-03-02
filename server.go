package rd

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func NewServer() {
	queue := new(WorkQueue)
	rpc.Register(queue)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
	for {
	}
}
