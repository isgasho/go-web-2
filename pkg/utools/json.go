package utools

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"go-web/common"
)

// Json 转结构体网站：https://app.quicktype.io/
// Json 解析常用库：https://www.cnblogs.com/luozhiyun/p/14875066.html

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

// Struct2MapStringInterface 将结构体转换成 map[string]interface
func Struct2MapStringInterface(struct1 interface{}) map[string]interface{} {
	return structs.Map(struct1)
}

// MapStringInterface2Struct 将结构体转换成 map[string]interface，第二个参数用指针
func MapStringInterface2Struct(data map[string]interface{}, obj interface{}) {
	jsonStr, _ := json.Marshal(data)
	err := json.Unmarshal(jsonStr, obj)
	if err != nil {
		common.Logger.Error("map[string]interface 转结构体转换失败：", err.Error())
	}
}
