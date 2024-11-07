package src

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Chat(ch net.Conn) {
	for {
		mes, err := bufio.NewReader(ch).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
		fmt.Print("CHAT GLOBAL: ",mes)
	}
}