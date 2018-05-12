package main

import (
	"fmt"
)

func square(number int) int {
	return number*number
}

func appendString(title string) string {
	return title+" mind fucked"
}

func main(){
	numbers := []int{1, 2,3, 4}
	reuslt := EachSeries(Collection{numbers}, square)
	fmt.Println(reuslt)

	result := EachSeries(Collection{[]string{"power", "terrible", "quick"}}, appendString)
	fmt.Println(result)
}