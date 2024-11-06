package src

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Server() {
	fmt.Println("Servidor criado!")

	ln, erro1 := net.Listen("tcp", ":8081")
	if erro1 != nil {
		fmt.Println(erro1)
		os.Exit(3)
	}

	// aceitando conexões
	conexao, erro2 := ln.Accept()
	if erro2 != nil {
		fmt.Println(erro2)
		os.Exit(3)
	}

	defer ln.Close()

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
