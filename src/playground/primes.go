
package main;



import (
	"fmt"
	"math"
)


func main() {
	fmt.Println("Hello!");
	fmt.Println("-1: ", isPrime(-1));
	fmt.Println("1: ", isPrime(1));
	fmt.Println("2: ", isPrime(2));
	fmt.Println("13: ", isPrime(13));
	fmt.Println("1299827: ", isPrime(1299827));
}


// Prime numbers are natural numbers
// Integers > 1 that are only divisible by 1 and themselves

// n is <= 1 -> not a prime by definition
// n == 2 -> n is prime, only divisible by 1 and 2
// n >= 2 -> check if it is divisible by all numbers >= 2 up to sqrt(n)
func isPrime(n int) bool {

	if n <= 1 {
		return false;
	} else if n == 2 {
		return true;
	} else {
		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
			if (n % i == 0) {
				return false;
			}
		}
	}
	return true;
}
