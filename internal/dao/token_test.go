package dao

import (
	"github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToken(t *testing.T) {
	//查询token
	tk, err := dao.FindByUserID("qq")
	assert.Nil(t, tk)
	assert.Error(t, err, redis.Nil)
	//创建token
	token, err := dao.CreateToken("qq")
	assert.Nil(t, err)
	token2, err := dao.FindByUserID("qq")
	assert.Equal(t, token, token2.Token)
	assert.Nil(t, err)
	token3, _ := dao.CreateToken("wx2")
	_, err = dao.CreateToken("wx2")
	assert.Error(t, err, "已经拥有token了:"+token3)
	//绑定
	assert.Nil(t, dao.BindToken("wx", "qq", token2.Token))
	assert.Error(t, dao.BindToken("wx", "qq", "error"), "Token不匹配")
	assert.Error(t, dao.BindToken("wx2", "qq", token2.Token), "已经绑定过token了:"+token3)
	assert.Error(t, dao.BindToken("wx", "qq2", token2.Token), "没有找到token")
}
