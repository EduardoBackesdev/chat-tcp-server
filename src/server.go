package src

import (
	"fmt"
	"net"
	"time"
)

func Server() {
	fmt.Println("Servidor criado!")
	ch := make(chan net.Conn, 5)
	chMes := make(chan string)
	go ListenConnection(ch)
	time.Sleep(5 * time.Second)
	for {
		go Connection(chMes)
		mes := <-chMes
		con := <-ch
		Hub(mes, con)
	}
}
