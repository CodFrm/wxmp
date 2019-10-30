package main

import (
	"flag"
	"github.com/CodFrm/wxmp/api"
	"github.com/CodFrm/wxmp/internal/dao"
	"github.com/CodFrm/wxmp/internal/wchat"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"log"
)

var config = &wechat.Config{}
var addr string
var redisconf string

func init() {
	flag.StringVar(&config.AppID, "wx_app_id", "", "")
	flag.StringVar(&config.AppSecret, "wx_app_secret", "", "")
	flag.StringVar(&config.EncodingAESKey, "wx_encode_aes_key", "", "")
	flag.StringVar(&config.Token, "wx_token", "", "")
	flag.StringVar(&addr, "server_addr", ":8080", "")
	flag.StringVar(&redisconf, "redis", "127.0.0.1:6379", "")
	flag.Parse()
}

func main() {
	config.Cache = cache.NewMemory()
	wc := wechat.NewWechat(config)
	d := dao.New(&dao.DaoConfig{Redis: &redis.Options{
		Addr: redisconf,
	}})
	wchat.Init(wc)
	r := gin.Default()
	if err := api.Handel(r, d).Run(addr); err != nil {
		log.Fatal(err)
	}
}
