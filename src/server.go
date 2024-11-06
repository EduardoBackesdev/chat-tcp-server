package src

import (
	"bufio"
	"fmt"
	"os"
)

func Server() {
	fmt.Println("Servidor criado!")
	go ListenConnection()
	fmt.Println("Conexão aceita...")

	for {
		mes, err := bufio.NewReader(conexao).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
		fmt.Print("CHAT GLOBAL: ",mes)
	}

}
