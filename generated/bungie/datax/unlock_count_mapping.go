package datax

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shihtzu-systems/bright/generated/bungie/data"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
)

func LoadUnlockCountMapping(destiny *data.Content, redis tower.Redis, overwrite bool) {
	log.Infof("UnlockCountMapping [%d]", len(destiny.UnlockCountMapping.Values()))
	if !redis.Exists(data.UnlockCountMappingDefinitionName) {
		log.Infof("Loading %d UnlockCountMapping", len(destiny.UnlockCountMapping.Values()))
		for _, definition := range destiny.UnlockCountMapping.Values() {
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

func GetUnlockCountMapping(hash string, redis tower.Redis) (out data.UnlockCountMappingDefinition) {
	rawJson := redis.HGet(destinyContentKey+":"+out.Name(), hash)
	if rawJson != "" {
		err := json.Unmarshal([]byte(rawJson), &out)
		if err != nil {
			log.Fatal(err)
		}
	}
	return out
}
