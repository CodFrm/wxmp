package api

import (
	"github.com/gin-gonic/gin"
)

func Handel(r *gin.Engine) *gin.Engine {
	r.Any("/wchat", Wchat)

	return r
}
