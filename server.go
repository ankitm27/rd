package rd

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"sync"

	pd "github.com/kkdai/pd"
)

var Lock sync.RWMutex
var PdQueue *pd.PD
var workQ map[string]chan []byte
var workSlice map[string][]Data

func NewServer() {
	forever := make(chan int)

	//Init data
	PdQueue = pd.NewPubsub()
	workQ = make(map[string]chan []byte)
	workSlice = make(map[string][]Data)

	queue := new(WorkQueue)
	rpc.Register(queue)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	log.Println("start to listen:", l)

	if e != nil {
		log.Fatal("listen error:", e)
	}

	go http.Serve(l, nil)
	go inLoop()

	<-forever
}

func inLoop() {

	//Infinite to read data
	for {
		for k, v := range workQ {
			Lock.RLock()
			defer Lock.Unlock()

			select {
			case newV := <-v:
				log.Println("[rd][Loop] Get topic:", k, " data=", string(newV))
				if oriSlice, exist := workSlice[k]; exist {
					oriSlice = append(oriSlice, newV)
					workSlice[k] = oriSlice
					log.Println("[rd][Loop] Current ", k, " slice:", oriSlice)
				} else {
					var val []Data
					val = append(val, newV)
					workSlice[k] = val
					log.Println("[rd][Loop] Insert new slice to map")
				}

			default:
				continue
			}
		}
	}
}
