package main

import "fmt"

func main(){
	floor := 0
	var input string

	fmt.Scan(&input)

	for i, v := range input {
		switch v {
			case '(':
				floor++
			case ')':
				floor--
		}
		if floor == -1{
			fmt.Printf("entered basement at index %d\n", i)
		}
	}
	fmt.Println(floor)

}
