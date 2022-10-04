```sh
peer chaincode install -n auth6 -v 1.0 -p github.com/chaincode/auth
peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n auth6 -v 1.0 -c '{"Args":["init"]}' -P "OR ('Org1MSP.peer','Org2MSP.peer')"

peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n auth6 -c '{"Args":["register","魏延","3059301306072a8648ce3d020106082a8648ce3d030107034200040ae52ae29402727b22a4f777417e5b3061feadb8c26db457b8642102901d35eb6e8432cf15e5b29daaec1bc4031eb2a7fda84c8d9497b7c68fe406af9e697c4e","30257b61c3d7eec73edc14d30f82681b2ad8921201604867cac5b1cc4fed3633","3b66162b10e93ab6a2c0e214dd3eadcafd096cf309808fde1cb8826c3647380c"]}'
peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n auth6 -c '{"Args":["register","thor","3059301306072a8648ce3d020106082a8648ce3d03010703420004f8d0bae2e00dec4c38f9c759f957af7b0fdafc9bbf15ebe125ace99fffd31677343dc2a276abafbd2f457ad2ecd3630b2c8eda63acdf94971ed1096543cb6476","752830e687b9e9885fbe7b859813578165c23f2bbbd85a933b03109528ee4011","2aec87d1d25389071ae28e132abb4ba116e699c8c2868b807efb99264507c9a2"]}'
peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n auth6 -c '{"Args":["queryAllUsers"]}'
peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n auth6 -c '{"Args":["queryUsersWithPagination","1","g1AAAABUeJzLYWBgYMpgSmHgKy5JLCrJTq2MT8lPzkzJBYoLJZaWZFiVFqcW-SXmplqVZOQXgVRywFTiUJMFADQtGx8"]}'
peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n auth6 -c '{"Args":["queryUserByUserName","魏延"]}'
peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n auth6 -c '{"Args":["queryToMessagesWithPagination","thor","10",""]}'
peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n auth6 -c '{"Args":["queryFromMessagesWithPagination","魏延","10",""]}'
peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n auth6 -c '{"Args":["sendMessage","thor","测试魏延发送给thor","魏延","9e99dcb172483240c1528921ea0d3d54ddbb470fd519bf459f8a4acab175523a","a94c1cc7ee8ffb50ed28602bc55bd8dd8fc0e535ea9a1c7dfea6e7c8051c2f95"]}'

```
