package main

import (
	"bufio"
	"fmt"
	"os"

	"./piglatin"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	encryptedText := piglatin.Encrypt(text)
	fmt.Print("Encrypted text: ", encryptedText)
}
