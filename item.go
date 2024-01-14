package slice

// 数字类型特有的
type Item interface {
	int | float32 | float64 | uint32 | uint64 | string
}

type SliceItem[T Item] []T

func NewNumber[T Item](a []T) SliceItem[T] {
	return a
}

// 求和
func (s SliceItem[T]) Sum() T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

// 取最大值。
func (s SliceItem[T]) Max() T {
	var max T = s[0]
	for _, item := range s {
		if item > max {
			max = item
		}
	}
	return max
}
