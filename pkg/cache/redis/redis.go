package redis

import (
	"reflect"
	"sync"

	"github.com/gomodule/redigo/redis"
)

var mu sync.Mutex
var single RedisProxy

type RedisProxy struct {
	conn redis.Conn
}

func (rp *RedisProxy) ZAdd(key string, score float64, value string) error {
	conn, e := rp.Connect()
	if e != nil {
		return e
	}
	_, err := conn.Do("zadd", key, score, value)
	return err
}

func (rp *RedisProxy) SetEx(key string, value string, ex int64) error {
	conn, e := rp.Connect()
	if e != nil {
		return e
	}
	_, err := conn.Do("setex", key, ex, value)
	return err
}

func (rp *RedisProxy) Get(key string) (string, error) {
	conn, e := rp.Connect()
	if e != nil {
		return "", e
	}
	rs, err := conn.Do("get", key)
	if err != nil {
		return "", err
	}
	return string(rs.([]byte)), nil
}

func (rp *RedisProxy) Del(key string) (int64, error) {
	conn, e := rp.Connect()
	if e != nil {
		return 0, e
	}
	rs, err := conn.Do("Del", key)
	if err != nil {
		return 0, err
	}
	return rs.(int64), nil
}

func (rp *RedisProxy) Incrby(key string, v int64) (int64, error) {
	conn, e := rp.Connect()
	if e != nil {
		return 0, e
	}
	rs, err := conn.Do("INCRBY", key, v)
	if err != nil {
		return 0, err
	}
	return rs.(int64), err
}

func (rp *RedisProxy) Exist(key string) (bool, error) {
	conn, e := rp.Connect()
	if e != nil {
		return false, e
	}
	rs, err := conn.Do("EXISTS", key)
	if err != nil {
		return false, err
	}
	return rs.(int64) == 1, err
}

func (rp *RedisProxy) Expire(key string, expire int64) (bool, error) {
	conn, e := rp.Connect()
	if e != nil {
		return false, e
	}
	rs, err := conn.Do("EXPIRE", key, expire)
	if err != nil {
		return false, err
	}
	return rs.(int64) == 1, err
}

func (rp *RedisProxy) Connect() (redis.Conn, error) {
	if rp.conn == nil || rp.conn.Err() != nil {
		mu.Lock()
		defer mu.Unlock()
		if rp.conn == nil || rp.conn.Err() != nil {
			conn, err := GetRedisConn()
			if err != nil {
				return nil, err
			}
			rp.conn = conn
		}
	}
	return rp.conn, nil
}

func (rp *RedisProxy) Close() {
	if rp.conn != nil && rp.conn.Err() == nil {
		rp.conn.Close()
	}
}

func (rp RedisProxy) IsEmpty() bool {
	return reflect.DeepEqual(rp, RedisProxy{})
}

//GetProxy get redis oper proxy
func GetProxy() *RedisProxy {
	return &RedisProxy{}
}
