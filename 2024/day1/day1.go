package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sum(nums []int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}

func part_one() string {
	input, err := os.ReadFile("./day1.input")

	if err != nil {
		fmt.Println("Error Reading Input File!", input)
	}

	lines := strings.Split(string(input), "\n")

	var leftlist, rightlist []int

	for _, line := range lines {

		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("error converting to int", err)
			continue
		}

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("error converting to int", err)
			return ""
		}
		leftlist = append(leftlist, left)
		rightlist = append(rightlist, right)

	}
	sort.Ints(leftlist)
	sort.Ints(rightlist)

	result := []int{}

	for i := 0; i < len(leftlist); i++ {
		diff := leftlist[i] - rightlist[i]
		differences := (math.Abs(float64(diff)))
		result = append(result, int(differences))
	}

	answer := sum(result)

	return fmt.Sprintf("%v\n", answer)
}

var my_map = make(map[int]int)
var leftlist []int

func part_two_helper() string {
	input, err := os.ReadFile("./day1.input")

	if err != nil {
		fmt.Println("Error Reading the file ")
	}

	lines := strings.Split(string(input), "\n")

	var rightlist []int

	for _, line := range lines {

		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("error converting to int", err)
			continue
		}

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("error converting to int", err)
			return ""
		}
		leftlist = append(leftlist, left)
		rightlist = append(rightlist, right)

	}

	for _, num := range rightlist {
		my_map[num]++
	}

	return fmt.Sprintf("%v\n%v", leftlist, rightlist)
}

var rizz = []int{}

func part_two() int {
	for _, num := range leftlist {

		if count, exists := my_map[num]; exists {
			rizz = append(rizz, num*count)
		} else {
			rizz = append(rizz, 0)
		}
	}

	return sum(rizz)

}

func main() {
	res1 := part_one()
	fmt.Printf("Part One: \n%v", res1)

	part_two_helper()
	res2 := part_two()
	fmt.Printf("Part Two: %d\n", res2)

}
