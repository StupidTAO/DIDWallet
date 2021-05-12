package model

import (
	"testing"
)

func TestCredentialSubjectMarshal(t *testing.T) {
	credentialSubjectEntry := new(CredentialSubject)
	credentialSubjectEntry.DID = "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	credentialSubjectEntry.ShortDescription = "342225199509082432"
	credentialSubjectEntry.LongDescription = "身份证号"
	credentialSubjectEntry.TypeCliam = "IDCardAuthentication"

	bs, err := CredentialSubjectMarshal(*credentialSubjectEntry)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(string(bs))
}

func TestCredentialSubjectUnmarshal(t *testing.T) {
	credentialSubjectEntry := new(CredentialSubject)
	credentialSubjectEntry.DID = "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	credentialSubjectEntry.ShortDescription = "342225199509082432"
	credentialSubjectEntry.LongDescription = "身份证号"
	credentialSubjectEntry.TypeCliam = "IDCardAuthentication"

	bs, err := CredentialSubjectMarshal(*credentialSubjectEntry)
	if err != nil {
		t.Error(err.Error())
		return
	}

	credentialSubjectNewEntry := new(CredentialSubject)
	err = CredentialSubjectUnmarshal(string(bs), credentialSubjectNewEntry)
	if err != nil {
		t.Error(err.Error())
		return
	}
}
