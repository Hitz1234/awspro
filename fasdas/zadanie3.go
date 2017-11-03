package fasdas

import "fmt"

type User struct {
	name string
	age  int
}

func zadanie3() {
	r := new(User)
	fmt.Println(&User{name: r.name, age: r.age})
}
func zaoplName(r *User) string {
	fmt.Print("Давай сюда имя: ")
	fmt.Scanln(&r.name)
	return r.name
}
func zaoplAge(r *User) int {
	fmt.Print("Давай сюда старость: ")
	fmt.Scanln(&r.age)
	return r.age
}
