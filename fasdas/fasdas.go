package fasdas

import "fmt"

func Fasdas() {
	r := input()
	fmt.Println(fuck(r))

}

func fuck(s int) int {
	var temp = 1
	for i := 1; i <= s; i++ {
		temp *= i
	}
	return temp
}

func input() int {
	var s int
	fmt.Scanln(&s)
	return s
}
