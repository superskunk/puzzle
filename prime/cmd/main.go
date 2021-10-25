package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/superskunk/puzzle/prime"
)

type action int

const (
	isPrimeAction action = iota + 1
	listOfPrimesAction
	exit
)

var stdin *bufio.Reader

func menu() action {
	option := 0
	for option < 1 || option > 3 {
		fmt.Println()
		fmt.Println("1. Create Face List")
		fmt.Println("2. List of Faces lists")
		fmt.Println("3. Exit")
		fmt.Printf("\nChoose an option....:")
		if _, err := fmt.Fscanf(stdin, "%d", &option); err != nil {
			// In case of not introducing a number
			option = 0
		}
		stdin.ReadLine() //This line is necessary to flush the buffer because there is a "\n" left

	}
	return action(option)
}

func main() {
	stdin = bufio.NewReader(os.Stdin)

	var option action
	for option != exit {
		option = menu()
		switch option {
		case isPrimeAction:
			var n uint64
			fmt.Print("Introduce a Number: ")
			fmt.Scanf("%d", &n)
			isPrime(n)
		case listOfPrimesAction:
			fmt.Println("Generating a list of primes")

		}
	}

}

func isPrime(n uint64) {
	if prime.IsPrime(n) {
		fmt.Printf("Number %d is a prime number\n", n)
	} else {
		fmt.Printf("Number %d is not a prime number\n", n)
	}
}
