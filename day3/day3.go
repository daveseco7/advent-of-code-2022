package day3

import (
	"unicode"
)

func runeSliceToSet(r []rune) map[rune]struct{} {
	s := make(map[rune]struct{}, 0)
	for i := 0; i < len(r); i++ {
		s[r[i]] = struct{}{}
	}

	return s
}

func commonFrom2(r1, r2 []rune) rune {
	set1 := runeSliceToSet(r1)
	set2 := runeSliceToSet(r2)

	var r rune
	for k := range set1 {
		if _, ok := set2[k]; ok {
			r = k
		}
	}

	return r
}

func commonFrom3(r1, r2, r3 []rune) rune {
	set1 := runeSliceToSet(r1)
	set2 := runeSliceToSet(r2)
	set3 := runeSliceToSet(r3)

	var r rune
	for k := range set1 {
		if _, ok2 := set2[k]; ok2 {
			if _, ok3 := set3[k]; ok3 {
				r = k
			}
		}
	}

	return r
}

func totalPriority(commonItems []rune) int {
	// since rune is alias for int32 that saves ASCII code
	// use a offset to convert the value to the appropriate challenge scale
	const lowercaseOffset = 96 // reset counter to 0, so we can have 1-26 scale
	const uppercaseOffset = 64 // reset counter to 0 and add the 26 (lowercase scale), so we can have 27-52 scale

	priority := 0
	for _, v := range commonItems {
		if unicode.IsUpper(v) {
			priority += int((v - uppercaseOffset) + 26)
		} else {
			priority += int((v - lowercaseOffset))
		}
	}

	return priority
}

func Exe1(lines []string) int {
	commonItems := make([]rune, 0)

	for _, l := range lines {
		size := len(l) / 2

		str1, str2 := l[:size], l[size:]
		r1, r2 := []rune(str1), []rune(str2)

		commonItems = append(commonItems, commonFrom2(r1, r2))
	}

	return totalPriority(commonItems)
}

func Exe2(lines []string) int {
	commonItems := make([]rune, 0)

	for i := 0; i < len(lines); i += 3 {
		r1, r2, r3 := []rune(lines[i]), []rune(lines[i+1]), []rune(lines[i+2])

		commonItems = append(commonItems, commonFrom3(r1, r2, r3))
	}

	return totalPriority(commonItems)
}
