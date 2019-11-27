package tower

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
)

func (r *Redis) Connect() {
	dial := fmt.Sprintf("%s:%v", r.Address, r.Port)
	log.Debugf("redis: %s [%d]", dial, r.Database)
	conn, err := redis.Dial("tcp", dial, redis.DialDatabase(r.Database))
	if err != nil {
		log.Fatal(err)
	}
	r.connection = conn
}

func (r *Redis) Disconnect() {
	if err := r.connection.Close(); err != nil {
		log.Fatal(err)
	}
	r.connection = nil
}

func (r Redis) Exists(key string) bool {
	reply, err := r.connection.Do("Exists", key)
	if err != nil {
		log.Fatal(err)
	}
	return reply.(int64) == int64(1)
}

func (r Redis) Set(key string, jsonValue []byte) {
	reply, err := r.connection.Do("SET", key, jsonValue)
	log.Debug(reply)
	if err != nil {
		log.Fatal(err)
	}
}

func (r Redis) HSet(key, field string, jsonValue []byte) {
	reply, err := r.connection.Do("HSET", key, field, jsonValue)
	log.Debug(reply)
	if err != nil {
		log.Fatal(err)
	}
}

func (r Redis) Get(key string) (out string) {
	reply, err := r.connection.Do("GET", key)
	if err != nil {
		log.Fatal(err)
	}
	if reply == nil {
		return ""
	} else {
		return string(reply.([]byte))
	}
}

func (r Redis) HGet(key, field string) (out string) {
	reply, err := r.connection.Do("HGET", key, field)
	if err != nil {
		log.Fatal(err)
	}
	if reply == nil {
		return ""
	} else {
		return string(reply.([]byte))
	}
}
