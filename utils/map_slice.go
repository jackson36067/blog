package utils

// MapSlice 泛型 Map 函数：把 []T 映射成 []R
func MapSlice[T any, R any](src []T, mapper func(T) R) []R {
	result := make([]R, len(src))
	for i, v := range src {
		result[i] = mapper(v)
	}
	return result
}
