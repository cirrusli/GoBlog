package utils

import (
	"SummerProject/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// GetRequestJsonParams 获取Json参数
func GetRequestJsonParams(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("parse form failed:", err.Error())
	}
	err = json.Unmarshal(body, &params)
	if err != nil {
		log.Println("Decode error:", err.Error())
	}
	return params
}

// Error 返回错误信息
func Error(w http.ResponseWriter, err error) {
	var resmsg model.ResMsg
	resmsg.Code = "0000"
	resmsg.Msg = err.Error()
	resJson, _ := json.Marshal(resmsg)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resJson)
	if err != nil {
		log.Println(err)
	}

}

// Success 返回成功信息
func Success(w http.ResponseWriter, data interface{}) {
	var resmsg model.ResMsg
	resmsg.Code = "0001"
	resmsg.Msg = "success"
	resmsg.Data = data
	resJson, _ := json.Marshal(resmsg) //没有tag的结构体字段无法传入Json
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(resJson)
	if err != nil {
		log.Println(err)
	}
}
