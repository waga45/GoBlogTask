package component

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"sync"
	"time"
)

type RedisClient struct {
	cli *redis.Client
}

var (
	once     sync.Once
	instance *RedisClient
)

func GetRedisClient() *RedisClient {
	return instance
}

/*
*
链接redis
*/
func InitRedis(host string, port int, password string) (*RedisClient, error) {
	var err error
	once.Do(func() {
		defer func() {
			if r := recover(); r != nil {
				// 记录错误信息
				err = fmt.Errorf("Redis初始化单例失败: %v", r)
			}
		}()
		//避免超时
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", host, port),
			Password: password,
			DB:       0,
		})
		status := client.Ping(ctx)
		if _, e := status.Result(); e != nil {
			err = fmt.Errorf("Redis连接失败: %v", e)
			return
		}
		instance = &RedisClient{cli: client}
	})
	if err != nil {
		return nil, err
	}
	return instance, nil
}

/*
*
存值
*/
func (c *RedisClient) Save(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.cli.Set(ctx, key, value, expiration).Err()
}

/*
*
字符串保存
*/
func (c *RedisClient) SaveString(ctx context.Context, key string, value string, expiration time.Duration) error {
	return c.cli.Set(ctx, key, value, expiration).Err()
}

/*
*
获取数据
*/
func (c *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return c.cli.Get(ctx, key).Result()
}

/*
移除Key
*/
func (c *RedisClient) RemoveKey(ctx context.Context, key string) error {
	return c.cli.Del(ctx, key).Err()
}

/*
*
向list尾部添加元素
*/
func (c *RedisClient) Push2List(ctx context.Context, key string, value string, expiration time.Duration) (int64, error) {
	i, e := c.cli.RPush(ctx, key, value).Result()
	if i > 0 {
		c.cli.Expire(ctx, key, expiration)
	}
	return i, e
}

/*
*
从最顶部取一个元素出来
*/
func (c *RedisClient) GetOneFromList(ctx context.Context, key string) (string, error) {
	return c.cli.LPop(ctx, key).Result()
}

/*
*添加到Map
 */
func (c *RedisClient) AddAllToMap(ctx context.Context, mapKey string, values map[string]interface{}, expiration time.Duration) error {
	_, err := c.cli.HSet(ctx, mapKey, values, expiration).Result()
	return err
}

/*
*
添加多个元素到
*/
func (c *RedisClient) AddOneToMap(ctx context.Context, mapKey string, fieldKey string, value string, expiration time.Duration) error {
	hashMap := make(map[string]string)
	hashMap[fieldKey] = value
	i, err := c.cli.HSet(ctx, mapKey, hashMap).Result()
	if i > 0 {
		c.cli.Expire(ctx, mapKey, expiration)
	}
	return err
}

/*
*
删除
*/
func (c *RedisClient) RemoveFromMap(ctx context.Context, mapKey string, fieldKey string) (int64, error) {
	return c.cli.HDel(ctx, mapKey, fieldKey).Result()
}

/*
* 从map中获取一个
 */
func (c *RedisClient) GetFromMap(ctx context.Context, mapKey string, fieldKey string) (string, error) {
	return c.cli.HGet(ctx, mapKey, fieldKey).Result()
}

/*
*
获取整个map
*/
func (c *RedisClient) GetMap(ctx context.Context, mapKey string) map[string]string {
	v, _ := c.cli.HGetAll(ctx, mapKey).Result()
	return v
}

// 自增
func (c *RedisClient) AutoIncr(ctx context.Context, key string) int64 {
	v, _ := c.cli.Incr(ctx, key).Result()
	return v
}

// 自增加
func (c *RedisClient) AutoIncrBy(ctx context.Context, key string, value int64) int64 {
	v, _ := c.cli.IncrBy(ctx, key, value).Result()
	return v
}

/*
*
锁
*/
func (c *RedisClient) Lock(ctx context.Context, key string, expireTime time.Duration) bool {
	uid, _ := uuid.NewUUID()
	b, _ := c.cli.SetNX(ctx, key, uid.String(), expireTime).Result()
	return b
}

const unlockScript = `
if redis.call("exists", KEYS[1]) == 1 then
    return redis.call("del", KEYS[1])
else
    return 0
end`

/*
* 解锁
 */
func (c *RedisClient) UnLock(ctx context.Context, key string) error {
	script := redis.NewScript(unlockScript)
	_, e := script.Run(ctx, c.cli, []string{key}).Int()
	return e
}

/*
*
关闭
*/
func (c *RedisClient) Close() error {
	return c.cli.Close()
}
