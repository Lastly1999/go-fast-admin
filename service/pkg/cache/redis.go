package cache

import (
	"fast-admin-service/global"
	"fast-admin-service/pkg/redis"
	"time"
)

const CAPTCHA = "captcha:"

type RedisStore struct {
}

func (r RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	err := redis.Rdb.Set(key, value, time.Minute*2).Err()
	if err != nil {
		global.ZAP_LOG.Error(err.Error())
		return err
	}
	return nil
}

func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := redis.Rdb.Get(key).Result()
	if err != nil {
		global.ZAP_LOG.Error(err.Error())
		return ""
	}
	if clear {
		err = redis.Rdb.Del(key).Err()
		if err != nil {
			return ""
		}
	}
	return val
}

func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	return v == answer
}
