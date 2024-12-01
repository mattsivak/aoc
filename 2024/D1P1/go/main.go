package main

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	list1 := make([]int, 0)
	list2 := make([]int, 0)

	for _, line := range lines {
		if line == "" {
			continue
		}
		nums := strings.Split(line, "   ")

		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	output := 0

	slices.Sort(list1)
	slices.Sort(list2)

	for i := 0; i < len(list1); i++ {
		if list1[i] > list2[i] {
			output += list1[i] - list2[i]
		} else {
			output += list2[i] - list1[i]
		}
	}

	println(output)
}
