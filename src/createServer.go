package src

import (
	"net"
)

func CreateServer() {
	_, erro1 := net.Dial("tcp", "192.168.0.105:8081")
	if erro1 != nil {
		go Server()
	}
}
