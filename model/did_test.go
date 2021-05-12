package model

import (
	"fmt"
	"testing"
)

func TestDIDMarshal(t *testing.T) {
	specificId := "2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	did := DID{"did", "welfare", specificId}
	strByte, err := DIDMarshal(did)
	if err != nil {
		t.Error("DIDMarshal Failed")
	}
	fmt.Println(string(strByte))
	t.Log("DIDMarshal Passed")
}
func TestDIDUnmarshal(t *testing.T) {
	specificId := "2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	did := DID{"did", "welfare", specificId}
	strByte, err := DIDMarshal(did)
	if err != nil {
		t.Error("DIDMarshal Failed")
	}

	DIDEntry := new(DID)
	err = DIDUnmarshal(string(strByte), DIDEntry)
	if err != nil {
		t.Error("DIDUnmarshal Failed")
	}
	t.Log("DIDUnmarshal Passed")
}
