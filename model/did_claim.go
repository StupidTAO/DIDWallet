package model

import "encoding/json"

const RealNameAuthentication = "RealNameAuthentication"
const IDCardAuthentication = "IDCardAuthentication"
const FingerprintAuthentication = "FingerprintAuthentication"
const EnterpriseAuthentication = "EnterpriseAuthentication"
const BusinessAuthentication = "BusinessAuthentication"
const VIPAuthentication =  "VIPAuthentication"

type CredentialSubject struct {
	DID string					`json:"id"`	//被签发方的ID
	ShortDescription string		`json:"shortDescription"`
	LongDescription	string		`json:"longDescription"`
	TypeCliam string			`json:"typeClaim"`
}

func CredentialSubjectUnmarshal(CredentialSubjectStr string, CredentialSubjectEntry * CredentialSubject) error {
	return json.Unmarshal([]byte(CredentialSubjectStr), &CredentialSubjectEntry)
}

func CredentialSubjectMarshal(CredentialSubjectEntry CredentialSubject) ([]byte, error) {
	return json.Marshal(CredentialSubjectEntry)
}

func GetCredentialSubject(did string, idNumber string) (CredentialSubject, error) {
	cs := CredentialSubject{did, idNumber, "ID Card", IDCardAuthentication}
	return cs, nil
}
