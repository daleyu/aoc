package set

// we need both an int set and also a string set
type IntSet struct {
	list map[int]struct{}
}

//methods: size, clear, remove, add, has

func (s *IntSet) size() int {
	return len(s.list)
}

func (s *IntSet) add(num int) bool {
	s.list[num] = struct{}{}
	return true
}

func NewIntSet() *IntSet {
	s := &IntSet{}
	s.list = make(map[int]struct{})
	return s
}

type StringSet struct {
	list map[int]struct{}
}
