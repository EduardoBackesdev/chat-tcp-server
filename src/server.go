package src

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Server() {
	ch := make(chan net.Conn, 5)
	fmt.Println("Servidor criado!")
	go ListenConnection(ch)
	Connection()
	for {
		fmt.Println("Entrou para digitar")
		conexao := <-ch
		leitor := bufio.NewReader(os.Stdin)
		texto, textoErr := leitor.ReadString('\n')
		if textoErr != nil {
			fmt.Println(textoErr)
			os.Exit(3)
		}
		fmt.Fprintf(conexao, texto+"\n")
		fmt.Println("")
	}
}
