package mapReduceFilter

func Map[T comparable](iterable []T, cb func(el T) T) []T {
	var res = make([]T, 0, len(iterable))

	for _, el := range iterable {
		res = append(res, cb(el))
	}

	return res
}

func Filter[T comparable](iterable []T, cb func(el T) bool) []T {
	var res = make([]T, 0, len(iterable))

	for _, el := range iterable {
		if cb(el) {
			res = append(res, el)
		}
	}

	return res
}

func Reduce[T comparable](iterable []T, cb func(acc T, cur T) T, initVal T) T {
	res := initVal

	for _, el := range iterable {
		res = cb(res, el)
	}

	return res
}
