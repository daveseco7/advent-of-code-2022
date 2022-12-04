package day2_test

import (
	"testing"

	"github.com/daveseco7/advent-of-code-2022/day2"
	"github.com/daveseco7/advent-of-code-2022/util"
)

func TestExe1(t *testing.T) {
	t.Run("puzzle sample", func(T *testing.T) {
		const expected = 15
		const filePath = "sample.txt"

		lineArray, err := util.ReadLines(filePath)
		if err != nil {
			t.Fatal(err)
		}

		ans := day2.Exe1(lineArray)
		if ans != expected {
			t.Errorf("invalid response %d", ans)
		}
	})
	t.Run("puzzle input", func(T *testing.T) {
		const expected = 12679
		const filePath = "input1.txt"

		lineArray, err := util.ReadLines(filePath)
		if err != nil {
			t.Fatal(err)
		}

		ans := day2.Exe1(lineArray)
		if ans != expected {
			t.Errorf("invalid response %d", ans)
		}
	})
}

func TestExe2(t *testing.T) {
	t.Run("puzzle sample", func(T *testing.T) {
		const expected = 12
		const filePath = "sample.txt"

		lineArray, err := util.ReadLines(filePath)
		if err != nil {
			t.Fatal(err)
		}

		ans := day2.Exe2(lineArray)
		if ans != expected {
			t.Errorf("invalid response %d", ans)
		}
	})
	t.Run("puzzle input", func(T *testing.T) {
		const expected = 14470
		const filePath = "input1.txt"

		lineArray, err := util.ReadLines(filePath)
		if err != nil {
			t.Fatal(err)
		}

		ans := day2.Exe2(lineArray)
		if ans != expected {
			t.Errorf("invalid response %d", ans)
		}
	})
}
