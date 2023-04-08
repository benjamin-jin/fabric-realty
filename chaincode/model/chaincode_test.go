package model

//docker exec cli bash -c "CORE_PEER_ADDRESS=peer0.taobao.com:7051 CORE_PEER_LOCALMSPID=TaobaoMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/taobao.com/users/Admin@taobao.com/msp peer chaincode invoke -C appchannel -n fabric-realty -c '{\"Args\":[\"queryAccountList\", \"4e07408562be\"]}'"

//docker exec cli bash -c "CORE_PEER_ADDRESS=peer0.taobao.com:7051 CORE_PEER_LOCALMSPID=TaobaoMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/taobao.com/users/Admin@taobao.com/msp peer chaincode invoke -C appchannel -n fabric-realty -c '{\"Args\":[\"getDemandList\", \"10\", \"\"]}'"

//docker exec cli bash -c "CORE_PEER_ADDRESS=peer0.taobao.com:7051 CORE_PEER_LOCALMSPID=TaobaoMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/taobao.com/users/Admin@taobao.com/msp peer chaincode invoke -C appchannel -n fabric-realty -c '{\"Args\":[\"getDemandsByLocation\", \"330105\"]}'"

//docker exec cli bash -c "CORE_PEER_ADDRESS=peer0.taobao.com:7051 CORE_PEER_LOCALMSPID=TaobaoMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/taobao.com/users/Admin@taobao.com/msp peer chaincode invoke -C appchannel -n fabric-realty -c '{\"Args\":[\"getStorageList\", \"5\", \"\"]}'"
//
//docker exec cli bash -c "CORE_PEER_ADDRESS=peer0.taobao.com:7051 CORE_PEER_LOCALMSPID=TaobaoMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/taobao.com/users/Admin@taobao.com/msp peer chaincode invoke -C appchannel -n fabric-realty -c '{\"Args\":[\"getStoragesByGoodCode\", \"6922266470912\"]}'"
//
//docker exec cli bash -c "CORE_PEER_ADDRESS=peer0.taobao.com:7051 CORE_PEER_LOCALMSPID=TaobaoMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/taobao.com/users/Admin@taobao.com/msp peer chaincode invoke -C appchannel -n fabric-realty -c '{\"Args\":[\"getTransportByLocation\", \"330101\", \"330105\"]}'"

//docker exec cli bash -c "CORE_PEER_ADDRESS=peer0.taobao.com:7051 CORE_PEER_LOCALMSPID=TaobaoMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/taobao.com/users/Admin@taobao.com/msp peer chaincode invoke -C appchannel -n fabric-realty -c '{\"Args\":[\"testBigEvent\"]}'"

//docker exec cli bash -c "CORE_PEER_ADDRESS=peer0.taobao.com:7051  CORE_PEER_LOCALMSPID=TaobaoMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/taobao.com/users/Admin@taobao.com/msp peer chaincode invoke -C appchannel -n fabric-realty --peerAddresses peer0.taobao.com:7051 -c '{\"Args\":[\"testAddOne\", \"test000001\"]}'"

//docker exec cli bash -c "CORE_PEER_ADDRESS=peer0.taobao.com:7051 CORE_PEER_LOCALMSPID=TaobaoMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/taobao.com/users/Admin@taobao.com/msp peer chaincode invoke -C appchannel -n fabric-realty -c '{\"Args\":[\"testAddOneCache\", \"test000001\"]}'"

//docker exec cli bash -c "CORE_PEER_ADDRESS=peer0.taobao.com:7051 CORE_PEER_LOCALMSPID=TaobaoMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/taobao.com/users/Admin@taobao.com/msp peer chaincode invoke -C appchannel -n fabric-realty -c '{\"Args\":[\"testWriteBack\"]}'"

//docker exec cli bash -c "CORE_PEER_ADDRESS=peer0.taobao.com:7051 CORE_PEER_LOCALMSPID=TaobaoMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/taobao.com/users/Admin@taobao.com/msp peer chaincode invoke -C appchannel -n fabric-realty -c '{\"Args\":[\"createProposal\", \"p123456\", \"{"create":"proposal"\"]}'"
