package api

import (
	"github.com/CodFrm/wxmp/internal/wchat"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Validate(c *gin.Context) {
	server := wchat.WChat.GetServer(c.Request, c.Writer)
	if server.Validate() {
		c.String(http.StatusOK, c.Param("nonce"))
	} else {
		c.String(http.StatusOK, "error")
	}
}
