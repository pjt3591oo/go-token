# chaincode

# 목차

* account
* transaction
* receipt
* utils

### dependencies

```bash
$ go get github.com/withmandala/go-log       # log lib
$ go get github.com/syndtr/goleveldb/leveldb # levelDB
```

### run & build

* build

```
$ go build cli.go
$ ./cli
```

* run

```
$ go run cli.go
```

### How to use

0. 빌드
1. 사용법 조회
2. 계정 생성
3. 계정 조회
4. 토큰 전송(transaction 발생)
5. tx 조회
6. 특정 receipt 조회
7. 특정 account에 발생된 receipt 조회

* **```0. build```***

```bash
$ go build cli.go
$ chmod 777 ./cli
```

* **```1. 사용법 조회```**

```
$ ./cli # OR ./cli --help

	-accountreceipt string
			account의 receipt내역, root, last 조회
	-createaccount string
			create account
	-from string
			sender
	-istransfer string
			해당 옵션에 1을 전달할 경우 from, to, value 인자를 전달해야 함
	-searchaccount string
			account 조회
	-searchreceipt string
			receipt 조회
	-searchtx string
			transaction 조회
	-to string
			receiver
	-value string
			token value
```

* **```2. 계정 생성```**

```bash
$ ./cli --createaccount "key" # key에 따라서 account 생성
```

```bash
$ ./cli --createaccount 12341qe
[INFO]  used key: {  12341qe  }
[INFO]  created account: {  695c6dd6d52f560b0ec77f3a4ca5c7ba1c590ecf217c95683d42ced73e30496f  }
```

정상적으로 출력될 경우 account를 만들기 위해 사용한 key와 생성된 account 반환

```bash
$ ./cli --createaccount 12341qe
[ERROR] main.main:cli.go:46 account 생성실패:  이미 account가 존재합니다.
```

account가 존재한다면 생성하지 않음.

*  **```3.계정 조회```**

```bash
$ ./cli --searchaccount "address"
```

```bash
$ ./cli --searchaccount 695c6dd6d52f560b0ec77f3a4ca5c7ba1c590ecf217c95683d42ced73e30496f
[INFO]  {"Value":100,"Kind":"user","Address":"695c6dd6d52f560b0ec77f3a4ca5c7ba1c590ecf217c95683d42ced73e30496f"}
```

./cli.go 코드중 다음 코드에서 두 번째 인자를 수정하여 초기할당량을 수정할 수 있다.

```go
account, accountCreateErr := account.Allocation(*accountKey, 100, "user")
```

```bash
$ ./cli --searchaccount 695c6dd6d52f560b0ec77f3a4ca5c7ba1c590ecf217c9568
[ERROR] main.main:cli.go:88 invalid address

```

만약, 전달된 인자가 전확한 address가 아니라면 에러출력.

* **```4. 토큰 전송(transaction 발생)```**

```bash
$ ./cli --istransfer 1 --from "보내는 address" --to "받는 address" --value "토큰 전송량"
```

```
$ ./cli --istransfer 1 --from 958d51602bbfbd18b2a084ba848a827c29952bfef170c936419b0922994c0589 --to 075c32345dd82047e1d3da2aa68f4c62581e20f7154b7bc10fba791ceef5f3cd --value 10
[INFO]  transactionId:  53f1f74d45ff6cb235e18bc603febbe97a0cc3c32b62854a7390aec10722dbcd
```

전송이 정상적으로 완료됬다면 transaction Id 출력

```bash
$ $ ./cli --istransfer 1 --from 958d51602bbfbd18b2a084ba848a827c29952bfef170c936419b0922994c0589 --to 075c32345dd82047e1d3da2aa68f4c62581e20f7154b7bc10fba791ceef5f3cd --value 1000
[ERROR] main.main:cli.go:59 잔액이 부족합니다.
```

만약 잔액이 부족하거나 옳바른 account가 아니라면 에러출력

```
[ERROR] main.main:cli.go:59 잔액이 부족합니다.
[ERROR] main.main:cli.go:59 invalid to account
[ERROR] main.main:cli.go:59 to address가 존재하지 않습니다.
```

잔액부족, 잘못된 address, 존재하지 않는 address

* **```5. tx 조회```**

```bash
$ ./cli --searchtx "transaction Id"
```

```bash
$ ./cli --searchtx 53f1f74d45ff6cb235e18bc603febbe97a0cc3c32b62854a7390aec10722dbcd
[INFO]  {"TxId":"53f1f74d45ff6cb235e18bc603febbe97a0cc3c32b62854a7390aec10722dbcd","From":"958d51602bbfbd18b2a084ba848a827c29952bfef170c936419b0922994c0589","To":"075c32345dd82047e1d3da2aa68f4c62581e20f7154b7bc10fba791ceef5f3cd","Value":10,"Timestamp":"1537942727848904297"}
```

```
$ ./cli --searchtx 53f1f74d45ff6cb235e18bc603febbe97a0cc3c32b62854a7390aec10722dbc
[ERROR] main.main:cli.go:80 transaction ID가 옳바르지 않습니다.
```

만약 tx가 옳바르지 않거나 존재하지 않는다면 에러출력

* **```6. 특정 receipt 조회```**

```bash
$ ./cli --searchreceipt "receipt Id"
```

```bash
$ ./cli --searchreceipt 926b8711f45a72e6797f695de63b78147e1bacd14450c3a48978c20c4af90b65
[INFO]  {"receiptId":"926b8711f45a72e6797f695de63b78147e1bacd14450c3a48978c20c4af90b65","txId":"9a4032c16712c9a22968485b86b26fc17a32d414df2a4041e6ae0fbf6b2d91b9","nextReceiptId":"ba83f07af0a3c5a514a8d6e9d41d1bb9e604ec2ab194a499dc9dffeb5b179ecb","prevReceiptId":"","status":"OUT"}
```

출력 결과중 ```status```는 **IN** 또는 **OUT** 값을 갖는다.


* **```7. 특정 account에 발생된 receipt 조회```**

```bash
$ ./cli --accountreceipt "address"
```

```bash
$ ./cli --accountreceipt 958d51602bbfbd18b2a084ba848a827c29952bfef170c936419b0922994c0589
[INFO]  ROOT RECEIPT :  {"receiptId":"926b8711f45a72e6797f695de63b78147e1bacd14450c3a48978c20c4af90b65","txId":"9a4032c16712c9a22968485b86b26fc17a32d414df2a4041e6ae0fbf6b2d91b9","nextReceiptId":"","prevReceiptId":"","status":"OUT"}
[INFO]  LAST RECEIPT :  {"receiptId":"0b90b82c5cf567652ce84a024b19918b515dae3243a4c2fd5546ed295d220a36"}
[INFO]  All RECEIPT  ( 3 ):  [{926b8711f45a72e6797f695de63b78147e1bacd14450c3a48978c20c4af90b65 9a4032c16712c9a22968485b86b26fc17a32d414df2a4041e6ae0fbf6b2d91b9 ba83f07af0a3c5a514a8d6e9d41d1bb9e604ec2ab194a499dc9dffeb5b179ecb  OUT} {ba83f07af0a3c5a514a8d6e9d41d1bb9e604ec2ab194a499dc9dffeb5b179ecb c5c1e44654d8812d73090ad8bb517bf022997f76f2b2e71f0a9ab2d244450d6e 0b90b82c5cf567652ce84a024b19918b515dae3243a4c2fd5546ed295d220a36 926b8711f45a72e6797f695de63b78147e1bacd14450c3a48978c20c4af90b65 OUT} {0b90b82c5cf567652ce84a024b19918b515dae3243a4c2fd5546ed295d220a36 53f1f74d45ff6cb235e18bc603febbe97a0cc3c32b62854a7390aec10722dbcd  ba83f07af0a3c5a514a8d6e9d41d1bb9e604ec2ab194a499dc9dffeb5b179ecb OUT}]
```

ROOT RECEIPT는 가장 처음에 발생된 txId를 가짐
LAST RECEIPT는 가장 마지막에 발생된 txId를 가짐
ALL RECEIPT는 가장 처음 발생된 txId부터 가장 마지막 txId까지 리스트로 반환

ALL RECEIPT의 각각의 아이템음 다음과 같은 데이터 구조를 갖는다.

```
{"receiptId":"","txId":"","nextReceiptId":"","prevReceiptId":"","status":""}
```

> 여기서는 따로 validation 하지 않음

### database

database는 leveldb 사용

**account**, **transaction**, **receipt는** 각각 **./db/account**, **./dbtransaction**, **./db/receipt**에서 관리

### chaincode 

[코드보러가기](https://github.com/pjt3591oo/hyperledger-fabric-token)
