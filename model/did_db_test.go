package model

import (
	"DIDWallet/utils"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestInsertDBDIDDoc(t *testing.T) {
	doc := new(DBDIDDoc)
	doc.Did = "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	doc.DidDoc = "{\"Context\":\"https://w3id.org/did/v1\",\"PublicKey\":{\"Id\":\"#keys-1\",\"Type\":\"Secp256k1\",\"PublicKeyHex\":\"02b97c30de767f084ce3080168ee293053ba33b235d7116a3263d29f1450936b71\"},\"Authentication\":\"#key-1\"}\n"
	doc.IsAvailable = 1
	doc.CreateTime = time.Now().Add(8 * time.Hour)
	doc.UpdateTime = time.Now().Add(8 * time.Hour)
	err := InsertDBDIDDoc(*doc)
	if err != nil {
		t.Error("error is ", err)
	}
}

func TestFindDBDIDDoc(t *testing.T) {
	did := "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	docs, err := FindDBDIDDoc(did)
	if err != nil {
		t.Error("find db did document error ", err)
		return
	}
	t.Log("docs counts : ", len(docs))
}

func TestInsertDBDIDClaim(t *testing.T) {
	claim := new(DBDIDClaim)
	claim.Did = "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	claim.DidClaim = "claim-test"
	bs := utils.GetRipemd160HashCode([]byte(claim.DidClaim))
	claim.ClaimId = utils.Base58Encode(bs)
	claim.IsAvailable = 1
	claim.CreateTime = time.Now().Add(8 * time.Hour)
	claim.UpdateTime = time.Now().Add(8 * time.Hour)
	err := InsertDBDIDClaim(*claim)
	if err != nil {
		t.Error("error is ", err)
	}
}

func TestFindDBDIDClaim(t *testing.T) {
	did := "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	claims, err := FindDBDIDClaim(did)
	if err != nil {
		t.Error("find db did claims error ", err)
		return
	}
	t.Log("claims counts : ", len(claims))
}

func TestInsertDBDIDPublicKey(t *testing.T) {
	publicKey := new(DBDIDPublicKey)
	publicKey.Did = "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	publicKey.DidPublicKey = "publickey-test"
	publicKey.IsAvailable = 1
	publicKey.CreateTime = time.Now().Add(8 * time.Hour)
	publicKey.UpdateTime = time.Now().Add(8 * time.Hour)
	err := InsertDBDIDPublicKey(*publicKey)
	if err != nil {
		t.Error("error is ", err)
	}
}

func TestFindDBDIDPublicKey(t *testing.T) {
	did := "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	publickey, err := FindDBDIDPublicKey(did)
	if err != nil {
		t.Error("find db did publickey error ", err)
		return
	}
	t.Log("publickey counts : ", len(publickey))
}

func TestInsertDBDIDChainAddr(t *testing.T) {
	chainAddr := new(DBDIDChainAddr)
	chainAddr.Did = "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	chainAddr.DidChainAddr = "chainAddr-test"
	chainAddr.IsAvailable = 1
	chainAddr.CreateTime = time.Now().Add(8 * time.Hour)
	chainAddr.UpdateTime = time.Now().Add(8 * time.Hour)
	err := InsertDBDIDChainAddr(*chainAddr)
	if err != nil {
		t.Error("error is ", err)
	}
}

func TestFindDBDIDChainAddr(t *testing.T) {
	did := "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	_, err := FindDBDIDChainAddr(did)
	if err != nil {
		t.Error("find db did chain address error ", err)
		return
	}
}

func TestInsertTransaction(t *testing.T) {
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		amount := uint32(rand.Intn(1000))
		tx := new(Transaction)
		tx.TxId = strconv.Itoa(i)
		tx.FromAddr = "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
		tx.ToAddr = "did:welfare:123456789abcdefghijklmnopqrs"
		tx.Amount = amount
		tx.ProjectPriority = 1.4
		tx.Contribution = 0
		tx.CreateTime = time.Now().Add((24*14+8) * time.Hour)
		tx.UpdateTime = time.Now().Add((24*14+8) * time.Hour)
		tx.HasCaculate = 0
		tx.Payload = "test-0"

		err := InsertTransaction(tx)
		if err != nil {
			t.Error("error is ", err)
		}
	}

}

func TestFindTransaction(t *testing.T) {
	txId := "1"
	txs, err := FindTransaction(txId)
	if err != nil {
		t.Error("find db transaction error ", err)
		return
	}
	t.Log("len txs is ", len(txs))
}

func TestFindTransactionByFromAddrOrderByAmount(t *testing.T) {
	addr := "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	txs, err := FindTransactionByFromAddrOrderByAmount(addr)
	if err != nil {
		t.Error("find db transaction by from addr order by amount: ", err)
	}
	t.Log("len txs is ", len(txs))
}

func TestFindTransactionByFromAddrOrderByPriority(t *testing.T) {
	addr := "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	txs, err := FindTransactionByFromAddrOrderByPriority(addr)
	if err != nil {
		t.Error("find db transaction by from addr order by priority: ", err)
	}
	t.Log("len txs is ", len(txs))
}

func TestGetMiddleAmount(t *testing.T) {
	addr := "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	txs, err := FindTransactionByFromAddrOrderByPriority(addr)
	if err != nil {
		t.Error("find db transaction by from addr order by priority: ", err)
	}
	amount, err := GetMiddleAmount(txs)
	if err != nil {
		t.Error("find db transaction by from addr order by priority: ", err)
	}
	fmt.Println("amount = ", amount)
}

func TestGetMiddlePriority(t *testing.T) {
	addr := "did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1"
	txs, err := FindTransactionByFromAddrOrderByPriority(addr)
	if err != nil {
		t.Error("find db transaction by from addr order by priority: ", err)
		return
	}
	priority, err := GetMiddlePriority(txs)
	if err != nil {
		t.Error("find db transaction by from addr order by priority: ", err)
		return
	}
	fmt.Println("priority = ", priority)
}

func TestUpdateContribution(t *testing.T) {
	err := UpdateTransactionContribution("1", 1.25)
	if err != nil {
		t.Error("update transaction contribution: ", err)
	}
}

func TestSetContributionByAddr(t *testing.T) {
	err := SetContributionByAddr("did:welfare:2z7tBiNoYRTCGGNyKcxatEmYxuN1")
	if err != nil {
		t.Error("set contribution error: ", err.Error())
		return
	}

	fmt.Println("set contribution normally")
}
