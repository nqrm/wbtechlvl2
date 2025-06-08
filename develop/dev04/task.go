package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Написать функцию поиска всех множеств анаграмм по словарю.

Например:
'пятак', 'пятка' и 'тяпка' — принадлежат одному множеству;
'листок', 'слиток' и 'столик' — другому.


"пятак" - ["пятка", "тяпка"]
"листок" - ["слиток", "столик"]

Требования
Входные данные для функции: ссылка на массив, каждый элемент которого — слово на русском языке в кодировке utf8.
Выходные данные: ссылка на мапу множеств анаграмм.
Ключ — первое встретившееся в словаре слово из множества. Значение — ссылка на массив, каждый элемент которого, слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
*/

func main() {
	//v := []string{"докер", "ЛИСТОК", "ГОланг", "каБАН", "декор", "пЯтак", "пяткА", "слиток", "тяпка", "банка", "столик", "КРедо"}
	v := []string{"докер", "докер", "докер"}
	m := GetAnagrams(v)
	fmt.Println(m)

	for k, v := range m {
		fmt.Printf("%v:\t%v\n", k, v)
	}
}

func getUnique(words []string) []string {
	m := make(map[string]bool)
	result := make([]string, 0, len(words))
	for _, word := range words {
		if m[word] {
			continue
		}
		m[word] = true
		result = append(result, word)
	}

	return result
}

func GetAnagrams(dict []string) map[string][]string {
	// "пятак", "пятка", "тяпка", "листок", "слиток", "столик"

	dict = getUnique(dict) // если все слова будут одинаковыми
	anagrams := make(map[string][]string)
	sortedWords := make(map[string][]string)

	for i := range dict {
		dict[i] = strings.ToLower(dict[i])
		sourceArray := []rune(dict[i])
		sort.SliceStable(sourceArray, func(i, j int) bool {
			return sourceArray[i] < sourceArray[j]
		})
		if _, ok := sortedWords[string(sourceArray)]; ok {
			sortedWords[string(sourceArray)] = append(sortedWords[string(sourceArray)], dict[i])
		} else {
			sortedWords[string(sourceArray)] = []string{dict[i]}
		}
	}
	for _, v := range sortedWords {
		if len(v) > 2 {
			anagrams[v[0]] = v[1:]
		}
	}

	for _, v := range anagrams {
		sort.Strings(v)
	}

	return anagrams
}
