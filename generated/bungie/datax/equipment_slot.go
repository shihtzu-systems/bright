package datax

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shihtzu-systems/bright/generated/bungie/data"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
)

func LoadEquipmentSlot(destiny *data.Content, redis tower.Redis, overwrite bool) {
	log.Infof("EquipmentSlot [%d]", len(destiny.EquipmentSlot.Values()))
	if !redis.Exists(data.EquipmentSlotDefinitionName) {
		log.Infof("Loading %d EquipmentSlot", len(destiny.EquipmentSlot.Values()))
		for _, definition := range destiny.EquipmentSlot.Values() {
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

func GetEquipmentSlot(hash string, redis tower.Redis) (out data.EquipmentSlotDefinition) {
	rawJson := redis.HGet(destinyContentKey+":"+out.Name(), hash)
	if rawJson != "" {
		err := json.Unmarshal([]byte(rawJson), &out)
		if err != nil {
			log.Fatal(err)
		}
	}
	return out
}
