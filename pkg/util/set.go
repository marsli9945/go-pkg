package util

type Set[T string | int] map[T]struct{}

func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}

func (s Set[T]) Remove(item T) {
	delete(s, item)
}

func (s Set[T]) Contains(item T) bool {
	_, exists := s[item]
	return exists
}

func (s Set[T]) ToList() []T {
	var list []T
	for k, _ := range s {
		list = append(list, k)
	}
	return list
}
