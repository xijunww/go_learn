package slice

import "fmt"

func RemoveTest() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []string{"1", "2", "3", "4", "5", "6"}
	s11, err1 := Remove(s1, 3)
	s22, err2 := Remove(s2, 3)
	fmt.Println(s11, err1)
	fmt.Println(s22, err2)
}
