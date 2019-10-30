package api

import (
	"github.com/CodFrm/wxmp/internal/dao"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
)

func Handel(r *gin.Engine) *gin.Engine {
	d := dao.New(&dao.DaoConfig{Redis: &redis.Options{
		Addr: "127.0.0.1:6379",
	}})
	NewWechat(r, d)

	return r
}
