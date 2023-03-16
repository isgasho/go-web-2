package utools

import (
	"encoding/json"
	"fmt"
	"go-web/common"
)

// Struct2Json 结构体转为 JSON
func Struct2Json(obj interface{}) string {
	jsonStr, err := json.Marshal(obj)
	if err != nil {
		common.Logger.Error(fmt.Sprintf("结构体转JSON发送异常 [Struct2Json]：%v", err))
	}
	return string(jsonStr)
}

// Json2Struct JSON 转为结构体
func Json2Struct(jsonStr string, obj interface{}) {
	err := json.Unmarshal([]byte(jsonStr), obj)
	if err != nil {
		common.Logger.Error(fmt.Sprintf("JSON转结构体发送异常 [Json2Struct]：%v", err))
	}
}

// JsonInterface2Struct JSON Interface 转为结构体
func JsonInterface2Struct(str interface{}, obj interface{}) {
	jsonStr, _ := str.(string)
	Json2Struct(jsonStr, obj)
}

// Struct2StructByJson 结构体转结构体, JSON 为中间桥梁, struct2 必须以指针方式传递, 否则可能获取到空数据
func Struct2StructByJson(struct1 interface{}, struct2 interface{}) {
	jsonStr := Struct2Json(struct1)
	Json2Struct(jsonStr, struct2)
}
