package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/shihtzu-systems/bright/generated/bungie/data"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
	"html/template"
	"os"
	"strings"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "", "")
	flag.Parse()

	t := template.Must(template.New("").Parse(loaderTemplate))

	if err := t.Execute(os.Stdout, struct {
		Name string
	}{
		strings.Title(name),
	}); err != nil {
		log.Fatal(err)
	}
}

func Load(destiny *data.Content, redis tower.Redis, overwrite bool) {
	log.Infof("Achievement [%d]", len(destiny.Achievement.Values()))
	if !redis.Exists(data.AchievementDefinitionName) {
		log.Infof("Loading %d Achievements", len(destiny.Achievement.Values()))
		for _, definition := range destiny.Achievement.Values() {
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

var loaderTemplate = `package datax

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shihtzu-systems/bright/generated/bungie/data"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
)

func Load{{ .Name }}(destiny *data.Content, redis tower.Redis, overwrite bool)  {
	log.Infof("{{ .Name }} [%d]", len(destiny.{{ .Name }}.Values()))
	if !redis.Exists(data.{{ .Name }}DefinitionName) {
		log.Infof("Loading %d {{ .Name }}", len(destiny.{{ .Name }}.Values()))
		for _, definition := range destiny.{{ .Name }}.Values() {
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

func Get{{ .Name }}(hash string, redis tower.Redis) (out data.{{ .Name }}Definition) {
	rawJson := redis.HGet(out.Name(), hash)
	if rawJson != "" {
		err := json.Unmarshal([]byte(rawJson), &out)
		if err != nil {
			log.Fatal(err)
		}
	}
	return out
}
`
