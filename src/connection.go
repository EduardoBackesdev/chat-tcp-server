package src

import (
	"fmt"
	"net"
)

func Connection() {
	_, erro1 := net.Dial("tcp", "127.0.0.1:8081")
	if erro1 != nil {
		fmt.Println("Servidor não criado!")
		go Server()
		return
	}
}