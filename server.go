package rd

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"sync"

	pd "github.com/kkdai/pd"
)

var rlock sync.RWMutex
var pubsubObj *pd.PD
var queueMapChan map[string]chan []byte
var queueMapData map[string][]Data

func NewServer(port string) {
	forever := make(chan int)

	//Init data
	pubsubObj = pd.NewPubsub()
	queueMapChan = make(map[string]chan []byte)
	queueMapData = make(map[string][]Data)

	queue := new(WorkQueue)
	rpc.Register(queue)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", port)
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

		for k, v := range queueMapChan {
			select {
			case newV := <-v:
				log.Println("[rd][Loop] Get topic:", k, " data=", string(newV))
				if oriSlice, exist := queueMapData[k]; exist {
					oriSlice = append(oriSlice, newV)
					queueMapData[k] = oriSlice
					log.Println("[rd][Loop] Current ", k, " slice:", oriSlice)
				} else {
					var val []Data
					val = append(val, newV)
					queueMapData[k] = val
					log.Println("[rd][Loop] Insert new slice to map")
				}

			default:
				continue
			}
		}
	}
	log.Fatal("Loop Exit")
}
