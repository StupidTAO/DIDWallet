package main

import (
	"DIDWallet/model"
	"DIDWallet/utils"
	"fmt"
	"testing"
)

const DID_BASE_DOC = "did_base_document"
const key=`{"address":"12769c3419a7f491cf4e576e2e983e009d579076","crypto":{"cipher":"aes-128-ctr","ciphertext":"215430a18ab1132c6eaecdf966bc0d878a3be06cff5dce173d801afec5002db5","cipherparams":{"iv":"d41d87954da3dfca1f38e14111169fb8"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"d5268e70fbf8666435bf82ee53850f14486810f1944110de2aede933ae97fff1"},"mac":"a80e95fc657473f543bb989da9e8e2cde73e2b8bde52e6b05355b44a40c165ec"},"id":"5e034459-6e81-4208-9e26-c91641d20f5d","version":3}`


func TestReadFile(t *testing.T) {
	fileInfo, err := utils.ReadFile(".wallet")
	if err != nil {
		t.Error("read file ERROR", err)
	}
	fmt.Println(fileInfo)
	t.Log("read file PASSED")
}

func TestCreateDID(t *testing.T) {
	publicKey := model.PublicKey{
		"#keys-1",
		"Secp256k1",
		"12b97c30de767f084ce3080168ee293053ba33b235d7116a3263d29f1450936b71",
	}

	baseDoc := model.BaseDIDDoc{"https://w3id.org/did/v1",publicKey, "#key-1"}
	strByte, err := model.BaseDIDDocMarshal(baseDoc)
	if err != nil {
		t.Error("BaseDIDDocUnmarshal failed")
	}

	createDID(string(strByte))
	t.Log("create did PASSED")
}

func TestCreateDIDByFile(t *testing.T) {
	err := createDIDByFile(DID_BASE_DOC, UNSYNC)
	if err != nil {
		t.Error("error is ", err)
		return
	}
	t.Log("CreateDIDByFile PASSED")
}

func TestNet(t *testing.T) {
	/**urls := "http://localhost:8000/getForm?"
	req, _ := http.NewRequest("GET", urls, nil)
	//req.Header.Add("cache-control", "no-cache")
	//req.Header.Add("postman-token", "0bffa3e0-568e-b7c7-1a22-1c98081ff27b")
	//获取返回值
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))*/

}

func TestCreateDIDByPriFile(t *testing.T) {
	priFile := "private_key3"
	err := createDIDByPriFile(priFile)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestGetRequestClaimUrl (t *testing.T) {
	//获得参数
	prifile := "private_key"
	did, _ := model.GetDIDByPrivateFile(prifile)

	//生成rawclaim
	idNumber := "342225200108015678"
	cs, _ := model.GetCredentialSubject(did, idNumber)
	bs, _ := model.CredentialSubjectMarshal(cs)
	rawClaim := string(bs)

	//获取签名以及地址
	prikey, _ := utils.GetPrivateKeyByFile(prifile)
	rawClaimSig, _ := utils.SignText(rawClaim, prikey)
	addr := utils.GetAddressByPublicKey(prikey.PublicKey)

	//获得url
	url, err := getRequestClaimUrl(rawClaim, rawClaimSig, addr)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println("url is", url)

}

func TestGetUrl(t *testing.T) {
	//获得参数
	prifile := "private_key"
	did, _ := model.GetDIDByPrivateFile(prifile)

	//生成rawclaim
	idNumber := "342225200108015678"
	cs, _ := model.GetCredentialSubject(did, idNumber)
	bs, _ := model.CredentialSubjectMarshal(cs)
	rawClaim := string(bs)

	//获取签名以及地址
	prikey, _ := utils.GetPrivateKeyByFile(prifile)
	rawClaimSig, _ := utils.SignText(rawClaim, prikey)
	addr := utils.GetAddressByPublicKey(prikey.PublicKey)

	//获得url
	url, err := getRequestClaimUrl(rawClaim, rawClaimSig, addr)
	if err != nil {
		t.Error(err.Error())
		return
	}

	response, _ := getUrl(url)
	fmt.Println(url)
	fmt.Println(response)
}
