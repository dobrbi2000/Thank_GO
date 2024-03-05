/* Футбольный турнир
Четыре команды (A, B, C и D) сыграли футбольный турнир. Каждая команда сыграла с каждой по одному разу, так что всего прошло 6 матчей. Рассчитайте, сколько очков получила каждая команда.

Победа приносит 3 очка, ничья — 1, поражение — 0.

Результаты игр подаются на вход программы строкой вида:

ABW DCD DAW CBL BDL ACW
Каждая тройка букв означает одну игру. Внутри тройки первая буква — название первой команды, вторая буква — название второй, третья — код результата (W — первая выиграла, L — первая проиграла, D — ничья).

Код, который считывает и парсит результаты игр, уже написан. Ваша задача — реализовать структуры данных и логику подсчета результата.

Sample Input:

ABW DCD DAW CBL BDL ACW
Sample Output:

A6 B3 C1 D7
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

// result представляет результат матча
type result byte

// возможные результаты матча
const (
	win  result = 'W' //победа - 3
	draw result = 'D' //ничья - 1
	loss result = 'L' //поражение - 0
)

// team представляет команду
type team byte

// match представляет матч
// содержит три поля:
// - first (первая команда)
// - second (вторая команда)
// - result (результат матча)
// например, строке BAW соответствует
// first=B, second=A, result=W
type match struct {
	first  team
	second team
	result result
}

// rating представляет турнирный рейтинг команд -
// количество набранных очков по каждой команде
type rating map[team]int

// tournament представляет турнир
type tournament []match

// calcRating считает и возвращает рейтинг турнира
func (trn *tournament) calcRating() rating {
	rating := make(rating)

	for _, match := range *trn {
		switch match.result {
		case win:
			rating[match.first] += 3
		case draw:
			rating[match.first] += 1
			rating[match.second] += 1
		case loss:
			rating[match.second] += 3
		}

	}
	return rating

}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// код, который парсит результаты игр, уже реализован
// код, который печатает рейтинг, уже реализован
// ваша задача - реализовать недостающие структуры и методы выше
func main() {
	src := readString()
	trn := parseTournament(src)
	rt := trn.calcRating()
	rt.print()
}

// readString считывает и возвращает строку из os.Stdin
func readString() string {

	str := "ABW DCD DAW CBL BDL ACW"
	// rdr := bufio.NewReader(os.Stdin)
	// str, err := rdr.ReadString('\n')
	// if err != nil && err != io.EOF {
	// 	log.Fatal(err)
	// }
	return str
}

// parseTournament парсит турнир из исходной строки
func parseTournament(s string) tournament {
	descriptions := strings.Split(s, " ")
	trn := tournament{}
	for _, descr := range descriptions {
		m := parseMatch(descr)
		trn.addMatch(m)
	}
	return trn
}

// parseMatch парсит матч из фрагмента исходной строки
func parseMatch(s string) match {
	return match{
		first:  team(s[0]),
		second: team(s[1]),
		result: result(s[2]),
	}
}

// addMatch добавляет матч к турниру
func (t *tournament) addMatch(m match) {
	*t = append(*t, m)
}

// print печатает результаты турнира
// в формате Aw Bx Cy Dz
func (r *rating) print() {
	var parts []string
	for team, score := range *r {
		part := fmt.Sprintf("%c%d", team, score)
		parts = append(parts, part)
	}
	sort.Strings(parts)
	fmt.Println(strings.Join(parts, " "))
}
