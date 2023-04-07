package v1

import (
	bc "application/blockchain"
	"application/pkg/app"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"net/http"
)

type TransportRequestBody struct {
	StartLocationCode string `json:"start_location_code" form:"start_location_code"`
	EndLocationCode   string `json:"end_location_code" form:"end_location_code"`
}

func GetTransportList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(TransportRequestBody)
	//解析Body参数
	err := c.ShouldBind(body)
	if err != nil {
		appG.Response(http.StatusBadRequest, "查询运力失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	var resp channel.Response
	bodyBytes = append(bodyBytes, []byte(body.StartLocationCode))
	bodyBytes = append(bodyBytes, []byte(body.EndLocationCode))
	resp, err = bc.ChannelQuery("getTransportsByLocation", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "按地点查询运力失败", err.Error())
		return
	}

	// todo: 反序列化json 可能存在问题
	var data []map[string]interface{}
	if err = json.Unmarshal(resp.Payload, &data); err != nil {
		appG.Response(http.StatusInternalServerError, "查询运力失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "查询运力成功", data)
}
