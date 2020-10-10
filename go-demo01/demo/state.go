package main

import (
	"fmt"
	"time"
)

func main() {
	// const a string = "1"
	// var b int = 2
	// c := 2
	// fmt.Println("Hello world")
	// fmt.Println(a, "--a")
	// fmt.Println(b, "--b")
	// fmt.Println(c, "--c")

	// const d, e, f string = "dd", "ee", "ff"
	// fmt.Println(d, e, f)
	// const name, age = "lili", 20
	// fmt.Println(name, age)

	// array()
	// slice()
	time.Sleep(time.Second * 3)
	fmt.Println("main")
}

func array() {
	//一维数组
	var arr_1 [5]int
	fmt.Println(arr_1)

	var arr_2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr_2)

	arr_3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr_3)

	arr_4 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr_4)

	arr_5 := [5]int{0: 3, 1: 5, 4: 6}
	fmt.Println(arr_5)

	//二维数组
	var arr_6 = [3][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	fmt.Println(arr_6)

	arr_7 := [3][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	fmt.Println(arr_7)

	arr_8 := [...][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {0: 3, 1: 5, 4: 6}}
	fmt.Println(arr_8)
}

func slice() {
	var sli_5 []int = make([]int, 5, 8)
	fmt.Println(sli_5)

	sli := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("sli[0:3:2] ==", sli[0:3:3])
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli[0:3:3]), cap(sli[0:3:3]), sli[0:3:3])
}
