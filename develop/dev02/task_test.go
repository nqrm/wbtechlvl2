package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnpackingString(t *testing.T) {
	testTable := []struct {
		inputString    string
		expectedString string
		expectedError  error
	}{
		{
			inputString:    "a4bc2d5e",
			expectedString: "aaaabccddddde",
			expectedError:  nil,
		},
		{
			inputString:    "abcd",
			expectedString: "abcd",
			expectedError:  nil,
		},
		{
			inputString:    "45",
			expectedString: "",
			expectedError:  fmt.Errorf("Error: Before the digit at index 0, there must be a character (not digit)."),
		},
		{
			inputString:    "",
			expectedString: "",
			expectedError:  nil,
		},
		{
			inputString:    `qwe\4\5`,
			expectedString: "qwe45",
			expectedError:  nil,
		},
		{
			inputString:    `qwe\45`,
			expectedString: "qwe44444",
			expectedError:  nil,
		},
		{
			inputString:    `qwe\\5`,
			expectedString: `qwe\\\\\`,
			expectedError:  nil,
		},
		{
			inputString:    `\\\\`,
			expectedString: `\\`,
			expectedError:  nil,
		},
		{
			inputString:    `\\q2\31we\\5\\\\\2`,
			expectedString: `\qq3we\\\\\\\2`,
			expectedError:  nil,
		},
		{
			inputString:    `\`,
			expectedString: ``,
			expectedError:  fmt.Errorf("Error: Аfter the backslash there must be either a digit or a backslash"),
		},
		{
			inputString:    `\a`,
			expectedString: ``,
			expectedError:  fmt.Errorf("Error: Unable to escape character at index 1 (a)"),
		},
	}

	for _, testCase := range testTable {
		result, err := UnpackingString(testCase.inputString)
		if result != testCase.expectedString {
			t.Errorf("Incorrect result. Expect %s, got %s", testCase.expectedString, result)
		}
		if !reflect.DeepEqual(err, testCase.expectedError) {
			t.Errorf("Incorrect error. Expect %s, got %s", testCase.expectedError, err)
		}
	}

}
