package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("day-1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var rightArr []int
	var leftArr []int

	for scanner.Scan() {
		text := scanner.Text()
		splits := strings.Split(text, "  ")

		left := strings.TrimSpace(splits[0])
		right := strings.TrimSpace(splits[1])

		leftInt, err := strconv.Atoi(left)
		if err != nil {
			log.Fatal(err)
		}

		rightInt, err := strconv.Atoi(right)
		if err != nil {
			log.Fatal(err)
		}

		leftArr = append(leftArr, leftInt)
		rightArr = append(rightArr, rightInt)

	}

	slices.Sort(leftArr)
	slices.Sort(rightArr)

	var diff int

	for i := 0; i < len(leftArr); i++ {
		left := leftArr[i]
		right := rightArr[i]

		diffRaw := float64(left - right)
		diff += int(math.Abs(diffRaw))
	}

	log.Printf("Diff: %s", fmt.Sprint(diff))

}
