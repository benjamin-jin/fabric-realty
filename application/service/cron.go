/**
 * @Author: 夜央 Oh oh oh oh oh oh (https://github.com/togettoyou)
 * @Email: zoujh99@qq.com
 * @Date: 2020/3/13 1:01 上午
 * @Description: 定时任务-每天晚上0点扫描一下处理过期交易
 */
package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	bc "github.com/togettoyou/blockchain-real-estate/application/blockchain"
	"github.com/togettoyou/blockchain-real-estate/application/lib"
	"log"
	"time"
)

const spec = "0 0 0 * * ?" // 每天0点执行
//const spec = "*/10 * * * * ?" //10秒执行一次，用于测试

func Init() {
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc(spec, GoRun)
	if err != nil {
		log.Printf("start cron failed %s", err)
	}
	c.Start()
	select {}
}

func GoRun() {
	log.Printf("定时任务已启动")
	//先把所有销售查询出来
	resp, err := bc.ChannelQuery("querySellingList", [][]byte{}) //调用智能合约
	if err != nil {
		log.Printf("定时任务-querySellingList失败%s", err.Error())
		return
	}
	// 反序列化json
	var data []lib.Selling
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		log.Printf("定时任务-反序列化json失败%s", err.Error())
		return
	}
	for _, v := range data {
		//把状态为销售中和交付中的筛选出来
		if v.SellingStatus == lib.SellingStatusConstant()["saleStart"] ||
			v.SellingStatus == lib.SellingStatusConstant()["delivery"] {
			//有效期天数
			day, _ := time.ParseDuration(fmt.Sprintf("%dh", v.SalePeriod*24))
			vTime := v.CreateTime.Add(day)
			//如果 time.Now()大于 vTime 说明过期
			if time.Now().After(vTime) {
				//将状态更改为已过期
				var bodyBytes [][]byte
				bodyBytes = append(bodyBytes, []byte(v.ObjectOfSale))
				bodyBytes = append(bodyBytes, []byte(v.Seller))
				bodyBytes = append(bodyBytes, []byte(v.Buyer))
				bodyBytes = append(bodyBytes, []byte("expired"))
				//调用智能合约
				resp, err := bc.ChannelExecute("updateSelling", bodyBytes)
				if err != nil {
					return
				}
				var data map[string]interface{}
				if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
					return
				}
				fmt.Println(data)
			}
		}
	}
}
