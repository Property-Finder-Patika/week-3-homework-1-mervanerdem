//The  original problem : Implement prime factorization algorithm

/* Project Euler same problem:
Largest Prime Faktor
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

package main

import "fmt"

func main() {
	val := 600851475143 //Given number
	i := 2
	for i = 2; i < val; i++ {
		/* The algorithm
		count start from smallest prime factor 2
		if given number can divide by `i` it's a prime number.
		if `i` has many factor in the given number in `for` loop we divided number.
		Lastly we divide given number for keep range lower.
		*/
		for val%i == 0 {
			val /= i
		}
	}
	if val == 1 { //if value is factor of a prime number value will be 1 in this case we use this method.
		fmt.Println(i - 1)
	} else {
		fmt.Println(val)
	}

}
