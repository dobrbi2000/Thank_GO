package main

import (
	"fmt"
)

// element - интерфейс элемента последовательности
type element interface{}

// weightFunc - функция, которая возвращает вес элемента
type weightFunc func(element) int

// iterator - интерфейс, который умеет
// поэлементно перебирать последовательность
type iterator interface {
	next() bool
	val() element
	// чтобы понять сигнатуры методов - посмотрите,
	// как они используются в функции max() ниже
}

// intIterator - итератор по целым числам
// (реализует интерфейс iterator)
type intIterator struct {
	src   []int
	index int

	// поля структуры
}

// методы intIterator, которые реализуют интерфейс iterator
func (it *intIterator) next() bool {
	if it.index < len(it.src) {
		return true
	}
	return false
}

func (it *intIterator) val() element {

	if it.next() {
		el := it.src[it.index]
		it.index++
		return el
	}
	return nil
}

// конструктор intIterator
func newIntIterator(src []int) *intIterator {
	return &intIterator{
		src: src,
	}

}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// main находит максимальное число из переданных на вход программы.
func main() {
	nums := readInput()
	it := newIntIterator(nums)
	weight := func(el element) int {
		return el.(int)
	}
	m := max(it, weight)
	fmt.Println(m)
}

// max возвращает максимальный элемент в последовательности.
// Для сравнения элементов используется вес, который возвращает
// функция weight.
func max(it iterator, weight weightFunc) element {
	var maxEl element
	for it.next() {
		curr := it.val()
		if maxEl == nil || weight(curr) > weight(maxEl) {
			maxEl = curr
		}
	}
	return maxEl
}

// readInput считывает последовательность целых чисел из os.Stdin.
func readInput() []int {
	nums := []int{1, 2, 4, 7, 3, 9}
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Split(bufio.ScanWords)
	// for scanner.Scan() {
	// 	num, err := strconv.Atoi(scanner.Text())
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	nums = append(nums, num)
	// }
	return nums
}
