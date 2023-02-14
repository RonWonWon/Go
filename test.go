package main

import (
	"fmt"
	"math/rand"
	"time"
)

var X = 1

func main() {
	rand.Seed(time.Now().Unix())
	//x := 10
	X = 10
	//var y = 0
	{
		y := 11
		fmt.Println(y)
	}
	//fmt.Println(x, y)
	var luck = rand.Intn(2)
	names := [3]string{"Rachit", "Ujjawal", "Vidit"}
	if luck == 1 {
		fmt.Println(names[rand.Intn(3)], "tum rahoge Hyderabad")
	} else {
		fmt.Println(names[rand.Intn(3)], "tum jaoge Chennai")
	}
}

func fun(a int, b int) int {
	fmt.Println(a + b)
	return a + b
}
