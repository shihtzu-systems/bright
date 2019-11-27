package brightsvc

import (
	"github.com/shihtzu-systems/bright/pkg/bungo"
	"github.com/shihtzu-systems/bright/pkg/ghost"
	"math/rand"
	"time"
)

func NewTryUser(dad bool) (out ghost.TryUser) {
	return ghost.TryUser{
		Name:         NewName(dad),
		MembershipId: "0",
		Characters: []bungo.Character{
			newTitan(),
			newHunter(),
			newWarlock(),
		},
	}
}

func newTitan() bungo.Character {
	randomGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	return bungo.Character{
		Id:             "titan",
		MembershipType: 0,
		Class:          "Titan",
		Level:          50,
		Light:          750 + randomGen.Intn(170),
	}
}

func newHunter() bungo.Character {
	randomGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	return bungo.Character{
		Id:             "hunter",
		MembershipType: 0,
		Class:          "Hunter",
		Level:          50,
		Light:          750 + randomGen.Intn(170),
	}
}

func newWarlock() bungo.Character {
	randomGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	return bungo.Character{
		Id:             "warlock",
		MembershipType: 0,
		Class:          "Warlock",
		Level:          50,
		Light:          750 + randomGen.Intn(170),
	}
}
