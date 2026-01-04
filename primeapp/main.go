package main

import "fmt"

func main() {
  n :=121
  _ ,msg :=isPrime(n)
  fmt.Println(msg)
}

func isPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition", n)

	}
	if n < 0 {
		return false, fmt.Sprintf("%d is not prime, by definition Negative number are not prime", n)

	}

	// use the modulas operator  repeatedly if we have a prime number

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime  number
			return false, fmt.Sprintf("%d is not prime, because it is divisible by %d", n, i)
		}
	}
  return true, fmt.Sprintf("%d is  prime number !", n)
}


fmt.Println