package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print("Digite: ")

		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			continue
		} else if text == "exit" {
			fmt.Println("Exiting the prompt")
			break
		}
		fmt.Println(">", ft_reverse(text))
	}
}
