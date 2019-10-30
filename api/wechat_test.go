package api

import (
	"github.com/silenceper/wechat/message"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWechat_WxHandel(t *testing.T) {
	handel := wechat.WxHandel()

	//创建token
	reply := handel(message.MixMessage{
		CommonToken: message.CommonToken{
			FromUserName: "88888888",
			MsgType:      message.MsgTypeEvent,
		},
		Event:    message.EventClick,
		EventKey: "create",
	})
	qqtoken := reply.MsgData.(*message.Text).Content

	reply = handel(message.MixMessage{
		CommonToken: message.CommonToken{
			FromUserName: "sq",
			MsgType:      message.MsgTypeText,
		},
		Content: "申请token",
	})
	assert.Equal(t, len(reply.MsgData.(*message.Text).Content), 11)
	//绑定token
	reply = handel(message.MixMessage{
		CommonToken: message.CommonToken{
			FromUserName: "wx",
			MsgType:      message.MsgTypeText,
		},
		Content: "88888888+" + qqtoken,
	})
	assert.Equal(t, "绑定成功", reply.MsgData.(*message.Text).Content)
	//重新绑定
	reply = handel(message.MixMessage{
		CommonToken: message.CommonToken{
			FromUserName: "wx",
			MsgType:      message.MsgTypeText,
		},
		Content: "88888888+" + qqtoken,
	})
	assert.NotEqual(t, "绑定成功", reply.MsgData.(*message.Text).Content)
	//token信息
	reply = handel(message.MixMessage{
		CommonToken: message.CommonToken{
			FromUserName: "wx",
			MsgType:      message.MsgTypeText,
		},
		Content: "token",
	})
	assert.Equal(t, "你的token为:"+qqtoken+" 剩余:200", reply.MsgData.(*message.Text).Content)
	reply = handel(message.MixMessage{
		CommonToken: message.CommonToken{
			FromUserName: "88888888",
			MsgType:      message.MsgTypeEvent,
		},
		Event:    message.EventClick,
		EventKey: "balance",
	})
	assert.Equal(t, "你的token为:"+qqtoken+" 剩余:200", reply.MsgData.(*message.Text).Content)
	reply = handel(message.MixMessage{
		CommonToken: message.CommonToken{
			FromUserName: "88888888",
			MsgType:      message.MsgTypeText,
		},
		Content: "token",
	})
	assert.Equal(t, "你的token为:"+qqtoken+" 剩余:200", reply.MsgData.(*message.Text).Content)

	reply = handel(message.MixMessage{
		CommonToken: message.CommonToken{
			FromUserName: "0000",
			MsgType:      message.MsgTypeText,
		},
		Content: "token",
	})
	assert.Equal(t, "你还没有token,请可以发送 qq号码+token 将原有的token和微信公众号进行绑定,例如: 88888888+ilyedbyd (绑定成功自动增加100点),发送 \"申请token\" 可申请一个新的token.", reply.MsgData.(*message.Text).Content)

}
