package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var safe int
var checked = false

func part_one_day2(file string) {
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(f), "\n")
	arr := []int{}

	for _, line := range lines {

		if line == "" {
			continue
		}
		parts := strings.Fields(line)

		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				log.Fatal("error converting to num")
			}
			arr = append(arr, num)
		}
		checksafe(arr)
		if checked {
			part_one_differ(arr)
		}
		arr = []int{}
	}
	fmt.Println(safe)

}

func checksafe(arr []int) {
	if arr[0] > arr[1] {
		gradual_decrease_checker(arr)
	} else {
		gradual_increase_checker(arr)
	}
}

func part_one_differ(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		diff := math.Abs(float64(arr[i] - arr[i+1]))

		if diff > 3 || diff == 0 {
			return
		}
	}
	safe++
}

func gradual_increase_checker(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			checked = false
			return
		}
	}
	checked = true
	return
}

func gradual_decrease_checker(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] <= arr[i+1] {
			checked = false
			return
		}
	}
	checked = true
	return
}

func main() {
	part_one_day2("day2.input")
}
