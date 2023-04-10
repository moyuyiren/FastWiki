package main

import (
	"FastWiki/Controller/reciprocal"
	"FastWiki/Dao/ElasticSearch"
	"FastWiki/Dao/MySql"
	"FastWiki/Dao/Redis"
	"FastWiki/Logger"
	"FastWiki/RabbitMQ"
	"FastWiki/RouterS"
	"FastWiki/Setting"
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//1.加载配置
	if err := Setting.Init(); err != nil {
		fmt.Printf("配置加载失败")
		return
	}
	//2.初始化日志模块
	if err := Logger.Init(Setting.Conf.Logconfig, Setting.Conf.Mode); err != nil {
		fmt.Println("Logger Init fail")
		return
	}
	//3.初始化Mysql连接
	if err := MySql.Init(Setting.Conf.MysqlConfig); err != nil {
		zap.L().Error("数据库连接异常，err:", zap.Error(err))
		return
	}
	defer MySql.CloseEngine()
	//4.初始化Redis连接
	if err := Redis.Init(Setting.Conf.RedisConfig); err != nil {
		zap.L().Error("Redis数据库链接失败：,", zap.Error(err))
		return
	}
	defer Redis.Close()
	//5.初始化ElasticSearch连接
	if err := ElasticSearch.Init(Setting.Conf.ElaSearchConfig); err != nil {
		zap.L().Error("ElasticSearch链接失败", zap.Error(err))
		return
	}
	//6.初始化RabbitMQ
	if err := RabbitMQ.InitRabbitMq(); err != nil {
		zap.L().Error("RabbitMQ 连接失败：,", zap.Error(err))
		return
	}
	defer RabbitMQ.Close()

	//7.初始化路由，配置自己的日志收集工具替换Gin原生的日志工具
	r := RouterS.Setup(Setting.Conf.Mode)
	r.Run(":8088")

	//8.初始化翻译器
	if err := reciprocal.InitTrans("zh"); err != nil {
		zap.L().Error("init Translator failed, err:%v\n", zap.Error(err))
		return
	}

	//10.启用优雅关机
	srv := &http.Server{
		Addr:    fmt.Sprintf("%d", Setting.Conf.Port),
		Handler: r,
	}
	/*开启一个GO协程启动服务*/
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}
	zap.L().Info("Server exiting")

}
