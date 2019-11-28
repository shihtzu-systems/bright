package bright

import (
	"fmt"
	haikunator "github.com/atrox/haikunatorgo/v2"
	"github.com/go-resty/resty/v2"
	"strings"
)

type NewBungieClientArgs struct {
	Host       string
	Base       string
	ApiKey     string
	AppVersion string
	AppId      string
	AppUrl     string
	AppEmail   string
}

func NewBungieClient(args NewBungieClientArgs) *resty.Client {
	rClient := resty.New()
	rClient.HostURL = bungieUrl
	// global
	rClient.SetHeaders(map[string]string{
		"X-Api-Key":   args.ApiKey,
		"Gamer-Agent": fmt.Sprintf("Bright/%s AppId/%s (+%s;%s)", args.AppVersion, args.AppId, args.AppUrl, args.AppEmail),
	})
	rClient.SetOutputDirectory("out/responses")
	return rClient
}

func GenerateName(dad bool) string {
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
