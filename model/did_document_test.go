package model

import (
	"fmt"
	"testing"
)

func TestBaseDIDDocMarshal(t *testing.T) {
	publicKey := PublicKey{
		"#keys-1",
		"Secp256k1",
		"02b97c30de767f084ce3080168ee293053ba33b235d7116a3263d29f1450936b71",
	}

	baseDoc := BaseDIDDoc{"https://w3id.org/did/v1",publicKey, "#key-1"}
	strByte, err := BaseDIDDocMarshal(baseDoc)
	if err != nil {
		t.Error("BaseDIDDocMarshal failed")
	}
	fmt.Println(string(strByte))
	t.Log("BaseDIDDocMarshal pass")
}

func TestBaseDIDDocUnmarshal(t *testing.T) {
	publicKey := PublicKey{
		"#keys-1",
		"Secp256k1",
		"02b97c30de767f084ce3080168ee293053ba33b235d7116a3263d29f1450936b71",
	}

	baseDoc := BaseDIDDoc{"https://w3id.org/did/v1",publicKey, "#key-1"}
	strByte, err := BaseDIDDocMarshal(baseDoc)
	if err != nil {
		t.Error("BaseDIDDocUnmarshal failed")
	}

	BaseDIDDocEntry := new(BaseDIDDoc)
	err = BaseDIDDocUnmarshal(string(strByte), BaseDIDDocEntry)
	if err != nil {
		t.Error("BaseDIDDocUnmarsha failed")
	}
	t.Log("BaseDIDDocMarshal pass")
}

func TestGetSpecificId(t *testing.T) {
	publicKey := PublicKey{
		"#keys-1",
		"Secp256k1",
		"02b97c30de767f084ce3080168ee293053ba33b235d7116a3263d29f1450936b71",
	}

	baseDoc := BaseDIDDoc{"https://w3id.org/did/v1",publicKey, "#key-1"}
	strByte, err := BaseDIDDocMarshal(baseDoc)
	if err != nil {
		t.Error("BaseDIDDocUnmarshal failed")
	}

	specificId, _ := GetSpecificId(string(strByte))
	fmt.Printf("%s", specificId)
	t.Log("GetSpecificId pass")
}

func TestGetDIDByPrivateFile(t *testing.T) {
	filename := "/Users/oker/go/src/github.com/DIDWallet/cmd/private_key"
	did, err := GetDIDByPrivateFile(filename)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%s\n", did)
}
