package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main(){
	var lxwxh string
	totalArea, ribbon := 0, 0

	for range 1000 {
		fmt.Scan(&lxwxh)
		lwhSlice := strings.Split(lxwxh, "x")
		l, _ := strconv.Atoi(lwhSlice[0])
		w, _ := strconv.Atoi(lwhSlice[1])
		h, _ := strconv.Atoi(lwhSlice[2])

		intSlice := []int{l,w,h}
		sort.Ints(intSlice)
		smallestSide := intSlice[0] * intSlice[1]

		surfaceArea := 2*l*w + 2*w*h + 2*h*l

		totalArea += smallestSide + surfaceArea
		ribbon += 2 * (intSlice[0] + intSlice[1]) +  (intSlice[0] * intSlice[1] * intSlice[2])

	}

	fmt.Println(totalArea)
	fmt.Println(ribbon)

}
