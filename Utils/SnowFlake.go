package Utils

import (
	"FastWiki/Setting"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
	"time"
)

/*
雪花算法生成用户id
*/

var node *snowflake.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return
}

// GenID生成ID
func GenID() (id int64) {
	if err := Init(Setting.Conf.StartTime, Setting.Conf.MachingeID); err != nil {
		zap.L().Error("init snowflake failed,err:%v\n", zap.Error(err))
		return 0
	}
	return node.Generate().Int64()
}
