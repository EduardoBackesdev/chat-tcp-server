package src

import (
	"fmt"
	"net"
)

func Hub(a string, c net.Conn) {
	var arrayConection = []net.Conn{c}
	fmt.Println("aqio")
	for _, i := range arrayConection {
		fmt.Println("aqio 2")
		_, err := fmt.Fprintf(i, "CHAT GLOBAL: %s\n", a)
		if err != nil {
			fmt.Println("Erro ao enviar mensagem:", err)
		}
		fmt.Println("aqio 3")
	}
}
