package slice

import "fmt"

func FindTest() {
	s1 := []int{1, 3, 5, 6}
	fun1 := func(int2 int) bool {
		if int2 == 5 {
			return true
		}
		return false
	}
	v, res := Find(s1, fun1)
	fmt.Println(v, res)

	s2 := []string{"小明", "小明", "大明"}
	fun2 := func(string2 string) bool {
		if string2 == "大明" {
			return true
		}
		return false
	}
	v2, res2 := Find(s2, fun2)
	fmt.Println(v2, res2)

}

func FindAllTest() {
	s1 := []int{1, 3, 5, 6}
	fun1 := func(int2 int) bool {
		if int2 >= 5 {
			return true
		}
		return false
	}
	res := FindAll(s1, fun1)
	fmt.Println(res)

	s2 := []string{"小明", "小明", "大明", "xiaoming"}
	fun2 := func(string2 string) bool {
		if string2 != "大明" {
			return true
		}
		return false
	}
	res2 := FindAll(s2, fun2)
	fmt.Println(res2)
}
