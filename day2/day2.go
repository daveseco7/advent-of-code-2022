package day2

type abcFormat string

const (
	A abcFormat = "A"
	B abcFormat = "B"
	C abcFormat = "C"
)

type xyzFormat string

const (
	X xyzFormat = "X"
	Y xyzFormat = "Y"
	Z xyzFormat = "Z"
)

func (xyz xyzFormat) score() int {
	switch xyz {
	case X:
		return 1
	case Y:
		return 2
	case Z:
		return 3
	}

	// if valid input is provided this should never occur
	panic("xyzFormat: invalid input provided")
}

func (xyz xyzFormat) misinterpretation(abc abcFormat) int {
	score := xyz.score()

	switch abc {
	case A:
		switch xyz {
		case X:
			// draw
			score += 3
		case Y:
			// xyz win
			score += 6
		case Z:
			// xyz lost
			score += 0
		}
	case B:
		switch xyz {
		case X:
			// xyz lost
			score += 0
		case Y:
			// draw
			score += 3
		case Z:
			// xyz win
			score += 6
		}
	case C:
		switch xyz {
		case X:
			// xyz win
			score += 6
		case Y:
			// xyz lost
			score += 0
		case Z:
			// draw
			score += 3
		}
	}

	return score
}

func (xyz xyzFormat) correctInterpretation(abc abcFormat) int {
	score := 0
	switch xyz {
	case X:
		// xyz MUST lost
		score += 0

		switch abc {
		case A:
			score += 3
		case B:
			score += 1
		case C:
			score += 2
		}
	case Y:
		// xyz MUST draw
		score += 3

		switch abc {
		case A:
			score += 1
		case B:
			score += 2
		case C:
			score += 3
		}

	case Z:
		// xyz MUST win
		score += 6

		switch abc {
		case A:
			score += 2
		case B:
			score += 3
		case C:
			score += 1
		}

	}

	return score
}

type play struct {
	p1 abcFormat
	p2 xyzFormat
}

type game []play

func buildGame(lines []string) game {
	g := make(game, 0, len(lines))
	for _, playToParse := range lines {
		p := play{
			p1: abcFormat(playToParse[0]),
			p2: xyzFormat(playToParse[2]),
		}

		g = append(g, p)
	}

	return g
}

func Exe1(lines []string) int {
	g := buildGame(lines)

	total := 0
	for _, p := range g {
		total += p.p2.misinterpretation(p.p1)
	}

	return total
}

func Exe2(lines []string) int {
	g := buildGame(lines)

	total := 0
	for _, p := range g {
		total += p.p2.correctInterpretation(p.p1)
	}

	return total
}
