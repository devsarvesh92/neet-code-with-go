package main

import (
	"fmt"
)

func main() {

	// array
	var myarray = [3]int32{1, 2, 3}
	fmt.Println(myarray[0])

	myarray1 := []int32{1123, 2, 3}
	fmt.Println(myarray1)

	// slice
	var sl []int32
	sl = append(sl, 7)
	sl = append(sl, myarray1...)
	fmt.Println(sl)

	// map
	var mymap = map[string]int32{"Adam": 32}
	res, ok := mymap["Jason"]

	var arr = [3]int32{1, 2, 3}
	var slice = []int32{2, 34}

	fmt.Print(arr)
	fmt.Print(slice)
	fmt.Print(mymap)

	if ok {
		fmt.Println(res)
	} else {
		fmt.Println("Go fuck yourself")
	}

	for idx, val := range myarray1 {
		fmt.Println(idx, val)
	}

	var i int32 = 0
	for i < 10 {
		fmt.Println(i)
		i += 1
	}
}
