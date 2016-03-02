package rd

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	pd "github.com/kkdai/pd"
)

var PdQueue *pd.PD
var workQ map[string]chan []byte

func NewServer() {
	queue := new(WorkQueue)
	log.Println("New work queue")
	PdQueue = pd.NewPubsub()
	log.Println("New pubsub")
	workQ = make(map[string]chan []byte)
	log.Println("init")
	rpc.Register(queue)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	log.Println("start to listen:", l)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
	for {
	}
}
