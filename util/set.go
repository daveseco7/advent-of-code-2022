package util

type IntSet map[int]struct{}

func (s *IntSet) AddValue(value int) {
	if *s == nil {
		m := make(map[int]struct{})
		*s = m
	}
	(*s)[value] = struct{}{}
}

func (s *IntSet) GetValues() []int {
	if *s == nil {
		return nil
	}
	values := make([]int, 0, len(*s))
	for v := range *s {
		values = append(values, v)
	}
	return values
}

func (s *IntSet) Contains(value int) bool {
	if *s == nil {
		return false
	}
	_, ok := (*s)[value]
	return ok
}

func (s *IntSet) Delete(value ...int) {
	for _, v := range value {
		delete(*s, v)
	}
}
