package src

import (
	"bufio"
	"fmt"
	"net"
	"os"
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
			mes, err := bufio.NewReader(conexao).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				os.Exit(3)
			}
			Hub(mes, conexao)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
