package brightsvc

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type NewBungieClientArgs struct {
	BungieNetUrl string
	Host         string
	Base         string
	ApiKey       string
	AppVersion   string
	AppId        string
	AppUrl       string
	AppEmail     string
}

func NewBungieClient(args NewBungieClientArgs) *resty.Client {
	rClient := resty.New()
	rClient.HostURL = args.BungieNetUrl
	rClient.SetHeaders(map[string]string{
		"X-Api-Key":  args.ApiKey,
		"User-Agent": fmt.Sprintf("Bright/%s AppId/%s (+%s;%s)", args.AppVersion, args.AppId, args.AppUrl, args.AppEmail),
	})
	return rClient
}
