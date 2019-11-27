package datax

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shihtzu-systems/bright/generated/bungie/data"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
)

func LoadMaterialRequirementSet(destiny *data.Content, redis tower.Redis, overwrite bool) {
	log.Infof("MaterialRequirementSet [%d]", len(destiny.MaterialRequirementSet.Values()))
	if !redis.Exists(data.MaterialRequirementSetDefinitionName) {
		log.Infof("Loading %d MaterialRequirementSet", len(destiny.MaterialRequirementSet.Values()))
		for _, definition := range destiny.MaterialRequirementSet.Values() {
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

func GetMaterialRequirementSet(hash int, redis tower.Redis) (out data.MaterialRequirementSetDefinition) {
	rawJson := redis.HGet(destinyContentKey+":"+out.Name(), fmt.Sprint(hash))
	if rawJson != "" {
		err := json.Unmarshal([]byte(rawJson), &out)
		if err != nil {
			log.Fatal(err)
		}
	}
	return out
}
