package tower

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/shihtzu-systems/bright/pkg/sorry"
	"testing"
)

func TestRedis_Connect(t *testing.T) {
	// setup
	sut := newTestRedis(t)

	// test
	sut.Connect()

	// verify
	if sut.connection == nil {
		sorry.UnexpectedNil(t)
		t.Fail()
	}
}

func TestRedis_Disconnect(t *testing.T) {
	// setup
	sut := newTestRedis(t)
	sut.Connect()

	// test
	sut.Disconnect()

	// verify
	if sut.connection != nil {
		sorry.UnexpectedNil(t)
		t.Fail()
	}
}

func TestRedis_Get(t *testing.T) {
	// setup
	sut := newTestRedis(t)

	sutKey := "key"
	sutValue := "value"
	sut.Set(sutKey, []byte(sutValue))
	sut.Connect()
	defer sut.Disconnect()

	// test
	result := sut.Get(sutKey)

	// verify
	if result != sutValue {
		sorry.Unexpected(t, result, sutValue)
		t.Fail()
	}
}

func newTestRedis(t *testing.T) (out Redis) {
	out = Redis{
		Address:  "localhost",
		Port:     6379,
		Database: 10,
	}
	dial := fmt.Sprintf("%s:%v", out.Address, out.Port)
	conn, err := redis.Dial("tcp", dial, redis.DialDatabase(out.Database))
	if err != nil {
		sorry.UnexpectedError(t, err)
		t.Fail()
	}

	reply, err := conn.Do("ping")
	if err != nil {
		sorry.UnexpectedError(t, err)
		t.Fail()
	}
	if reply == nil {
		sorry.UnexpectedNil(t)
		t.Fail()
	}

	pong := reply.(string)
	if reply != "PONG" {
		sorry.Unexpected(t, pong, "PONG")
		t.Fail()
	}
	return out
}
