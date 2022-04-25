package mapReduceFilter

func Map(iterable []any, cb func(el any) any) []any {
	var res = make([]any, len(iterable))

	for _, el := range iterable {
		res = append(res, cb(el))
	}

	return res
}

func Filter(iterable []any, cb func(el any) bool) []any {
	var res = make([]any, len(iterable))

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
