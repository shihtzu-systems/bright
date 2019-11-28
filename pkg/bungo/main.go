package bungo

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type Gamer struct {
	Name         string
	MembershipId string
	Guardians    []Guardian
}

type Guardian struct {
	Id             string `yaml:"id" json:"id"`
	MembershipType int    `yaml:"membershipType" json:"membership_type"`

	Class string `yaml:"class" json:"class"`
	Level int    `yaml:"level" json:"level"`
	Light int    `yaml:"light" json:"light"`

	Emblem Emblem `yaml:"emblem" json:"emblem"`

	Bag      Bag    `yaml:"bag" json:"bag"`
	Equipped Outfit `yaml:"equipped" json:"equipped"`
}

func (g Guardian) String() string {
	return fmt.Sprintf("id=%s\nemblem=%s\nclass=%s\nkinetic=%s\nenergy=%s\npower=%s\nclass-item=%s\narmors=%d\nweapons=%d",
		g.Id, g.Emblem.Name, g.Class, g.Equipped.KineticWeapon.Name, g.Equipped.EnergyWeapon.Name, g.Equipped.PowerWeapon.Name, g.Equipped.Class.Name, len(g.Bag.Armors), len(g.Bag.Weapons))
}

func (g Guardian) Differences(other Guardian) (out []string) {
	if g.Equipped.KineticWeapon.InstanceId != other.Equipped.KineticWeapon.InstanceId {
		out = append(out, g.Equipped.KineticWeapon.InstanceId)
	}
	if g.Equipped.EnergyWeapon.InstanceId != other.Equipped.EnergyWeapon.InstanceId {
		out = append(out, g.Equipped.EnergyWeapon.InstanceId)
	}
	if g.Equipped.PowerWeapon.InstanceId != other.Equipped.PowerWeapon.InstanceId {
		out = append(out, g.Equipped.PowerWeapon.InstanceId)
	}

	if g.Equipped.Helmet.InstanceId != other.Equipped.Helmet.InstanceId {
		out = append(out, g.Equipped.Helmet.InstanceId)
	}
	if g.Equipped.Gauntlets.InstanceId != other.Equipped.Gauntlets.InstanceId {
		out = append(out, g.Equipped.Gauntlets.InstanceId)
	}
	if g.Equipped.Chest.InstanceId != other.Equipped.Chest.InstanceId {
		out = append(out, g.Equipped.Chest.InstanceId)
	}
	if g.Equipped.Leg.InstanceId != other.Equipped.Leg.InstanceId {
		out = append(out, g.Equipped.Leg.InstanceId)
	}
	if g.Equipped.Class.InstanceId != other.Equipped.Class.InstanceId {
		out = append(out, g.Equipped.Class.InstanceId)
	}
	return out
}

func (g *Guardian) SwapWeapon(swapOut Weapon, swapIn Weapon) {
	switch swapOut.Slot.Name {
	case g.Equipped.KineticWeapon.Slot.Name:
		g.Bag.StowWeapon(swapOut)
		g.Bag.TakeWeapon(swapIn)
		log.Debug("KINETECIDJfoidsjaoiajsdoifja ", len(g.Bag.Weapons))
		g.Equipped.KineticWeapon = swapIn
	case g.Equipped.EnergyWeapon.Slot.Name:
		g.Bag.StowWeapon(swapOut)
		g.Bag.TakeWeapon(swapIn)
		g.Equipped.EnergyWeapon = swapIn
	case g.Equipped.PowerWeapon.Slot.Name:
		g.Bag.StowWeapon(swapOut)
		g.Bag.TakeWeapon(swapIn)
		g.Equipped.PowerWeapon = swapIn
	}
}

func (g *Guardian) SwapArmor(swapOut Armor, swapIn Armor) {
	switch swapOut.Slot.Name {
	case g.Equipped.Helmet.Slot.Name:
		g.Bag.StowArmor(swapOut)
		g.Bag.TakeArmor(swapIn)
		g.Equipped.Helmet = swapIn
	case g.Equipped.Gauntlets.Slot.Name:
		g.Bag.StowArmor(swapOut)
		g.Bag.TakeArmor(swapIn)
		g.Equipped.Gauntlets = swapIn
	case g.Equipped.Chest.Slot.Name:
		g.Bag.StowArmor(swapOut)
		g.Bag.TakeArmor(swapIn)
		g.Equipped.Chest = swapIn
	case g.Equipped.Leg.Slot.Name:
		g.Bag.StowArmor(swapOut)
		g.Bag.TakeArmor(swapIn)
		g.Equipped.Leg = swapIn
	case g.Equipped.Class.Slot.Name:
		g.Bag.StowArmor(swapOut)
		g.Bag.TakeArmor(swapIn)
		g.Equipped.Class = swapIn
	}
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

func (o Outfit) FindWeapon(instanceId string) (out Weapon, found bool) {
	log.Debug("Outfit->FindWeapon: ", instanceId)
	switch instanceId {
	case o.KineticWeapon.InstanceId:
		out = o.KineticWeapon
		found = true
	case o.EnergyWeapon.InstanceId:
		out = o.EnergyWeapon
		found = true
	case o.PowerWeapon.InstanceId:
		out = o.PowerWeapon
		found = true
	}
	return out, found
}

func (o Outfit) FindArmor(instanceId string) (out Armor, found bool) {
	log.Debug("Outfit->FindArmor: ", instanceId)
	switch instanceId {
	case o.Helmet.InstanceId:
		out = o.Helmet
		found = true
	case o.Gauntlets.InstanceId:
		out = o.Gauntlets
		found = true
	case o.Chest.InstanceId:
		out = o.Chest
		found = true
	case o.Leg.InstanceId:
		out = o.Leg
		found = true
	case o.Class.InstanceId:
		out = o.Class
		found = true
	}
	return out, found
}

type Bag struct {
	Weapons []Weapon `yaml:"weapons" json:"weapons"`
	Armors  []Armor  `yaml:"armors" json:"armor"`
}

func (b Bag) FindWeapon(instanceId string) (out Weapon, found bool) {
	log.Debug("Bag->FindWeapon: ", instanceId)
	for _, weapon := range b.Weapons {
		if weapon.InstanceId == instanceId {
			out = weapon
			found = true
			break
		}
	}
	return out, found
}

func (b *Bag) TakeWeapon(instance Weapon) {
	log.Debug("Bag->TakeWeapon: ", instance)
	var weapons []Weapon
	for _, w := range b.Weapons {
		if w.InstanceId == instance.InstanceId {
			continue
		}
		weapons = append(weapons, w)
	}
	b.Weapons = weapons
}

func (b *Bag) StowWeapon(instance Weapon) {
	log.Debug("Bag->StowWeapon: ", instance.InstanceId)
	var found bool
	for _, w := range b.Weapons {
		if w.InstanceId == instance.InstanceId {
			found = true
			break
		}
	}
	if !found {
		b.Weapons = append(b.Weapons, instance)
	}
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

func (b Bag) FindArmor(instanceId string) (out Armor, found bool) {
	log.Debugf("Bag->FindArmor: %s [%d]", instanceId, len(b.Armors))
	for _, armor := range b.Armors {
		if armor.InstanceId == instanceId {
			out = armor
			found = true
			break
		}
	}
	return out, found
}
func (b *Bag) TakeArmor(instance Armor) {
	log.Debug("Bag->TakeArmor: ", instance.InstanceId)
	var armors []Armor
	for _, a := range b.Armors {
		if a.InstanceId == instance.InstanceId {
			continue
		}
		armors = append(armors, a)
	}
	b.Armors = armors
}

func (b *Bag) StowArmor(instance Armor) {
	log.Debug("Bag->StowArmor: ", instance.InstanceId)
	var found bool
	for _, a := range b.Armors {
		if a.InstanceId == instance.InstanceId {
			found = true
			break
		}
	}
	if !found {
		b.Armors = append(b.Armors, instance)
	}
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
