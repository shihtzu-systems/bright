package datax

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shihtzu-systems/bright/generated/bungie/data"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
)

func LoadRewardAdjusterProgressionMap(destiny *data.Content, redis tower.Redis, overwrite bool) {
	log.Infof("RewardAdjusterProgressionMap [%d]", len(destiny.RewardAdjusterProgressionMap.Values()))
	if !redis.Exists(data.RewardAdjusterProgressionMapDefinitionName) {
		log.Infof("Loading %d RewardAdjusterProgressionMap", len(destiny.RewardAdjusterProgressionMap.Values()))
		for _, definition := range destiny.RewardAdjusterProgressionMap.Values() {
			if overwrite {
				redis.HSet(destinyContentKey+":"+definition.Name(), fmt.Sprint(definition.Hash), []byte{})
			}

			definitionJson := redis.HGet(destinyContentKey+":"+definition.Name(), fmt.Sprint(definition.Hash))
			if definitionJson != "" {
				continue
			}

			dout, err := json.Marshal(definition)
			if err != nil {
				log.Fatal(err)
			}

			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, dout, "", "\t"); err != nil {
				log.Fatal(err)
			}

			redis.HSet(destinyContentKey+":"+definition.Name(), fmt.Sprint(definition.Hash), prettyJSON.Bytes())
		}
	}
}

func GetRewardAdjusterProgressionMap(hash int, redis tower.Redis) (out data.RewardAdjusterProgressionMapDefinition) {
	rawJson := redis.HGet(destinyContentKey+":"+out.Name(), fmt.Sprint(hash))
	if rawJson != "" {
		err := json.Unmarshal([]byte(rawJson), &out)
		if err != nil {
			log.Fatal(err)
		}
	}
	return out
}
