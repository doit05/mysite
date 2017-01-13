package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

type Redis struct {
	p        *redis.Pool // redis connection pool
	conninfo string
	dbNum    int
	key      string
	password string
	prefix   string
}

func NewRedisCache() *Redis {
	return &Redis{}
}

// actually do the redis cmds
func (rc *Redis) do(commandName string, args ...interface{}) (reply interface{}, err error) {
	c := rc.p.Get()
	defer c.Close()

	return c.Do(commandName, args...)
}

// 获取缓存键名
func (rc *Redis) GetCacheKey(key string) string {
	return rc.prefix + ":" + key
}

// 是否存在
func (rc *Redis) IsExist(key string) bool {
	v, err := redis.Bool(rc.do("EXISTS", rc.GetCacheKey(key)))
	if err != nil {
		return false
	}
	return v
}

func (rc *Redis) Get(key string) interface{} {
	if v, err := rc.do("GET", rc.GetCacheKey(key)); err == nil {
		return v
	}
	return nil
}

func (rc *Redis) Set(key string, val interface{}, timeout time.Duration) error {
	var err error
	if _, err = rc.do("SETEX", rc.GetCacheKey(key), int64(timeout/time.Second), val); err != nil {
		return err
	}
	return err
}

// 设置string值
func (rc *Redis) SetString(key string, value string, timeout time.Duration) error {
	_, err := rc.do("SETEX", rc.GetCacheKey(key), int64(timeout/time.Second), value)
	return err
}

// 设置int值
func (rc *Redis) SetInt(key string, value int, timeout time.Duration) error {
	_, err := rc.do("SETEX", rc.GetCacheKey(key), int64(timeout/time.Second), value)
	return err
}

// 设置int64值
func (rc *Redis) SetInt64(key string, value int64, timeout time.Duration) error {
	_, err := rc.do("SETEX", rc.GetCacheKey(key), int64(timeout/time.Second), value)
	return err
}

// 设置float64值
func (rc *Redis) SetFloat64(key string, value float64, timeout time.Duration) error {
	_, err := rc.do("SETEX", rc.GetCacheKey(key), int64(timeout/time.Second), value)
	return err
}

// 设置bool值
func (rc *Redis) SetBool(key string, value bool, timeout time.Duration) error {
	_, err := rc.do("SETEX", rc.GetCacheKey(key), int64(timeout/time.Second), value)
	return err
}

// 设置StringMap值
func (rc *Redis) SetStringMap(key string, value map[string]string, timeout time.Duration) error {
	var err error

	key = rc.GetCacheKey(key)

	// 设置hash表
	for k, v := range value {
		_, err = rc.do("HSET", key, k, v)
		if err != nil {
			break
		}
	}

	// 设置成功再为hash表设置有效期
	if err == nil { // 设置成功
		_, err = rc.do("EXPIRE", key, int64(timeout/time.Second))
		if err != nil {
			fmt.Printf("SetStringMap:设置redis键名为 %s 的hash表有效时间失败,err:%v\n", key, err)
		}
	}

	return err
}

// 设置IntMap值
func (rc *Redis) SetIntMap(key string, value map[string]int, timeout time.Duration) error {
	var err error

	key = rc.GetCacheKey(key)

	// 设置hash表
	for k, v := range value {
		_, err = rc.do("HSET", key, k, v)
		if err != nil {
			break
		}
	}

	// 设置成功再为hash表设置有效期
	if err == nil { // 设置成功
		_, err = rc.do("EXPIRE", key, int64(timeout/time.Second))
		if err != nil {
			fmt.Printf("SetIntMap:设置redis键名为 %s 的hash表有效时间失败,err:%v\n", key, err)
		}
	}

	return err
}

// 设置Int64Map值
func (rc *Redis) SetInt64Map(key string, value map[string]int64, timeout time.Duration) error {
	var err error

	key = rc.GetCacheKey(key)

	// 设置hash表
	for k, v := range value {
		_, err = rc.do("HSET", key, k, v)
		if err != nil {
			break
		}
	}

	// 设置成功再为hash表设置有效期
	if err == nil { // 设置成功
		_, err = rc.do("EXPIRE", key, int64(timeout/time.Second))
		if err != nil {
			fmt.Printf("SetInt64Map:设置redis键名为 %s 的hash表有效时间失败,err:%v\n", key, err)
		}
	}

	return err
}

// 获取字符串值
// 如果err不等于nil，string返回空字符串
func (rc *Redis) GetString(key string) (string, error) {
	return redis.String(rc.do("GET", rc.GetCacheKey(key)))
}

// 获取int值
// 如果err不等于nil，int返回0
func (rc *Redis) GetInt(key string) (int, error) {
	return redis.Int(rc.do("GET", rc.GetCacheKey(key)))
}

// 获取int64值
// 如果err不等于nil，int64返回0
func (rc *Redis) GetInt64(key string) (int64, error) {
	return redis.Int64(rc.do("GET", rc.GetCacheKey(key)))
}

// 获取float64值
// 如果err不等于nil，float64返回0
func (rc *Redis) GetFloat64(key string) (float64, error) {
	return redis.Float64(rc.do("GET", rc.GetCacheKey(key)))
}

// 获取bool值
// 如果err不等于nil，bool返回false
func (rc *Redis) GetBool(key string) (bool, error) {
	return redis.Bool(rc.do("GET", rc.GetCacheKey(key)))
}

// 获取StringMap值
func (rc *Redis) GetStringMap(key string) (map[string]string, error) {
	return redis.StringMap(rc.do("HGETALL", rc.GetCacheKey(key)))
}

// 获取IntMap值
func (rc *Redis) GetIntMap(key string) (map[string]int, error) {
	return redis.IntMap(rc.do("HGETALL", rc.GetCacheKey(key)))
}

// 获取Int64Map值
func (rc Redis) GetInt64Map(key string) (map[string]int64, error) {
	return redis.Int64Map(rc.do("HGETALL", rc.GetCacheKey(key)))
}

// 删除
func (rc *Redis) Delete(key string) error {
	_, err := rc.do("DEL", rc.GetCacheKey(key))
	return err
}

func (rc *Redis) Connect(config string) error {
	var cf map[string]string
	json.Unmarshal([]byte(config), &cf)

	/*if _, ok := cf["key"]; !ok {
		cf["key"] = DefaultKey
	}*/

	if _, ok := cf["prefix"]; !ok {
		cf["prefix"] = ""
	}

	if _, ok := cf["conn"]; !ok {
		return errors.New("config has no conn key")
	}
	if _, ok := cf["dbNum"]; !ok {
		cf["dbNum"] = "0"
	}
	if _, ok := cf["password"]; !ok {
		cf["password"] = ""
	}
	rc.key = cf["key"]
	rc.conninfo = cf["conn"]
	rc.dbNum, _ = strconv.Atoi(cf["dbNum"])
	rc.password = cf["password"]
	rc.prefix = cf["prefix"]

	rc.connectInit()

	c := rc.p.Get()
	defer c.Close()

	return c.Err()
}

func (rc *Redis) connectInit() {
	dialFunc := func() (c redis.Conn, err error) {
		c, err = redis.Dial("tcp", rc.conninfo)
		if err != nil {
			return nil, err
		}

		if rc.password != "" {
			if _, err := c.Do("AUTH", rc.password); err != nil {
				c.Close()
				return nil, err
			}
		}

		_, selecterr := c.Do("SELECT", rc.dbNum)
		if selecterr != nil {
			c.Close()
			return nil, selecterr
		}
		return
	}
	// initialize a new pool
	rc.p = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 180 * time.Second,
		Dial:        dialFunc,
	}
}
