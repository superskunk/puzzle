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
	var n uint64
	var option action
	printPrimes(5)

	for option != exit {
		option = menu()
		switch option {
		case isPrimeAction:
			fmt.Print("Introduce a Number: ")
			fmt.Scanf("%d", &n)
			isPrime(n)
		case listOfPrimesAction:
			fmt.Println("Generating a list of primes")
			fmt.Print("Introduce the number up to you want to generate prime numbers: ")
			fmt.Scanf("%d", &n)
			printPrimes(n)
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

func printPrimes(n uint64) {
	fmt.Println()
	numbers := prime.GeneratePrimes(n)
	for number, isPrime := range numbers {
		if isPrime {
			fmt.Printf("%d ", number)
		}
	}
}
