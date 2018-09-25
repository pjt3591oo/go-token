# chaincode

# 목차

* account
* transaction
* receipt

### dependencies

```bash
$ go get github.com/withmandala/go-log
```

### run & build

* build

```
$ go build main.go
$ ./main
```

* run

```
$ go run main.go
```

### chaincode 

[코드보러가기](https://github.com/pjt3591oo/hyperledger-fabric-token)

### chaincode query/invoke

```bash
# 토큰정보 조회
$ peer chaincode query -n token -c '{"Args":["get_token_info"]}' -C myc

# 어카운트 생성 1
peer chaincode invoke -n token -c '{"Args":["create_account", "1", "100"]}' -C myc
# 어카운트 생성 2
peer chaincode invoke -n token -c '{"Args":["create_account", "2", "300"]}' -C myc

# 어카운트1 정보 조회
peer chaincode query -n token -c '{"Args":["get_account", "6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b"]}' -C myc
# 어카운트2 정보 조회
peer chaincode query -n token -c '{"Args":["get_account", "d4735e3a265e16eee03f59718b9b5d03019c07d8b6c51f90da3a666eec13ab35"]}' -C myc

# 토큰전송 어카운트1 -> 어카운트2
$ peer chaincode invoke -n token -c '{"Args":["transfer", "6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b","d4735e3a265e16eee03f59718b9b5d03019c07d8b6c51f90da3a666eec13ab35","30"]}' -C myc
# 토큰전송 어카운트2 -> 어카운트1
$ peer chaincode invoke -n token -c '{"Args":["transfer", "d4735e3a265e16eee03f59718b9b5d03019c07d8b6c51f90da3a666eec13ab35","6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b","10"]}' -C myc

# root receipt account1 조회
$ peer chaincode query -n token -c '{"Args":["get_root_receipt", "6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b"]}' -C myc
# root receipt account2 조회
$ peer chaincode query -n token -c '{"Args":["get_root_receipt", "d4735e3a265e16eee03f59718b9b5d03019c07d8b6c51f90da3a666eec13ab35"]}' -C myc

# last receipt account1 조회
$ peer chaincode query -n token -c '{"Args":["get_last_receipt", "6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b"]}' -C myc
# last receipt account2 조회
$ peer chaincode query -n token -c '{"Args":["get_last_receipt", "d4735e3a265e16eee03f59718b9b5d03019c07d8b6c51f90da3a666eec13ab35"]}' -C myc

# 특정 receipt 조회
$ peer chaincode query -n token -c '{"Args":["get_receipt", "acf6aad4d454a9a80ef1c32231f6ff4728d982a66b877b156f168bb693596020"]}' -C myc

# tx 조회
$ peer chaincode query -n token -c '{"Args":["get_receipt", "e445b62fafae008942acdedb2f5043889b5692d3fcb757d5e5743d7ee516d90c"]}' -C myc
```