package datax

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shihtzu-systems/bright/generated/bungie/data"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
)

func LoadProgression(destiny *data.Content, redis tower.Redis, overwrite bool) {
	log.Infof("Progression [%d]", len(destiny.Progression.Values()))
	if !redis.Exists(data.ProgressionDefinitionName) {
		log.Infof("Loading %d Progression", len(destiny.Progression.Values()))
		for _, definition := range destiny.Progression.Values() {
			if overwrite {
				redis.HSet(definition.Name(), fmt.Sprint(definition.Hash), []byte{})
			}

			definitionJson := redis.HGet(definition.Name(), fmt.Sprint(definition.Hash))
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

			redis.HSet(definition.Name(), fmt.Sprint(definition.Hash), prettyJSON.Bytes())
		}
	}
}

func GetProgression(hash string, redis tower.Redis) (out data.ProgressionDefinition) {
	rawJson := redis.HGet(out.Name(), hash)
	if rawJson != "" {
		err := json.Unmarshal([]byte(rawJson), &out)
		if err != nil {
			log.Fatal(err)
		}
	}
	return out
}
