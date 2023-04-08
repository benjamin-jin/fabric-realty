package api

import (
	"chaincode/model"
	"chaincode/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"time"
)

func TestBigEvent(stub shim.ChaincodeStubInterface) pb.Response {
	var testIds = [20]string{
		"test000001",
		"test000002",
		"test000003",
		"test000004",
		"test000005",
		"test000006",
		"test000007",
		"test000008",
		"test000009",
		"test000010",
		"test000011",
		"test000012",
		"test000013",
		"test000014",
		"test000015",
		"test000016",
		"test000017",
		"test000018",
		"test000019",
		"test000020",
	}
	var testDataList []model.TestData
	for _, v := range testIds {
		testData := model.TestData{}
		testDataString, err := stub.GetState(v)
		if err != nil {
			return shim.Error(fmt.Sprintf("获取测试数据 %s 失败", v))
		} else if testDataString == nil {
			return shim.Error(fmt.Sprintf("测试数据 %s 不存在", v))
		}
		err = json.Unmarshal(testDataString, &testData)
		if err != nil {
			return shim.Error(fmt.Sprintf("测试数据 %s 反序列化失败", v))
		}
		testDataList = append(testDataList, testData)
	}
	for i := range testDataList {
		testDataList[i].Balance += 100000
		testDataString, err := json.Marshal(testDataList[i])
		if err != nil {
			return shim.Error(fmt.Sprintf("测试数据大事务-序列化json数据失败出错: %s", err))
		}
		if err = stub.PutState(testDataList[i].Key, testDataString); err != nil {
			return shim.Error("测试数据初始化失败")
		}
		time.Sleep(200 * time.Millisecond)
	}

	return shim.Success(nil)
}

func TestAddOne(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("参数数量错误，需要1个参数")
	}
	testID := args[0]
	old, err := stub.GetState(testID)
	if err != nil {
		return shim.Error("获取测试数据失败: " + err.Error())
	} else if old == nil {
		return shim.Error("测试数据不存在: " + testID)
	}
	testData := model.TestData{}
	err = json.Unmarshal(old, &testData)
	if err != nil {
		return shim.Error("测试数据反序列化失败: " + testID)
	}
	testData.Balance += 1
	var testDataString []byte
	testDataString, err = json.Marshal(testData)
	if err != nil {
		return shim.Error("测试数据序列化失败: " + testID)
	}
	err = stub.PutState(testID, testDataString)
	if err != nil {
		return shim.Error("测试数据更新失败： " + err.Error())
	}
	return shim.Success(nil)
}

func TestAddOneCache(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("参数数量错误，需要1个参数")
	}
	testID := args[0]
	key := "cache" + testID
	testData := model.TestData{}
	old, err := stub.GetState(testID)
	if err != nil {
		return shim.Error("获取测试数据失败: " + err.Error())
	} else if old == nil {
		testData = model.TestData{
			Key:        testID,
			ObjectType: "cachetest",
			UserName:   testID,
			Balance:    0,
		}
	} else {
		err = json.Unmarshal(old, &testData)
		if err != nil {
			return shim.Error("测试数据反序列化失败: " + testID)
		}
	}

	testData.Balance += 1
	var testDataString []byte
	testDataString, err = json.Marshal(testData)
	if err != nil {
		return shim.Error("测试数据序列化失败: " + testID)
	}
	err = stub.PutState(key, testDataString)
	if err != nil {
		return shim.Error("测试数据更新失败： " + err.Error())
	}
	return shim.Success(nil)
}

func TestWriteBack(stub shim.ChaincodeStubInterface) pb.Response {
	var testDataList []model.TestData
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"cachetest\"}")
	queryResults, err := utils.GetQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error("获取测试数据缓存列表失败: " + err.Error())
	}
	err = json.Unmarshal(queryResults, &testDataList)
	if err != nil {
		return shim.Error("测试数据缓存反序列化失败: " + err.Error())
	}
	for i := range testDataList {
		testDataList[i].ObjectType = "test"
		var old, testDataString []byte
		old, err = stub.GetState(testDataList[i].Key)
		if err != nil {
			return shim.Error("获取测试数据失败: " + err.Error())
		} else if old != nil {
			oldData := model.TestData{}
			err = json.Unmarshal(old, &oldData)
			if err != nil {
				return shim.Error("测试数据缓存元数据反序列化失败: " + err.Error())
			}
			testDataList[i].Balance += oldData.Balance
		}
		testDataString, err = json.Marshal(testDataList[i])
		if err != nil {
			return shim.Error("测试数据缓存序列化失败: " + testDataList[i].Key)
		}
		err = stub.PutState(testDataList[i].Key, testDataString)
		if err != nil {
			return shim.Error("测试数据重写失败: " + testDataList[i].Key)
		}
	}
	return shim.Success(nil)
}