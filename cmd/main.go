package main

import (
	"fmt"
	"github.com/sergeidovydenko/num2word"
)

func main() {
	var a int
	n, err := fmt.Scan(&a)
	if err != nil || n == 0 {
		fmt.Println("Ошибка ввода: введите целое число")
		return
	}
	if a >= 12307 {
		fmt.Printf("Введенное число %d должно быть меньше 12307\n", a)
		return
	}
	if a < 0 {
		a = -a
	}

	if a%13 == 0 && a%9 == 0 {
		fmt.Println("service error")
		return
	}
	switch {
	case a < 0:

		a = a * -1
	case a%7 == 0:

		a = a * 39
	case a%9 == 0:
		a = a*13 + 1
	default:

		a = (a + 2) * 3
	}
	a = a + 1

	str := num2word.Convert(a, false)

	fmt.Printf("Итоговое число: %d\n", a)

	fmt.Printf("Итоговое число прописью: %s\n", str)

}
