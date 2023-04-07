package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	pb "github.com/hyperledger/fabric/protos/peer"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// WriteLedger 写入账本
func WriteLedger(obj interface{}, stub shim.ChaincodeStubInterface, objectType string, keys []string) error {
	//创建复合主键
	var key string
	if val, err := stub.CreateCompositeKey(objectType, keys); err != nil {
		return errors.New(fmt.Sprintf("%s-创建复合主键出错 %s", objectType, err))
	} else {
		key = val
	}
	bytes, err := json.Marshal(obj)
	if err != nil {
		return errors.New(fmt.Sprintf("%s-序列化json数据失败出错: %s", objectType, err))
	}
	//写入区块链账本
	if err := stub.PutState(key, bytes); err != nil {
		return errors.New(fmt.Sprintf("%s-写入区块链账本出错: %s", objectType, err))
	}
	return nil
}

// DelLedger 删除账本
func DelLedger(stub shim.ChaincodeStubInterface, objectType string, keys []string) error {
	//创建复合主键
	var key string
	if val, err := stub.CreateCompositeKey(objectType, keys); err != nil {
		return errors.New(fmt.Sprintf("%s-创建复合主键出错 %s", objectType, err))
	} else {
		key = val
	}
	//写入区块链账本
	if err := stub.DelState(key); err != nil {
		return errors.New(fmt.Sprintf("%s-删除区块链账本出错: %s", objectType, err))
	}
	return nil
}

// GetStateByPartialCompositeKeys 根据复合主键查询数据(适合获取全部，多个，单个数据)
// 将keys拆分查询
func GetStateByPartialCompositeKeys(stub shim.ChaincodeStubInterface, objectType string, keys []string) (results [][]byte, err error) {
	if len(keys) == 0 {
		// 传入的keys长度为0，则查找并返回所有数据
		// 通过主键从区块链查找相关的数据，相当于对主键的模糊查询
		resultIterator, err := stub.GetStateByPartialCompositeKey(objectType, keys)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%s-获取全部数据出错: %s", objectType, err))
		}
		defer resultIterator.Close()

		//检查返回的数据是否为空，不为空则遍历数据，否则返回空数组
		for resultIterator.HasNext() {
			val, err := resultIterator.Next()
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-返回的数据出错: %s", objectType, err))
			}

			results = append(results, val.GetValue())
		}
	} else {
		// 传入的keys长度不为0，查找相应的数据并返回
		for _, v := range keys {
			// 创建组合键
			key, err := stub.CreateCompositeKey(objectType, []string{v})
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-创建组合键出错: %s", objectType, err))
			}
			// 从账本中获取数据
			bytes, err := stub.GetState(key)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-获取数据出错: %s", objectType, err))
			}

			if bytes != nil {
				results = append(results, bytes)
			}
		}
	}

	return results, nil
}

// GetStateByPartialCompositeKeys2 根据复合主键查询数据(适合获取全部或指定的数据)
func GetStateByPartialCompositeKeys2(stub shim.ChaincodeStubInterface, objectType string, keys []string) (results [][]byte, err error) {
	// 通过主键从区块链查找相关的数据，相当于对主键的模糊查询
	resultIterator, err := stub.GetStateByPartialCompositeKey(objectType, keys)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s-获取全部数据出错: %s", objectType, err))
	}
	defer resultIterator.Close()

	//检查返回的数据是否为空，不为空则遍历数据，否则返回空数组
	for resultIterator.HasNext() {
		val, err := resultIterator.Next()
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%s-返回的数据出错: %s", objectType, err))
		}

		results = append(results, val.GetValue())
	}
	return results, nil
}

func GetQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	//fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}

func GetQueryResultForQueryStringWithPagination(stub shim.ChaincodeStubInterface, queryString string, pageSize int32, bookmark string) ([]byte, error) {

	//fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, responseMetadata, err := stub.GetQueryResultWithPagination(queryString, pageSize, bookmark)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	bufferWithPaginationInfo := addPaginationMetadataToQueryResults(buffer, responseMetadata)

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", bufferWithPaginationInfo.String())

	return buffer.Bytes(), nil
}

func addPaginationMetadataToQueryResults(buffer *bytes.Buffer, responseMetadata *pb.QueryResponseMetadata) *bytes.Buffer {

	buffer.WriteString("[{\"ResponseMetadata\":{\"RecordsCount\":")
	buffer.WriteString("\"")
	buffer.WriteString(fmt.Sprintf("%v", responseMetadata.FetchedRecordsCount))
	buffer.WriteString("\"")
	buffer.WriteString(", \"Bookmark\":")
	buffer.WriteString("\"")
	buffer.WriteString(responseMetadata.Bookmark)
	buffer.WriteString("\"}}]")

	return buffer
}
