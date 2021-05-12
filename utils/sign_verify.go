package utils

import (
	"bytes"
	"compress/gzip"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strings"
)

/**
  通过一个随机key创建公钥和私钥
  随机key至少为36位
*/
func GetEcdsaKey(randKey string) (*ecdsa.PrivateKey, ecdsa.PublicKey, error){

	var err error

	var prk *ecdsa.PrivateKey
	var puk ecdsa.PublicKey
	var curve elliptic.Curve

	lenth := len(randKey)
	if lenth < 224/8 {
		err =errors.New("私钥长度太短，至少为36位！")
		return prk,puk,err
	}

	if lenth > 521/8 + 8 {
		curve = elliptic.P521()
	}else if lenth > 384/8 + 8 {
		curve = elliptic.P384()
	}else if lenth > 256/8 + 8 {
		curve = elliptic.P256()
	}else if lenth > 224/8 + 8 {
		curve = elliptic.P224()
	}

	prk, err = ecdsa.GenerateKey(curve,strings.NewReader(randKey))
	if err != nil {
		return prk, puk, err
	}

	puk = prk.PublicKey

	return prk, puk, err

}

/**
  对text加密，text必须是一个hash值，例如md5、sha1等
  使用私钥prk
  使用随机熵增强加密安全，安全依赖于此熵，randsign
  返回加密结果，结果为数字证书r、s的序列化后拼接，然后用hex转换为string
*/
func sign(text []byte,randSign string,prk *ecdsa.PrivateKey) (string, error) {
	r, s, err := ecdsa.Sign(strings.NewReader(randSign), prk, text)
	if err != nil {
		return "", err
	}
	rt, err := r.MarshalText()
	if err != nil {
		return "", err
	}
	st, err := s.MarshalText()
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	_, err = w.Write([]byte(string(rt) + "+" + string(st)))
	if err != nil {
		return "", err
	}
	w.Flush()
	return hex.EncodeToString(b.Bytes()), nil
}

/**
  证书分解
  通过hex解码，分割成数字证书r，s
*/
func getSign( signature string) (rint, sint big.Int, err error)  {
	byterun, err := hex.DecodeString(signature)
	if err != nil {
		err = errors.New("decrypt error, "+ err.Error())
		return
	}
	r, err := gzip.NewReader(bytes.NewBuffer(byterun))
	if err !=  nil {
		err = errors.New("decode error,"+err.Error())
		return
	}
	defer r.Close()
	buf := make([]byte, 1024)
	count, err := r.Read(buf)
	if err != nil {
		fmt.Println("decode = ",err)
		err = errors.New("decode read error," + err.Error())
		return
	}
	rs := strings.Split(string(buf[:count]),"+")
	if len(rs) != 2 {
		err = errors.New("decode fail")
		return
	}
	err = rint.UnmarshalText([]byte(rs[0]))
	if err != nil {
		err = errors.New("decrypt rint fail, "+ err.Error())
		return
	}
	err = sint.UnmarshalText([]byte(rs[1]))
	if err != nil {
		err = errors.New("decrypt sint fail, "+ err.Error())
		return
	}
	return

}
/**
  校验文本内容是否与签名一致
  使用公钥校验签名和文本内容
*/
func verify(text []byte, signature string, key ecdsa.PublicKey) (bool, error) {

	rint, sint, err :=getSign(signature)
	if err != nil {
		return false, err
	}
	result := ecdsa.Verify(&key,text,&rint,&sint)
	return result, nil

}

/**
  hash加密
  使用md5加密
*/
func hashtext(text, salt string) ([]byte) {

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(text))
	result := Md5Inst.Sum([]byte(salt))

	return result
}

//产生一个新的私钥
func GenerateKey() (*ecdsa.PrivateKey, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, errors.New("failed GenerateKey")
	}
	return key, nil
}

//从文件导入私钥
func GetPrivateKeyByFile(filename string) (*ecdsa.PrivateKey, error) {
	prk, err := ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if len(prk) < 64 {
		return nil, errors.New("private file content too short")
	}
	//省去最后一个换行符
	ecdsaKey, err := crypto.HexToECDSA(prk[:64])
	if err != nil {
		return nil, err
	}
	return ecdsaKey, nil
}

//将私钥写入文件
func WritePrivateKeyToFile(filename string, key *ecdsa.PrivateKey) error {
	priStr := fmt.Sprintf("%s", hexutil.Encode(crypto.FromECDSA(key)))
	//省去前缀0x
	WriteToFile(filename, priStr[2:])
	return nil
}

//通过私钥导出公钥
func GetPublicKeyHexByPrivateKey(key *ecdsa.PrivateKey) (string, error) {
	pukStr := fmt.Sprintf("%s", hexutil.Encode(crypto.FromECDSAPub(&key.PublicKey)))
	return pukStr, nil
}

//通过公钥导出16进制
func GetPublickKeyHex(key ecdsa.PublicKey) (string, error) {
	pukStr := fmt.Sprintf("%s", hexutil.Encode(crypto.FromECDSAPub(&key)))
	return pukStr, nil
}

//通过公钥推导出地址
func GetAddressByPublicKey(key ecdsa.PublicKey) string {
	address := crypto.PubkeyToAddress(key)
	return  address.String()
}

//字节转地址
func BytesToAddress(bs []byte) string {
	addr := common.BytesToAddress(bs)
	return addr.String()
}

//字节转hash
func BytesToHash(bs []byte) string {
	hash := common.BytesToHash(bs)
	return hash.String()
}

//使用私钥进行签名
func SignText(digestHash string, prv *ecdsa.PrivateKey) ([]byte, error) {
	msg := crypto.Keccak256([]byte(digestHash))
	sig, err := crypto.Sign(msg, prv)
	return sig, err
}

//使用原数据和签名恢复公钥地址
func VerifyToAddress(digestHash string, sig []byte) (string, error) {
	msg := crypto.Keccak256([]byte(digestHash))
	recoveredPub, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return "", err
	}
	recoveredAddr := crypto.PubkeyToAddress(*recoveredPub)
	return recoveredAddr.String(), nil
}
