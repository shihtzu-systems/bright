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
	log.Debugf("ghost materialized: id=%s session_id=%s", g.Id, g.SessionId)
}

func (s *Soul) Embody(gamer bungo.Gamer) {
	s.Gamer = gamer
	if len(gamer.Guardians) >= 1 {
		s.Possessed = gamer.Guardians[0].Id
		s.One = gamer.Guardians[0]
	}
	if len(gamer.Guardians) >= 2 {
		s.Two = gamer.Guardians[1]
	}
	if len(gamer.Guardians) == 3 {
		s.Three = gamer.Guardians[2]
	}
}

func (s *Soul) Possess(id string) {
	s.Possessed = id
}

func (s *Soul) Charge(guardian bungo.Guardian) {
	switch guardian.Id {
	case s.One.Id:
		s.One = guardian
	case s.Two.Id:
		s.Two = guardian
	case s.Three.Id:
		s.Three = guardian
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

func (s Soul) Summon() (out bungo.Guardian) {
	switch s.Possessed {
	case s.One.Id:
		out = s.One
	case s.Two.Id:
		out = s.Two
	case s.Three.Id:
		out = s.Three
	}
	return out
}

func (s Soul) SummonOthers() (out []bungo.Guardian) {
	if s.Possessed != s.One.Id {
		out = append(out, s.One)
	}
	if s.Possessed != s.Two.Id {
		out = append(out, s.Two)
	}
	if s.Possessed != s.Three.Id {
		out = append(out, s.Three)
	}
	return out
}
func (s Soul) Call(id string) (out bungo.Guardian, exists bool) {
	guardians, exists := s.Callx(id)
	if exists {
		out = guardians[0]
		exists = true
	} else if len(guardians) > 1 {
		log.Fatal("found too many guardians for id: ", id)
	}
	return out, exists
}

func (s Soul) Callx(ids ...string) (out []bungo.Guardian, exists bool) {
	for _, id := range ids {
		switch id {
		case s.One.Id:
			out = append(out, s.One)
			exists = true
		case s.Two.Id:
			out = append(out, s.Two)
			exists = true
		case s.Three.Id:
			out = append(out, s.Three)
			exists = true
		}
	}
	return out, exists
}

func (g Ghost) Save(redis tower.Redis) {
	prettyJson := g.PrettyJson()

	redis.Connect()
	defer redis.Disconnect()

	redis.Set(g.Key(), prettyJson)

	log.Debugf("ghost save: id=%s session_id=%s", g.Id, g.SessionId)
}

func (g Ghost) PrettyJson() (out []byte) {
	gout, _ := json.Marshal(g)
	var prettyJSON bytes.Buffer
	_ = json.Indent(&prettyJSON, gout, "", "\t")
	return prettyJSON.Bytes()
}
