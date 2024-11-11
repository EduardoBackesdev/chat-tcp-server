package src

import (
	"net"
	"time"
)

func CreateServer() {
	_, erro1 := net.Dial("tcp", "192.168.0.105:8081")
	if erro1 != nil {
		go Server()
	}
	time.Sleep(5 * time.Second)
	//Connection()
}
