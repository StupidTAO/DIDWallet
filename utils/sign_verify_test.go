package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestAll(t *testing.T) {

	//随机熵，用于加密安全
	randSign := "20180619zafes20180619zafes20180619zafessss"//至少36位
	//随机key，用于创建公钥和私钥
	randKey := "fb0f7279c18d4394594fc9714797c9680335a320"
	//创建公钥和私钥
	prk, puk, err := GetEcdsaKey(randKey)
	if err != nil {
		fmt.Println(err)
	}

	//hash加密使用md5用到的salt
	salt := "131ilzaw"
	//待加密的明文
	text := "hlloaefaefaefaefaefaefaefhelloaefaefaefaefaefaefaefhelloaefaefaefaefaefaefaef"
	//text1 := "hlloaefaefaefaefaefaefaefhelloaefaefaefaefaefaefaefhelloaefaefaefaefaefaefaef1"

	//hash取值
	htext := hashtext(text,salt)
	//htext1 := hashtext(text1,salt)
	//hash值编码输出
	fmt.Println(hex.EncodeToString(htext))

	//hash值进行签名
	result, err := sign(htext,randSign,prk)
	if err != nil {
		fmt.Println(err)
	}
	//签名输出
	fmt.Println(result)

	//签名与hash值进行校验
	tmp, err := verify(htext,result,puk)
	fmt.Println(tmp)
}

func TestGetEcdsaKey(t *testing.T) {
	content, err := ReadFile("seed")
	if err != nil {
		t.Error(err.Error())
		return
	}
	_, _, err = GetEcdsaKey(content)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestSign(t *testing.T) {
	text := "caohaitao"
	textBytes := GetSHA256HashCode([]byte(text))

	//随机熵要大于36个字符
	randomSign := "20180619zafes20180619zafes20180619zafessss"

	content, err := ReadFile("seed")

	if err != nil {
		t.Error(err.Error())
		return
	}
	prk, _, err := GetEcdsaKey(content)
	if err != nil {
		t.Error(err.Error())
		return
	}
	signResult, err := sign(textBytes, randomSign, prk)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(signResult)
}

func TestVerify (t *testing.T) {
	text := "caohaitao"
	textBytes := GetSHA256HashCode([]byte(text))

	//随机熵要大于36个字符
	randomSign := "20180619zafes20180619zafes20180619zafessss"

	content, err := ReadFile("seed")

	if err != nil {
		t.Error(err.Error())
		return
	}
	prk, pub, err := GetEcdsaKey(content)
	if err != nil {
		t.Error(err.Error())
		return
	}

	signResult, err := sign(textBytes, randomSign, prk)
	if err != nil {
		t.Error(err.Error())
		return
	}

	//将验证签名结果
	ok, err := verify(textBytes, signResult, pub)
	if err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Println(ok)
}

func TestGenerateKey(t *testing.T) {
	_, err := GenerateKey()
	if err != nil {
		t.Error(err.Error())
		return
	}

}

func TestGetPrivateKeyByFile(t *testing.T) {
	_, err := GetPrivateKeyByFile("writePrk1")
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestWritePrivateKeyToFile(t *testing.T) {
	prv, _ :=  GenerateKey()
	err := WritePrivateKeyToFile("writePrk1", prv)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestGetPublicKeyHexByPrivateKey(t *testing.T) {
	prk, err := GenerateKey()
	if err != nil {
		t.Error(err.Error())
		return
	}

	pukAddr, err := GetPublicKeyHexByPrivateKey(prk)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(pukAddr)
}

func TestGetAddressByPublicKey(t *testing.T) {
	prk, err := GenerateKey()
	if err != nil {
		t.Error(err.Error())
		return
	}
	addr := GetAddressByPublicKey(prk.PublicKey)
	fmt.Println(addr)
}

func TestBytesToAddress(t *testing.T) {
	addr := BytesToAddress([]byte("caohaitao"))
	fmt.Println(addr)
}

func TestBytesToHash(t *testing.T) {
	hashStr := BytesToHash([]byte("caohaitao"))
	fmt.Println(hashStr)
}

func TestSignText(t *testing.T) {
	prk, err := GenerateKey()
	if err != nil {
		t.Error(err.Error())
		return
	}

	result, err := SignText("caohaitao", prk)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%v", result)
}

func TestVerifyToAddress(t *testing.T) {
	//prk, err := GenerateKey()
	prk, err := GetPrivateKeyByFile("private_key")
	if err != nil {
		t.Error(err.Error())
		return
	}

	text := "caohaitao"
	result, err := SignText(text, prk)
	if err != nil {
		t.Error(err.Error())
		return
	}

	addr, err := VerifyToAddress(text, result)
	if err != nil {
		t.Error(err.Error())
		return
	}
	addrRaw := GetAddressByPublicKey(prk.PublicKey)
	fmt.Println(addrRaw)
	fmt.Println(addr)
}

func TestGenerateKey1(t *testing.T) {
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("failed GenerateKey with %s.", err)
	}

	fmt.Println("private key have", hexutil.Encode(crypto.FromECDSA(key)))
	fmt.Println("private key no 0x", hex.EncodeToString(crypto.FromECDSA(key)))

	if err := crypto.SaveECDSA("privatekey", key); err != nil {
		t.Error(fmt.Sprintf("Failed to persist node key: %v", err))
	}

	fmt.Println("public key have 0x   n", hexutil.Encode(crypto.FromECDSAPub(&key.PublicKey)))
	fmt.Println("public key no 0x n", hex.EncodeToString(crypto.FromECDSAPub(&key.PublicKey)))

	//由私钥字符串转换私钥
	acc1Key, _ := crypto.HexToECDSA("8a1f9a8f95be41cd7ccb6168179afb4504aefe388d1e14474d32c45c72ce7b7a")
	address1 := crypto.PubkeyToAddress(acc1Key.PublicKey)
	fmt.Println("address ", address1.String())

	dummyAddr := common.HexToAddress("9b2055d370f73ec7d8a03e965129118dc8f5bf83")
	fmt.Println("dummyAddr",dummyAddr.String())

	//字节转地址
	addr3      := common.BytesToAddress([]byte("ethereum"))
	fmt.Println("address ",addr3.String())

	//字节转hash
	hash1 := common.BytesToHash([]byte("topic1"))
	fmt.Println("hash ",hash1.String())


	var testAddrHex = "970e8128ab834e8eac17ab8e3812f010678cf791"
	var testPrivHex = "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"
	key1, _ := crypto.HexToECDSA(testPrivHex)
	addrtest := common.HexToAddress(testAddrHex)

	msg := crypto.Keccak256([]byte("foo"))
	sig, err := crypto.Sign(msg, key1)
	recoveredPub, err := crypto.Ecrecover(msg, sig)
	pubKey, _ := crypto.UnmarshalPubkey(recoveredPub)
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	// should be equal to SigToPub
	recoveredPub2, _ := crypto.SigToPub(msg, sig)
	recoveredAddr2 := crypto.PubkeyToAddress(*recoveredPub2)

	fmt.Println("addrtest ",addrtest.String())
	fmt.Println("recoveredAddr ",recoveredAddr.String())
	fmt.Println("recoveredAddr2 ",recoveredAddr2.String())
}
