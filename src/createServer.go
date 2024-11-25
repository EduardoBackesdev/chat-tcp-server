package src

import (
	"net"
)

func CreateServer() {
	_, erro1 := net.Dial("tcp", "192.168.0.105:8081")
	if erro1 != nil {
		go Server()
	}
	// ch := make(chan net.Conn, 5)
	// chMes := make(chan string)
	// go ListenConnection(ch)
	// time.Sleep(5 * time.Second)
	// for {
	// 	go Connection(chMes)
	// 	mes := <-chMes
	// 	fmt.Println(mes)
	// 	con := <-ch
	// 	fmt.Println(con)
	// 	Hub(mes)
	// }

}
