package src

import (
	"fmt"
	"net"
)

func Server() {
	ch := make(chan net.Conn, 5)
	chMes := make(chan string)
	fmt.Println("Servidor criado!")
	go ListenConnection(ch)
	go Connection(chMes)
	for {

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
