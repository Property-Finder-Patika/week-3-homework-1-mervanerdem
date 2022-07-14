/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
Find the sum of all the primes below two million.
*/
package main

import "fmt"

func SieveOfEratosthenes(n int) []int { //get a function for solve the problem
	integers := make([]bool, n+1) // make a slice for all number under the given number and slice is bool
	for i := 2; i < n+1; i++ {
		integers[i] = true // make all slice true
	}
	for p := 2; p*p < n; p++ {
		/*
			Algoritm: if the number is prime number don't touch it and set all coefficients to false
			if the number is false skip it
			in this step we find prime number in slice bool-shaped(true)
		*/
		if integers[p] {
			for i := p * 2; i <= n; i += p {
				integers[i] = false
			}
		}
	}

	var primes []int
	for p := 2; p <= n; p++ { //here we find bool-shaped prime numbers and add it to a new slice called prime
		if integers[p] {
			primes = append(primes, p)
		}
	}
	return primes // and return it
}

func main() {
	primes := SieveOfEratosthenes(2000000) //using the function
	fmt.Println(len(primes))
	sum := 0
	for _, prime := range primes { // sum prime numbers
		sum1 := sum + prime
		if sum1 < sum {
			fmt.Println(sum)
		} else {
			sum = sum1
		}
	}
	fmt.Println("Result:", sum) // print to screen
}
