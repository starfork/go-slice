package slice

// 直接使用的函数
func Unique[T comparable](s Slice[T]) Slice[T] {

	tmp := map[T]T{}

	for _, item := range s {
		tmp[item] = item
	}
	var res []T
	for _, k := range tmp {
		res = append(res, k)
	}
	s = res //给后续用
	return s
}

func Rand[T comparable](s Slice[T]) T {
	max := len(s)
	idx, _ := RangeRand(0, int64(max-1))
	return s[idx : idx+1][0]
}
