package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
// Print a welcome message
  intro()
// create a channel to indicate when the use  want to quite
  doneChan :=make(chan bool)


//start a goroutine to read user input and run program
go readUerInput(doneChan)

// block until the doneChan gets a value
<-doneChan

//close the channel
close(doneChan)
// say goodbuy

fmt.Println("Goodbye. ")
}
 func readUerInput(doneChan chan bool){
	scanner :=bufio.NewScanner(os.Stdin)
	for {
		res , done := checkNumbers(scanner)
		if done {
			doneChan <- false
			return
		}

		fmt.Println(res)
		promt()
	}
 }

func checkNumbers(scanner *bufio.Scanner)(string , bool){
	// read user input
scanner.Scan()
	//check to see if the user want to quite
	if strings.EqualFold(scanner.Text(), "q"){
		return "", true
	}
	// try to convert what use type into an int
	numtoCheck , err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter an whole number", false
	}

	_ ,msg := isPrime(numtoCheck)
	return msg, false
	
}

func intro(){
	fmt.Println("Is it Prime?")
	fmt.Println("--------------")
	fmt.Println("Enter a whole number , and we'll tell you if it is prime or not .Enter q to quite")
	promt()
}

func promt(){
	fmt.Print("->")
}
func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime, by definition!"
	}

	// use the modulus operator repeatedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime number
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d!", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}