package MySql

import (
	"FastWiki/Setting"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var dbEngine *xorm.Engine

func Init(cfg *Setting.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	fmt.Println(dsn)
	dbEngine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		zap.L().Error("Xorm link mysql faild", zap.Error(err))
		return
	}
	//添加统一前缀
	tbMapper := names.NewPrefixMapper(names.SnakeMapper{}, "tb_")
	dbEngine.SetTableMapper(tbMapper)
	//连接测试
	if err := dbEngine.Ping(); err != nil {
		zap.L().Error("Xorm Ping mysql faild", zap.Error(err))
		return err
	}
	//defer dbEngine.Close()
	return
}

func CloseEngine() {
	dbEngine.Close()

}
