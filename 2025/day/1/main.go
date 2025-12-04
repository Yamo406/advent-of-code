package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	dial := make([]int, 100)
	for i := range 99{
		dial[i] = i
	}

	pointer := dial[50]

	zeroCounts := 0

	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rotation := scanner.Text()
		direction := rotation[0]

		degree, _:= strconv.Atoi(rotation[1:])

		for range degree{
			switch direction {
				case 'L':
					pointer -= 1
				case 'R':
					pointer += 1
				}
				if (pointer % 100) == 0{
					zeroCounts++
				}

		}
	}


	fmt.Println(zeroCounts)

}