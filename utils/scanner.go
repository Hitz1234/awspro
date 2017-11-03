package utils

import (
	"bufio"
	"os"
)

//FIXME Функции должны называтся глаголами
func Scanner() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
