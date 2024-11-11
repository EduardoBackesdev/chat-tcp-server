package src

import (
	"fmt"
	"net"
)

func Hub(a string, conn net.Conn) {
	arrayConections := []net.Conn{}
	arrayConections = append(arrayConections, conn)
	for _, i := range arrayConections {
		_, err := fmt.Fprintf(i, "CHAT GERAL: %s\n", a)
		if err != nil {
			fmt.Println(err)
		}
	}
}
