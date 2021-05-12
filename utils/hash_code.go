package utils

import (
	"crypto/sha256"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

func GetSHA256HashCode(message []byte) []byte {
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	hash.Write(message)
	//计算哈希值
	bytes := hash.Sum(nil)
	//返回哈希值
	return bytes
}

func GetRipemd160HashCode(message []byte) []byte {
	hasher := ripemd160.New()
	hasher.Write(message)
	hashBytes := hasher.Sum(nil)
	return hashBytes
}

func Base58Encode(message []byte) string {
	return base58.Encode(message)
}

func Base58Decode(msgStr string) []byte {
	return base58.Decode(msgStr)
}
