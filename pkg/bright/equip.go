package bright

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/shihtzu-systems/bright/generated/bungie/data"

	log "github.com/sirupsen/logrus"
)

type EquipItemsArgs struct {
	CharacterId    string
	InstanceIds    []string
	MembershipType int

	OathClientId     string
	OathClientSecret string
	Auth             BungieToken
	BungieClient     *resty.Client
	Destiny          data.Content
}

func EquipItems(args EquipItemsArgs) (err error) {
	freshToken, err := refresh(args.BungieClient, args.OathClientId, args.OathClientSecret, args.Auth)
	if err != nil {
		return err
	}

	body := struct {
		ItemIds        []string `json:"itemIds"`
		CharacterId    string   `json:"characterId"`
		MembershipType int      `json:"membershipType"`
	}{
		args.InstanceIds,
		args.CharacterId,
		args.MembershipType,
	}
	jout, _ := json.Marshal(&body)

	log.Infof("POST %s/Destiny2/Actions/Items/EquipItems/", bungieApiUri)
	log.Debug(string(jout))

	equipItemsResponse, err := args.BungieClient.R().
		EnableTrace().
		SetAuthToken(freshToken.AccessToken).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
		}).
		SetBody(jout).
		Post(bungieApiUri + "/Destiny2/Actions/Items/EquipItems/")

	if err != nil {
		return err
	}
	log.Trace(string(equipItemsResponse.Body()))

	return nil
}

type EquipItemArgs struct {
	CharacterId    string
	InstanceId     string
	MembershipType int

	OathClientId     string
	OathClientSecret string
	Auth             BungieToken
	BungieClient     *resty.Client
	Destiny          data.Content
}

func EquipItem(args EquipItemArgs) (err error) {
	return EquipItems(EquipItemsArgs{
		CharacterId:      args.CharacterId,
		InstanceIds:      []string{args.InstanceId},
		MembershipType:   args.MembershipType,
		OathClientId:     args.OathClientId,
		OathClientSecret: args.OathClientSecret,
		Auth:             args.Auth,
		BungieClient:     args.BungieClient,
		Destiny:          args.Destiny,
	})
}
