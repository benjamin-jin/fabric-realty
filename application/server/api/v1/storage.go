package v1

import (
	bc "application/blockchain"
	"application/pkg/app"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"net/http"
	"strconv"
)

type StorageRequestBody struct {
	PageSize int    `json:"pageSize" form:"pageSize"`
	Bookmark string `json:"bookmark" form:"bookmark"`
	GoodCode string `json:"good_code" form:"good_code"`
}

func GetStorageList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(StorageRequestBody)
	//解析Body参数
	err := c.ShouldBind(body)
	if err != nil {
		appG.Response(http.StatusBadRequest, "查询仓储失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	var resp channel.Response
	if body.GoodCode == "" {
		bodyBytes = append(bodyBytes, []byte(strconv.FormatInt(int64(body.PageSize), 10)))
		bodyBytes = append(bodyBytes, []byte(body.Bookmark))
		resp, err = bc.ChannelQuery("getStorageList", bodyBytes)
		if err != nil {
			appG.Response(http.StatusInternalServerError, "查询所有仓储失败", err.Error())
			return
		}
		var data map[string]interface{}
		if err = json.Unmarshal(resp.Payload, &data); err != nil {
			appG.Response(http.StatusInternalServerError, "查询仓储失败", err.Error())
			return
		}
		appG.Response(http.StatusOK, "查询仓储成功", data)
	} else {
		bodyBytes = append(bodyBytes, []byte(body.GoodCode))
		resp, err = bc.ChannelQuery("getStoragesByGoodCode", bodyBytes)
		if err != nil {
			appG.Response(http.StatusInternalServerError, "按物资查询仓储失败", err.Error())
			return
		}
		var data []map[string]interface{}
		if err = json.Unmarshal(resp.Payload, &data); err != nil {
			appG.Response(http.StatusInternalServerError, "查询仓储反序列化失败", err.Error())
			return
		}
		appG.Response(http.StatusOK, "查询仓储成功", data)
	}
	// todo: 反序列化json 可能存在问题

}
