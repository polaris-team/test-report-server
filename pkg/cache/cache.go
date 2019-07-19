package cache

type Cache interface {
	Set(key string, value string) error
	SetEx(key string, value string, milli int64) error
	Get(key string) (string, error)
	Del(key string) (int64, error)
	Exist(key string) (bool, error)
	Expire(key string, expire int64) (bool, error)
	Incrby(key string, v int64) (int64, error)
}
