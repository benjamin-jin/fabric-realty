package api

import (
	"chaincode/pkg/utils"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

func GetDemandList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("参数数量错误，需要2个参数")
	}
	pageSize, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		return shim.Error(err.Error())
	}
	bookmark := args[1]
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"demand\"}}")
	queryResults, err := utils.GetQueryResultForQueryStringWithPagination(stub, queryString, int32(pageSize), bookmark)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func GetDemandsByLocation(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("参数数量错误，需要1个参数")
	}
	locationCode := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"demand\",\"location_code\":\"%s\"}}", locationCode)
	queryResults, err := utils.GetQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}
