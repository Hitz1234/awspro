package fasdas

import (
	"fmt"
	"strings"
)

func zadanie4() {
	a := zapola()
	fmt.Println(a)

}
func zapola() string {
	var a string
	fmt.Scanln(&a)
	s := strings.Split(a, "")
	naoborot(s)
	return strings.Join(s, "")
}
func naoborot(x []string) {
	temp := len(x) - 1
	for i := 0; i < len(x)/2; i++ {
		x[i], x[temp-i] = x[temp-i], x[i]
	}
}
