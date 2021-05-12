package model

import (
	"DIDWallet/db"
	"DIDWallet/utils"
	"errors"
	"fmt"
	"time"
)

type DBDIDDoc struct {
	Id uint					`db:"id"`
	Did string				`db:"did"`
	DidDoc string			`db:"did_doc"`
	CreateTime time.Time	`db:"createTime"`
	UpdateTime time.Time	`db:"updateTime"`
	IsAvailable	uint		`db:"isAvailable"`
}

type DBDIDClaim struct {
	Id uint					`db:"id"`
	Did string				`db:"did"`
	ClaimId string			`db:"claimId"`
	DidClaim string			`db:"didClaim"`
	CreateTime time.Time	`db:"createTime"`
	UpdateTime time.Time	`db:"updateTime"`
	IsAvailable	uint		`db:"isAvailable"`
}

type DBDIDPublicKey struct {
	Id uint					`db:"id"`
	Did string				`db:"did"`
	DidPublicKey string		`db:"did_public_key"`
	CreateTime time.Time	`db:"createTime"`
	UpdateTime time.Time	`db:"updateTime"`
	IsAvailable	uint		`db:"isAvailable"`
}

type DBDIDChainAddr struct {
	Id uint					`db:"id"`
	Did string				`db:"did"`
	DidChainAddr string		`db:"did_chain_addr"`
	CreateTime time.Time	`db:"createTime"`
	UpdateTime time.Time	`db:"updateTime"`
	IsAvailable	uint		`db:"isAvailable"`
}

type Transaction struct {
	Id uint					`db:"id"`
	TxId string				`db:"txId"`
	FromAddr string			`db:"fromAddr"`
	ToAddr string			`db:"toAddr"`
	Payload string			`db:"payload"`
	Amount uint32			`db:"amount"`
	ProjectPriority float32 `db:"projectPriority"`
	Contribution float32	`db:"contribution"`
	CreateTime time.Time	`db:"createTime"`
	UpdateTime time.Time	`db:"updateTime"`
	HasCaculate uint		`db:"hasCaculate"`
}

func InsertDBDIDDoc(didDoc DBDIDDoc) error {
	sql := "insert into did_document(did, did_doc, createTime, updateTime, isAvailable)values (?,?,?,?,?)"

	//执行SQL语句
	db.InitDB()
	_, err := db.DB.Exec(sql, didDoc.Did, didDoc.DidDoc, didDoc.CreateTime, didDoc.UpdateTime, didDoc.IsAvailable)
	if err != nil {
		fmt.Println("exec failed,", err)
		return err
	}

	return nil
}

func FindDBDIDDoc(did string) ([]DBDIDDoc, error){
	DB := db.InitDB()

	var docs []DBDIDDoc
	sql := "select id, did, did_doc, createTime, updateTime, isAvailable from did_document where did=?"
	err := DB.Select(&docs, sql, did)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return nil, err
	}
	return docs, nil
}

func InsertDBDIDClaim(didClaim DBDIDClaim) error {
	sql := "insert into did_claim(did, claimId, didClaim, createTime, updateTime, isAvailable)values (?,?,?,?,?,?)"

	//执行SQL语句
	db.InitDB()
	_, err := db.DB.Exec(sql, didClaim.Did, didClaim.ClaimId, didClaim.DidClaim, didClaim.CreateTime, didClaim.UpdateTime, didClaim.IsAvailable)
	if err != nil {
		fmt.Println("exec failed,", err)
		return err
	}

	return nil
}

func FindDBDIDClaim(did string) ([]DBDIDClaim, error){
	DB := db.InitDB()

	var claims []DBDIDClaim
	sql := "select id, did, claimId, didClaim, createTime, updateTime, isAvailable from did_claim where did=?"
	err := DB.Select(&claims, sql, did)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return nil, err
	}
	return claims, nil
}

func InsertDBDIDPublicKey(didPublicKey DBDIDPublicKey) error {
	sql := "insert into did_publickey(did, did_public_key, createTime, updateTime, isAvailable)values (?,?,?,?,?)"

	//执行SQL语句
	db.InitDB()
	_, err := db.DB.Exec(sql, didPublicKey.Did, didPublicKey.DidPublicKey, didPublicKey.CreateTime, didPublicKey.UpdateTime, didPublicKey.IsAvailable)
	if err != nil {
		fmt.Println("exec failed,", err)
		return err
	}

	return nil
}

func FindDBDIDPublicKey(did string) ([]DBDIDPublicKey, error){
	DB := db.InitDB()

	var publickeys []DBDIDPublicKey
	sql := "select id, did, did_public_key, createTime, updateTime, isAvailable from did_publickey where did=?"
	err := DB.Select(&publickeys, sql, did)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return nil, err
	}
	return publickeys, nil
}

func InsertDBDIDChainAddr(didChainAddr DBDIDChainAddr) error {
	sql := "insert into did_chain_addr(did, did_chain_addr, createTime, updateTime, isAvailable)values (?,?,?,?,?)"

	//执行SQL语句
	db.InitDB()
	_, err := db.DB.Exec(sql, didChainAddr.Did, didChainAddr.DidChainAddr, didChainAddr.CreateTime, didChainAddr.UpdateTime, didChainAddr.IsAvailable)
	if err != nil {
		fmt.Println("exec failed,", err)
		return err
	}

	return nil
}

func FindDBDIDChainAddr(did string) ([]DBDIDChainAddr, error){
	DB := db.InitDB()

	var chainAddrs []DBDIDChainAddr
	sql := "select id, did, did_chain_addr, createTime, updateTime, isAvailable from did_chain_addr where did=?"
	err := DB.Select(&chainAddrs, sql, did)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return nil, err
	}
	return chainAddrs, nil
}

func InsertTransaction(tx *Transaction) error {
	sql := "insert into transactions(txId, fromAddr, toAddr, payload, amount, projectPriority, contribution, createTime, updateTime, hasCaculate)values (?,?,?,?,?,?,?,?,?,?)"

	//执行SQL语句
	db.InitDB()
	_, err := db.DB.Exec(sql, tx.TxId, tx.FromAddr, tx.ToAddr, tx.Payload, tx.Amount, tx.ProjectPriority, tx.Contribution, tx.CreateTime, tx.UpdateTime, tx.HasCaculate)
	if err != nil {
		fmt.Println("exec failed,", err)
		return err
	}

	return nil
}

func FindTransaction(txId string) ([]Transaction, error){
	DB := db.InitDB()

	var txs []Transaction
	var sql = "select id, txId, fromAddr, toAddr, payload, amount, projectPriority, contribution, createTime, updateTime, hasCaculate from transactions where id=?"
	err := DB.Select(&txs, sql, txId)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return nil, err
	}
	return txs, nil
}

func FindTransactionByFromAddrOrderByAmount(fromAddr string) ([]Transaction, error){
	DB := db.InitDB()

	var txs []Transaction
	var sql = "select id, txId, fromAddr, toAddr, payload, amount, projectPriority, contribution, createTime, updateTime, hasCaculate from transactions where fromAddr=? and hasCaculate=? ORDER BY amount"
	err := DB.Select(&txs, sql, fromAddr, 0)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return nil, err
	}
	return txs, nil
}

func FindTransactionByFromAddrOrderByPriority(fromAddr string) ([]Transaction, error){
	DB := db.InitDB()

	var txs []Transaction
	var sql = "select id, txId, fromAddr, toAddr, payload, amount, projectPriority, contribution, createTime, updateTime, hasCaculate from transactions where fromAddr=? and hasCaculate=? ORDER BY projectPriority"
	err := DB.Select(&txs, sql, fromAddr, 0)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return nil, err
	}
	return txs, nil
}

func GetMiddleAmount(txs []Transaction) (amount uint32, err error) {
	if (len(txs) == 0) {
		return 0, errors.New("txs len is 0")
	}
	if (len(txs) % 2) == 0 {
		index1 := len(txs) / 2 - 1
		index2 := len(txs) / 2
		return (txs[index2].Amount + txs[index1].Amount) / 2, nil
	} else {
		index := len(txs) / 2
		return txs[index].Amount, nil
	}
}

func GetMiddlePriority(txs []Transaction) (priority float32, err error) {
	if (len(txs) == 0) {
		return 0, errors.New("txs len is 0")
	}

	if (len(txs) % 2) == 0 {
		index1 := len(txs) / 2 - 1
		index2 := len(txs) / 2
		return (txs[index2].ProjectPriority + txs[index1].ProjectPriority) / 2, nil
	} else {
		index := len(txs) / 2
		return txs[index].ProjectPriority, nil
	}
}

func UpdateTransactionContribution(txId string, contribution float64) (error) {
	sql := "UPDATE transactions SET contribution=?, hasCaculate=1 WHERE txId=? and hasCaculate=0"

	//执行SQL语句
	db.InitDB()
	_, err := db.DB.Exec(sql, contribution, txId)
	if err != nil {
		fmt.Println("exec failed,", err)
		return err
	}

	return nil
}

//设置没有计算贡献度的交易
func SetContributionByAddr(fromAddr string) (err error) {
	//获取中值，出度
	txs1, _ := FindTransactionByFromAddrOrderByAmount(fromAddr)
	txs2, _ := FindTransactionByFromAddrOrderByPriority(fromAddr)
	amountMid, _ := GetMiddleAmount(txs1)
	prioritMid, _ := GetMiddlePriority(txs2)
	outDgree := len(txs1)

	//计算每条记录的贡献度指标
	for _, tx := range txs1 {
		selfAmountMid := int32(tx.Amount - amountMid)
		slefPriorityMid := float32(tx.ProjectPriority - prioritMid)
		contribution, err := utils.GetContribution(tx.Amount, selfAmountMid, tx.ProjectPriority, slefPriorityMid, outDgree)
		if err != nil {
			fmt.Println("get contribution err: ", err.Error())
			return err
		}
		err = UpdateTransactionContribution(tx.TxId, contribution)
		if err != nil {
			fmt.Println("update transaction contribution err: ", err.Error())
			return err
		}
	}

	return nil
}
