package Redis

import (
	"FastWiki/Model"
	"context"
	"encoding/json"
	"time"
)

// FlushMenuLv1 刷新一级菜单缓存
func FlushMenuLv1(p *[]Model.Community) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	pipeline := rdb.TxPipeline()
	pipeline.Del(ctx, "MenuLv1")
	menuLv1Byte, err1 := json.Marshal(p)
	if err1 != nil {
		return err1
	}
	menuLv1Str := string(menuLv1Byte)
	pipeline.RPush(ctx, "MenuLv1", menuLv1Str)
	_, err = pipeline.Exec(ctx)
	defer cancel()
	return err
}

// FlushMenuLv1 刷新二级菜单缓存
func FlushMenuLv2(p *[]Model.Seccommunity, s string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	pipeline := rdb.TxPipeline()
	pipeline.Del(ctx, s)
	menuLv2Byte, err1 := json.Marshal(p)
	if err1 != nil {
		return err1
	}
	menuLv2Str := string(menuLv2Byte)
	pipeline.RPush(ctx, s, menuLv2Str)
	_, err = pipeline.Exec(ctx)
	defer cancel()
	return err
}
