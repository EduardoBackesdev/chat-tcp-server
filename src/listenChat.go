package src

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func ListenChat(a net.Conn) {
	for {
		// Assim que receber o controle de nova linha (\n), processa a mensagem recebida
		mensagem, erro3 := bufio.NewReader(a).ReadString('\n')
		if erro3 != nil {
			fmt.Println(erro3)
			os.Exit(3)
		}

		// escreve no terminal a mensagem recebida
		fmt.Print("Mensagem recebida:", string(mensagem))
	}
}
