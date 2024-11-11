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
		//Hub(mes, conn)
		// fmt.Println("Entrou para digitar")
		// conexao := <-ch
		// leitor := bufio.NewReader(os.Stdin)
		// texto, textoErr := leitor.ReadString('\n')
		// if textoErr != nil {
		// 	fmt.Println(textoErr)
		// 	os.Exit(3)
		// }
		// fmt.Fprintf(conexao, texto+"\n")
		// fmt.Println("")
	}
}
