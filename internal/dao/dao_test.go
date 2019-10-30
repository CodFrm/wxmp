package dao

import (
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v7"
	"testing"
)

var dao *Dao

func TestMain(m *testing.M) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	dao = New(&DaoConfig{
		Redis: &redis.Options{
			Addr: mr.Addr(),
		}})
	m.Run()
}
