package day4_test

import (
	"testing"

	"github.com/daveseco7/advent-of-code-2022/day4"
	"github.com/daveseco7/advent-of-code-2022/util"
)

func TestExe1(t *testing.T) {
	t.Run("puzzle sample", func(T *testing.T) {
		const expected = 2
		const filePath = "sample.txt"

		lineArray, err := util.ReadLines(filePath)
		if err != nil {
			t.Fatal(err)
		}

		ans := day4.Exe1(lineArray)
		if ans != expected {
			t.Errorf("invalid response %d", ans)
		}
	})
	t.Run("puzzle input", func(T *testing.T) {
		const expected = 550
		const filePath = "input1.txt"

		lineArray, err := util.ReadLines(filePath)
		if err != nil {
			t.Fatal(err)
		}

		ans := day4.Exe1(lineArray)
		if ans != expected {
			t.Errorf("invalid response %d", ans)
		}
	})
}

func TestExe2(t *testing.T) {
	t.Run("puzzle sample", func(T *testing.T) {
		const expected = 4
		const filePath = "sample.txt"

		lineArray, err := util.ReadLines(filePath)
		if err != nil {
			t.Fatal(err)
		}

		ans := day4.Exe2(lineArray)
		if ans != expected {
			t.Errorf("invalid response %d", ans)
		}
	})
	t.Run("puzzle input", func(T *testing.T) {
		const expected = 931
		const filePath = "input1.txt"

		lineArray, err := util.ReadLines(filePath)
		if err != nil {
			t.Fatal(err)
		}

		ans := day4.Exe2(lineArray)
		if ans != expected {
			t.Errorf("invalid response %d", ans)
		}
	})
}
