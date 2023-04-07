package api

import (
	"chaincode/pkg/utils"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func GetProposal(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("参数数量错误，需要1个参数")
	}
	proposalID := args[0]
	proposalString, err := stub.GetState(proposalID)
	if err != nil {
		return shim.Error(fmt.Sprintf("获取方案 %s 失败", proposalID))
	} else if proposalString == nil {
		return shim.Error(fmt.Sprintf("方案 %s 不存在", proposalID))
	}
	return shim.Success(proposalString)
}

func GetProposalsByUserID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("参数数量错误，需要1个参数")
	}
	userID := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"proposal\",\"admin_user_id\":\"%s\"}}", userID)
	queryResults, err := utils.GetQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func CreateProposal(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("参数数量错误，需要2个参数")
	}
	proposalID, proposalString := args[0], args[1]
	oldProposal, err := stub.GetState(proposalID)
	if err != nil {
		return shim.Error("获取方案失败: " + err.Error())
	} else if oldProposal != nil {
		return shim.Error("方案已存在: " + proposalID)
	}

	err = stub.PutState(proposalID, []byte(proposalString))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func PostProposal(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("参数数量错误，需要2个参数")
	}
	proposalID, proposalString := args[0], args[1]
	oldProposal, err := stub.GetState(proposalID)
	if err != nil {
		return shim.Error("获取方案失败: " + err.Error())
	} else if oldProposal == nil {
		return shim.Error("方案不存在: " + proposalID)
	}

	err = stub.PutState(proposalID, []byte(proposalString))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func DeleteProposal(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("参数数量错误，需要2个参数")
	}
	proposalID := args[0]
	oldProposal, err := stub.GetState(proposalID)
	if err != nil {
		return shim.Error("获取方案失败: " + err.Error())
	} else if oldProposal == nil {
		return shim.Error("方案不存在: " + proposalID)
	}
	err = stub.DelState(proposalID)
	if err != nil {
		return shim.Error("删除方案失败" + err.Error())
	}
	return shim.Success(nil)
}
