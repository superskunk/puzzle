package prime

import "math"

// IsPrime is a func to know if a number is prime. It bases its algorithm in that every composite number has a prime factor less than or equal to its square root.
func IsPrime(n uint64) bool {
	// 1 is not considered a prime number, because it is needed to be divided by two numbers to be a prime number (the number itself and 1)
	if n < 2 {
		return false
	}

	sqrtN := uint64(math.Floor(math.Sqrt(float64(n))))
	for i := uint64(2); i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// GeneratePrimes follows de Sieve of Eratosthenes to calculate a list with the prime numbers between 2 and n https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func GeneratePrimes(n uint64) []bool {
	primeList := make([]bool, n+1)
	for i := range primeList {
		primeList[i] = true
	}
	primeList[0], primeList[1] = false, false

	prime := uint64(2)
	sqrtN := uint64(math.Floor(math.Sqrt(float64(n))))
	for prime <= sqrtN {
		// The multiples of each primer shall be marked as non prime.
		for i := prime * prime; i <= n; i += prime {
			primeList[i] = false
		}

		prime = getNextPrime(prime, &primeList)
	}

	return primeList
}

// getNextPrime returns the next number marked as non prime in the primeLis
func getNextPrime(currentPrime uint64, primeList *[]bool) uint64 {
	n := uint64(len(*primeList))
	nextPrime := currentPrime + 1
	for nextPrime < n && !(*primeList)[nextPrime] {
		nextPrime++
	}
	return nextPrime
}
