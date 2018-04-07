package main

import(
	"fmt"
)

const LIM = 40

func main() {
	colors := map[string]string {
	"AliceBlue": "#f0f8ff",
	"coral": "#ff7F50",
	"DarkGray": "#a9a9a9",
	"ForesGreen": "#228b22",
	}
	for key,value := range colors {
		fmt.Printf("key: %s , value: %s\n",key,value)
	}
	delete(colors,"coral")
	for key , value := range colors{
		fmt.Printf("key: %s , value: %s\n",key,value)
	}
	value := colors["Red"]
	if value != " " {
		fmt.Println(value)
	} else{
		fmt.Println("无红色键值对")
	}
	// dict := make(map[string]int)
	fmt.Println("---------------------next----------------")

	var array [LIM]int
	for i := 0; i< LIM; i++ {
		array[i] = fibonacci(i)	
	}
	fmt.Println(array)
}

func fibonacci (n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1)+fibonacci(n-2)
	}
	return res
}
