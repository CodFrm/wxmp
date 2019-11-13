package api

import (
	"bytes"
	"encoding/json"
	"github.com/CodFrm/wxmp/internal/dao"
	"github.com/CodFrm/wxmp/internal/model"
	"github.com/CodFrm/wxmp/internal/wchat"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/silenceper/wechat/message"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
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
	return func(msg message.MixMessage) (ret *message.Reply) {
		content := "发送查+题目内容即可查询题目答案(eg.查 我们通常所说的历史就是二阶历史),访问地址: http://cx.icodef.com/query.html 也可以进行查询哦,发送token可以查看token相关命令"
		defer func() {
			ret = &message.Reply{
				MsgType: message.MsgTypeText,
				MsgData: message.NewText(content),
			}
		}()
		switch msg.MsgType {
		case message.MsgTypeText:
			{
				cnt := strings.Trim(msg.Content, " ")
				if cnt == "token" {
					content = w.getToken(msg.FromUserName)
				} else if cnt == "申请token" {
					if token, err := w.token.CreateToken(msg.FromUserName); err != nil {
						content = err.Error()
					} else {
						content = token
					}
				} else {
					if regex, err := regexp.Compile(`^(\d+)\+(\w+)`); err != nil {
						content = err.Error()
						return
					} else {
						str := regex.FindStringSubmatch(msg.Content)
						if str != nil {
							if err := w.token.BindToken(msg.FromUserName, str[1], str[2]); err != nil {
								content = err.Error()
							} else {
								content = "绑定成功"
							}
							return
						}
					}
					if regex, err := regexp.Compile(`^查 (.*?)$`); err != nil {
						content = err.Error()
						return
					} else {
						str := regex.FindStringSubmatch(msg.Content)
						if str != nil {
							req, err := http.NewRequest("POST", "http://cx.icodef.com/v2/answer", bytes.NewBuffer([]byte(
								"topic[0]="+url.QueryEscape(str[1]),
							)))
							req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
							client := &http.Client{}
							resp, err := client.Do(req)
							if err != nil {
								content = err.Error()
								return
							}
							defer resp.Body.Close()
							b, err := ioutil.ReadAll(resp.Body)
							if err != nil {
								content = err.Error()
								return
							}
							data := make([]model.Answer, 0)
							if err := json.Unmarshal(b, &data); err != nil {
								content = err.Error()
								return
							}
							if len(data[0].Result) <= 0 {
								content = "未找到答案"
								return
							}
							for _, v := range data[0].Result {
								if v.Type == 3 {
									if v.Correct[0].Content.(bool) {
										content = "正确"
									} else {
										content = "错误"
									}
								} else {
									content = ""
									for _, correct := range v.Correct {
										content = content + correct.Option.(string) + ":" + correct.Content.(string) + "\n"
									}
								}
							}
						}
						return
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
					default:
						{
							content = "发送查+题目内容即可查询题目答案(eg.查 我们通常所说的历史就是二阶历史),访问地址: http://cx.icodef.com/query.html 也可以进行查询哦,发送token可以查看token相关命令"
						}
					}
				} else if msg.Event == message.EventSubscribe {
					content = "欢迎关注icodef.com"
				}
			}
		}
		return
	}
}

func (w *Wechat) getToken(uid string) string {
	if t, err := w.token.FindByUserID(uid); err != nil {
		if err == redis.Nil {
			return "你还没有token,请可以发送 qq号码+token 将原有的token和微信公众号进行绑定,例如: 88888888+ilyedbyd (绑定成功自动增加100点),发送 \"申请token\" 可申请一个新的token."
		}
		return err.Error()
	} else {
		return "你的token为:" + t.Token + " 剩余:" + strconv.Itoa(int(t.Num))
	}
}
