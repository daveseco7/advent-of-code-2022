package day3_test

import (
	"testing"

	"github.com/daveseco7/advent-of-code-2022/day3"
	"github.com/daveseco7/advent-of-code-2022/util"
)

func TestExe1(t *testing.T) {
	t.Run("puzzle sample", func(T *testing.T) {
		const expected = 157
		const filePath = "sample.txt"

		lineArray, err := util.ReadLines(filePath)
		if err != nil {
			t.Fatal(err)
		}

		ans := day3.Exe1(lineArray)
		if ans != expected {
			t.Errorf("invalid response %d", ans)
		}
	})
	t.Run("puzzle input", func(T *testing.T) {
		const expected = 7875
		const filePath = "input1.txt"

		lineArray, err := util.ReadLines(filePath)
		if err != nil {
			t.Fatal(err)
		}

		ans := day3.Exe1(lineArray)
		if ans != expected {
			t.Errorf("invalid response %d", ans)
		}
	})
}

func TestExe2(t *testing.T) {
	t.Run("puzzle sample", func(T *testing.T) {
		const expected = 70
		const filePath = "sample.txt"

		lineArray, err := util.ReadLines(filePath)
		if err != nil {
			t.Fatal(err)
		}

		ans := day3.Exe2(lineArray)
		if ans != expected {
			t.Errorf("invalid response %d", ans)
		}
	})
	t.Run("puzzle input", func(T *testing.T) {
		const expected = 2479
		const filePath = "input1.txt"

		lineArray, err := util.ReadLines(filePath)
		if err != nil {
			t.Fatal(err)
		}

		ans := day3.Exe2(lineArray)
		if ans != expected {
			t.Errorf("invalid response %d", ans)
		}
	})
}
