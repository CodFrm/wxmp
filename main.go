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
	flag.StringVar(&config.AppID, "wc_app_id", "", "")
	flag.StringVar(&config.AppSecret, "wc_app_secret", "", "")
	flag.StringVar(&config.EncodingAESKey, "wc_encode_aes_key", "", "")
	flag.StringVar(&config.Token, "wc_token", "", "")
	flag.StringVar(&addr, "server_addr", ":8080", "")
	flag.Parse()
}

func main() {
	wc := wechat.NewWechat(config)
	wchat.Init(wc)
	r := gin.Default()
	api.Handel(r)
	if err := api.Handel(r).Run(addr); err != nil {
		log.Fatal(err)
	}
}
