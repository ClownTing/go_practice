package main 

import (
	"fmt"
	"time"
)

func MY_PL1(ss, ss1 []int) []int {
	ss=append(ss, ss1...)
	return ss
}

func MY_PL2(ss []int, index int) []int {
	ss = append(ss[:index], ss[index+1:]...)
	fmt.Printf("---------ss=%v ----------\n", ss)
	return ss
}	

func MY_PL3(ss, a []int, index int) []int {
	rear := append([]int{}, ss[index:]...)
	ss = append(ss[:index], a...)
	ss = append(ss, rear...)
	return ss
}
	
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
	ss:=[]int{0,1,2,3,4}
	ss1:=[]int{5,6,7,8,9}
	fmt.Println("ss is",MY_PL1(ss, ss1))
	fmt.Println(ss)
	index := 3
	ss = MY_PL2(ss, index)
	fmt.Printf("ss is %v\t len(ss) is %v\n", ss,len(ss))
	fmt.Println(ss)	
	a :=[]int{78,96,85,64}
	fmt.Println("ss is",MY_PL3(ss, a, index ) )
	fmt.Println(ss)
}
