package main

import (
	"fmt"
	"time"
)

func sum (slice []int, ch chan int) {

	result := 0

	for _, element := range slice {

		result = result + element
	} 

	ch <- result
}

func main () {

	fmt.Println(time.Now())

	tunel := make(chan int)

	numbers := []int{1, 2, 3, 4, 5, 6}
	
	go sum(numbers[:len(numbers)/2], tunel)   	 // 1 + 2 + 3 = 6
	go sum(numbers[len(numbers)/2:], tunel)      // 4 + 5 + 6 = 15

	sum1, sum2 := <- tunel, <- tunel

	fmt.Println(sum1, sum2, sum1 + sum2)

	fmt.Println(time.Now())
	
}


