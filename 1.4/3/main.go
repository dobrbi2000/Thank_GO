/*
Напишите функцию filter(), которая фильтрует срез целых чисел с помощью функции-предиката и возвращает отфильтрованный срез.
Функция-предикат вызывается для каждого элемента исходного среза. Если она возвращает true, элемент попадает в отфильтрованный срез.
Если возвращает false — не попадает.

Считайте исходный срез из стандартного ввода с помощью готовой функции readInput(). Затем выполните на нем filter().
В качестве предиката используйте функцию, которая возвращает true только для четных чисел. Напечатайте отфильтрованный срез.

Гарантируется, что на вход подаются только целые числа.

Sample Input:
1 2 3 4 5 6

Sample Output:
[2 4 6]
*/
package main

import (
	"fmt"
)

func filter(predicate func(int) bool, iterable []int) []int {
	var result []int
	for _, num := range iterable {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result

}

func main() {
	src := readInput()

	res := filter(func(num int) bool { return num%2 == 0 }, src)

	fmt.Println(res)
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
func readInput() []int {
	nums := []int{1, 2, 3, 4, 5}
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Split(bufio.ScanWords)
	// for scanner.Scan() {
	// 	num, _ := strconv.Atoi(scanner.Text())
	// 	nums = append(nums, num)
	// }
	return nums
}
