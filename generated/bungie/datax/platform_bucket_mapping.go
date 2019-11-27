package datax

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shihtzu-systems/bright/generated/bungie/data"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
)

func LoadPlatformBucketMapping(destiny *data.Content, redis tower.Redis, overwrite bool) {
	log.Infof("PlatformBucketMapping [%d]", len(destiny.PlatformBucketMapping.Values()))
	if !redis.Exists(data.PlatformBucketMappingDefinitionName) {
		log.Infof("Loading %d PlatformBucketMapping", len(destiny.PlatformBucketMapping.Values()))
		for _, definition := range destiny.PlatformBucketMapping.Values() {
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

func GetPlatformBucketMapping(hash string, redis tower.Redis) (out data.PlatformBucketMappingDefinition) {
	rawJson := redis.HGet(destinyContentKey+":"+out.Name(), hash)
	if rawJson != "" {
		err := json.Unmarshal([]byte(rawJson), &out)
		if err != nil {
			log.Fatal(err)
		}
	}
	return out
}
