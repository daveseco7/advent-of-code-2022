package day1

import (
	"testing"

	"github.com/daveseco7/advent-of-code-2022/util"
)

func TestExe1(t *testing.T) {
	const expected = 1624

	lineArray, err := util.ReadLinesAsInt(filePath)
	if err != nil {
		t.Fatal(err)
	}

	ans := exe1(lineArray)
	if ans != expected {
		t.Errorf("invalid response %d", ans)
	}
}

func TestExe2(t *testing.T) {
	const expected = 1653

	lineArray, err := util.ReadLinesAsInt(filePath)
	if err != nil {
		t.Fatal(err)
	}

	ans := exe2(lineArray)
	if ans != expected {
		t.Errorf("invalid response %d", ans)
	}
}
