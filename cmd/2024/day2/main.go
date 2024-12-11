package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func good(d []int, s int) bool {
	for i := 0; i < len(d)-1; i++ {
		diff := d[i] - d[i+1]

		if diff < 1 || diff > 3 {
			if s == 0 {
				return false
			}

			for j := i; j <= i+1; j++ {

				newSlice := make([]int, 0, len(d)-1)
				newSlice = append(newSlice, d[:j]...)
				newSlice = append(newSlice, d[j+1:]...)

				if good(newSlice, 0) {
					return true
				}
			}
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("day2.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		numbers := make([]int, len(fields))

		for i, field := range fields {
			numbers[i], err = strconv.Atoi(field)
			if err != nil {
				panic(err)
			}
		}
		data = append(data, numbers)
	}

	for s := 0; s <= 1; s++ {
		safe := 0

		for _, d := range data {
			if good(d, s) {
				safe++
				continue
			}

			reversed := make([]int, len(d))
			for i := range d {
				reversed[i] = d[len(d)-1-i]
			}

			if good(reversed, s) {
				safe++
			}
		}
		fmt.Println(safe)
	}
}
