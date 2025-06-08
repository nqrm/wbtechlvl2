package main

import (
	"reflect"
	"testing"
)

func TestGetAnagrams(t *testing.T) {
	testCases := []struct {
		input    []string
		expected map[string][]string
	}{
		{
			[]string{"докер", "ЛИСТОК", "ГОланг", "каБАН", "декор", "пЯтак", "пяткА", "слиток", "тяпка", "банка", "столик", "КРедо"},
			map[string][]string{
				"докер":  {"декор", "кредо"},
				"листок": {"слиток", "столик"},
				"пятак":  {"пятка", "тяпка"},
			},
		},
		{
			[]string{"докер"},
			map[string][]string{},
		},
		{
			[]string{""},
			map[string][]string{},
		},
		{
			[]string{"докер", "докер", "докер"},
			map[string][]string{},
		},
	}

	for _, testCase := range testCases {
		result := GetAnagrams(testCase.input)
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Incorrect result. Expect %s, got %s", testCase.expected, result)
		}
	}

}
