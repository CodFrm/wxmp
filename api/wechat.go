package api

import (
	"github.com/CodFrm/wxmp/internal/wchat"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/message"
)

func Wchat(c *gin.Context) {
	server := wchat.WChat.GetServer(c.Request, c.Writer)

	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		if msg.MsgType == message.MsgTypeText {
			if msg.Content == "token" {
				return &message.Reply{
					MsgType: message.MsgTypeText,
					MsgData: message.NewText("test"),
				}
			}
		}

		return &message.Reply{
			MsgType: message.MsgTypeText,
			MsgData: message.NewText("^_^"),
		}
	})

	server.Serve()
	server.Send()
}
