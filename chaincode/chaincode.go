package main

import (
	"chaincode/api"
	"chaincode/model"
	"chaincode/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"time"
)

type MyChaincode struct {
}

func (c *MyChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	// 链码初始化
	fmt.Println("链码初始化")
	//初始化默认数据
	var accountIds = [6]string{
		"5feceb66ffc8",
		"6b86b273ff34",
		"d4735e3a265e",
		"4e07408562be",
		"4b227777d4dd",
		"ef2d127de37b",
	}
	var userNames = [6]string{"管理员", "①号业主", "②号业主", "③号业主", "④号业主", "⑤号业主"}
	var balances = [6]float64{0, 5000000, 5000000, 5000000, 5000000, 5000000}
	//初始化账号数据
	for i, val := range accountIds {
		account := &model.Account{
			AccountId: val,
			UserName:  userNames[i],
			Balance:   balances[i],
		}
		// 写入账本
		if err := utils.WriteLedger(account, stub, model.AccountKey, []string{val}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}
	// 初始化运力数据
	var transportID = []string{
		"t400001001330101330105",
		"t400001001330101330106",
		"t400001001330101330107",
		"t400001001330102330105",
		"t400001001330102330106",
		"t400001001330102330107",
		"t400001001330103330105",
		"t400001001330103330106",
		"t400001001330103330107",
		"t400001001330104330105",
		"t400001001330104330106",
		"t400001001330104330107",
		"t400001002330101330105",
		"t400001002330101330106",
		"t400001002330101330107",
		"t400001002330102330105",
		"t400001002330102330106",
		"t400001002330102330107",
		"t400001002330103330105",
		"t400001002330103330106",
		"t400001002330103330107",
		"t400001002330104330105",
		"t400001002330104330106",
		"t400001002330104330107",
		"t400001003330101330105",
		"t400001003330101330106",
		"t400001003330101330107",
		"t400001003330102330105",
		"t400001003330102330106",
		"t400001003330102330107",
		"t400001003330103330105",
		"t400001003330103330106",
		"t400001003330103330107",
		"t400001003330104330105",
		"t400001003330104330106",
		"t400001003330104330107",
	}
	transports := [36]model.Transport{}
	types := [3]string{"货车", "火车", "货机"}
	startLocationCodes := [4]string{"330101", "330102", "330103", "330104"}
	startLocationNames := [4]string{"S1", "S2", "S3", "S4"}
	endLocationCodes := [4]string{"330105", "330106", "330107"}
	endLocationNames := [4]string{"D1", "D2", "D3"}
	capacities := [3]int{20, 1200, 40}
	tonnages := [3]float64{15, 800, 30}
	transportCounts := [12]int{60, 65, 50, 40, 4, 2, 2, 1, 10, 6, 5, 4}
	costs := [36]int{
		3000, 3200, 2700,
		2400, 2500, 2650,
		2600, 3100, 2700,
		3400, 3000, 2600,
		4500, 4300, 4800,
		3000, 3460, 2600,
		3800, 3500, 3450,
		3850, 3650, 4000,
		18000, 16000, 19000,
		16000, 18000, 15000,
		17000, 16000, 15000,
		19000, 18000, 20000,
	}
	times := [36]float64{
		12.5, 13, 11.5,
		9.5, 9.8, 11.3,
		10.4, 12.5, 11.5,
		14, 12.4, 10.8,
		20.3, 19.1, 22.4,
		13, 15.6, 11.3,
		17.5, 15.6, 15.3,
		17.6, 15.5, 18,
		1.6, 1.4, 1.7,
		1.4, 1.6, 1.3,
		1.5, 1.4, 1.3,
		1.7, 1.6, 1.8,
	}
	for i, _ := range transports {
		transports[i].ObjectType = "transport"
		transports[i].OrgID = "400001"
		transports[i].OrgName = "物流企业01"
		transports[i].TransportStatus = 0
		transports[i].TransportCapacity = capacities[i/12]
		transports[i].TransportTonnage = tonnages[i/12]
		transports[i].TransportType = types[i/12]
		transports[i].StartLocationCode = startLocationCodes[(i/3)%4]
		transports[i].StartLocationName = startLocationNames[(i/3)%4]
		transports[i].EndLocationCode = endLocationCodes[i%3]
		transports[i].EndLocationName = endLocationNames[i%3]
		transports[i].TransportCount = transportCounts[i/3]
		transports[i].TransportCost = costs[i]
		transports[i].TransportDuration = times[i]
	}

	for i, v := range transportID {
		transportString, err := json.Marshal(transports[i])
		if err != nil {
			return shim.Error(fmt.Sprintf("运力数据-序列化json数据失败出错: %s", err))
		}
		if err := stub.PutState(v, transportString); err != nil {
			return shim.Error("运力初始化失败")
		}
	}

	// 初始化仓储
	goodWeights := [2]float64{500, 750}
	goodCodes := [2]string{"6922266470912", "6920118553615"}
	goodNames := [2]string{"清风纸巾", "金味麦片"}
	storageIDs := [8]string{
		"s20000110000169222664709120",
		"s20000110000169201185536150",
		"s20000110000269222664709120",
		"s20000110000269201185536150",
		"s20000110000369222664709120",
		"s20000110000369201185536150",
		"s20000110000469222664709120",
		"s20000110000469201185536150",
	}
	storageCounts := [8]int{
		360, 1050,
		250, 520,
		170, 450,
		120, 310,
	}
	storageUserIDs := [4]string{"200001100001", "200001100002", "200001100003", "200001100004"}
	storageUserNames := [4]string{"仓储01", "仓储02", "仓储03", "仓储04"}
	storageLocationCodes := [4]string{"330101", "330102", "330103", "330104"}
	storageLocationNames := [4]string{"S1", "S2", "S3", "S4"}
	storages := [8]model.Storage{}
	for i := range storageIDs {
		storages[i].ObjectType = "storage"
		storages[i].StorageStatus = 0
		storages[i].OrgID = "200001"
		storages[i].OrgName = "供应机构01"
		storages[i].Count = storageCounts[i]
		storages[i].UserID = storageUserIDs[i/2]
		storages[i].UserName = storageUserNames[i/2]
		storages[i].UserTel = "133xxxx3333"
		storages[i].UserAddress = "xx省xx市xx区xx街道xx"
		storages[i].GoodCode = goodCodes[i%2]
		storages[i].GoodName = goodNames[i%2]
		storages[i].GoodGrossWeight = goodWeights[i%2]
		storages[i].LocationCode = storageLocationCodes[i/2]
		storages[i].LocationName = storageLocationNames[i/2]
	}
	for i, v := range storageIDs {
		storageString, err := json.Marshal(storages[i])
		if err != nil {
			return shim.Error(fmt.Sprintf("仓储数据-序列化json数据失败出错: %s", err))
		}
		if err := stub.PutState(v, storageString); err != nil {
			return shim.Error("仓储初始化失败")
		}
	}

	// 初始化需求

	demandIDs := [6]string{
		"s30000110000169222664709120",
		"s30000110000169201185536150",
		"s30000110000269222664709120",
		"s30000110000269201185536150",
		"s30000110000369222664709120",
		"s30000110000369201185536150",
	}
	demandCounts := [6]int{
		420, 1260,
		200, 660,
		180, 200,
	}
	demandUserIDs := [3]string{"300001100001", "300001100002", "300001100003"}
	demandUserNames := [3]string{"需求点01", "需求点02", "需求点03"}
	demandLocationCodes := [3]string{"330105", "330106", "330107"}
	demandLocationNames := [3]string{"D1", "D2", "D3"}
	demands := [8]model.Demand{}
	for i := range demandIDs {
		demands[i].ObjectType = "demand"
		demands[i].DemandStatus = 0
		demands[i].OrgID = "300001"
		demands[i].OrgName = "需求机构01"
		demands[i].Count = demandCounts[i]
		demands[i].UserID = demandUserIDs[i/2]
		demands[i].UserName = demandUserNames[i/2]
		demands[i].UserTel = "133xxxx3333"
		demands[i].UserAddress = "xx省xx市xx区xx街道xx"
		demands[i].GoodCode = goodCodes[i%2]
		demands[i].GoodName = goodNames[i%2]
		demands[i].GoodGrossWeight = goodWeights[i%2]
		demands[i].LocationCode = demandLocationCodes[i/2]
		demands[i].LocationName = demandLocationNames[i/2]
	}
	for i, v := range demandIDs {
		demandString, err := json.Marshal(demands[i])
		if err != nil {
			return shim.Error(fmt.Sprintf("需求数据-序列化json数据失败出错: %s", err))
		}
		if err := stub.PutState(v, demandString); err != nil {
			return shim.Error("需求初始化失败")
		}
	}
	return shim.Success(nil)
}

func (c *MyChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	case "queryAccountList":
		return api.QueryAccountList(stub, args)
	case "getStorageList":
		return api.GetStorageList(stub, args)
	case "getStoragesByGoodCode":
		return api.GetStoragesByGoodCode(stub, args)
	case "getDemandList":
		return api.GetDemandList(stub, args)
	case "getDemandsByLocation":
		return api.GetDemandsByLocation(stub, args)
	case "getTransportsByLocation":
		return api.GetTransportsByLocation(stub, args)
	case "getProposal":
		return api.GetProposal(stub, args)
	case "getProposalsByUserID":
		return api.GetProposalsByUserID(stub, args)
	case "createProposal":
		return api.CreateProposal(stub, args)
	case "postProposal":
		return api.PostProposal(stub, args)
	case "deleteProposal":
		return api.DeleteProposal(stub, args)
	case "hello":
		return api.Hello(stub, args)
	case "createRealEstate":
		return api.CreateRealEstate(stub, args)
	case "queryRealEstateList":
		return api.QueryRealEstateList(stub, args)
	case "createSelling":
		return api.CreateSelling(stub, args)
	case "createSellingByBuy":
		return api.CreateSellingByBuy(stub, args)
	case "querySellingList":
		return api.QuerySellingList(stub, args)
	case "querySellingListByBuyer":
		return api.QuerySellingListByBuyer(stub, args)
	case "updateSelling":
		return api.UpdateSelling(stub, args)
	case "createDonating":
		return api.CreateDonating(stub, args)
	case "queryDonatingList":
		return api.QueryDonatingList(stub, args)
	case "queryDonatingListByGrantee":
		return api.QueryDonatingListByGrantee(stub, args)
	case "updateDonating":
		return api.UpdateDonating(stub, args)
	default:
		return shim.Error(fmt.Sprintf("没有该方法: %s", funcName))
	}
}

func main() {
	timeLocal, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	time.Local = timeLocal
	err = shim.Start(new(MyChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
