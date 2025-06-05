package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"testing"
)

func TestGetColumn(t *testing.T) {
	testCases := []struct {
		inputStr string
		k        int
		expected string
	}{
		{
			"Отсортировать строки в файле по аналогии с консольной утилитой sort",
			4,
			"файле",
		},
		{
			"Отсортировать",
			2,
			"",
		},
		{
			"",
			1,
			"",
		},
	}

	for _, testCase := range testCases {
		result := getColumn(testCase.inputStr, testCase.k)
		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %s, got %s", testCase.expected, result)
		}
	}

}

func TestGetUnique(t *testing.T) {
	testCases := []struct {
		inputLines []string
		expected   []string
	}{
		{
			[]string{"aaa", "bb", "222", "bb", "cc", "22", "cc"},
			[]string{"aaa", "bb", "222", "cc", "22"},
		},
		{
			[]string{"1", "2", "3", "4", "5", "6", "7"},
			[]string{"1", "2", "3", "4", "5", "6", "7"},
		},
		{
			[]string{"1", "1", "1", "1", "1", "1", "1"},
			[]string{"1"},
		},
		{
			[]string{"", "", "", "", "", "", ""},
			[]string{""},
		},
	}

	for _, testCase := range testCases {
		result := getUnique(testCase.inputLines)
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect %s, got %s", testCase.expected, result)
		}
	}
}

func TestSort(t *testing.T) {
	testCmdArgs := [][]string{
		{"run", "task.go", "-k", "9", "-f", "textFile.txt"},
		{"run", "task.go", "-k", "5", "-n", "-f", "textFile.txt"},
		{"run", "task.go", "-u", "-r", "-f", "textFile.txt"},
	}
	for i := range testCmdArgs {
		cmd := exec.Command("go", testCmdArgs[i]...)
		cmd.CombinedOutput()

		result, _ := os.ReadFile("sortedFile.txt")
		expected, _ := os.ReadFile(fmt.Sprintf("expected_%v.txt", i))

		if !bytes.Equal(result, expected) {
			t.Errorf("Incorrect result.\nExpect: %s\nGot: %s", expected, result)
		}
	}
}
