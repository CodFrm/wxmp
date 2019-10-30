package api

import (
	"github.com/CodFrm/wxmp/internal/dao"
	"github.com/gin-gonic/gin"
)

func Handel(r *gin.Engine, d *dao.Dao) *gin.Engine {

	NewWechat(r, d)

	return r
}
