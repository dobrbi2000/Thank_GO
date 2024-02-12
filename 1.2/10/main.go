/*
Напишите программу, которая определяет название языка по его коду. Правила:

en → English
fr → French
ru или rus → Russian
иначе → Unknown

Sample Input:
en

Sample Output:
English
*/
package main

import (
	"fmt"
)

func main() {
	var code string
	fmt.Scan(&code)
	var lang string
	// определите полное название языка по его коду
	// и запишите его в переменную `lang`
	// ...
	switch code {
	case "en":
		lang = "English"
	case "fr":
		lang = "French"
	case "ru", "rus":
		lang = "Russian"
	default:
		lang = "Unknown"
	}

	fmt.Println(lang)
}

// package main

// import "fmt"

// func main() {
// 	var arr [2][3]int
// 	for i := 0; i < 2; i++ { // 0 1
// 		for j := 0; j < 3; j++ { // 0 1 2
// 			arr[i][j] = i + j
// 		}
// 	}
// 	fmt.Println("2d:", arr)
// 	// 2d: [[0 1 2] [1 2 3]]

// }
