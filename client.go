package rd

import (
	"log"
	"net/rpc"
)

//ClientRPC This struct contain RPC call data structure include connection
type ClientRPC struct {
	conn *rpc.Client
}

//NewClientRPC Init related RPC call and dial to server by address<addr>
func NewClientRPC(addr string) *ClientRPC {

	var err error

	c := new(ClientRPC)
	c.conn, err = rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return c
}

//QueryRPC  Declare Queue, will create new queue if not exist
func (c *ClientRPC) QueryRPC(name string) {
	var err error
	args := &QueryArgs{QueueName: name}
	var reply int
	err = c.conn.Call("WorkQueue.QueueDeclare", args, &reply)
	if err != nil {
		log.Fatal("rpc error:", err)
	}
	log.Println("WorkQueue: ", name)
}

//PublishRPC Publish data to the Queue
func (c *ClientRPC) PublishRPC(name string, value []byte) {
	var err error
	var pReply int
	argsP := &PublishArgs{QName: name, QValue: value}
	err = c.conn.Call("WorkQueue.Publish", argsP, &pReply)
	if err != nil {
		log.Fatal("rpc error:", err)
	}
	log.Println("Publish ", string(value), " to ", name, " done.")
}

//ConsumeRPC Consume data from Queue until now, if there still not get any info return empty result
func (c *ClientRPC) ConsumeRPC(name string) {
	var err error
	argsC := &ConsumeArgs{QueueName: name}
	retC := &ConsumeRet{}

	err = c.conn.Call("WorkQueue.Consume", argsC, &retC)
	if err != nil {
		log.Fatal("rpc error:", err)
	}
	log.Println("Consume from ", name, ":", retC)

	for _, v := range retC.ReturnValue {
		log.Println(" value:", string(v))
	}
}
