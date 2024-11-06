package src

import (
	"net"
	"time"
)

func CreateServer() {
	_, erro1 := net.Dial("tcp", "127.0.0.1:8081")
	if erro1 != nil {
		go Server()
	}
	time.Sleep(5 * time.Second) 
	Connection()
}
