package wchat

import "github.com/silenceper/wechat"

var WChat *wechat.Wechat

func Init(wc *wechat.Wechat) {
	WChat = wc
}
