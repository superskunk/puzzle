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
	listOfPrimesWithoutEvenNumbersAction
	exit
)

var stdin *bufio.Reader

func menu() action {

	option := 0
	for option < 1 || option > 4 {
		fmt.Println()
		fmt.Println("1. Check if a number is prime")
		fmt.Println("2. List of prime numbers up to N")
		fmt.Println("3. List of prime numbers up to N using only Odd numbers")
		fmt.Println("4. Exit")
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

	for option != exit {
		option = menu()
		switch option {
		case isPrimeAction:
			fmt.Print("Introduce a Number: ")
			fmt.Scanf("%d", &n)
			isPrime(n)
		case listOfPrimesAction:
			fmt.Println("Generating a list of prime numbers")
			fmt.Print("Introduce the number up to you want to generate prime numbers: ")
			fmt.Scanf("%d", &n)
			numbers := prime.GeneratePrimes(n)
			printPrimes(numbers)
		case listOfPrimesWithoutEvenNumbersAction:
			fmt.Println("Generating a list of prime numbers using only Odd numbers")
			fmt.Print("Introduce the number up to you want to generate prime numbers: ")
			fmt.Scanf("%d", &n)
			numbers := prime.GeneratePrimesWithoutUsingEvenNumbers(n)
			printPrimes(numbers)
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

func printPrimes(numbers []bool) {
	for number, isPrime := range numbers {
		if isPrime {
			fmt.Printf("%d ", number)
		}
	}
}
