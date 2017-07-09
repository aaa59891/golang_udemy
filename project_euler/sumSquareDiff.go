package main

import "fmt"

func sumSquareDiff(num int) int{
	var squareSum int
	var sum int
	for i := 1; i <= num; i++{
		squareSum += i * i
		sum += i
	}
	return sum * sum - squareSum
}

func main() {
	fmt.Println(sumSquareDiff(100))
}
