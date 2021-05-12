package utils

import (
	"fmt"
	"testing"
)

func TestGetSHA256HashCode(t *testing.T) {
	strByte := "peter"
	hashCode := GetSHA256HashCode([]byte(strByte))
	fmt.Printf("%x", hashCode)
	t.Log("GetSHA256HashCode pass")
}

func TestGetRipemd160HashCode(t *testing.T) {
	strByte := "peter"
	hashCode := GetRipemd160HashCode([]byte(strByte))
	fmt.Printf("%x", hashCode)
	t.Log("GetSHA256HashCode pass")
}

func TestBase58Encode(t *testing.T) {
	str := Base58Encode([]byte("peter"))
	fmt.Println(str)
	t.Log("Base58Encode pass")
}

func TestBase58Decode(t *testing.T)  {
	strByte := Base58Decode("DgUwp2V")
	fmt.Println(string(strByte))
	t.Log("Base58Decode pass")
}
