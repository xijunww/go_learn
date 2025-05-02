package types

import "fmt"

func NewUser() {
	u := User{}
	fmt.Printf("%v\n", u)
	fmt.Printf("%+v\n", u) //仔细一些

	//up是一个指针
	up := &User{Name: "kiki"}
	fmt.Printf("%+v \n", up)
	fmt.Println(*up)
	up2 := new(User) //up和up2的效果差不多，up这个方法用的多
	fmt.Printf("%+v \n", up2)

	u4 := User{Name: "Tom", Age: 12}
	u4.Name = "Jerry"
	fmt.Printf("%+v \n", u4)

	var up3 *User
	fmt.Println(up3)
	//nil上访问字段
	//fmt.Println(up3.Age)
}

type User struct {
	Name string
	Age  int
}

func (u User) ChangeName(name string) {
	fmt.Printf("change name中 u的地址 %p\n", &u)
	u.Name = name
}

func (u *User) ChangeAge(age int) {
	fmt.Printf("change age中 u的地址 %p\n", u)
	u.Age = age
}

func ChangeUser() {
	u1 := User{Name: "Tom", Age: 18}
	fmt.Printf("u1 的地址 %p \n", &u1)
	u1.ChangeAge(53)
	u1.ChangeName("Jerry")
	fmt.Printf("%+v", u1)
}
