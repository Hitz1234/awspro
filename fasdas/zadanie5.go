package fasdas

import (
	"strconv"
	"fmt"
	"strings"
	"awspro/utils"
)

func zadanie5() {
	checkArray()
}

func checkArray() bool {
	var b bool
	w := inputArray()
	n := len(w)
	fmt.Println("Массив: ", w)
	for i := 0; i <= (n - 1); i++ {
		switch {
		case w[i]%2 == 0:
			b = true
			fmt.Println("Четные число массива: ", w[i])

		case w[i]%2 != 0:
			b = false
			fmt.Println("Нечетные число массива: ", w[i])
		}

	}
	return b
}

func inputArray() [] int {
	fmt.Println("Давай сюда числа через пробел: ")
	q := utils.Scanner()
	s := strings.Split(q, " ")

	var w = []int{}

	for _, i := range s {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		w = append(w, j)
	}
	return w
}
