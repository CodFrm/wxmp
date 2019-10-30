package api

import (
	"github.com/CodFrm/wxmp/internal/dao"
	"github.com/CodFrm/wxmp/internal/wchat"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/silenceper/wechat/message"
	"regexp"
	"strconv"
)

type Wechat struct {
	token dao.Token
}

func NewWechat(r *gin.Engine, token dao.Token) *Wechat {
	w := &Wechat{token: token}
	r.Any("/wchat", w.Wchat())
	return w
}

func (w *Wechat) Wchat() gin.HandlerFunc {
	return func(c *gin.Context) {
		server := wchat.WChat.GetServer(c.Request, c.Writer)
		server.SetMessageHandler(w.WxHandel())
		server.Serve()
		server.Send()
	}
}

func (w *Wechat) WxHandel() func(message.MixMessage) *message.Reply {
	return func(msg message.MixMessage) *message.Reply {
		content := "^_^"
		switch msg.MsgType {
		case message.MsgTypeText:
			{
				if msg.Content == "token" {
					content = w.getToken(msg.FromUserName)
				} else if regex, err := regexp.Compile(`^(\d+)\+(\w+)`); err == nil {
					str := regex.FindStringSubmatch(msg.Content)
					if str != nil {
						if err := w.token.BindToken(msg.FromUserName, str[1], str[2]); err != nil {
							content = err.Error()
						} else {
							content = "绑定成功"
						}
					}
				}
			}
		case message.MsgTypeEvent:
			{
				if msg.Event == message.EventClick {
					switch msg.EventKey {
					case "balance":
						{
							content = w.getToken(msg.FromUserName)
						}
					case "bind":
						{
							content = "请发送 qq号码+token 将原有的token和微信公众号进行绑定,例如: 88888888+ilyedbyd (绑定成功自动增加100点)"
						}
					case "create":
						{
							if token, err := w.token.CreateToken(msg.FromUserName); err != nil {
								content = err.Error()
							} else {
								content = token
							}
						}
					}
				}
			}
		}
		return &message.Reply{
			MsgType: message.MsgTypeText,
			MsgData: message.NewText(content),
		}
	}

}

func (w *Wechat) getToken(uid string) string {
	if t, err := w.token.FindByUserID(uid); err != nil {
		if err == redis.Nil {
			return "你还没有token,请选择绑定token或者申请token"
		}
		return err.Error()
	} else {
		return "你的token为:" + t.Token + " 剩余:" + strconv.Itoa(int(t.Num))
	}
}
