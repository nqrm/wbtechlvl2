package main

import "testing"

func TestNonExistentAddress(t *testing.T) {
	ntpServer = "123"
	err := getCurrentTime()
	if err == nil {
		t.Errorf("Should not produce an error")
	}

}
