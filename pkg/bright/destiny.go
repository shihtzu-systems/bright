package bright

import (
	"encoding/json"
	"github.com/shihtzu-systems/bright/generated/bungie/data"
	"github.com/shihtzu-systems/bright/pkg/fsystem"
	"path"
)

func DestinyContentPath(workingBasePath string) (out string, err error) {
	fs, _, err := fsystem.FileExists(path.Join(workingBasePath, "current"))
	if err != nil {
		return "", err
	}

	if fs {
		cout, err := fsystem.ReadRoFile(path.Join(workingBasePath, "current"))
		if err != nil {
			return "", err
		}
		out = path.Join(workingBasePath, string(cout))
	}
	return out, nil
}

type Destiny struct {
	basePath string
}

func (d Destiny) FindInventoryItem(id string) (out data.InventoryItemDefinition, err error) {
	jsonPath := path.Join(d.basePath, id+".json")
	fs, _, err := fsystem.FileExists(jsonPath)
	if err != nil {
		return data.InventoryItemDefinition{}, err
	}

	if fs {
		jout, err := fsystem.ReadRoFile(jsonPath)
		if err != nil {
			return data.InventoryItemDefinition{}, err
		}

		if err := json.Unmarshal(jout, &out); err != nil {
			return data.InventoryItemDefinition{}, err
		}
	}
	return out, nil
}
