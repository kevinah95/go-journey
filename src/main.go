package main

import (
	"bufio"
	"fmt"
	"github.com/kevinah95/go-journey/src/menu"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func main() {

loop:
	for {
		fmt.Println("Seleccione una opción")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) quit")
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			menu.PrintMenu()
		case "2":
			err := menu.AddItem()
			if err != nil {
				fmt.Println(fmt.Errorf("invalid input: %w", err))
			}

		case "q":
			break loop

		default:
			fmt.Println("Opción desconocida")

		}
	}

}
