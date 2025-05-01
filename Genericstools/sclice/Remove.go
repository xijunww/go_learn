package slice

import "fmt"

// 2. 删除（Remove）
// 功能：删除切片中的指定位置元素。
// 参数：切片、要删除的元素。
// 返回值：删除元素后的切片。如果元素不存在，原始切片不变。
func Remove[T any](s []T, index int) (res []T, err error) {
	//判断索引值是否合法
	if index < 0 || index >= len(s) {
		return nil, fmt.Errorf("索引值%d不合法，slice长度为%d", index, len(s))
	}
	//创建res切片
	res = make([]T, len(s)-1)
	//将index前面的数值放入res
	copy(res[:index], s[:index])
	//将index+1到最后的值放入res
	copy(res[index:], s[index+1:])
	return res, nil
}
