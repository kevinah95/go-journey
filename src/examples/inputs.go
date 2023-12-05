package examples

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func inputs() {
	fmt.Println("¿Qué me quieres decir?")
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	fmt.Println(s + "!")
}
