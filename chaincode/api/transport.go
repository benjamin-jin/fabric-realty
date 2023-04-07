package api

import (
	"chaincode/pkg/utils"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//func GetTransportList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
//	if len(args) < 2 {
//		return shim.Error("参数数量错误，需要2个参数")
//	}
//	pageSize, err := strconv.ParseInt(args[0], 10, 32)
//	if err != nil {
//		return shim.Error(err.Error())
//	}
//	bookmark := args[1]
//	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"storage\"}}")
//	queryResults, err := utils.GetQueryResultForQueryStringWithPagination(stub, queryString, int32(pageSize), bookmark)
//
//	if err != nil {
//		return shim.Error(err.Error())
//	}
//
//	return shim.Success(queryResults)
//}

func GetTransportsByLocation(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("参数数量错误，需要2个参数")
	}
	startLocationCode, endLocationCode := args[0], args[1]
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"transport\",\"start_location_code\":\"%s\",\"end_location_code\":\"%s\"}}", startLocationCode, endLocationCode)
	queryResults, err := utils.GetQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}
