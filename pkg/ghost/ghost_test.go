package ghost

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/shihtzu-systems/bright/pkg/bright"
	"github.com/shihtzu-systems/bright/pkg/sorry"
	"github.com/shihtzu-systems/bright/pkg/tower"
	"testing"
)

func TestGhost_Key(t *testing.T) {
	// setup
	sut := newTestGhost(t)
	sutSystemId := sut.Id
	sutSessionId := sut.SessionId

	// test
	testCases := []struct {
		pieces      []string
		expectedKey string
	}{
		{[]string{}, sutSystemId + ":" + sutSessionId + ":ghost"},
		{[]string{"aa"}, sutSystemId + ":" + sutSessionId + ":ghost:aa"},
		{[]string{"aa", "123"}, sutSystemId + ":" + sutSessionId + ":ghost:aa:123"},
	}
	for _, tc := range testCases {
		result := sut.Key(tc.pieces...)

		// verify
		if result != tc.expectedKey {
			sorry.Unexpected(t, result, tc.expectedKey)
			t.Fail()
		}
	}
}

func TestGhost_Save(t *testing.T) {
	// setup
	sut := newTestGhost(t)
	sutSystemId := sut.Id
	sutSessionId := sut.SessionId
	testRedis := newTestRedis(t)

	// test
	sut.Save(testRedis)

	// verify
	result := Ghost{
		Id:        sutSystemId,
		SessionId: sutSessionId,
	}
	result.Materialize(testRedis)

	if result.CurrentId != sut.CurrentId {
		sorry.Unexpected(t, result.CurrentId, sut.CurrentId)
		t.Fail()
	}
}

func TestGhost_Materialize_NotFound(t *testing.T) {
	// setup
	sut := newTestGhost(t)
	testRedis := newTestRedis(t)

	// test
	sut.Materialize(testRedis)

	// verify
	if sut.CurrentId == "" {
		sorry.UnexpectedBlank(t)
		t.Fail()
	}
}

func TestGhost_Materialize_Found(t *testing.T) {
	// setup
	sut := newTestGhost(t)
	testRedis := newTestRedis(t)

	sut.Save(testRedis)

	// test
	sut.Materialize(testRedis)

	// verify
	if sut.CurrentId == "" {
		sorry.UnexpectedBlank(t)
		t.Fail()
	}
}

func newTestGhost(t *testing.T) Ghost {
	testUser := newTestUser(t)
	testGhost := Ghost{
		User:      testUser,
		CurrentId: "1",
		Bright:    newTestBright(t, testUser),
		Shadow:    newTestShadow(t, testUser),
		Try:       newTestTry(t, testUser),
		Token:     newTestToken(t),
	}
	return testGhost
}

func newTestUser(t *testing.T) (out User) {
	testUser, err := bright.TryUser(bright.TryUserArgs{
		Dad: false,
	})
	if err != nil {
		sorry.UnexpectedError(t, err)
		t.Fail()
	}
	return User(testUser)
}

func newTestBright(t *testing.T, user User) (out Bright) {
	return Bright{
		One:   BrightCharacter(user.Characters[0]),
		Two:   BrightCharacter(user.Characters[1]),
		Three: BrightCharacter(user.Characters[2]),
	}
}

func newTestShadow(t *testing.T, user User) (out Shadow) {
	return Shadow{
		One:   ShadowCharacter(user.Characters[0]),
		Two:   ShadowCharacter(user.Characters[1]),
		Three: ShadowCharacter(user.Characters[2]),
	}
}

func newTestTry(t *testing.T, user User) (out Try) {
	return Try{
		One:   TryCharacter(user.Characters[0]),
		Two:   TryCharacter(user.Characters[1]),
		Three: TryCharacter(user.Characters[2]),
	}
}

func newTestToken(t *testing.T) (out BungieToken) {
	return BungieToken{
		AccessToken:      "access-token",
		TokenType:        "Bearer",
		ExpiresIn:        3600,
		RefreshToken:     "refresh-token",
		RefreshExpiresIn: 7776000,
		MembershipId:     "12345",
	}
}

func newTestRedis(t *testing.T) (out tower.Redis) {
	out = tower.Redis{
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
