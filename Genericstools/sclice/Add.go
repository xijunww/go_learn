package slice

import "fmt"

// 1. 添加（Add）
// 功能：向切片中添加一个元素或多个元素。
// 参数：切片、要添加的元素或元素列表。
// 返回值：更新后的切片。
func Add[T any](s []T, element T, index int) ([]T, error) {
	//检查索引是否有效
	if index < 0 || index > len(s) {
		return nil, fmt.Errorf("索引index越界：%d (slice长度：%d)", index, len(s))
	}
	//创建切片，长度为之前切片长度加一
	res := make([]T, len(s)+1)
	//复制索引index之前的元素
	copy(res[:index], s[:index])
	//在index放入新元素
	res[index] = element
	//添加剩下元素
	copy(res[index+1:], s[index:])
	return res, nil
}
