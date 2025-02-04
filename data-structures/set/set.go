package set

// we need both an int set and also a string set
type IntSet map[int]struct{}

//methods: size, clear, remove, add, has, keys

func (s IntSet) size() int {
	return len(s)
}

func (s IntSet) remove(num int) bool {
	delete(s, num)
	return true
}

func (s IntSet) has(num int) bool {
	_, present := s[num]
	return present
}

func (s IntSet) clear() bool {
	clear(s)
}

func (s IntSet) keys() []int {
	var keys []int
	for v := range s {
		keys = append(keys, v)
	}
	return keys
}

func (s IntSet) add(num int) bool {
	s[num] = struct{}{}
	return true
}

func NewIntSet(input []int) IntSet {
	s := IntSet{}
	for _, v := range input {
		s[v] = struct{}{}
	}
	return s
}

type StringSet map[string]struct{}

func (s StringSet) add(input string) bool {
	s[input] = struct{}{}
	return true
}

func (s StringSet) remove(input string) bool {
	delete(s, input)
	return true
}

func (s StringSet) clear() {
	clear(s)
}

func (s StringSet) has(input string) bool {
	_, present := s[input]
	return present
}

func (s StringSet) keys() []string {
	var keys []string
	for v := range s {
		keys = append(keys, v)
	}
	return keys
}

func (s StringSet) size() int {
	return len(s)
}

func NewStringSet(input ...[]string) StringSet {
	s := StringSet{}
	for _, v := range input {
		s[v] = struct{}{}
	}
	return s
}
