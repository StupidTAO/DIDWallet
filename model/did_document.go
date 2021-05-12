package model

import (
	"encoding/json"
	"DIDWallet/utils"
)

type PublicKey struct {
	Id string			`json: "id"`
	Type string			`json: "type"`
	PublicKeyHex string	`json: "publicKeyHex"`
}

type DIDService struct {
	Id string				`json: "id"`
	Type string				`json: "type"`
	serviceEndpoint string	`json: "serviceEndpoint"`
}

type BaseDIDDoc struct {
	Context string			`json: "@context"`
	PublicKey PublicKey 	`json: "publicKey"`
	Authentication string	`json: "authentication"`
}

type DIDDoc struct {
	Context string			`json: "@context"`
	Id string				`json: "id"`
	Version uint			`json: "version"`
	CreatedTime string		`json: "created"`
	UpdatedTime string		`json: "updated"`
	PublicKey PublicKey 	`json: "publicKey"`
	Authentication string	`json: "authentication"`
	Service	DIDService		`json: "service"`
}

func BaseDIDDocUnmarshal(BaseDIDDocStr string, BaseDIDDocEntry * BaseDIDDoc) error {
	return json.Unmarshal([]byte(BaseDIDDocStr), &BaseDIDDocEntry)
}

func BaseDIDDocMarshal(BaseDIDDocEntry BaseDIDDoc) ([]byte, error) {
	return json.Marshal(BaseDIDDocEntry)
}

func GetSpecificId(BaseDIDDocStr string) (string, error) {
	tmpBytes := utils.GetSHA256HashCode([]byte(BaseDIDDocStr))
	tmpBytes = utils.GetRipemd160HashCode(tmpBytes)
	specificId := utils.Base58Encode(tmpBytes)
	return specificId, nil
}

func GetDIDByPrivateFile(filename string) (string, error) {
	//pri, err := utils.GetPrivateKeyByFile("/Users/oker/go/src/github.com/DIDIssuer/private_key")
	pri, err := utils.GetPrivateKeyByFile(filename)
	publicKeyHex, err := utils.GetPublicKeyHexByPrivateKey(pri)
	if err != nil {
		return "", err
	}

	publicKey := PublicKey{
		"#keys-1",
		"ecdsa",
		publicKeyHex,
	}

	baseDoc := BaseDIDDoc{"https://w3id.org/did/v1",publicKey, "#key-1"}
	strByte, err := BaseDIDDocMarshal(baseDoc)
	if err != nil {
		return "", err
	}

	specificId, _ := GetSpecificId(string(strByte))
	did := "did:" + "welfare:" + specificId
	return did, nil
}
