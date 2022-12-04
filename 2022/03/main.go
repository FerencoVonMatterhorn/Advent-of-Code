package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	leftCompartmentList, rightCompartmentList := splitRucksackInv(lines)
	priorities := createPrioritieMap()
	prioritieSum := 0
	elfGroupPrioritiesSum := 0

	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
		fmt.Println(leftCompartmentList[i])
		fmt.Println(rightCompartmentList[i])
		duplicateItem := findDuplicateItemInTwoStrings(leftCompartmentList[i], rightCompartmentList[i])
		fmt.Printf("Duplicate Item is %s with prioritie %d\n", duplicateItem, priorities[duplicateItem])
		prioritieSum += priorities[duplicateItem]
	}
	fmt.Printf("The sum of all priorities is %d\n", prioritieSum)

	for i := 0; i < len(lines); i += 3 {
		fmt.Println(lines[i])
		fmt.Println(lines[i+1])
		fmt.Println(lines[i+2])
		duplicateItem := findDuplicateItemInThreeStrings(lines[i], lines[i+1], lines[i+2])
		fmt.Printf("Duplicate Item in Group is %s with prioritie %d\n", duplicateItem, priorities[duplicateItem])
		elfGroupPrioritiesSum += priorities[duplicateItem]
		fmt.Println()
	}

	fmt.Printf("The sum of all Elf group priorities is %d\n", elfGroupPrioritiesSum)
}

func splitRucksackInv(inventorylist []string) ([]string, []string) {
	var lC, rC []string
	for _, line := range inventorylist {
		lC = append(lC, line[0:len(line)/2])
		rC = append(rC, line[len(line)/2:len(line)])
	}
	return lC, rC
}

func findDuplicateItemInTwoStrings(leftString string, rightString string) string {
	for i := 0; i < len(leftString); i++ {
		for j := 0; j < len(rightString); j++ {
			if leftString[i] == rightString[j] {
				return string(leftString[i])
			}
		}
	}
	return ""
}

func findDuplicateItemInThreeStrings(firstString string, secondString string, thridString string) string {
	for i := 0; i < len(firstString); i++ {
		for j := 0; j < len(secondString); j++ {
			for k := 0; k < len(thridString); k++ {
				if firstString[i] == secondString[j] {
					if firstString[i] == thridString[k] {
						return string(firstString[i])
					}
				}
			}
		}
	}
	return ""
}

func createPrioritieMap() map[string]int {
	var pMap = make(map[string]int)
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i := 0; i < len(alphabet); i++ {
		pMap[string(alphabet[i])] = i + 1
	}

	return pMap
}

func createElfGroups(allLines []string) []string {
	var elfGroups []string
	for i := 0; i < len(allLines); i++ {
		for j := 0; j < 3; j++ {
			elfGroups = append(elfGroups)
		}
	}
	return elfGroups
}
