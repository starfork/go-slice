package slice

import "reflect"

type Col []any

func Colunm(source any, key string) Col {
	var rs []any
	val := reflect.ValueOf(source)
	for i := 0; i < val.Len(); i++ {
		k := val.Index(i).Elem()
		f := k.FieldByName(key).Interface()
		rs = append(rs, f)
	}
	return rs
}

func (e Col) Unique() Col {
	tmp := map[any]any{}

	for _, v := range e {
		tmp[v] = v
	}
	var rs Col
	for _, v := range tmp {
		rs = append(rs, v)
	}

	return rs
}

// 此处没有判定类型是否相同，使用时自行判定
func (e Col) String() []string {
	rs := []string{}
	for _, v := range e {
		rs = append(rs, v.(string))
	}
	return rs
}

// 此处没有判定类型是否相同，使用时自行判定
func (e Col) Uint32() []uint32 {
	rs := []uint32{}
	for _, v := range e {
		rs = append(rs, v.(uint32))
	}
	return rs
}

// 此处没有判定类型是否相同，使用时自行判定
func (e Col) Uint64() []uint64 {
	rs := []uint64{}
	for _, v := range e {
		rs = append(rs, v.(uint64))
	}
	return rs
}
