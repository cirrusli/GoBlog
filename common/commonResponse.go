package common

import (
	"GoBlog/internal/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// GetRequestJsonParams 获取POST请求body中的JSON参数
func GetRequestJsonParams(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	ErrorLog("GetRequestJsonParams: parse form failed:", err)
	log.Println("GetRequestJsonParams: 请求体的数据:", r.Body)
	err = json.Unmarshal(body, &params)
	ErrorLog("GetRequestJsonParams: Decode error:", err)

	return params
}

// Error 返回错误信息
func Error(w http.ResponseWriter, err error) {
	var resmsg model.ResMsg
	resmsg.Code = "0000"
	resmsg.Msg = err.Error()
	resJson, err := json.Marshal(resmsg)
	ErrorLog("Error: JSON marshal failed:", err)

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resJson)
	ErrorLog("Error: write json data failed:", err)

}

// Success 返回成功信息
func Success(w http.ResponseWriter, data interface{}) {
	var resmsg model.ResMsg
	resmsg.Code = "0001"
	resmsg.Msg = "success"
	resmsg.Data = data
	resJson, err := json.Marshal(resmsg) //没有tag的结构体字段无法传入Json
	ErrorLog("Success: JSON marshal failed:", err)

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resJson)
	ErrorLog("Success: write json data failed:", err)

}

// ErrorLog 普通错误
func ErrorLog(msg string, err error) {
	if err != nil {
		log.Println(msg, err.Error())
	}
}

// PanicLog 崩溃错误
func PanicLog(msg string, err error) {
	if err != nil {
		log.Panicln(msg, err.Error())
	}
}
