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

type DemandRequestBody struct {
	PageSize     int    `form:"pageSize" json:"pageSize"`
	Bookmark     string `form:"bookmark" json:"bookmark"`
	LocationCode string `form:"location_code" json:"location_code"`
}

func GetDemandList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(DemandRequestBody)
	//解析Body参数
	err := c.ShouldBind(body)
	if err != nil {
		appG.Response(http.StatusBadRequest, "查询需求失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	var resp channel.Response
	if body.LocationCode == "" {
		bodyBytes = append(bodyBytes, []byte(strconv.FormatInt(int64(body.PageSize), 10)))
		bodyBytes = append(bodyBytes, []byte(body.Bookmark))
		resp, err = bc.ChannelQuery("getDemandList", bodyBytes)
		if err != nil {
			appG.Response(http.StatusInternalServerError, "查询所有需求失败", err.Error())
			return
		}
		var data map[string]interface{}
		if err = json.Unmarshal(resp.Payload, &data); err != nil {
			appG.Response(http.StatusInternalServerError, "查询所有需求失败", err.Error())
			return
		}
		appG.Response(http.StatusOK, "查询所有需求成功", data)
	} else {
		bodyBytes = append(bodyBytes, []byte(body.LocationCode))
		resp, err = bc.ChannelQuery("getDemandsByLocation", bodyBytes)
		if err != nil {
			appG.Response(http.StatusInternalServerError, "按地点查询需求失败", err.Error())
			return
		}
		var data []map[string]interface{}
		if err = json.Unmarshal(resp.Payload, &data); err != nil {
			appG.Response(http.StatusInternalServerError, "按地点查询需求反序列化失败", err.Error())
			return
		}
		appG.Response(http.StatusOK, "按地点查询需求成功", data)
	}
}
