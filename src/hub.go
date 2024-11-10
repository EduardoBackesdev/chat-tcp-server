package src

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Hub(a string, c net.Conn) {
	var arrayConection = []net.Conn{c}
	for _, i := range arrayConection {
		a, err := bufio.NewReader(i).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
		fmt.Print("CHAT GLOBAL: ", a)
	}
	// for _, i := range arrayConection {
	// 	fmt.Println(i)
	// 	a, err := bufio.NewReader(i).ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(3)
	// 	}
	// 	fmt.Print("CHAT GLOBAL: ", a)
	// 	fmt.Println("OPAAA")
	// }
}
