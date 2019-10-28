package main

import (
	"flag"
	"github.com/CodFrm/wxmp/api"
	"github.com/CodFrm/wxmp/internal/wchat"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat"
	"log"
)

var config = &wechat.Config{}
var addr string

func init() {
	flag.StringVar(&config.AppID, "wx_app_id", "", "")
	flag.StringVar(&config.AppSecret, "wx_app_secret", "", "")
	flag.StringVar(&config.EncodingAESKey, "wx_encode_aes_key", "", "")
	flag.StringVar(&config.Token, "wx_token", "", "")
	flag.StringVar(&addr, "server_addr", ":8080", "")
	flag.Parse()
}

func main() {
	wc := wechat.NewWechat(config)
	wchat.Init(wc)
	r := gin.Default()
	if err := api.Handel(r).Run(addr); err != nil {
		log.Fatal(err)
	}
}
