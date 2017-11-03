package userprof

import "fmt"

type User struct {
	id        int
	firstName string
	surname   string
	age       int
	pass      string
}

func NewUser(id, age int, firstName, surname, pass string) *User {
	return &User{
		id:        id,
		age:       age,
		surname:   surname,
		pass:      pass,
		firstName: firstName,
	}
}

func (u *User) SendEmail(text, addr string) bool {

	fmt.Printf("%v \n", u)
	fmt.Println(u.firstName)
	u.firstName = "Петр"

	return true
}
