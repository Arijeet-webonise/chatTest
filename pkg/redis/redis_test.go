package redis

import (
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	redis, _, err := InitRedis("localhost:6379", "", 0)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if redis == nil {
		t.Error(err)
		t.Fail()
		return
	}
	key := "testString"
	value := "something"
	if err := redis.Set(key, value, time.Second*10); err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	result, err := redis.Get(key)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if _, err := result.Result(); err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	str := ""
	result.Scan(&str)
	if str != value {
		t.Error(err)
		t.Fail()
		return
	}

	if count, err := redis.Delete(key); err != nil {
		t.Error(err)
		t.Fail()
		return
	} else if count != 1 {
		t.Error(err)
		t.Fail()
		return
	}
}

func TestRedisError(t *testing.T) {
	redis, _, err := InitRedis("localhost:6379", "", 0)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if redis == nil {
		t.Error(err)
		t.Fail()
		return
	}
	key := "testString"
	value := "something"
	if err := redis.Set(key, value, time.Second*10); err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if _, err := redis.Get(key + "d"); err == nil {
		t.Error(err)
		t.Fail()
		return
	}

	if count, err := redis.Delete(key); err != nil {
		t.Error(err)
		t.Fail()
		return
	} else if count != 1 {
		t.Error(err)
		t.Fail()
		return
	}
}
