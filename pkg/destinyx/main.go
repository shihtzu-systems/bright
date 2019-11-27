package destinyx

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/shihtzu-systems/bright/generated/bungie/data"
	"github.com/shihtzu-systems/bright/generated/bungie/datax"
	"github.com/shihtzu-systems/bright/pkg/bungo"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
)

type NewDestinyContentArgs struct {
	Language      string
	BungieNetHost string
	BungieClient  *resty.Client
}

func NewDestinyContent(args NewDestinyContentArgs) (destiny *data.Content, contentId string, err error) {
	getDestinyManifestResponse, err := args.BungieClient.R().
		Get("/Platform/Destiny2/Manifest/")
	if err != nil {
		return nil, "", err
	}
	var getDestinyManifest bungo.GetDestinyManifest
	if err := json.Unmarshal(getDestinyManifestResponse.Body(), &getDestinyManifest); err != nil {
		return nil, "", err
	}

	contentPath := getDestinyManifest.Response.JSONWorldContentPaths.En
	contentFilename := filepath.Base(contentPath)
	contentId = strings.TrimSuffix(contentFilename, ".json")
	log.Info("destiny content from bungie: GET https://www.bungie.net/Platform" + contentPath)
	contentResp, err := args.BungieClient.R().
		Get("/" + contentPath)
	if err != nil {
		return nil, contentId, err
	}
	err = json.Unmarshal(contentResp.Body(), &destiny)
	if err != nil {
		return nil, contentId, err
	}
	return destiny, contentId, nil
}

func Load(destiny *data.Content, redis tower.Redis, overwrite bool) {
	redis.Connect()
	datax.LoadAchievement(destiny, redis, overwrite)
	datax.LoadActivity(destiny, redis, overwrite)
	datax.LoadActivityGraph(destiny, redis, overwrite)
	datax.LoadActivityInteractable(destiny, redis, overwrite)
	datax.LoadActivityMode(destiny, redis, overwrite)
	datax.LoadActivityModifier(destiny, redis, overwrite)
	datax.LoadActivityType(destiny, redis, overwrite)
	datax.LoadArtDyeChannel(destiny, redis, overwrite)
	datax.LoadArtDyeReference(destiny, redis, overwrite)
	datax.LoadArtifact(destiny, redis, overwrite)
	datax.LoadBond(destiny, redis, overwrite)
	datax.LoadBreakerType(destiny, redis, overwrite)
	datax.LoadCharacterCustomizationCategory(destiny, redis, overwrite)
	datax.LoadCharacterCustomizationOption(destiny, redis, overwrite)
	datax.LoadChecklist(destiny, redis, overwrite)
	datax.LoadClass(destiny, redis, overwrite)
	datax.LoadCollectible(destiny, redis, overwrite)
	datax.LoadDamageType(destiny, redis, overwrite)
	datax.LoadDestination(destiny, redis, overwrite)
	datax.LoadEnemyRace(destiny, redis, overwrite)
	datax.LoadEnergyType(destiny, redis, overwrite)
	datax.LoadEntitlementOffer(destiny, redis, overwrite)
	datax.LoadEquipmentSlot(destiny, redis, overwrite)
	datax.LoadFaction(destiny, redis, overwrite)
	datax.LoadGender(destiny, redis, overwrite)
	datax.LoadInventoryBucket(destiny, redis, overwrite)
	datax.LoadInventoryItem(destiny, redis, overwrite)
	datax.LoadItemCategory(destiny, redis, overwrite)
	datax.LoadItemTierType(destiny, redis, overwrite)
	datax.LoadLocation(destiny, redis, overwrite)
	datax.LoadLore(destiny, redis, overwrite)
	datax.LoadMaterialRequirementSet(destiny, redis, overwrite)
	datax.LoadMedalTier(destiny, redis, overwrite)
	datax.LoadMilestone(destiny, redis, overwrite)
	//datax.LoadNodeStepSummary(destiny, redis, overwrite)
	datax.LoadObjective(destiny, redis, overwrite)
	datax.LoadPlace(destiny, redis, overwrite)
	datax.LoadPlatformBucketMapping(destiny, redis, overwrite)
	datax.LoadPlugSet(destiny, redis, overwrite)
	datax.LoadPresentationNode(destiny, redis, overwrite)
	datax.LoadProgression(destiny, redis, overwrite)
	datax.LoadProgressionLevelRequirement(destiny, redis, overwrite)
	datax.LoadProgressionMapping(destiny, redis, overwrite)
	datax.LoadRace(destiny, redis, overwrite)
	datax.LoadRecord(destiny, redis, overwrite)
	datax.LoadReportReasonCategory(destiny, redis, overwrite)
	datax.LoadRewardAdjusterPointer(destiny, redis, overwrite)
	datax.LoadRewardAdjusterProgressionMap(destiny, redis, overwrite)
	datax.LoadRewardItemList(destiny, redis, overwrite)
	datax.LoadRewardMapping(destiny, redis, overwrite)
	datax.LoadRewardSheet(destiny, redis, overwrite)
	//datax.LoadRewardSource(destiny, redis, overwrite)
	datax.LoadSackRewardItemList(destiny, redis, overwrite)
	datax.LoadSandboxPattern(destiny, redis, overwrite)
	datax.LoadSandboxPerk(destiny, redis, overwrite)
	datax.LoadSeason(destiny, redis, overwrite)
	datax.LoadSeasonPass(destiny, redis, overwrite)
	datax.LoadSocketCategory(destiny, redis, overwrite)
	datax.LoadSocketType(destiny, redis, overwrite)
	datax.LoadStat(destiny, redis, overwrite)
	datax.LoadStatGroup(destiny, redis, overwrite)
	datax.LoadTalentGrid(destiny, redis, overwrite)
	datax.LoadUnlockCountMapping(destiny, redis, overwrite)
	datax.LoadUnlock(destiny, redis, overwrite)
	datax.LoadUnlockEvent(destiny, redis, overwrite)
	datax.LoadUnlockExpressionMapping(destiny, redis, overwrite)
	datax.LoadUnlockValue(destiny, redis, overwrite)
	datax.LoadVendor(destiny, redis, overwrite)
	datax.LoadVendorGroup(destiny, redis, overwrite)
	redis.Disconnect()
}
