package src

import (
	"fmt"
	"net"
)

func Server() {
	ch := make(chan net.Conn, 5)
	fmt.Println("Servidor criado!")
	go ListenConnection(ch)
	fmt.Println("Conexão aceita...")
	conexao := <-ch
	go Chat(conexao)
}
