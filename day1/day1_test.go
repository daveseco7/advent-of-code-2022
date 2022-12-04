package day1

import (
	"testing"

	"github.com/daveseco7/advent-of-code-2022/util"
)

func TestExe1(t *testing.T) {
	t.Run("puzzle sample", func(T *testing.T) {
		const expected = 24000
		const filePath = "sample.txt"

		lineArray, err := util.ReadLines(filePath)
		if err != nil {
			t.Fatal(err)
		}

		tch := make(topCalorieHolder, 1) // preload the holder struct with zeros

		ans := exe(lineArray, tch)
		if ans != expected {
			t.Errorf("invalid response %d", ans)
		}
	})
	t.Run("puzzle input", func(T *testing.T) {
		const expected = 66719
		const filePath = "input1.txt"

		lineArray, err := util.ReadLines(filePath)
		if err != nil {
			t.Fatal(err)
		}

		tch := make(topCalorieHolder, 1) // preload the holder struct with zeros

		ans := exe(lineArray, tch)
		if ans != expected {
			t.Errorf("invalid response %d", ans)
		}
	})
}

func TestExe2(t *testing.T) {
	t.Run("puzzle sample", func(T *testing.T) {
		const expected = 45000
		const filePath = "sample.txt"

		lineArray, err := util.ReadLines(filePath)
		if err != nil {
			t.Fatal(err)
		}

		tch := make(topCalorieHolder, 3) // preload the holder struct with zeros

		ans := exe(lineArray, tch)
		if ans != expected {
			t.Errorf("invalid response %d", ans)
		}
	})
	t.Run("puzzle input", func(T *testing.T) {
		const expected = 66719
		const filePath = "input1.txt"

		lineArray, err := util.ReadLines(filePath)
		if err != nil {
			t.Fatal(err)
		}

		tch := make(topCalorieHolder, 3) // preload the holder struct with zeros

		ans := exe(lineArray, tch)
		if ans != expected {
			t.Errorf("invalid response %d", ans)
		}
	})
}
