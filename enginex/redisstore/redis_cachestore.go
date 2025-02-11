package redisstore

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

type CacheStoreConnections map[string]*RedisCacheStore

type CacheSerializer interface {
	Deserialize(d []byte, cache *Cache) error
	Serialize(cache *Cache) ([]byte, error)
}

type Cache struct {
	ID     string
	MaxAge int64
	Values map[string]interface{}
}

// GobSerializer uses gob package to encode the session map
type CacheGobSerializer struct{}

// Serialize using gob
func (s CacheGobSerializer) Serialize(cache *Cache) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(cache.Values)
	if err == nil {
		return buf.Bytes(), nil
	}
	return nil, err
}

// Deserialize back to map[interface{}]interface{}
func (s CacheGobSerializer) Deserialize(d []byte, cache *Cache) error {
	dec := gob.NewDecoder(bytes.NewBuffer(d))
	return dec.Decode(&cache.Values)
}

type RedisCacheStore struct {
	Pool          *redis.Pool
	defaultMaxAge int
	keyPrefix     string
	maxLength     int
	serializer    CacheSerializer
}

func NewRedisCacheStore(size int, maxlength int, network, address, password string) (*RedisCacheStore, error) {
	return NewRedisCacheStoreWithPool(maxlength, &redis.Pool{
		MaxIdle:     size,
		IdleTimeout: 240 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			return dial(network, address, password)
		},
	})
}

func NewRedisCacheStoreWithPool(maxlength int, pool *redis.Pool) (*RedisCacheStore, error) {
	rs := &RedisCacheStore{
		keyPrefix:     "cache_",
		Pool:          pool,
		defaultMaxAge: 60 * 20,
		maxLength: maxlength,
		serializer:    CacheGobSerializer{},
	}
	_, err := rs.Ping()
	return rs, err
}

func (s *RedisCacheStore) Close() error {
	return s.Pool.Close()
}

func (s *RedisCacheStore) Ping() (bool, error) {
	conn := s.Pool.Get()
	defer conn.Close()
	data, err := conn.Do("PING")
	if err != nil || data == nil {
		return false, err
	}
	return (data == "PONG"), nil
}

// save stores the cache in redis.
func (s *RedisCacheStore) Save(cache *Cache) error {
	b, err := s.serializer.Serialize(cache)
	if err != nil {
		return err
	}
	if s.maxLength != 0 && len(b) > s.maxLength {
		return errors.New("SessionStore: the value to store is too big")
	}
	conn := s.Pool.Get()
	defer conn.Close()
	if err = conn.Err(); err != nil {
		return err
	}

	age := cache.MaxAge
	if age <= 0 {
		age = int64(s.defaultMaxAge)
	}
	_, err = conn.Do("SETEX", s.keyPrefix+cache.ID, age, b)
	return err
}

// delete cache
func (s *RedisCacheStore) Del(cache *Cache) error {
	conn := s.Pool.Get()
	defer conn.Close()
	if _, err := conn.Do("DEL", s.keyPrefix+cache.ID); err != nil {
		return err
	}
	return nil
}

// get the cache from redis.
func (s *RedisCacheStore) Read(id string) (*Cache, error) {
	conn := s.Pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return nil, err
	}
	data, err := conn.Do("GET", s.keyPrefix+id)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New("no data was associated with this key")
	}
	b, err := redis.Bytes(data, err)
	if err != nil {
		return nil, err
	}

	cache := &Cache{
		ID: id,
	}

	err = s.serializer.Deserialize(b, cache)
	if err != nil {
		return nil, err
	}

	//get  ttl
	ittl, err := conn.Do("TTL", s.keyPrefix+id)
	if err != nil {
		return nil, err
	}

	ttl, ok := ittl.(int64)
	if ! ok {
		return nil, errors.New("cannot read cache ttl")
	}

	cache.MaxAge = ttl

	return cache, nil
}

func (s *RedisCacheStore) del(key string) error {
	conn := s.Pool.Get()
	defer conn.Close()
	if _, err := conn.Do("DEL", key); err != nil {
		return err
	}
	return nil
}

func (s *RedisCacheStore) delLike(key string) error {
	conn := s.Pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return err
	}
	ikeyArr, err := conn.Do("KEYS", key+"*")
	if err != nil {
		return err
	}
	ikeys, ok := ikeyArr.([]interface{})
	if ok {
		for _, ik := range ikeys {
			k := string(ik.([]uint8))
			if ok {
				err := s.del(k)
				if err != nil {
					return err
				}
			}

		}
	}

	return nil
}

func (s *RedisCacheStore) get(key string) (interface{}, error) {
	conn := s.Pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return false, err
	}
	data, err := conn.Do("GET", key)
	if err != nil {
		return false, err
	}
	return data, nil
}

func (s *RedisCacheStore) set(key string, data interface{}) error {
	conn := s.Pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return err
	}
	data, err := conn.Do("SET", key, data)
	if err != nil {
		return err
	}
	return nil
}

func (s *RedisCacheStore) expire(key string, ttl int) error {
	conn := s.Pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return err
	}
	_, err := conn.Do("EXPIRE", key, ttl)
	if err != nil {
		return err
	}
	return nil
}

func (s *RedisCacheStore) ttl(key string) (int64, error) {
	conn := s.Pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return -1, err
	}
	data, err := conn.Do("TTL", key)
	if err != nil {
		return -1, err
	}

	iData, ok := data.(int64)
	if !ok {
		return -1, fmt.Errorf("not number result %v", data)
	}
	return iData, nil
}

// SetSerializer sets the serializer
func (s *RedisCacheStore) SetSerializer(ss CacheSerializer) {
	s.serializer = ss
}

func (s *RedisCacheStore) SetKeyPrefix(p string) {
	s.keyPrefix = p
}
