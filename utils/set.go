package utils

type Set[T comparable] map[T]struct{}

func MakeSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

func (s Set[T]) Rm(value T) {
	delete(s, value)
}

func (s Set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Equal(other Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	for elem := range s {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}
