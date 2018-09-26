package main

import (
	"fmt"
	"os"
	"strconv"

	"./src/account"
	"./src/receipt"
	"github.com/syndtr/goleveldb/leveldb"
	log "github.com/withmandala/go-log" // https://github.com/Mandala/go-log
)

const DB_DIRECTORY_ACCOUNT = "db/account"
const DB_DIRECTORY_TRANSACTION = "db/transaction"
const DB_DIRECTORY_RECEIPT = "db/receipt"

const ROOT_RECEIPT = "ROOT_"
const NODE_RECEIPT = "NODE_"
const LAST_RECEIPT = "LAST_"
const RECEIPT = "RECEIPT"

func main() {
	logger := log.New(os.Stderr).WithColor().WithDebug().WithoutTimestamp()

	logger.Info("**************** 테스트 시작 - 테스트 데이터 생성****************")
	addressKey1 := "test-address 1"
	addressKey2 := "test-address 2"

	// account 동작 테스트
	logger.Debug("------------- account 동작 테스트 -------------")
	logger.Trace("============= create account =============")
	address1, _ := account.Allocation(addressKey1, 100, "user")
	address2, _ := account.Allocation(addressKey2, 300, "user")
	fmt.Println("created address1: ", address1)
	fmt.Println("created address2: ", address2)

	logger.Trace("============= duplication create account =============")
	duplicateAddress1, duplicatedCreateAccountErr1 := account.Allocation(addressKey1, 100, "user")
	duplicateAddress2, duplicatedCreateAccountErr2 := account.Allocation(addressKey2, 300, "user")
	fmt.Println("duplicated created address1: ", duplicateAddress1, "msg: ", duplicatedCreateAccountErr1)
	fmt.Println("duplicated created address2: ", duplicateAddress2, "msg: ", duplicatedCreateAccountErr2)

	logger.Trace("============= search account =============")
	ac1, _ := account.GetAccountInfo(address1)
	ac2, _ := account.GetAccountInfo(address2)
	fmt.Println(ac1)
	fmt.Println(ac2)

	logger.Trace("============= search invalid account  =============")
	_, invalidAccountErr1 := account.GetAccountInfo(address2 + "1234")
	_, invalidAccountErr2 := account.GetAccountInfo(address2 + "123")
	fmt.Println("invalid1 account: ", invalidAccountErr1)
	fmt.Println("invalid2 account: ", invalidAccountErr2)

	// transaction 동작 테스트
	fmt.Println("\n")
	logger.Debug("------------- transaction 동작 테스트 -------------")
	logger.Trace("============= create transaction =============")
	from := address1
	to := address2
	value, _ := strconv.Atoi("10")
	txId1, receiptIdIn, receiptIdOut, _ := account.Transfer(from, to, value)
	fmt.Println("created transactionId1 :", txId1)

	logger.Trace("============= search account after transaction =============")
	accountAfterTransaction1, _ := account.GetAccountInfo(address1)
	accountAfterTransaction2, _ := account.GetAccountInfo(address2)
	fmt.Println("account1 조회 :", accountAfterTransaction1)
	fmt.Println("account2 조회 :", accountAfterTransaction2)

	logger.Trace("============= overflow transaction =============")
	overflowFrom := address1
	overflowTo := address2
	overflowValue, _ := strconv.Atoi("1000")
	_, _, _, overflowErr := account.Transfer(overflowFrom, overflowTo, overflowValue)
	fmt.Println("overflow transaction :", overflowErr)

	logger.Trace("============= invalid account transaction case 1. from =============")
	invalidFrom1 := address1 + "123"
	invalidTo1 := address2
	invalidValue1, _ := strconv.Atoi("1")
	_, _, _, invalidTransactionErr1 := account.Transfer(invalidFrom1, invalidTo1, invalidValue1)
	fmt.Println("invalid transaction1 :", invalidTransactionErr1)

	logger.Trace("============= invalid account transaction case 2. to =============")
	invalidFrom2 := address1
	invalidTo2 := address2 + "123"
	invalidValue2, _ := strconv.Atoi("1")
	_, _, _, invalidTransactionErr2 := account.Transfer(invalidFrom2, invalidTo2, invalidValue2)
	fmt.Println("invalid transaction1 :", invalidTransactionErr2)

	logger.Trace("============= invalid account transaction case 3. not exist from =============")
	invalidFrom3 := "200d483437be5779fb20d4c195cf219fe1774cfe6df1f460168cfadde51e7a1l"
	invalidTo3 := address2
	invalidValue3, _ := strconv.Atoi("1")
	_, _, _, invalidTransactionErr3 := account.Transfer(invalidFrom3, invalidTo3, invalidValue3)
	fmt.Println("invalid transaction1 :", invalidTransactionErr3)

	logger.Trace("============= invalid account transaction case 4. not exist to =============")
	invalidFrom4 := address1
	invalidTo4 := "200d483437be5779fb20d4c195cf219fe1774cfe6df1f460168cfadde51e7a1l"
	invalidValue4, _ := strconv.Atoi("1")
	_, _, _, invalidTransactionErr4 := account.Transfer(invalidFrom4, invalidTo4, invalidValue4)
	fmt.Println("invalid transaction1 :", invalidTransactionErr4)

	// receipt 동작 테스트
	fmt.Println("\n")
	logger.Debug("------------- receipt 동작 테스트 -------------")

	logger.Trace("============= search receipt IN =============")
	receiptIdIn1 := receipt.GetSpecificReceipt(receiptIdIn)
	fmt.Println("In case 1.", receiptIdIn1)

	logger.Trace("============= search receipt OUT =============")
	receiptIdOut1 := receipt.GetSpecificReceipt(receiptIdOut)
	fmt.Println("Out case 1.", receiptIdOut1)

	logger.Trace("============= search account1 receipt =============")
	rr1 := receipt.GetRooteceipt(address1)
	lr1 := receipt.GetLastReceipt(address1)
	ar1 := receipt.GetAllReceipt(address1)

	fmt.Println("root receipt: ", rr1)
	fmt.Println("last receipt: ", lr1)
	fmt.Println("all  receipt: ", ar1)
	fmt.Println("len  receipt: ", len(ar1))

	logger.Trace("============= search account2 receipt =============")
	rr2 := receipt.GetRooteceipt(address2)
	lr2 := receipt.GetLastReceipt(address2)
	ar2 := receipt.GetAllReceipt(address2)

	fmt.Println("root receipt: ", rr2)
	fmt.Println("last receipt: ", lr2)
	fmt.Println("all  receipt: ", ar2)
	fmt.Println("len  receipt: ", len(ar2))

	// tx 추가한 후 조회
	fmt.Println("\n")
	logger.Debug("------------- tx 추가한 후 동작 테스트 -------------")
	logger.Trace("============= 추가적으로 transaction 발생 =============")
	addFrom := address1
	addTo := address2
	addValue, _ := strconv.Atoi("10")
	addTxId1, addReceiptIdIn, addReceiptIdOut, _ := account.Transfer(addFrom, addTo, addValue)
	fmt.Println("created transactionId1 :", addTxId1)

	accountAfterAddTransaction1, _ := account.GetAccountInfo(address1)
	accountAfterAddTransaction2, _ := account.GetAccountInfo(address2)
	fmt.Println("추가적으로 tx 발생 후 account1 조회 :", accountAfterAddTransaction1)
	fmt.Println("추가적으로 tx 발생 후 account2 조회 :", accountAfterAddTransaction2)

	logger.Trace("============= tx 추가한 후 search receipt IN =============")
	addReceiptIdIn1 := receipt.GetSpecificReceipt(addReceiptIdIn)
	fmt.Println("In case 1.", addReceiptIdIn1)

	logger.Trace("============= tx 추가한 후 search receipt OUT =============")
	addReceiptIdOut1 := receipt.GetSpecificReceipt(addReceiptIdOut)
	fmt.Println("Out case 1.", addReceiptIdOut1)

	logger.Trace("============= tx 추가한 후 search account1 receipt =============")
	addRr1 := receipt.GetRooteceipt(address1)
	addLr1 := receipt.GetLastReceipt(address1)
	addAr1 := receipt.GetAllReceipt(address1)

	fmt.Println("root receipt: ", addRr1)
	fmt.Println("last receipt: ", addLr1)
	fmt.Println("all  receipt: ", addAr1)
	fmt.Println("len  receipt: ", len(addAr1))

	// 테스트 종료
	fmt.Println("\n")
	logger.Info("**************** 테스트 종료 - 테스트에서 사용한 데이터 삭제 ****************")

	accountDb, _ := leveldb.OpenFile(DB_DIRECTORY_ACCOUNT, nil)
	transactionDb, _ := leveldb.OpenFile(DB_DIRECTORY_TRANSACTION, nil)
	receiptDb, _ := leveldb.OpenFile(DB_DIRECTORY_RECEIPT, nil)

	defer accountDb.Close()
	defer transactionDb.Close()
	defer receiptDb.Close()

	logger.Trace("============= delete account from database =============")
	accountDb.Delete([]byte(address1), nil)
	fmt.Println("deleted address1: ", address1)

	accountDb.Delete([]byte(address2), nil)
	fmt.Println("deleted address2: ", address2)

	logger.Trace("============= delete transaction from database =============")
	transactionDb.Delete([]byte(txId1), nil)
	fmt.Println("deleted transaction: ", txId1)

	logger.Trace("============= delete added transaction from database =============")
	transactionDb.Delete([]byte(addTxId1), nil)
	fmt.Println("deleted added transaction: ", addTxId1)

	logger.Trace("============= delete root receipt from database =============")
	receiptDb.Delete([]byte(ROOT_RECEIPT+address1+RECEIPT), nil)
	receiptDb.Delete([]byte(ROOT_RECEIPT+address2+RECEIPT), nil)
	fmt.Println("root receiptId1: ", ROOT_RECEIPT+address1+RECEIPT)
	fmt.Println("root receiptId2: ", ROOT_RECEIPT+address2+RECEIPT)

	logger.Trace("============= delete last receipt from database =============")
	receiptDb.Delete([]byte(LAST_RECEIPT+address1+RECEIPT), nil)
	receiptDb.Delete([]byte(LAST_RECEIPT+address2+RECEIPT), nil)
	fmt.Println("last receipt1: ", LAST_RECEIPT+address1+RECEIPT)
	fmt.Println("last receipt2: ", LAST_RECEIPT+address2+RECEIPT)

	logger.Trace("============= delete node receipt from database =============")
	receiptDb.Delete([]byte(NODE_RECEIPT+receiptIdIn+RECEIPT), nil)
	receiptDb.Delete([]byte(NODE_RECEIPT+receiptIdOut+RECEIPT), nil)
	fmt.Println("node receipt In: ", NODE_RECEIPT+receiptIdIn+RECEIPT)
	fmt.Println("node receipt OUT: ", NODE_RECEIPT+receiptIdOut+RECEIPT)

	logger.Trace("============= delete node added receipt from database =============")
	receiptDb.Delete([]byte(NODE_RECEIPT+addReceiptIdIn+RECEIPT), nil)
	receiptDb.Delete([]byte(NODE_RECEIPT+addReceiptIdOut+RECEIPT), nil)
	fmt.Println("node added receipt IN: ", NODE_RECEIPT+addReceiptIdIn+RECEIPT)
	fmt.Println("node added receipt OUT: ", NODE_RECEIPT+addReceiptIdOut+RECEIPT)
}
