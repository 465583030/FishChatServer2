package dao

import (
	"github.com/garyburd/redigo/redis"
	"github.com/golang/glog"
	"github.com/oikomi/FishChatServer2/common/dao/xredis"
	"github.com/oikomi/FishChatServer2/server/manager/conf"
	"golang.org/x/net/context"
)

const (
	_keyExceptionMsg = "mge_"
	_keyNormalMsg    = "mgn_"
)

func keyExceptionMsg(msgID string) string {
	return _keyExceptionMsg + msgID
}

func keyNormalMsg(msgID string) string {
	return _keyNormalMsg + msgID
}

type Dao struct {
	redis *xredis.Pool
}

func NewDao() (dao *Dao) {
	dao = &Dao{
		redis: xredis.NewPool(conf.Conf.Redis.Redis),
	}
	return
}

func (dao *Dao) SetExceptionMsg(ctx context.Context, msgID string, data string) (err error) {
	conn := dao.redis.Get(ctx)
	defer conn.Close()
	_, err = conn.Do("SET", keyExceptionMsg(msgID), data)
	if err != nil {
		glog.Error(err)
	}
	return
}

func (dao *Dao) ExceptionMsg(ctx context.Context, msgID string) (res string, err error) {
	conn := dao.redis.Get(ctx)
	defer conn.Close()
	res, err = redis.String(conn.Do("GET", keyExceptionMsg(msgID)))
	if err != nil {
		glog.Error(err)
	}
	return
}

func (dao *Dao) SetNormalMsg(ctx context.Context, msgID string, data string) (err error) {
	conn := dao.redis.Get(ctx)
	defer conn.Close()
	_, err = conn.Do("SET", keyNormalMsg(msgID), data)
	if err != nil {
		glog.Error(err)
	}
	return
}
