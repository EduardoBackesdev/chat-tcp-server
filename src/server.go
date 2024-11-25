package src

import (
	"fmt"
)

func Server() {
	fmt.Println("Servidor criado!")
	for {
		go ListenConnection()
	}
}
