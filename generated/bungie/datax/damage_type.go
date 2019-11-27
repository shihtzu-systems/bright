package datax

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shihtzu-systems/bright/generated/bungie/data"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
)

func LoadDamageType(destiny *data.Content, redis tower.Redis, overwrite bool) {
	log.Infof("DamageType [%d]", len(destiny.DamageType.Values()))
	if !redis.Exists(data.DamageTypeDefinitionName) {
		log.Infof("Loading %d DamageType", len(destiny.DamageType.Values()))
		for _, definition := range destiny.DamageType.Values() {
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

func GetDamageType(hash int, redis tower.Redis) (out data.DamageTypeDefinition) {
	rawJson := redis.HGet(destinyContentKey+":"+out.Name(), fmt.Sprint(hash))
	if rawJson != "" {
		err := json.Unmarshal([]byte(rawJson), &out)
		if err != nil {
			log.Fatal(err)
		}
	}
	return out
}
