package set

// we need both an int set and also a string set
type IntSet struct {
	list map[int]struct{}
}

//methods: size, clear, remove, add, has, keys

func (s *IntSet) Size() int {
	return len(s.list)
}

func (s *IntSet) Remove(num int) bool {
	delete(s.list, num)
	return true
}

func (s *IntSet) Has(num int) bool {
	_, present := s.list[num]
	return present
}

func (s *IntSet) Clear() {
	s.list = make(map[int]struct{})
}

func (s *IntSet) Keys() []int {
	var keys []int
	for v := range s.list {
		keys = append(keys, v)
	}
	return keys
}

func (s *IntSet) Add(num int) bool {
	s.list[num] = struct{}{}
	return true
}

func NewIntSet(input ...int) *IntSet {
	s := IntSet{}
	s.list = make(map[int]struct{})
	for _, v := range input {
		s.list[v] = struct{}{}
	}
	return &s
}

type StringSet struct {
	list map[string]struct{}
}

func (s *StringSet) Add(input string) bool {
	s.list[input] = struct{}{}
	return true
}

func (s *StringSet) Remove(input string) bool {
	delete(s.list, input)
	return true
}

func (s *StringSet) Clear() {
	s.list = make(map[string]struct{})
}

func (s *StringSet) Has(input string) bool {
	_, present := s.list[input]
	return present
}

func (s *StringSet) Keys() []string {
	var keys []string
	for v := range s.list {
		keys = append(keys, v)
	}
	return keys
}

func (s *StringSet) Size() int {
	return len(s.list)
}

func NewStringSet(input ...string) *StringSet {
	s := &StringSet{}
	s.list = make(map[string]struct{})
	for _, v := range input {
		s.list[v] = struct{}{}
	}
	return s
}
