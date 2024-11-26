package src

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func ListenChat(a []net.Conn) {
	for {
		for _, i := range a {
			mensagem, erro3 := bufio.NewReader(i).ReadString('\n')
			if erro3 != nil {
				fmt.Println(erro3)
				os.Exit(3)
			}

			// escreve no terminal a mensagem recebida
			fmt.Print("Mensagem recebida:", string(mensagem))
		}
	}
}
