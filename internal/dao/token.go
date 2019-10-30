package dao

import (
	"errors"
	"github.com/CodFrm/wxmp/internal/model"
	"github.com/CodFrm/wxmp/utils"
	"github.com/go-redis/redis/v7"
)

type Token interface {
	FindByUserID(userid string) (*model.Token, error)
	CreateToken(userid string) (string, error)
	BindToken(userid string, olduserid string, token string) error
}

func (d *Dao) FindByUserID(userid string) (*model.Token, error) {
	ret := &model.Token{}
	m := d.redis.HGet("cxmooc:genuser", userid)
	if m.Err() != nil {
		return nil, m.Err()
	}
	ret.Token = m.Val()
	m = d.redis.Get("cxmooc:vtoken:" + ret.Token)
	if m.Err() != nil {
		return nil, m.Err()
	}
	ret.Num, _ = m.Int64()
	return ret, nil
}

func (d *Dao) CreateToken(userid string) (string, error) {
	token, err := d.FindByUserID(userid)
	if err != nil && err != redis.Nil {
		return "", err
	}
	if token != nil {
		return "", errors.New("已经拥有token了:" + token.Token)
	}
	randToken := utils.RandStringRunes(11)
	if m := d.redis.HSet("cxmooc:genuser", userid, randToken); m.Err() != nil {
		return "", m.Err()
	}
	if m := d.redis.Set("cxmooc:vtoken:"+randToken, 100, 0); m.Err() != nil {
		return "", m.Err()
	}
	return randToken, nil
}

func (d *Dao) BindToken(userid string, olduserid string, token string) error {
	t, err := d.FindByUserID(olduserid)
	if err != nil {
		if err == redis.Nil {
			return errors.New("没有找到token")
		}
		return err
	}
	if t.Token != token {
		return errors.New("Token不匹配")
	}
	t, err = d.FindByUserID(userid)
	if t != nil {
		return errors.New("已经绑定过token了:" + t.Token)
	}
	if err != nil && err != redis.Nil {
		return err
	}
	if m := d.redis.HSet("cxmooc:genuser", userid, token); m.Err() != nil {
		return m.Err()
	}
	if m := d.redis.IncrBy("cxmooc:vtoken:"+token, 100); m.Err() != nil {
		return m.Err()
	}
	return nil
}
