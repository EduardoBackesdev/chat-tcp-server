package src

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Connection() {
	conexao, erro1 := net.Dial("tcp", "127.0.0.1:8081")
	if erro1 != nil {
		fmt.Println("Servidor não criado!")
		go Server()
	}
	nome := ""
	if nome == "" {
		nome = User()
	}
	for {
		leitor := bufio.NewReader(os.Stdin)
		fmt.Println("Digite sua mensagem:")
		texto, textoErr := leitor.ReadString('\n')
		if textoErr != nil {
			fmt.Println(textoErr)
			os.Exit(3)
		}

		fmt.Fprintf(conexao, nome+": "+texto+"\n")

		mensagem, err3 := bufio.NewReader(conexao).ReadString('\n')
		if err3 != nil {
			fmt.Println(err3)
			os.Exit(3)
		}
		fmt.Print("Resposta do servidor: " + mensagem)
	}
}
