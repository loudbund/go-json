package main

import (
	"fmt"
	"github.com/loudbund/go-json/json_v1"
	log "github.com/sirupsen/logrus"
)

func main() {
	// map格式数据
	J := map[string]interface{}{
		"haha": map[string]interface{}{
			"YY": []interface{}{
				map[string]interface{}{
					"number":       123,
					"stringNumber": "456",
					"name":         "I'm here!",
				},
				map[string]interface{}{
					"DFloat": 123.123,
				},
			},
		},
	}
	// 示例1、map格式数据转换成字符串
	SData, err := json_v1.JsonEncode(J)
	if err != nil {
		log.Panic(err)
	}

	// 示例2、反解开
	JData, err := json_v1.JsonDecode(SData)
	if err != nil {
		log.Panic(err)
	}

	// 示例3、获取不存在的节点
	if D1, err := json_v1.GetJsonInterface(JData, "wawa"); err != nil {
		log.Error("查找 JData.wawa 失败 ：", err)
	} else {
		fmt.Println("找到 JData.wawa", D1)
	}

	// 示例4、获取存在的节点
	if D1, err := json_v1.GetJsonInterface(JData, "haha"); err != nil {
		log.Error("查找 JData.haha 失败 ：", err)
	} else {
		fmt.Println("找到 JData.haha :", D1)
	}

	// 示例5、查找字符串节点
	if D1, err := json_v1.GetJsonString(JData, "haha", "YY", 0, "name"); err != nil {
		log.Error("查找 JData.haha.YY[0].name 失败 ：", err)
	} else {
		fmt.Println("找到 JData.haha.YY[0].name :", D1)
	}

	// 示例6、查找int64节点
	if D1, err := json_v1.GetJsonInt64(JData, "haha", "YY", 0, "number"); err != nil {
		log.Error("查找 JData.haha.YY[0].number 失败 ：", err)
	} else {
		fmt.Println("找到 JData.haha.YY[0].number :", D1)
	}
	// 示例7、查找数字字符节点，转int64数据
	if D1, err := json_v1.GetJsonInt64Force(JData, "haha", "YY", 0, "stringNumber"); err != nil {
		log.Error("查找 JData.haha.YY[0].stringNumber 失败 ：", err)
	} else {
		fmt.Println("找到 JData.haha.YY[0].stringNumber :", D1)
	}

	// 示例8、查找字符串
	if D1, err := json_v1.GetJsonString(JData, "haha", "YY", 1, "DFloat"); err != nil {
		log.Error("查找 JData.haha.YY[1].DFloat 失败 ：", err)
	} else {
		fmt.Println("找到 JData.haha.YY[1].DFloat :", D1)
	}
}
