package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры):
на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:
-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел);
-n — сортировать по числовому значению;
-r — сортировать в обратном порядке;
-u — не выводить повторяющиеся строки.

Дополнительно
Реализовать поддержку утилитой следующих ключей:
-M — сортировать по названию месяца;
-b — игнорировать хвостовые пробелы;
-c — проверять отсортированы ли данные;
-h — сортировать по числовому значению с учетом суффиксов.
*/

func main() {
	var kFlag = flag.Int("k", 1, "Specify a sort field that consists of the part of the line")
	var nFlag = flag.Bool("n", false, "Sort numerically")
	var rFlag = flag.Bool("r", false, "Reverse the result of comparison, so that lines with greater key values appear earlier in the output instead of later.")
	var uFlag = flag.Bool("u", false, "Output only the first of a sequence of lines that compare equal")
	var filePathFlag = flag.String("f", "", "File path")
	flag.Parse()

	if *filePathFlag == "" {
		fmt.Println("Specify the filepath using the flag -f")
		flag.Usage()
		return
	}

	file, err := os.Open(*filePathFlag)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("File reading error:", err)
		return
	}

	if *uFlag {
		lines = getUnique(lines)
	}

	compare := func(i, j int) bool {
		strI := getColumn(lines[i], *kFlag)
		strJ := getColumn(lines[j], *kFlag)
		if *nFlag {
			numI, err1 := strconv.Atoi(strI)
			numJ, err2 := strconv.Atoi(strJ)
			if err1 != nil || err2 != nil {
				if *rFlag {
					return strI > strJ
				}
				return strI < strJ
			}

			if *rFlag {
				return numI > numJ
			}

			return numI < numJ

		}

		if *rFlag {
			return strI > strJ
		}
		return strI < strJ
	}
	sort.SliceStable(lines, compare)

	outputFileName := "sortedFile.txt"
	outputFile, err := os.Create(outputFileName)
	writer := bufio.NewWriter(outputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()

	for i, row := range lines {
		writer.WriteString(row)
		if i == len(lines)-1 {
			break
		}
		writer.WriteString("\r\n")
	}
	writer.Flush()
}

func getColumn(line string, k int) string {
	words := strings.Fields(line)
	if k > len(words) {
		return ""
	}
	return words[k-1]
}

func getUnique(lines []string) []string {
	m := make(map[string]bool)
	result := make([]string, 0, len(lines))
	for _, line := range lines {
		if m[line] {
			continue
		}
		m[line] = true
		result = append(result, line)
	}

	return result

}
