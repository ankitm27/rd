package main

import (
	"log"

	rd "github.com/kkdai/rd"
)

func main() {
	log.Println("Server Starting...")
	rd.NewServer()
}
