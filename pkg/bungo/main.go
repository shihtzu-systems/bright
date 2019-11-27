package bungo

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type CurrentUser struct {
	Name         string
	MembershipId string
	Characters   []Character
}

type Character struct {
	Id             string `yaml:"id" json:"id"`
	MembershipType int    `yaml:"membershipType" json:"membership_type"`

	Class string `yaml:"class" json:"class"`
	Level int    `yaml:"level" json:"level"`
	Light int    `yaml:"light" json:"light"`

	Emblem Emblem `yaml:"emblem" json:"emblem"`

	Bag      Bag    `yaml:"bag" json:"bag"`
	Equipped Outfit `yaml:"equipped" json:"equipped"`
}

func (c Character) String() string {
	return fmt.Sprintf("id=%s\nemblem=%s\nclass=%s\nkinetic=%s\nenergy=%s\npower=%s\nclass-item=%s\narmors=%d\nweapons=%d",
		c.Id, c.Emblem.Name, c.Class, c.Equipped.KineticWeapon.Name, c.Equipped.EnergyWeapon.Name, c.Equipped.PowerWeapon.Name, c.Equipped.Class.Name, len(c.Bag.Armors), len(c.Bag.Weapons))
}

func (c Character) Differences(other Character) (out []string) {
	if c.Equipped.KineticWeapon.InstanceId != other.Equipped.KineticWeapon.InstanceId {
		out = append(out, c.Equipped.KineticWeapon.InstanceId)
	}
	if c.Equipped.EnergyWeapon.InstanceId != other.Equipped.EnergyWeapon.InstanceId {
		out = append(out, c.Equipped.EnergyWeapon.InstanceId)
	}
	if c.Equipped.PowerWeapon.InstanceId != other.Equipped.PowerWeapon.InstanceId {
		out = append(out, c.Equipped.PowerWeapon.InstanceId)
	}

	if c.Equipped.Helmet.InstanceId != other.Equipped.Helmet.InstanceId {
		out = append(out, c.Equipped.Helmet.InstanceId)
	}
	if c.Equipped.Gauntlets.InstanceId != other.Equipped.Gauntlets.InstanceId {
		out = append(out, c.Equipped.Gauntlets.InstanceId)
	}
	if c.Equipped.Chest.InstanceId != other.Equipped.Chest.InstanceId {
		out = append(out, c.Equipped.Chest.InstanceId)
	}
	if c.Equipped.Leg.InstanceId != other.Equipped.Leg.InstanceId {
		out = append(out, c.Equipped.Leg.InstanceId)
	}
	if c.Equipped.Class.InstanceId != other.Equipped.Class.InstanceId {
		out = append(out, c.Equipped.Class.InstanceId)
	}
	return out
}

type Outfit struct {
	KineticWeapon Weapon `yaml:"kineticWeapon" json:"kinetic_equipped"`
	EnergyWeapon  Weapon `yaml:"energyWeapon" json:"energy_equipped"`
	PowerWeapon   Weapon `yaml:"powerWeapon" json:"power_weapon"`

	Helmet    Armor `yaml:"helmet" json:"helmet"`
	Gauntlets Armor `yaml:"gauntlets" json:"gauntlets"`
	Chest     Armor `yaml:"chest" json:"chest"`
	Leg       Armor `yaml:"leg" json:"leg"`
	Class     Armor `yaml:"class" json:"class"`
}

func (o Outfit) FindWeapon(instanceId string) Weapon {
	log.Debug("Outfit->FindWeapon: ", instanceId)
	switch instanceId {
	case o.KineticWeapon.InstanceId:
		return o.KineticWeapon
	case o.EnergyWeapon.InstanceId:
		return o.EnergyWeapon
	case o.PowerWeapon.InstanceId:
		return o.PowerWeapon
	}
	return Weapon{}
}

func (o Outfit) FindArmor(instanceId string) Armor {
	log.Debug("Outfit->FindArmor: ", instanceId)
	switch instanceId {
	case o.Helmet.InstanceId:
		return o.Helmet
	case o.Gauntlets.InstanceId:
		return o.Gauntlets
	case o.Chest.InstanceId:
		return o.Chest
	case o.Leg.InstanceId:
		return o.Leg
	case o.Class.InstanceId:
		return o.Class
	}
	return Armor{}
}

type Bag struct {
	Weapons []Weapon `yaml:"weapons" json:"weapons"`
	Armors  []Armor  `yaml:"armors" json:"armor"`
}

func (b Bag) FindWeapon(instanceId string) Weapon {
	log.Debug("Bag->FindWeapon: ", instanceId)
	for _, weapon := range b.Weapons {
		if weapon.InstanceId == instanceId {
			return weapon
		}
	}
	return Weapon{}
}

func (b Bag) TakeWeapon(instanceId string) (Weapon, Bag) {
	log.Debug("Bag->TakeWeapon: ", instanceId)
	var found Weapon
	var out Bag
	for _, weapon := range b.Weapons {
		if weapon.InstanceId == instanceId {
			found = weapon
		} else {
			out.Weapons = append(out.Weapons, weapon)
		}
	}
	for _, armor := range b.Armors {
		out.Armors = append(out.Armors, armor)
	}
	return found, out
}

func (b Bag) StowWeapon(piece Weapon) Bag {
	log.Debug("Bag->StowWeapon: ", piece.InstanceId)
	var out Bag
	found := false
	for _, weapon := range b.Weapons {
		if weapon.InstanceId == piece.InstanceId {
			found = true
			break
		}
	}
	if !found {
		out.Weapons = append(b.Weapons, piece)
	}
	for _, armor := range b.Armors {
		out.Armors = append(out.Armors, armor)
	}
	return out
}

func (b Bag) FindWeapons(slot string) (out []Weapon) {
	log.Debugf("Bag->FindWeapons: %s [%d]", slot, len(b.Weapons))
	for _, weapon := range b.Weapons {
		if weapon.Slot.Name == slot {
			out = append(out, weapon)
		}
	}
	return out
}

func (b Bag) FindArmor(instanceId string) Armor {
	log.Debugf("Bag->FindArmor: %s [%d]", instanceId, len(b.Armors))
	for _, armor := range b.Armors {
		if armor.InstanceId == instanceId {
			return armor
		}
	}
	return Armor{}
}

func (b Bag) TakeArmor(instanceId string) (Armor, Bag) {
	log.Debug("Bag->TakeArmor: ", instanceId)
	var found Armor
	var out Bag
	for _, armor := range b.Armors {
		if armor.InstanceId == instanceId {
			found = armor
		} else {
			out.Armors = append(out.Armors, armor)
		}
	}
	for _, weapon := range b.Weapons {
		out.Weapons = append(out.Weapons, weapon)
	}
	return found, out
}

func (b Bag) StowArmor(piece Armor) Bag {
	log.Debug("Bag->StowArmor: ", piece.InstanceId)
	var out Bag
	found := false
	for _, armor := range b.Armors {
		if armor.InstanceId == piece.InstanceId {
			found = true
			break
		}
	}
	if !found {
		out.Armors = append(b.Armors, piece)
	}
	for _, weapon := range b.Weapons {
		out.Weapons = append(out.Weapons, weapon)
	}
	return out
}

func (b Bag) FindArmors(slot string) (out []Armor) {
	log.Debugf("Bag->FindArmors: %s [%d]", slot, len(b.Armors))

	for _, armor := range b.Armors {
		log.Debug(armor.Slot.Name, "==", slot)
		if armor.Slot.Name == slot {
			log.Debug("found: " + armor.Name)
			out = append(out, armor)
		}
	}
	return out
}

type Weapon Item

func (w Weapon) SwapType() string {
	log.Trace("swap type: ", w.Bucket)
	switch w.Bucket {
	case "Kinetic":
		fallthrough
	case "Kinetic Weapons":
		return "kinetic"
	case "Energy":
		fallthrough
	case "Energy Weapons":
		return "energy"
	case "Power":
		fallthrough
	case "Power Weapons":
		return "power"
	default:
		return "unknown"
	}
}

type Armor Item

func (w Armor) SwapType() string {
	log.Debug("swap type: ", w.Bucket)
	switch w.Bucket {
	case "Helmet":
		return "helmet"
	case "Gauntlets":
		return "gauntlets"
	case "Chest Armor":
		return "chest"
	case "Leg Armor":
		return "legs"
	case "Class Armor":
		return "class"
	default:
		return "unknown"
	}
}

type Emblem struct {
	Name            string `yaml:"name" json:"name"`
	InventoryItemId int    `yaml:"inventoryItemId" json:"inventory_item_id"`
	IconUri         string `yaml:"iconUri" json:"icon_uri"`
	BannerUri       string `yaml:"bannerUri" json:"banner_uri"`
	IconOverlayUri  string `yaml:"iconOverlayUri" json:"icon_overlay_uri"`
	Background      Color  `yaml:"background" json:"background"`
}

type Color struct {
	Red   int `yaml:"red" json:"red"`
	Green int `yaml:"green" json:"green"`
	Blue  int `yaml:"blue" json:"blue"`
	Alpha int `yaml:"alpha" json:"alpha"`
}

type Item struct {
	MembershipType  int    `yaml:"membershipType" json:"membership_type"`
	CharacterId     string `yaml:"characterId" json:"character_id"`
	InstanceId      string `yaml:"instanceId" json:"instance_id"`
	InventoryItemId int    `yaml:"inventoryItemId" json:"inventory_item_id"`

	Tier        string `yaml:"tier" json:"tier"`
	Bucket      string `yaml:"bucket" json:"bucket"`
	Name        string `yaml:"name" json:"name"`
	Type        string `yaml:"type" json:"type"`
	IconEnabled bool   `yaml:"iconEnabled" json:"icon_enabled"`
	IconUri     string `yaml:"iconUri" json:"icon_uri"`
	Slot        Slot   `yaml:"slot" json:"slot"`
	Equippable  bool   `yaml:"equippable" json:"equippable"`
}

type Slot struct {
	Name string `yaml:"name" json:"name"`
}
