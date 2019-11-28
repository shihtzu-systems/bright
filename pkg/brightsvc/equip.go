package brightsvc

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/shihtzu-systems/bright/pkg/ghost"
	"github.com/shihtzu-systems/bright/pkg/tower"

	log "github.com/sirupsen/logrus"
)

type EquipItemsArgs struct {
	InstanceIds []string

	Tower        tower.Tower
	Ghost        ghost.Ghost
	BungieClient *resty.Client
}

func EquipItems(args EquipItemsArgs) {
	freshToken := Refresh(args.BungieClient, args.Ghost, args.Tower)

	soul := args.Ghost.Bright.Summon()

	body := struct {
		ItemIds        []string `json:"itemIds"`
		CharacterId    string   `json:"characterId"`
		MembershipType int      `json:"membershipType"`
	}{
		args.InstanceIds,
		soul.Id,
		soul.MembershipType,
	}
	jout, _ := json.Marshal(&body)

	equipItemsResponse, err := args.BungieClient.R().
		EnableTrace().
		SetAuthToken(freshToken.AccessToken).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
		}).
		SetBody(jout).
		Post("/Platform/Destiny2/Actions/Items/EquipItems/")
	if err != nil {
		log.Fatal(err)
	}
	log.Trace(string(equipItemsResponse.Body()))
}

type EquipItemArgs struct {
	InstanceId string

	Tower        tower.Tower
	Ghost        ghost.Ghost
	BungieClient *resty.Client
}

func EquipItem(args EquipItemArgs) {
	EquipItems(EquipItemsArgs{
		InstanceIds:  []string{args.InstanceId},
		Tower:        args.Tower,
		Ghost:        args.Ghost,
		BungieClient: args.BungieClient,
	})
}

func Refresh(bungieClient *resty.Client, ghost ghost.Ghost, tower tower.Tower) (out ghost.BungieToken) {
	resp, err := bungieClient.R().
		EnableTrace().
		SetBasicAuth(tower.Oauth.ClientId, tower.Oauth.ClientSecret).
		SetFormData(map[string]string{
			"grant_type":    "refresh_token",
			"refresh_token": ghost.Token.RefreshToken,
		}).
		SetContentLength(true).
		Post("/Platform/App/OAuth/Token/")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(resp.Body(), &out); err != nil {
		log.Warn(string(resp.Body()))
		log.Fatal(err)
	}
	return out
}
