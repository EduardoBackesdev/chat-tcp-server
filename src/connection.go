package src

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"net"
	"os"
)

func Connection() {
	cores := []string{
		"\033[91m",  
		"\033[92m",  
		"\033[93m",  
		"\033[94m",  
		"\033[95m",  
		"\033[96m",  
	}
	conexao, err := net.Dial("tcp", "192.168.1.113:8081")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	nome := ""
	if nome == "" {
		nome = User()
	}
	ran := cores[rand.IntN(len(cores))]
	for {
		leitor := bufio.NewReader(os.Stdin)
		texto, textoErr := leitor.ReadString('\n')
		if textoErr != nil {
			fmt.Println(textoErr)
			os.Exit(3)
		}

		fmt.Fprintf(conexao, ran+nome+": "+"\033[0m"+texto+"\n")
		fmt.Println("")

		// _, err3 := bufio.NewReader(conexao).ReadString('\n')
		// if err3 != nil {
		// 	fmt.Println(err3)
		// 	os.Exit(3)
		// }
		
	}
}