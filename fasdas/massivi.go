package fasdas

import (
	"fmt"
	"strings"
	"strconv"
)

func massivi() {
	r := inputmas()
	fmt.Println(izv(r))
}

func inputmas() string {
	var s string
	fmt.Scanln(&s)
	return s
}

func izv(s string) []int {
	var a []string = strings.Split(s, ",")
	var chicla []int = []int{}

	for i := 0; i < len(a); i++ {
		temp, err := strconv.Atoi(a[i])
		if err != nil {
			panic(err)
		}
		chicla = append(chicla, temp)
	}
	return chicla
}
