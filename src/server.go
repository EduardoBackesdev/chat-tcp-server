package src

import (
	"fmt"
	"net"
	"time"
)

func Server() {

	ch := make(chan net.Conn, 5)
	fmt.Println("Servidor criado!")
	go ListenConnection(ch)
	fmt.Println("Conexão aceita...")
	for {
		select {
		case conexao := <-ch:
			go Chat(conexao)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
