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
