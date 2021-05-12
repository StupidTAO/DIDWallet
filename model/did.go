package model

import (
	"encoding/json"
)

type DID struct {
	Prefixes string		`json: "prefixes"`
	MethodName string	`json: "methodName"`
	SpecificId string	`json: "specificId"`
}

func DIDUnmarshal(DIDStr string, DIDEntry * DID) error {
	return json.Unmarshal([]byte(DIDStr), &DIDEntry)
}

func DIDMarshal(DIDEntry DID) ([]byte, error) {
	return json.Marshal(DIDEntry)
}

