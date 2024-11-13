package src

import (
	"fmt"
)

func Hub(a string) {
	for _, i := range arrayConections {
		_, err := fmt.Fprintf(i, "CHAT GERAL: %s\n", a)
		if err != nil {
			fmt.Println(err)
		}
	}
}
