package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	lines := []string{}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	tempCalSum := 0
	calSums := []int{}

	for _, line := range lines {
		if len(line) != 0 {
			intValue, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			tempCalSum += intValue
		} else {
			calSums = append(calSums, tempCalSum)
			tempCalSum = 0
		}
	}

	fmt.Println("The largest calorie number is: ", getHighestCalories(calSums))
	fmt.Println("Top three elves calories are: ", getTopThreeCalories(calSums))
	fmt.Printf("Those three elves have %d calories in total.", sumTopThreeCalories(calSums))
}

func getHighestCalories(calSlice []int) int {
	var largestNumber, temp int
	for _, calsum := range calSlice {
		if calsum > temp {
			temp = calsum
			largestNumber = temp
		}
	}
	return largestNumber
}

func getTopThreeCalories(calSlice []int) []int {
	sort.Ints(calSlice)
	return []int{calSlice[len(calSlice)-1], calSlice[len(calSlice)-2], calSlice[len(calSlice)-3]}
}

func sumTopThreeCalories(calSlice []int) int {
	var sum int
	for _, v := range getTopThreeCalories(calSlice) {
		sum += v
	}
	return sum
}
