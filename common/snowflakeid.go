package common

import (
	"github.com/bwmarrin/snowflake"
	"log"
	"time"
)

var node *snowflake.Node

func InitSnow() {
	const (
		startTime = "2003-06-17"
		machineID = 19
	)
	var st time.Time

	// 格式化 1月2号下午3时4分5秒  2006年
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		log.Println(err)
		return
	}

	snowflake.Epoch = st.UnixNano() / 1e6
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

// EncodeID 生成64位的雪花ID,转换为int
func EncodeID() int {
	return int(node.Generate().Int64())
}
