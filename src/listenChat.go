package src

import (
	"fmt"
	"net"
)

func ListenChat(a net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := a.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		msg := string(buf[:n])
		fmt.Printf("Received message: %s\n", msg)
	}
}
