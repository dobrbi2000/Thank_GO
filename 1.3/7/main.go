package main

/*
Аббревиатура
Напишите программу, которая принимает на вход фразу и составляет аббревиатуру по первым буквам слов:

Today I learned → TIL
Высшее учебное заведение → ВУЗ
Кот обладает талантом → КОТ
Если слово начинается не с буквы, игнорируйте его:

Ар 2 Ди #2 → АД
Разделителями слов считаются только пробельные символы. Дефис, дробь и прочие можно не учитывать:

Анна-Мария Волхонская → АВ

Sample Input:
Today I learned

Sample Output:
TIL



*/
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	//phrase := readString()
	phrase := "I 3earned today Good"
	words := strings.Fields(phrase)
	var abbr []rune

	for _, word := range words {
		letter := []rune(word)[0]
		if !unicode.IsLetter(letter) {
			continue
		}
		abbr = append(abbr, unicode.ToUpper(letter))

	}

	// for _, word := range words {
	// 	if len(word) > 0 && unicode.IsLetter([]rune(word)[0]) {
	// 		firstLetter := []rune(word)[0]
	// 		upperCaseFirstLetter := unicode.ToUpper(firstLetter)
	// 		abbr += string(upperCaseFirstLetter)
	// 	}

	// }

	// 1. Разбейте фразу на слова, используя `strings.Fields()`
	// 2. Возьмите первую букву каждого слова и приведите
	//    ее к верхнему регистру через `unicode.ToUpper()`
	// 3. Если слово начинается не с буквы, игнорируйте его
	//    проверяйте через `unicode.IsLetter()`
	// 4. Составьте слово из получившихся букв и запишите его
	//    в переменную `abbr`
	// ...

	fmt.Println(string(abbr))
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
