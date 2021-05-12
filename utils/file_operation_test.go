package utils

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	filename := ".wallet_test"
	content, err := ReadFile(filename)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(content)
	t.Log("ReadFile PASSED")
}

func TestAppendToFile(t *testing.T) {
	filename := ".wallet_test"
	AppendToFile(filename, "test")
}
