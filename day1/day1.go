package day1

import "github.com/daveseco7/advent-of-code-2022/util"

type TopCalorieHolder []int

func (tch TopCalorieHolder) setPossibleTopValue(currentCalories int) {
	for index, topCalories := range tch {
		if currentCalories > topCalories {
			// set new top value
			tch[index] = currentCalories

			// re-run the algorithm with the value that just got evicted
			tch.setPossibleTopValue(topCalories)

			// new value already set, break slice looping
			break
		}
	}
}

func Exe(lines []string, tch TopCalorieHolder) (topTotalCalories int) {
	var tempCal = 0

	for i := 0; i < len(lines); i++ {
		isNewLine := len(lines[i]) == 0
		isLastLine := len(lines)-1 == i

		if !isNewLine {
			tempCal += util.MustAtoi(lines[i])
		}

		// if no calories input is given, or if its the last vlaue to be processed
		// we finish processing a elf calorie package and run the top value algorithm
		if isNewLine || isLastLine {
			tch.setPossibleTopValue(tempCal)
			tempCal = 0
		}
	}

	for _, calories := range tch {
		topTotalCalories += calories
	}

	return topTotalCalories
}
