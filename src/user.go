package src

import (
	"bufio"
	"fmt"
	"os"
)

func User() string {
	var user string
	fmt.Println("Digite seu nome:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		user := scanner.Text()
		fmt.Println("Parabéns voce entrou no chat,  ", user)
	}
	return user
}
