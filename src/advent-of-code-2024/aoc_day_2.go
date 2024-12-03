package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func solveMainDay2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	countUnsafe := 0
	for scanner.Scan() {
		line := scanner.Text()
		stringArray := strings.Fields(line)

		intArray := make([]int, len(stringArray))
		for i, str := range stringArray {
			intArray[i], err = strconv.Atoi(str)
			if err != nil {
				log.Println("Error converting string to int:", err)
				return 0
			}
		}
		log.Println(intArray)
		currentSolution := solve(intArray)
		countUnsafe += currentSolution
		log.Println("current row returns :", currentSolution)
	}

	log.Println("countUnsafe:", countUnsafe)
	return countUnsafe
}

func solve(intArray []int) int {
	// 1 is safe, 0 unsafe to add to counter
	isIncreasing := false

	for i := 0; i < len(intArray); i++ {
		log.Println("current iteration:", i, "current value", intArray[i])

		if i == 0 {
			isIncreasing = intArray[i] < intArray[i+1]
		}
		if i < len(intArray)-1 {
			if isIncreasing && intArray[i] >= intArray[i+1] {
				log.Println("unsafe due to (", intArray[i], "&", intArray[i+1], "decreasing after increasing")
				return 0
			} else if !isIncreasing && intArray[i] <= intArray[i+1] {
				log.Println("unsafe due to either no change across iterations (", intArray[i], "&", intArray[i+1], "or increasing after decreasing")
				return 0
			}
		}

		if i+1 == len(intArray) {
			log.Println("safe end of array due to end of array")
			return 1
		}

		diff := intArray[i] - intArray[i+1]
		if diff < 0 {
			diff = -diff
		}
		if diff > 3 {
			log.Println("unsafe differing more than 3")
			return 0
		}
	}
	log.Println("safe due to default return")
	return 1
}
