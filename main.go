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

	/////////////////////////////////////////////////////////

	fmt.Print("\nEnter a, b and c (in this exact order) of the quadratic equation (ax^2 + bx + c = 0), to find its roots: ")

	var a, b, c float64
	fmt.Scanf("%f", &a)
	fmt.Scanf("%f", &b)
	fmt.Scanf("%f", &c)

	x1, x2, err := lib.RootsOfQuadraticEquation(a, b, c)

	if err != nil {
		fmt.Printf("coundn't find the roots of the quadratic equation: %s\n", err)
	} else {
		fmt.Printf("The roots of the quadratic equation are: x = %f, x2 = %f\n", x1, x2)
	}

	/////////////////////////////////////////////////////////
	fmt.Print("\nEnter the length of an UUID to create it: ")

	var length int
	fmt.Scanf("%d", &length)

	uuid, err := lib.CreateUUID(length)

	if err != nil {
		fmt.Printf("coundn't create UUID: %s\n", err)
	} else {
		fmt.Printf("UUID: %s\n", uuid)
	}
}
