package src

import (
	"fmt"
	"time"
)

func Server() {
	fmt.Println("Servidor criado!")
	go ListenConnection()
	go Get_user()
	go ListenChat()
	for {
		time.Sleep(1 * time.Second)
	}
}
