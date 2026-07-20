package advance

//The [T int | float64 | string] is a type constraint — T must be one of these
func Max[T int | float64 | string](a, b T) T {
	if a > b {
		return a
	}
	return b
}

//any is the new alias for interface{} — generic without restriction:
func Filter[T any](s []T, keep func(T) bool)[]T{
	var result []T
	for _, x := range s {
		if keep(x){
			result = append(result, x)
		}
	}
	return result
}

// Generic types:
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(x T) {
	s.items = append(s.items, x)
}

func (s *Stack[T]) Pop() (T, bool){
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	x := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return x, true
}