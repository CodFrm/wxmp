package api

import (
	"github.com/CodFrm/wxmp/internal/dao"
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v7"
	"testing"
)

var wechat *Wechat

func TestMain(m *testing.M) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	wechat = &Wechat{
		token: dao.New(&dao.DaoConfig{
			Redis: &redis.Options{
				Addr: mr.Addr(),
			}}),
	}
	m.Run()
}
