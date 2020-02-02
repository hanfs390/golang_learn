package main

import "fmt"

func sum(values [] int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	resultChan <- sum // send the result to channel
}

func main() {

	values := [] int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	resultChan := make(chan int, 2)
	go sum(values[:len(values)/2], resultChan) // the first five
	go sum(values[len(values)/2:], resultChan) // the end five
	sum1, sum2 := <-resultChan, <-resultChan // recv the result

	fmt.Println("Result: ", sum1, sum2, sum1+sum2)
}
