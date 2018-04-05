package main 

import (
	"fmt"
	"time"
)

func main() {
	var array = [5]int{5,4,3,2,1}
	var slice []int
	slice1 := []int{int('a'),int('b')}
	slice = array[1:4]
	fmt.Println("array len is",len(array))
	fmt.Printf("current Time is%v\n",time.Now())
	fmt.Println("slice is",slice)
	for index,value := range slice1 {
		if slice1[index] == 97 {
			fmt.Println(index,value)
		}	
	}
}
