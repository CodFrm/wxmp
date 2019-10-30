package wchat

import (
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/menu"
)

var WChat *wechat.Wechat

func Init(wc *wechat.Wechat) {
	mu := wc.GetMenu()
	mu.SetMenu([]*menu.Button{{
		Name: "token菜单",
		SubButtons: []*menu.Button{{
			Type: "click",
			Name: "查询余额",
			Key:  "balance",
		}, {
			Type: "click",
			Name: "绑定key",
			Key:  "bind",
		}, {
			Type: "click",
			Name: "申请token",
			Key:  "create",
		}},
	}})
	WChat = wc
}
