package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	lib "github.com/bazeeko/onelab_homework1_lib"
)

func main() {
	fmt.Print("Enter a string to change the case of its letters: ")
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		log.Fatalln("coundn't read the string from console: ", err)
	}

	fmt.Printf("Result: %s", lib.ChangeLetterCase(input))
}
