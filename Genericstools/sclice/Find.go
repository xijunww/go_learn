package slice

// Find
// 功能描述：查找切片中第一个满足条件的元素。
// 返回：匹配的元素、是否找到（bool）。
//
// FindAll
// 功能描述：查找切片中所有满足条件的元素。
// 返回：包含所有匹配元素的切片，永不为 nil。
func Find[T any](s []T, match func(T) bool) (element T, res bool) {
	//循环比对
	for _, val := range s {
		if match(val) {
			return val, true
		}
	}
	var t T
	return t, false
}

func FindAll[T any](s []T, match func(T) bool) (res []T) {
	//创建需要输出的切片
	res = make([]T, 0, len(s))
	//循环遍历原切片的val
	for _, val := range s {
		if match(val) {
			//如果符合方法需求，放入res
			res = append(res, val)
		}
		//不符合的丢弃
	}
	return res
}
