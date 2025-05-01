package slice

import "fmt"

func AddTest() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []string{"1", "2", "3", "4", "5"}
	s11, err1 := Add(s1, 123321, 3)
	s22, err2 := Add(s2, "123321", 3)
	fmt.Println(s11, err1)
	fmt.Println(s22, err2)
}
