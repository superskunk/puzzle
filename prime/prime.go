package prime

import "math"

// IsPrime is a func to know if a number is prime. It bases its algorithm in that every composite number has a prime factor less than or equal to its square root.
func IsPrime(n uint64) bool {
	// 1 is not considered a prime number, because it is needed to be divided by two numbers to be a prime number (the number itself and 1)
	if n < 2 {
		return false
	}

	sqrtN := uint64(math.Sqrt(float64(n)))
	for i := uint64(2); i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
