package util

import (
	"log"
	"strconv"
)

func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("invalid number in the puzzle input")
	}
	return n
}
