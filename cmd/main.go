package main

import (
	"fmt"

	//"github.com/sergeidovydenko/num2word"
	"github.com/Snezhana28/study2025/cmd/internal/containers"
)

//func main() {
//	var a int
//	n, err := fm.Scan(&a)
//	if err != nil || n == 0 {
//		fm.Println("Ошибка ввода: введите целое число")
//		return
//	}
//	if a >= 12307 {
//		fm.Printf("Введенное число %d должно быть меньше 12307\n", a)
//		return
//	}
//	if a < 0 {
//		a = -a
//	}
//
//	if a%13 == 0 && a%9 == 0 {
//		fm.Println("service error")
//		return
//	}
//	switch {
//	case a < 0:
//
//		a = a * -1
//	case a%7 == 0:
//
//		a = a * 39
//	case a%9 == 0:
//		a = a*13 + 1
//	default:
//
//		a = (a + 2) * 3
//	}
//	a = a + 1
//
//	str := num2word.Convert(a, false)
//
//	fm.Printf("Итоговое число: %d\n", a)
//
//	fm.Printf("Итоговое число прописью: %s\n", str)
//
//}

func main() {
	stack := containers.NewStack()

	fmt.Println("Размер:", stack.Size())
	stack.Push(1)
	fmt.Println("Размер после добавления:", stack.Size())
	stack.Pop()
	fmt.Println("Стек пустой?", stack.IsEmpty())
	stack.Push(15)
	fmt.Println("Cтек пустой сейчас?", stack.IsEmpty())
	stack.Pop()

	for i := 0; i < 10000000; i++ {
		stack.Push(i)
	}
	fmt.Println("Длина стека: ", len(stack.Values))
	topElem, _ := stack.Pop()
	fmt.Println("Верхний элемент стека: ", topElem)
	fmt.Println("Длина стека после удаления одного элемента: ", len(stack.Values))
	fmt.Println("Размер стека до очистки:", stack.Size())
	stack.Clear()
	fmt.Println("Размер стека после очистки: ", stack.Size())

	deque := containers.NewDeque()
	deque.PushBack(31)
	deque.PushBack(43)
	deque.PushBack(28)
	fmt.Println("После добавления элементов: ", deque.Values)
	frontElement, ok := deque.PopFront()
	if ok {
		fmt.Println("Добавление вперед: ", frontElement)
	}
	rearElement, ok := deque.PopBack()
	if ok {
		fmt.Println("Добавление в конец: ", rearElement)
	}
	fmt.Println("После извлечения элементов: ", deque.Values)
	fmt.Println("Размер:", deque.Size())
	deque.PushFront(3710)
	fmt.Println("Размер после добавления:", deque.Size())
	fmt.Println("В вашей очереди следующие значения>: ", deque.Values)
	deque.Clear()
	fmt.Println("Размер очереди после очистки: ", deque.Size())
	fmt.Println("Пустая очередь: ", deque.IsEmpty())v
}
