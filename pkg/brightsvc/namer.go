package brightsvc

import (
	haikunator "github.com/atrox/haikunatorgo/v2"
	"math/rand"
	"strings"
	"time"
)

func NewName(dad bool) string {
	randomGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	namer := haikunator.New()
	namer.TokenLength = randomGen.Intn(4)
	delimiters := []string{"-", "_", ".", ""}
	namer.Delimiter = delimiters[randomGen.Intn(len(delimiters))]
	name := namer.Haikunate()

	skat8r := "420"
	sexy := "69"
	satan := "666"

	if dad {
		skat8r = "360"
		satan = "777"
		sexy = "42"
	}

	switch randomGen.Intn(10) {
	case 0:
		name = strings.ToUpper(name)
	case 1:
		if namer.TokenLength == 0 {
			name = name + skat8r
		}
		name = strings.ReplaceAll(name, "t", "7")
	case 2:
		name = "x" + name + "x"
		name = strings.ReplaceAll(name, "e", "3")
	case 3:
		name = strings.ToTitle(name)
	case 4:
		name = strings.ToLower(name)
		name = strings.ReplaceAll(name, "l", "1")
	case 5:
		name = strings.ToLower(name)
		name = strings.ReplaceAll(name, "i", "!")
	case 6:
		if namer.TokenLength == 0 {
			name = name + sexy
		}
		name = strings.ReplaceAll(name, "o", "00")
	case 7:
		if namer.TokenLength == 0 {
			name = name + satan
		}
		name = strings.ReplaceAll(name, "ie", "ei")
	case 9:
		if namer.TokenLength == 0 {
			name = name + sexy
		}
		name = name + "!!"
	default:
	}
	return name
}
