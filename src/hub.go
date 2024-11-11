package src

import (
	"net"
)

func Hub(a string, c net.Conn) {
	var arrayConection = []net.Conn{c}
	for _, i := range arrayConection {
		Chat(i)
	}
}
