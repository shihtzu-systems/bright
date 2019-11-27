// Code generated; DANGER ZONE FOR EDITS

package data

import (
    "fmt"
	"bytes"
	"encoding/json"
	"gopkg.in/yaml.v2"
)

const ${prefix}DefinitionName = "${name}"

type ${prefix}Definitions map[string]${prefix}Definition

func (d ${prefix}Definitions) Keys() (out []string) {
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (d ${prefix}Definitions) Values() (out []${prefix}Definition) {
	for _, v := range d {
		out = append(out, v)
	}
	return out
}

func (d ${prefix}Definitions) Find(id int) (out ${prefix}Definition) {
	for k, v := range d {
		if k == fmt.Sprint(id) {
			return v
		}
	}
	return ${prefix}Definition{}
}

func (d ${prefix}Definitions) Name() string {
	return ${prefix}DefinitionName
}

${structsContent}

func (d ${prefix}Definition) Name() string {
	return ${prefix}DefinitionName
}

func (d ${prefix}Definition) Json() ([]byte, error) {
	return json.Marshal(d)
}

func (d ${prefix}Definition) PrettyJson() ([]byte, error) {
	jout, err := d.Json()
	if err != nil {
		return nil, err
	}
	var pretty bytes.Buffer
	if err := json.Indent(&pretty, jout, "", "  "); err != nil {
		return nil, err
	}
	return pretty.Bytes(), nil
}

func (d ${prefix}Definition) Yaml() ([]byte, error) {
	return yaml.Marshal(d)
}