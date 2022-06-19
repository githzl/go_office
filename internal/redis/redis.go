package redis

import "time"

func S(key string, value interface{}, expire time.Duration) (string, error) {
	return Client.Set(key, value, 1000*time.Second).Result()
}
