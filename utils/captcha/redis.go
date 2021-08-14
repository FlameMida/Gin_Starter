package captcha

import (
	"context"
	"gin-starter/global"
	"time"

	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

func NewDefaultRedisStore() base64Captcha.Store {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
	}
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
}

func (rs *RedisStore) Set(id string, value string) error {
	err := global.REDIS.Set(context.Background(), rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		global.LOG.Error("RedisStoreSetError!", zap.Error(err))
		return err
	}
	return nil
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := global.REDIS.Get(context.Background(), key).Result()
	if err != nil {
		global.LOG.Error("RedisStoreGetError!", zap.Error(err))
		return ""
	}
	if clear {
		err := global.REDIS.Del(context.Background(), key).Err()
		if err != nil {
			global.LOG.Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}
