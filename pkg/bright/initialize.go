package bright

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/shihtzu-systems/bright/generated/bungie/data"
	"github.com/shihtzu-systems/bright/pkg/bungo"
	"github.com/shihtzu-systems/bright/pkg/fsystem"
	log "github.com/sirupsen/logrus"
	"path"
	"path/filepath"
	"strings"
)

type InitializeArgs struct {
	Language        string
	SourceBasePath  string
	WorkingBasePath string
	BungieNetHost   string
	BungieClient    *resty.Client
	Overwrite       bool
}

func Initialize(args InitializeArgs) (err error) {
	destiny, contentId, err := NewDestinyContent(NewDestinyContentArgs{
		Language:      args.Language,
		BungieNetHost: args.BungieNetHost,
		BungieClient:  args.BungieClient,

		SourceBasePath:  args.SourceBasePath,
		WorkingBasePath: args.WorkingBasePath,
	})
	if err != nil {
		return err
	}

	workingBasePath := path.Join(args.WorkingBasePath, contentId)
	log.Infof("saving denormalized destiny content [%s] to %s", contentId, workingBasePath)
	if err := denormalizeDefinitions(destiny, workingBasePath, args.Overwrite); err != nil {
		return err
	}

	if err := fsystem.WriteFsFile(path.Join(args.WorkingBasePath, "current"), []byte(contentId)); err != nil {
		return err
	}
	return nil
}

type NewDestinyContentArgs struct {
	Language      string
	BungieNetHost string
	BungieClient  *resty.Client

	SourceBasePath  string
	WorkingBasePath string
}

func NewDestinyContent(args NewDestinyContentArgs) (destiny *data.Content, contentId string, err error) {
	getDestinyManifestResponse, err := args.BungieClient.R().
		Get(bungieApiUri + "/Destiny2/Manifest/")
	if err != nil {
		return nil, "", err
	}
	var getDestinyManifest bungo.GetDestinyManifest
	if err := json.Unmarshal(getDestinyManifestResponse.Body(), &getDestinyManifest); err != nil {
		return nil, "", err
	}

	currentDestinyContentPath := getDestinyManifest.Response.JSONWorldContentPaths.En
	currentDestinyContentFilename := filepath.Base(currentDestinyContentPath)

	contentId = strings.TrimSuffix(currentDestinyContentFilename, ".json")

	localSourceContentFile := path.Join(args.SourceBasePath, currentDestinyContentFilename)

	fsExists, _, _ := fsystem.FileExists(localSourceContentFile)
	if fsExists {
		log.Infof("destiny content file found: %s", localSourceContentFile)
		dout, err := fsystem.ReadFsFile(localSourceContentFile)
		if err != nil {
			return nil, contentId, err
		}

		err = json.Unmarshal(dout, &destiny)
		if err != nil {
			return nil, contentId, err
		}

	} else {
		log.Info("destiny content from bungie: GET https://www.bungie.net/Platform/" + currentDestinyContentPath)
		contentResp, err := args.BungieClient.R().
			Get("/" + currentDestinyContentPath)
		if err != nil {
			return nil, contentId, err
		}

		// load to memory
		if err := fsystem.WriteMemFile(currentDestinyContentFilename, contentResp.Body()); err != nil {
			return nil, contentId, err
		}
		memout, err := fsystem.ReadMemFile(currentDestinyContentFilename)
		if err != nil {
			return nil, contentId, err
		}
		// save to file
		if err := fsystem.WriteFsFile(localSourceContentFile, memout); err != nil {
			return nil, contentId, err
		}

		err = json.Unmarshal(memout, &destiny)
		if err != nil {
			return nil, contentId, err
		}
	}
	return destiny, contentId, nil
}

func denormalizeDefinitions(destiny *data.Content, basePath string, overwrite bool) (err error) {
	for _, definition := range destiny.Achievement.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)

		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "\t"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}

	for _, definition := range destiny.Activity.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.ActivityGraph.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.ActivityInteractable.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.ActivityMode.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.ActivityModifier.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.ActivityType.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.ArtDyeChannel.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.ArtDyeReference.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.Artifact.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.Bond.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.BreakerType.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.CharacterCustomizationCategory.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.CharacterCustomizationOption.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.Checklist.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.Class.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.Collectible.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.DamageType.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.Destination.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.EnemyRace.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.EnergyType.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.EntitlementOffer.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.EquipmentSlot.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.Faction.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.Gender.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.InventoryBucket.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.InventoryItem.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.ItemCategory.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.ItemTierType.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.Location.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.Lore.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.MaterialRequirementSet.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.MedalTier.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.Milestone.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.NodeStepSummary.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), 0)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.Objective.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.Place.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.PlatformBucketMapping.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.PlugSet.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.PresentationNode.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.Progression.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.ProgressionLevelRequirement.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.ProgressionMapping.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.Race.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.Record.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.ReportReasonCategory.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.RewardAdjusterPointer.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.RewardAdjusterProgressionMap.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.RewardItemList.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.RewardMapping.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.RewardSheet.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.RewardSource.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), 0)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.SackRewardItemList.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.SandboxPattern.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.SandboxPerk.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.Season.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.SeasonPass.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.SocketCategory.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.SocketType.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.Stat.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.StatGroup.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.TalentGrid.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.UnlockCountMapping.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.Unlock.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.UnlockEvent.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.UnlockExpressionMapping.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.UnlockValue.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.Vendor.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	for _, definition := range destiny.VendorGroup.Values() {
		definitionPath := fmt.Sprintf("%s/%s/%d.json", basePath, definition.Name(), definition.Hash)
		fs, _, err := fsystem.FileExists(definitionPath)
		if fs && !overwrite && err == nil {
			continue
		}

		dout, err := json.Marshal(definition)
		if err != nil {
			return err
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, dout, "", "	"); err != nil {
			return err
		}

		if err := fsystem.WriteFsFile(definitionPath, prettyJSON.Bytes()); err != nil {
			return err
		}
	}
	return nil
}
