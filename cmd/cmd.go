package main

import (
	"DIDWallet/model"
	"DIDWallet/utils"
	"errors"
	"fmt"
	hub "github.com/StupidTAO/DIDHub/model"
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/urfave/cli.v1"
	"io/ioutil"
	"log"
	"net/http"
	URL "net/url"
	"os"
	"sort"
	"strconv"
	"time"
)

const WALLET_FILE = ".wallet"
const SYNC = "sync"
const UNSYNC = "unsync"

func Run() {
	var language string

	app := cli.NewApp()
	app.Name = "DIDWallet"
	app.Usage = "hello world"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "port, p",
			Value: 8000,
			Usage: "listening port",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:     "did",
			Aliases:  []string{"d"},
			Usage:    "did operations",
			Category: "did",
			Subcommands: []cli.Command{
				{
					Name:  "create",
					Usage: "create data_file",
					Action: func(c *cli.Context) error {
						argsCount := len(os.Args)
						if (argsCount < 4) {
							fmt.Println("error: please enter compelte info")
							return nil
						}
						//读取文件内容并创建DID
						filename := os.Args[3]
						err := createDIDByPriFile(filename)
						if err != nil {
							fmt.Println("error is ", err)
							return err
						}
						return nil
					},
				},
				{
					Name:  "find",
					Usage: "find did",
					Action: func(c *cli.Context) error {
						argsCount := len(os.Args)
						if (argsCount < 4) {
							fmt.Println("error: please enter complete info")
							return nil
						}
						did := os.Args[3]
						err := findDocument(did)
						if err != nil {
							fmt.Println("error is ", err)
							return err
						}

						return nil
					},
				},
				{
					Name:  "list",
					Usage: "list",
					Action: func(c *cli.Context) error {
						fileInfo, err := utils.ReadFile(WALLET_FILE)
						if err != nil {
							fmt.Println("err is ", err)
						}
						fmt.Println(fileInfo)
						return nil
					},
				},
				{
					Name:  "claim",
					Usage: "claim private_key_file id_number",
					Action: func(c *cli.Context) error {

						argsCount := len(os.Args)
						if (argsCount < 5) {
							fmt.Println("error: please enter complete info")
							return nil
						}
						prifile := os.Args[3]
						did, _ := model.GetDIDByPrivateFile(prifile)

						//生成rawclaim
						idNumber := os.Args[4]
						cs, _ := model.GetCredentialSubject(did, idNumber)
						bs, _ := model.CredentialSubjectMarshal(cs)
						rawClaim := string(bs)

						//获取签名以及地址
						prikey, _ := utils.GetPrivateKeyByFile(prifile)
						rawClaimSig, _ := utils.SignText(rawClaim, prikey)
						addr := utils.GetAddressByPublicKey(prikey.PublicKey)
						url, err := getRequestClaimUrl(rawClaim, rawClaimSig, addr)
						if err != nil {
							fmt.Println("error is: ", err.Error())
							return nil
						}
						response, err := getUrl(url)
						if err != nil {
							fmt.Println("error is: ", err.Error())
							return nil
						}
						fmt.Println(response)

						return nil
					},
				},
				{
					Name:  "calc",
					Usage: "calc did",
					Action: func(c *cli.Context) error {
						argsCount := len(os.Args)
						if (argsCount < 4) {
							fmt.Println("error: please enter complete info")
							return nil
						}
						did := os.Args[3]
						err := model.SetContributionByAddr(did)
						if err != nil {
							fmt.Println("set contribution error: ", err.Error())
							return nil
						}

						return nil
					},
				},
				{
					Name:  "donate",
					Usage: "donate private_key_file claimId amount priority",
					Action: func(c *cli.Context) error {
						argsCount := len(os.Args)
						if (argsCount < 7) {
							fmt.Println("error: please enter complete info")
							return nil
						}

						prifile := os.Args[3]

						//生成rawDonate
						claimId := os.Args[4]
						amount := os.Args[5]
						priority := os.Args[6]
						wd, _ := model.GetWelfareDonate(claimId, amount, priority)
						bs, _ := model.WelfareDonateMarshal(wd)
						rawDonate := string(bs)

						//获取签名以及地址
						prikey, _ := utils.GetPrivateKeyByFile(prifile)
						rawDonateSig, _ := utils.SignText(rawDonate, prikey)
						addr := utils.GetAddressByPublicKey(prikey.PublicKey)
						url, err := getRequestDonateUrl(rawDonate, rawDonateSig, addr)
						if err != nil {
							fmt.Println("error is: ", err.Error())
							return nil
						}
						response, err := getUrl(url)
						if err != nil {
							fmt.Println("error is: ", err.Error())
							return nil
						}
						fmt.Println(response)
						return nil
					},
				},
			},
		},
		{
			Name:     "ballot",
			Aliases:  []string{"b"},
			Usage:    "ballot operations",
			Category: "ballot",
			Subcommands: []cli.Command{
				{
					Name:  "vote",
					Usage: "vote private_key_file proposalIndex",
					Action: func(c *cli.Context) error {
						begin := time.Now().UnixNano()
						argsCount := len(os.Args)
						if (argsCount < 4) {
							fmt.Println("error: please enter complete info")
							return nil
						}

						//获取提案ID，并给出投票
						index := os.Args[4]
						id, err := strconv.Atoi(index)
						if err != nil {
							fmt.Println("proposal index is not number!")
						}

						//读取私钥，并投票
						prk := os.Args[3]
						prk, err = utils.ReadFile(prk)
						if err != nil {
							fmt.Println("contract vote failed, error is ", err)
							return err
						}

						err = hub.ContractVote(prk, int64(id))
						if err != nil {
							fmt.Println("contract vote failed, error is ", err)
							return err
						}
						end := time.Now().UnixNano()
						fmt.Printf("spend time is: %d", end-begin)
						return nil
					},
				},
				{
					Name:  "right",
					Usage: "right address number",
					Action: func(c *cli.Context) error {
						argsCount := len(os.Args)
						if (argsCount < 5) {
							fmt.Println("error: please enter complete info")
							return nil
						}

						//为账户申请票权
						addr := os.Args[3]
						number := os.Args[4]
						num, err := strconv.Atoi(number)
						if err != nil {
							fmt.Println("right count is not number!")
						}

						err = hub.ContractGiveRightToVote(
							common.HexToAddress(addr),
							int64(num),
						)
						if err != nil {
							fmt.Println("contract get vote right failed, error is ", err)
							return err
						}
						return nil
					},
				},
				{	//撤销投票
					Name:  "revoke",
					Usage: "revoke private_key_file",
					Action: func(c *cli.Context) error {
						argsCount := len(os.Args)
						if (argsCount < 4) {
							fmt.Println("error: please enter complete info")
							return nil
						}

						//获取私钥并撤回投票
						prk := os.Args[3]
						prk, err := utils.ReadFile(prk)
						if err != nil {
							fmt.Println("contract revoke failed, error is ", err)
							return err
						}

						err = hub.ContractRevokeVote(prk)
						if err != nil {
							fmt.Println("contract get revoke failed, error is ", err)
							return err
						}
						return nil
					},
				},
				{	//winner 获取胜出者信息
					Name:  "winner",
					Usage: "winner",
					Action: func(c *cli.Context) error {
						argsCount := len(os.Args)
						if (argsCount < 3) {
							fmt.Println("error: please enter complete info")
							return nil
						}

						index, err := hub.ContractWinningProposal()
						if err != nil {
							fmt.Println("contract get proposal index failed, error is ", err)
							return err
						}
						proposalName, proposalIndex, needFunds, err := hub.ContractProposals(index)
						if err != nil {
							fmt.Println("contract get proposal info failed, error is ", err)
							return err
						}
						fmt.Printf("name: %s, vote count:%d, funds:%d\n", proposalName, proposalIndex, needFunds)

						return nil
					},
				},
				{	//voters 获取参与者信息
					Name:  "voters",
					Usage: "voters address",
					Action: func(c *cli.Context) error {
						argsCount := len(os.Args)
						if (argsCount < 4) {
							fmt.Println("error: please enter complete info")
							return nil
						}
						addr := os.Args[3]

						weight, voted, address, vote, err := hub.ContractVoters(common.HexToAddress(addr))
						if err != nil {
							fmt.Println("contract get proposal index failed, error is ", err)
							return err
						}
						fmt.Printf("voter weight is: %d, voted is: %t, delegate is: %s, vote index is: %d\n", weight, voted, address, vote)

						return nil
					},
				},
				{	//delaget 设置代理
					Name:  "delegate",
					Usage: "delegate address private_key_file",
					Action: func(c *cli.Context) error {
						argsCount := len(os.Args)
						if (argsCount < 5) {
							fmt.Println("error: please enter complete info")
							return nil
						}

						addr := os.Args[3]
						//获取私钥并设置代理
						prk := os.Args[4]
						prk, err := utils.ReadFile(prk)
						if err != nil {
							fmt.Println("contract revoke failed, error is ", err)
							return err
						}

						err = hub.ContractDelegate(common.HexToAddress(addr), prk)
						if err != nil {
							fmt.Println("contract set delegate failed, error is ", err)
							return err
						}

						return nil
					},
				},
			},
		},
	}
	app.Action = func(c *cli.Context) error {
		//fmt.Println(c.String("force"), c.String("f"))
		fmt.Println(language)

		// if c.Int("port") == 8000 {
		//     return cli.NewExitError("invalid port", 88)
		// }

		return nil
	}
	app.Before = func(c *cli.Context) error {
		return nil
	}
	app.After = func(c *cli.Context) error {
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	cli.HelpFlag = cli.BoolFlag{
		Name:  "help, h",
		Usage: "Help!Help!",
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "print-version, v",
		Usage: "print version",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}


func createDID(didDoc string) (string, error) {
	//将字符串解析为结构体，验证didDoc的合法性
	baseDIDDocEntry := new(model.BaseDIDDoc)
	err := model.BaseDIDDocUnmarshal(didDoc, baseDIDDocEntry)
	if err != nil {
		return "", err
	}

	//将结构体转换为字符串
	didDocBytes, err := model.BaseDIDDocMarshal(*baseDIDDocEntry)
	if err != nil {
		return "", err
	}

	//拼接did字符串
	specificId, err := model.GetSpecificId(string(didDocBytes))
	if err != nil {
		return "", err
	}
	did := "did:" + "welfare:" + specificId
	fmt.Println(did)

	//写入文件
	content := fmt.Sprintf("%s\n", did)
	utils.AppendToFile(WALLET_FILE, content)
	return did, nil
}

func getDIDBaseDocByPublicKeyHex(publicKeyHex string) (string, error){
	var baseDIDDoc model.BaseDIDDoc
	baseDIDDoc.Context = "https://w3id.org/did/v1"
	baseDIDDoc.Authentication = "#key-1"

	var publicKey model.PublicKey
	publicKey.Id = "#keys-1"
	publicKey.Type = "Secp256k1"
	publicKey.PublicKeyHex = publicKeyHex
	baseDIDDoc.PublicKey = publicKey

	bys, err := model.BaseDIDDocMarshal(baseDIDDoc)
	if err != nil {
		return "", err
	}
	return string(bys), nil
}

func createDIDByPriFile(priFile string) error {
	priKey, err := utils.GetPrivateKeyByFile(priFile)
	if err != nil {
		return err
	}

	publicKeyHex, err := utils.GetPublicKeyHexByPrivateKey(priKey)
	if err != nil {
		return err
	}

	content, err := getDIDBaseDocByPublicKeyHex(publicKeyHex)
	if err != nil {
		return err
	}

	did, err := model.GetDIDByPrivateFile(priFile)
	if err != nil {
		return err
	}
	//DID document同步到hub数据库和区块链中
	//写入数据库和区块链
	chainAddr := new(model.DBDIDChainAddr)
	chainAddr.Did = did
	chainAddr.DidChainAddr = utils.GetAddressByPublicKey(priKey.PublicKey)
	chainAddr.UpdateTime = time.Now().Add(8 * time.Hour)
	chainAddr.CreateTime = time.Now().Add(8 * time.Hour)
	chainAddr.IsAvailable = 1

	err = hub.InsertDBDIDChainAddr(hub.DBDIDChainAddr(*chainAddr))
	if err != nil {
		return err
	}
	fmt.Printf("did chain addr insert hub success")
	//查询数据库和区块链
	docs, err := model.FindDBDIDDoc(did)
	if err != nil {
		return err
	}
	if len(docs) > 0 {
		//已经存在则什么都不做
		return errors.New("did already exist")
	}

	//写入数据库和区块链
	doc := new(model.DBDIDDoc)
	doc.Did = did
	doc.DidDoc = content
	doc.UpdateTime = time.Now().Add(8 * time.Hour)
	doc.CreateTime = time.Now().Add(8 * time.Hour)
	doc.IsAvailable = 1
	return hub.InsertHubDIDDoc(hub.DBDIDDoc(*doc))

}

func createDIDByFile(docname string, sync string) error {
	content, err := utils.ReadFile(docname)
	if err != nil {
		fmt.Println("error is ", err)
		return err
	}

	//在本地创建DID，并写入.wallet文件中
	did, err := createDID(content)
	if err != nil {
		fmt.Println("error is ", err)
		return err
	}

	//DID document同步到hub数据库和区块链中
	if sync == SYNC {
		//查询数据库和区块链
		docs, err := model.FindDBDIDDoc(did)
		if err != nil {
			return err
		}
		if len(docs) > 0 {
			//已经存在则什么都不做
			return errors.New("did already exist")
		}

		//写入数据库和区块链
		doc := new(model.DBDIDDoc)
		doc.Did = did
		doc.DidDoc = content
		doc.UpdateTime = time.Now().Add(8 * time.Hour)
		doc.CreateTime = time.Now().Add(8 * time.Hour)
		doc.IsAvailable = 1
		return model.InsertDBDIDDoc(*doc)
	}
	return nil
}

func findDocument(did string) error {
	docs, err := model.FindDBDIDDoc(did)
	if err != nil {
		return err
	}
	if len(docs) == 0 {
		return errors.New("can't find docs")
	}
	fmt.Println("the did corresponding document is: ", docs[0].DidDoc)
	return nil
}

func getRequestClaimUrl(rawClaim string, sigClaim []byte, addr string) (string, error) {
	//对参数进行URL编码
	sigClaimBase58 := utils.Base58Encode(sigClaim)
	rawClaim = URL.QueryEscape(rawClaim)
	sigClaimBase58 = URL.QueryEscape(sigClaimBase58)

	url := fmt.Sprintf("http://127.0.0.1:8000/getClaim?rawClaim=%s&sigClaim=%s&addr=%s", rawClaim, sigClaimBase58, addr)
	return url, nil
}

func getRequestDonateUrl(rawDonate string, sigDonate []byte, addr string) (string, error) {
	//对参数进行URL编码
	sigDonateBase58 := utils.Base58Encode(sigDonate)
	rawDonate = URL.QueryEscape(rawDonate)
	sigDonateBase58 = URL.QueryEscape(sigDonateBase58)

	url := fmt.Sprintf("http://127.0.0.1:8001/transfer?rawDonate=%s&sigDonate=%s&addr=%s", rawDonate, sigDonateBase58, addr)
	return url, nil
}

func getUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	s, err:=ioutil.ReadAll(resp.Body)

	return string(s), nil
}
