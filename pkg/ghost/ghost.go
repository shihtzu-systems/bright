package ghost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shihtzu-systems/bright/pkg/bungo"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
	"strings"
)

func NewGhost(id, sessionId string) Ghost {
	return Ghost{
		Id:        id,
		SessionId: sessionId,
	}

}

func (g Ghost) Key(pieces ...string) string {
	key := "unknown"
	if len(pieces) > 0 {
		key = fmt.Sprintf("%s:%s:ghost:%s", g.Id, g.SessionId, strings.Join(pieces, ":"))
	} else {
		key = fmt.Sprintf("%s:%s:ghost", g.Id, g.SessionId)
	}
	log.Trace("ghost key: ", key)
	return key
}

func (g *Ghost) Materialize(redis tower.Redis) {
	redis.Connect()
	defer redis.Disconnect()

	prettyJson := redis.Get(g.Key())

	if prettyJson == "" {
		log.Debug("no ghost found: ", g.Key())
		return
	}

	if err := json.Unmarshal([]byte(prettyJson), &g); err != nil {
		log.Fatal(err)
	}
	log.Debug("ghost materialized: ", g.CurrentId)
}

func (g *Ghost) SetTryUser(user TryUser) {
	g.Try.User = user
	if len(user.Characters) >= 1 {
		g.Try.One = TryCharacter(user.Characters[0])
	}
	if len(user.Characters) >= 2 {
		g.Try.Two = TryCharacter(user.Characters[1])
	}
	if len(user.Characters) == 3 {
		g.Try.Three = TryCharacter(user.Characters[2])
	}
}

func (g *Ghost) SetTryCurrent(character bungo.Character) {
	switch character.Id {
	case g.Try.One.Id:
		g.Try.One = TryCharacter(character)
	case g.Try.Two.Id:
		g.Try.Two = TryCharacter(character)
	case g.Try.Three.Id:
		g.Try.Three = TryCharacter(character)
	}
}

func (g Ghost) FindTry(id string) TryCharacter {
	switch id {
	case g.Try.One.Id:
		return g.Try.One
	case g.Try.Two.Id:
		return g.Try.Two
	case g.Try.Three.Id:
		return g.Try.Three
	}
	return TryCharacter{}
}

func (g *Ghost) Save(redis tower.Redis) {
	prettyJson := g.PrettyJson()

	redis.Connect()
	defer redis.Disconnect()

	redis.Set(g.Key(), prettyJson)

	log.Debug("ghost save: ", g.CurrentId)
}

func (g *Ghost) PrettyJson() (out []byte) {
	gout, _ := json.Marshal(g)
	var prettyJSON bytes.Buffer
	_ = json.Indent(&prettyJSON, gout, "", "\t")
	return prettyJSON.Bytes()
}
