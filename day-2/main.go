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

	file, err := os.Open("day-2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var numbers [][]int

	for scanner.Scan() {
		text := scanner.Text()
		stringsOfNumbers := strings.Split(text, " ")

		//log.Println(stringsOfNumbers)

		var line []int
		for _, number := range stringsOfNumbers {
			num, _ := strconv.Atoi(number)
			line = append(line, num)
		}
		numbers = append(numbers, line)
	}
	lines, checks := countSafeLines(numbers)
	fmt.Printf("%s Safe Lines", fmt.Sprint(lines))
	fmt.Println("")
	fmt.Printf("%s Partially Safe Lines", fmt.Sprint(checks))

}

func countSafeLines(numbers [][]int) (safeLines int, partiallySafeLines int) {
	for _, nums := range numbers {
		var safetyChecks int
		for i := 0; i < len(nums); i++ {
			copied := make([]int, len(nums))
			copy(copied, nums)

			copied = append(copied[:i], copied[i+1:]...)

			//log.Println(copied)
			if checkSafety(copied) {
				safetyChecks += 1
			}

		}
		if checkSafety(nums) {
			safeLines += 1
		}
		//fmt.Println(safetyChecks)
		if safetyChecks >= 1 {
			partiallySafeLines += 1
		}
	}

	return safeLines, partiallySafeLines
}

func isDiffCorrect(nums []int) bool {
	var previousNum int
	firstRun := true

	for _, num := range nums {
		if firstRun == true {
			firstRun = false
			previousNum = num
			continue
		}

		diff := math.Abs(float64(previousNum - num))
		if diff > 3 || diff < 1 || previousNum == num {
			return false
		}

		previousNum = num
	}

	return true
}

func checkSafety(nums []int) bool {
	reversedNums := make([]int, len(nums))
	copy(reversedNums, nums)
	slices.Reverse(reversedNums)
	//log.Println(reversedNums)
	if (slices.IsSorted(nums) || slices.IsSorted(reversedNums)) && isDiffCorrect(nums) {
		//log.Printf("%s is safe", fmt.Sprint(nums))
		//log.Println(slices.IsSorted(nums), slices.IsSorted(reversedNums), isDiffCorrect(nums))
		return true
	}
	return false
}
