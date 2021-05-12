package utils

import (
	"fmt"
	"testing"
)

func TestGetContribution(t *testing.T) {
	index, err := GetContribution(1,1, 1,1,2)
	if err != nil {
		t.Error("contribution caculate is error: ", err.Error())
	}
	fmt.Println("index is: ", index)
}
