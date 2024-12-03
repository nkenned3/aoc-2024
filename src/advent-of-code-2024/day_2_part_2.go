package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func solveDay2Part2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	countSafe := 0
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
		isUnSafe := solvePart2(intArray)
		if !isUnSafe {
			countSafe++
		}
		log.Println("is current row unsafe?", isUnSafe)
	}

	log.Println("count safe:", countSafe)
	return countSafe
}

func solvePart2(intArray []int) bool {
	isIncreasing := false
	problems := 0
	for i := 0; i < len(intArray); i++ {
		log.Println("current iteration:", i, "current value", intArray[i])

		if i == 0 {
			isIncreasing = intArray[i] < intArray[i+1]
		}
		if i < len(intArray)-1 {
			if isIncreasing && intArray[i] >= intArray[i+1] {
				log.Println("unsafe due to (", intArray[i], "&", intArray[i+1], "decreasing after increasing")
				intArray, i = removeItem(intArray, i)
				problems = incrementProblems(problems)
				continue
			} else if !isIncreasing && intArray[i] <= intArray[i+1] {
				log.Println("unsafe due to either no change across iterations (", intArray[i], "&", intArray[i+1], "or increasing after decreasing")
				intArray, i = removeItem(intArray, i)
				problems = incrementProblems(problems)
				continue
			}
		}

		if i+1 == len(intArray) {
			log.Println("safe end of array due to end of array")
			break
		}

		diff := intArray[i] - intArray[i+1]
		if diff < 0 {
			diff = -diff
		}
		if diff > 3 {
			log.Println("unsafe differing more than 3")
			intArray, i = removeItem(intArray, i)
			problems = incrementProblems(problems)
			continue
		}
	}

	isUnsafe := problems > 1
	log.Println(" problems: " + strconv.Itoa(problems) + " isUnSafe: " + strconv.FormatBool(isUnsafe))
	return isUnsafe
}

func incrementProblems(problems int) int {
	problems += 1
	return problems
}
func removeItem(intArray []int, index int) ([]int, int) {
	result := append(intArray[:index], intArray[index+1:]...)
	index--
	return result, index
}
