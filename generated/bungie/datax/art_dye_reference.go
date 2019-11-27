package datax

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shihtzu-systems/bright/generated/bungie/data"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
)

func LoadArtDyeReference(destiny *data.Content, redis tower.Redis, overwrite bool) {
	log.Infof("ArtDyeReference [%d]", len(destiny.ArtDyeReference.Values()))
	if !redis.Exists(data.ArtDyeReferenceDefinitionName) {
		log.Infof("Loading %d ArtDyeReference", len(destiny.ArtDyeReference.Values()))
		for _, definition := range destiny.ArtDyeReference.Values() {
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

func GetArtDyeReference(hash int, redis tower.Redis) (out data.ArtDyeReferenceDefinition) {
	rawJson := redis.HGet(destinyContentKey+":"+out.Name(), fmt.Sprint(hash))
	if rawJson != "" {
		err := json.Unmarshal([]byte(rawJson), &out)
		if err != nil {
			log.Fatal(err)
		}
	}
	return out
}
