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
	case "ru":
		lang = "Russian"
	case "rus":
		lang = "Russian"
	default:
		lang = "Unknown"
	}

	fmt.Println(lang)
}
