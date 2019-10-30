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
	assert.Equal(t, "你还没有token,请选择绑定token或者申请token", reply.MsgData.(*message.Text).Content)

}
