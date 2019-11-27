package brightx

import (
	"github.com/go-resty/resty/v2"
	"github.com/shihtzu-systems/bright/pkg/destinyx"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
)

const (
	destinyContentIdKey = "destiny:content:id"
)

type DestinyArgs struct {
	Language        string
	SourceBasePath  string
	WorkingBasePath string
	BungieNetHost   string
	Overwrite       bool
	BungieClient    *resty.Client
	Tower           tower.Tower
}

func Destiny(args DestinyArgs) {
	destiny, contentId, err := destinyx.NewDestinyContent(destinyx.NewDestinyContentArgs{
		Language:      args.Language,
		BungieNetHost: args.BungieNetHost,
		BungieClient:  args.BungieClient,
	})
	if err != nil {
		log.Fatal(err)
	}
	args.Tower.Redis.Connect()
	loadedContentId := args.Tower.Redis.Get(destinyContentIdKey)
	args.Tower.Redis.Disconnect()

	overwrite := args.Overwrite || loadedContentId != contentId
	destinyx.Load(destiny, args.Tower.Redis, overwrite)

	args.Tower.Redis.Connect()
	args.Tower.Redis.Set(destinyContentIdKey, []byte(contentId))
	args.Tower.Redis.Disconnect()
}
