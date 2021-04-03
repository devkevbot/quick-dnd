package models

import (
	"encoding/json"
	"errors"
)

var (
	ErrNoRecord           = errors.New("models: no record was found")
	ErrUpdateSingleRecord = errors.New("models: number of records updated was not one")
	ErrDeleteSingleRecord = errors.New("models: number of records deleted was not one")
	ErrDuplicateUsername  = errors.New("models: duplicate player username")
	ErrDuplicateCharacter = errors.New("models: player cannot have two characters with the same name")
	ErrDuplicateSpell     = errors.New("models: spell names must be unique for a given character")
	ErrDuplicateItem      = errors.New("models: item names must be unique for a given character")
)

// JSON unmarshal errors for custom character data types.
var (
	ErrInvalidClassType     = errors.New("models: invalid character class type")
	ErrInvalidAlignmentType = errors.New("models: invalid character alignment type")
	ErrInvalidRaceType      = errors.New("models: invalid character race type")
	ErrInvalidSexType       = errors.New("models: invalid character sex type")
)

// JSON unmarshal errors for custom item data types.
var (
	ErrInvalidItemType   = errors.New("models: invalid item type")
	ErrInvalidRarityType = errors.New("models: invalid item rarity type")
)

// JSON unmarshal errors for custom spell data types.
var (
	ErrInvalidMagicSchoolType = errors.New("models: invalid spell magic school type")
)

// JSON unmarshal errors for class-specific data types.
var (
	ErrInvalidBarbarianPrimalPathType = errors.New("models: invalid barbarian primal path type")
	ErrInvalidBardCollegeType         = errors.New("models: invalid bard college type")
	ErrInvalidClericDomainType        = errors.New("models: invalid cleric domain type")
	ErrInvalidDruidCircleType         = errors.New("models: invalid druid circle type")
	ErrInvalidFighterArchetypeType    = errors.New("models: invalid fighter archetype type")
	ErrInvalidFightingStyleType       = errors.New("models: invalid fighting style type")
	ErrInvalidMonkTraditionType       = errors.New("models: invalid monk tradition type")
	ErrInvalidPaladinOathType         = errors.New("models: invalid paladin oath type")
	ErrInvalidRangerConclaveType      = errors.New("models: invalid ranger conclave type")
	ErrInvalidRogueArchetypeType      = errors.New("models: invalid rogue archetype type")
	ErrInvalidSorcererOriginType      = errors.New("models: invalid sorcerer origin type")
	ErrInvalidWarlockPatronType       = errors.New("models: invalid warlock patron type")
	ErrInvalidWizardTraditionType     = errors.New("models: invalid wizard tradition type")
)

// Player is the code representation of the "Player" relation in the
// database schema.
type Player struct {
	Username string `json:"username" db:"username"`
	Password string `json:"-" db:"password"`
	Name     string `json:"name" db:"name"`
}

type ClassType string

const (
	Barbarian ClassType = "Barbarian"
	Bard                = "Bard"
	Cleric              = "Cleric"
	Druid               = "Druid"
	Fighter             = "Fighter"
	Monk                = "Monk"
	Paladin             = "Paladin"
	Ranger              = "Ranger"
	Rogue               = "Rogue"
	Sorcerer            = "Sorcerer"
	Warlock             = "Warlock"
	Wizard              = "Wizard"
	None                = "None"
)

func (t *ClassType) UnmarshalJSON(b []byte) error {
	type T ClassType
	var r *T = (*T)(t)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	switch *t {
	case
		Barbarian,
		Bard,
		Cleric,
		Druid,
		Fighter,
		Monk,
		Paladin,
		Ranger,
		Rogue,
		Sorcerer,
		Warlock,
		Wizard,
		None:
		return nil
	}
	return ErrInvalidClassType
}

type AlignmentType string

const (
	LawfulGood     AlignmentType = "Lawful Good"
	NeutralGood                  = "Neutral Good"
	ChaoticGood                  = "Chaotic Good"
	LawfulNeutral                = "Lawful Neutral"
	TrueNeutral                  = "True Neutral"
	ChaoticNeutral               = "Chaotic Neutral"
	LawfulEvil                   = "Lawful Evil"
	NeutralEvil                  = "Neutral Evil"
	ChaoticEvil                  = "Chaotic Evil"
)

func (t *AlignmentType) UnmarshalJSON(b []byte) error {
	type T AlignmentType
	var r *T = (*T)(t)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	switch *t {
	case
		LawfulGood,
		NeutralGood,
		ChaoticGood,
		LawfulNeutral,
		TrueNeutral,
		ChaoticNeutral,
		LawfulEvil,
		NeutralEvil,
		ChaoticEvil:
		return nil
	}
	return ErrInvalidAlignmentType
}

type RaceType string

const (
	Dragonborn RaceType = "Dragonborn"
	Dwarf               = "Dwarf"
	Elf                 = "Elf"
	Gnome               = "Gnome"
	HalfElf             = "Half-Elf"
	Halfling            = "Halfling"
	HalfOrc             = "Half-Orc"
	Human               = "Human"
	Tiefling            = "Tiefling"
)

func (t *RaceType) UnmarshalJSON(b []byte) error {
	type T RaceType
	var r *T = (*T)(t)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	switch *t {
	case
		Dragonborn,
		Dwarf,
		Elf,
		Gnome,
		HalfElf,
		Halfling,
		HalfOrc,
		Human,
		Tiefling:
		return nil
	}
	return ErrInvalidRaceType
}

type SexType string

const (
	Male   SexType = "Male"
	Female         = "Female"
	Other          = "Other"
)

func (t *SexType) UnmarshalJSON(b []byte) error {
	type T SexType
	var r *T = (*T)(t)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	switch *t {
	case
		Male,
		Female,
		Other:
		return nil
	}
	return ErrInvalidSexType
}

// Character is the code representation of the "Character" relation in
// the database schema.
type Character struct {
	ID             int           `json:"id" db:"id"`
	Name           string        `json:"name" db:"name"`
	Weight         int           `json:"weight" db:"weight"`
	Height         int           `json:"height" db:"height"`
	Alignment      AlignmentType `json:"alignment" db:"alignment"`
	Sex            SexType       `json:"sex" db:"sex"`
	Background     string        `json:"background" db:"background"`
	Race           RaceType      `json:"race" db:"race"`
	Speed          int           `json:"speed" db:"speed"`
	Strength       int           `json:"strength" db:"strength"`
	Dexterity      int           `json:"dexterity" db:"dexterity"`
	Intelligence   int           `json:"intelligence" db:"intelligence"`
	Wisdom         int           `json:"wisdom" db:"wisdom"`
	Charisma       int           `json:"charisma" db:"charisma"`
	Constitution   int           `json:"constitution" db:"constitution"`
	HPMax          int           `json:"hp_max" db:"hp_max"`
	AbilityPoints  int           `json:"ability_points" db:"ability_points"`
	XPPoints       int           `json:"xp_points" db:"xp_points"`
	Class          ClassType     `json:"class" db:"class"`
	PlayerUsername string        `json:"player_username" db:"player_username"`
}

// CharacterAbility is the code representation of the
// "CharacterAbilities" relation in the database schema.
type CharacterAbility struct {
	CharacterID int    `json:"character_id" db:"character_id"`
	Ability     string `json:"ability" db:"ability"`
}

type ItemType string

const (
	Armor        ItemType = "Armor"
	Potion                = "Potion"
	Ring                  = "Ring"
	Rod                   = "Rod"
	Scroll                = "Scroll"
	Staff                 = "Staff"
	Wand                  = "Wand"
	Weapon                = "Weapon"
	WondrousItem          = "Wondrous Item"
)

func (t *ItemType) UnmarshalJSON(b []byte) error {
	type T ItemType
	var r *T = (*T)(t)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	switch *t {
	case
		Armor,
		Potion,
		Ring,
		Rod,
		Scroll,
		Staff,
		Wand,
		Weapon,
		WondrousItem:
		return nil
	}
	return ErrInvalidItemType
}

type RarityType string

const (
	Common    RarityType = "Common"
	Uncommon             = "Uncommon"
	Rare                 = "Rare"
	VeryRare             = "Very Rare"
	Legendary            = "Legendary"
	Artifact             = "Artifact"
)

func (t *RarityType) UnmarshalJSON(b []byte) error {
	type T RarityType
	var r *T = (*T)(t)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	switch *t {
	case
		Common,
		Uncommon,
		Rare,
		VeryRare,
		Legendary,
		Artifact:
		return nil
	}
	return ErrInvalidRarityType
}

// Item is the code representation of the "Items" relation in the
// database schema.
type Item struct {
	CharacterID int        `json:"character_id" db:"character_id"`
	ItemName    string     `json:"item_name" db:"item_name"`
	Type        ItemType   `json:"type" db:"type"`
	Rarity      RarityType `json:"rarity" db:"rarity"`
	Weight      int        `json:"weight" db:"weight"` // pounds
	GoldValue   int        `json:"gold_value" db:"gold_value"`
	Quantity    int        `json:"quantity" db:"quantity"`
	Description string     `json:"description" db:"description"`
}

type MagicSchoolType string

const (
	Abjuration   MagicSchoolType = "Abjuration"
	Conjuration                  = "Conjuration"
	Divination                   = "Divination"
	Enchantment                  = "Enchantment"
	Evocation                    = "Evocation"
	Illusion                     = "Illusion"
	Necromancy                   = "Necromancy"
	Transmuation                 = "Transmutation"
)

func (t *MagicSchoolType) UnmarshalJSON(b []byte) error {
	type T MagicSchoolType
	var r *T = (*T)(t)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	switch *t {
	case
		Abjuration,
		Conjuration,
		Divination,
		Enchantment,
		Evocation,
		Illusion,
		Necromancy,
		Transmuation:
		return nil
	}
	return ErrInvalidMagicSchoolType
}

// Spell is the code representation of the "Spells" relation in the
// database schema.
type Spell struct {
	CharacterID   int             `json:"character_id" db:"character_id"`
	SpellName     string          `json:"spell_name" db:"spell_name"`
	Level         int             `json:"level" db:"level"`
	School        MagicSchoolType `json:"school" db:"school"`
	Concentration bool            `json:"concentration" db:"concentration"`
	Description   string          `json:"description" db:"description"`
	CastingTime   int             `json:"casting_time" db:"casting_time"`
	Range         int             `json:"range" db:"range"` // feet
	Duration      int             `json:"duration" db:"duration"`
}

// SpellComponent is the code representation of the "SpellComponents"
// relation in the database schema.
type SpellComponent struct {
	CharacterID int    `json:"character_id" db:"character_id"`
	SpellName   string `json:"spell_name" db:"spell_name"`
	Component   string `json:"component" db:"component"`
}

// Campaign is the code representation of the "Campaign" relation in the
// database schema.
type Campaign struct {
	ID              int    `json:"id" db:"id"`
	Name            string `json:"name" db:"name"`
	CurrentLocation string `json:"current_location" db:"current_location"`
	State           string `json:"state" db:"state"`
	DungeonMaster   string `json:"dungeon_master" db:"dungeon_master"`
}

// CampaignMilestone is the code representation of the
// "CampaignMilestones" relation in the database schema.
type CampaignMilestone struct {
	CampaignID int    `json:"campaign_id" db:"campaign_id"`
	Milestone  string `json:"milestone" db:"milestone"`
}

// BelongsTo is the code representation of the "BelongsTo" relation in
// the database schema.
type BelongsTo struct {
	CharacterID int `json:"character_id" db:"character_id"`
	CampaignID  int `json:"campaign_id" db:"campaign_id"`
}

type BarbarianPrimalPathType string

const (
	AncestralGuardian BarbarianPrimalPathType = "Ancestral Guardian"
	Battlerager                               = "Battlerager"
	Beast                                     = "Beast"
	Berserker                                 = "Berserker"
	StormHerald                               = "Storm Herald"
	TotemWarrior                              = "Totem Warrior"
	WildMagic                                 = "Wild Magic"
	Zealot                                    = "Zealot"
)

func (t *BarbarianPrimalPathType) UnmarshalJSON(b []byte) error {
	type T BarbarianPrimalPathType
	var r *T = (*T)(t)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	switch *t {
	case
		AncestralGuardian,
		Battlerager,
		Beast,
		Berserker,
		StormHerald,
		TotemWarrior,
		WildMagic,
		Zealot:
		return nil
	}
	return ErrInvalidBarbarianPrimalPathType
}

type BarbarianClass struct {
	CharacterID int                     `json:"character_id" db:"character_id"`
	PrimalPath  BarbarianPrimalPathType `json:"primal_path" db:"primal_path"`
}

type BardCollegeType string

const (
	Creation  BardCollegeType = "Creation"
	Eloquence                 = "Eloquence"
	Glamour                   = "Glamour"
	Lore                      = "Lore"
	Swords                    = "Swords"
	Valor                     = "Valor"
	Whispers                  = "Whispers"
)

func (t *BardCollegeType) UnmarshalJSON(b []byte) error {
	type T BardCollegeType
	var r *T = (*T)(t)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	switch *t {
	case
		Creation,
		Eloquence,
		Glamour,
		Lore,
		Swords,
		Valor,
		Whispers:
		return nil
	}
	return ErrInvalidBardCollegeType
}

type BardClass struct {
	CharacterID int             `json:"character_id" db:"character_id"`
	College     BardCollegeType `json:"college" db:"college"`
}

type ClericDomainType string

const (
	Arcana    ClericDomainType = "Arcana"
	Death                      = "Death"
	Forge                      = "Forge"
	Grave                      = "Grave"
	Knowledge                  = "Knowledge"
	Life                       = "Life"
	Light                      = "Light"
	Nature                     = "Nature"
	Order                      = "Order"
	Peace                      = "Peace"
	Tempest                    = "Tempest"
	Trickery                   = "Trickery"
	Twilight                   = "Twilight"
	War                        = "War"
)

func (t *ClericDomainType) UnmarshalJSON(b []byte) error {
	type T ClericDomainType
	var r *T = (*T)(t)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	switch *t {
	case
		Arcana,
		Death,
		Forge,
		Grave,
		Knowledge,
		Life,
		Light,
		Nature,
		Order,
		Peace,
		Tempest,
		Trickery,
		Twilight,
		War:
		return nil
	}
	return ErrInvalidClericDomainType
}

type ClericClass struct {
	CharacterID  int              `json:"character_id" db:"character_id"`
	DivineDomain ClericDomainType `json:"divine_domain" db:"divine_domain"`
}

type DruidCircleType string

const (
	Dreams      DruidCircleType = "Dreams"
	TheLand                     = "The Land"
	TheMoon                     = "The Moon"
	TheShepherd                 = "The Shepherd"
	TheSpores                   = "The Spores"
	TheStars                    = "The Stars"
	Wildfire                    = "Wildfire"
)

func (t *DruidCircleType) UnmarshalJSON(b []byte) error {
	type T DruidCircleType
	var r *T = (*T)(t)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	switch *t {
	case
		Dreams,
		TheLand,
		TheMoon,
		TheShepherd,
		TheSpores,
		TheStars,
		Wildfire:
		return nil
	}
	return ErrInvalidDruidCircleType
}

type DruidClass struct {
	CharacterID int             `json:"character_id" db:"character_id"`
	DruidCircle DruidCircleType `json:"druid_circle" db:"druid_circle"`
}

type FighterArchetypeType string

const (
	ArcaneArcher   FighterArchetypeType = "Arcane Archer"
	Banneret                            = "Banneret"
	BattleMaster                        = "Battle Master"
	Cavalier                            = "Cavalier"
	Champion                            = "Champion"
	EchoKnight                          = "Echo Knight"
	EldritchKnight                      = "Eldritch Knight"
	PsiWarrior                          = "Psi Warrior"
	RuneKnight                          = "Rune Knight"
	Samurai                             = "Samurai"
)

func (t *FighterArchetypeType) UnmarshalJSON(b []byte) error {
	type T FighterArchetypeType
	var r *T = (*T)(t)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	switch *t {
	case
		ArcaneArcher,
		Banneret,
		BattleMaster,
		Cavalier,
		Champion,
		EchoKnight,
		EldritchKnight,
		PsiWarrior,
		RuneKnight,
		Samurai:
		return nil
	}
	return ErrInvalidFighterArchetypeType
}

type FightingStyleType string

const (
	Archery              FightingStyleType = "Archery"
	BlindFighting                          = "Blind Fighting"
	Defense                                = "Defense"
	DruidicWarrior                         = "Druidic Warrior"
	Dueling                                = "Dueling"
	GreatWeaponFighting                    = "Great Weapon Fighting"
	Interception                           = "Interception"
	Protection                             = "Protection"
	SuperiorTechnique                      = "Superior Technique"
	ThrownWeaponFighting                   = "Thrown Weapon Fighting"
	TwoWeaponFighting                      = "Two-Weapon Fighting"
	UnarmedFighting                        = "Unarmed Fighting"
)

func (t *FightingStyleType) UnmarshalJSON(b []byte) error {
	type T FightingStyleType
	var r *T = (*T)(t)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	switch *t {
	case
		Archery,
		BlindFighting,
		Defense,
		DruidicWarrior,
		Dueling,
		GreatWeaponFighting,
		Interception,
		Protection,
		SuperiorTechnique,
		ThrownWeaponFighting,
		TwoWeaponFighting,
		UnarmedFighting:
		return nil
	}
	return ErrInvalidFightingStyleType
}

type FighterClass struct {
	CharacterID   int                  `json:"character_id" db:"character_id"`
	Archetype     FighterArchetypeType `json:"archetype" db:"archetype"`
	FightingStyle FightingStyleType    `json:"fighting_style" db:"fighting_style"`
}

type MonkTraditionType string

const (
	AstralSelf    MonkTraditionType = "Astral Self"
	DrunkenMaster                   = "Drunken Master"
	FourElements                    = "Four Elements"
	Kensei                          = "Kensei"
	LongDeath                       = "Long Death"
	Mercy                           = "Mercy"
	OpenHand                        = "Open Hand"
	Shadow                          = "Shadow"
	SunSoul                         = "Sun Soul"
)

func (t *MonkTraditionType) UnmarshalJSON(b []byte) error {
	type T MonkTraditionType
	var r *T = (*T)(t)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	switch *t {
	case
		AstralSelf,
		DrunkenMaster,
		FourElements,
		Kensei,
		LongDeath,
		Mercy,
		OpenHand,
		Shadow,
		SunSoul:
		return nil
	}
	return ErrInvalidMonkTraditionType
}

type MonkClass struct {
	CharacterID       int               `json:"character_id" db:"character_id"`
	MonasticTradition MonkTraditionType `json:"monastic_tradition" db:"monastic_tradition"`
}

type PaladinOathType string

const (
	TheAncients PaladinOathType = "The Ancients"
	Conquest                    = "Conquest"
	TheCrown                    = "The Crown"
	Devotion                    = "Devotion"
	Glory                       = "Glory"
	Redemption                  = "Redemption"
	Vengeance                   = "Vengeance"
	TheWatchers                 = "The Watchers"
	Oathbreaker                 = "Oathbreaker"
)

func (t *PaladinOathType) UnmarshalJSON(b []byte) error {
	type T PaladinOathType
	var r *T = (*T)(t)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	switch *t {
	case
		TheAncients,
		Conquest,
		TheCrown,
		Devotion,
		Glory,
		Redemption,
		Vengeance,
		TheWatchers,
		Oathbreaker:
		return nil
	}
	return ErrInvalidPaladinOathType
}

type PaladinClass struct {
	CharacterID   int               `json:"character_id" db:"character_id"`
	Oath          PaladinOathType   `json:"oath" db:"oath"`
	FightingStyle FightingStyleType `json:"fighting_style" db:"fighting_style"`
}

type RangerConclaveType string

const (
	BeastMaster   RangerConclaveType = "Beast Master"
	FeyWanderer                      = "Fey Wanderer"
	GloomStalker                     = "Gloom Stalker"
	HorizonWalker                    = "Horizon Walker"
	Hunter                           = "Hunter"
	MonsterSlayer                    = "Monster Slayer"
	Swarmkeeper                      = "Swarmkeeper"
)

func (t *RangerConclaveType) UnmarshalJSON(b []byte) error {
	type T RangerConclaveType
	var r *T = (*T)(t)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	switch *t {
	case
		BeastMaster,
		FeyWanderer,
		GloomStalker,
		HorizonWalker,
		Hunter,
		MonsterSlayer,
		Swarmkeeper:
		return nil
	}
	return ErrInvalidRangerConclaveType
}

type RangerClass struct {
	CharacterID   int                `json:"character_id" db:"character_id"`
	Conclave      RangerConclaveType `json:"conclave" db:"conclave"`
	FavouredEnemy string             `json:"favoured_enemy" db:"favoured_enemy"`
}

type RogueArchetypeType string

const (
	ArcaneTrickster RogueArchetypeType = "Arcane Trickster"
	Assassin                           = "Assassin"
	Inquisitive                        = "Inquisitive"
	Mastermind                         = "Mastermind"
	Phantom                            = "Phantom"
	Scout                              = "Scout"
	Soulknife                          = "Soulknife"
	Swashbuckler                       = "Swashbuckler"
	Thief                              = "Thief"
)

func (t *RogueArchetypeType) UnmarshalJSON(b []byte) error {
	type T RogueArchetypeType
	var r *T = (*T)(t)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	switch *t {
	case
		ArcaneTrickster,
		Assassin,
		Inquisitive,
		Mastermind,
		Phantom,
		Scout,
		Soulknife,
		Swashbuckler,
		Thief:
		return nil
	}
	return ErrInvalidRogueArchetypeType
}

type RogueClass struct {
	CharacterID int                `json:"character_id" db:"character_id"`
	Archetype   RogueArchetypeType `json:"archetype" db:"archetype"`
}

type SorcererOriginType string

const (
	AberrantMind      SorcererOriginType = "Aberrant Mind"
	ClockworkSoul                        = "Clockwork Soul"
	DraconicBloodline                    = "Draconic Bloodline"
	DivineSoul                           = "Divine Soul"
	SorcShadow                           = "Shadow"
	Storm                                = "Storm"
	SorcWildMagic                        = "Wild Magic"
)

func (t *SorcererOriginType) UnmarshalJSON(b []byte) error {
	type T SorcererOriginType
	var r *T = (*T)(t)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	switch *t {
	case
		AberrantMind,
		ClockworkSoul,
		DraconicBloodline,
		DivineSoul,
		SorcShadow,
		Storm,
		SorcWildMagic:
		return nil
	}
	return ErrInvalidSorcererOriginType
}

type SorcererClass struct {
	CharacterID     int                `json:"character_id" db:"character_id"`
	SorcerousOrigin SorcererOriginType `json:"sorcerous_origin" db:"sorcerous_origin"`
}

type WarlockPatronType string

const (
	Archfey     WarlockPatronType = "Archfey"
	Celestial                     = "Celestial"
	Fathomless                    = "Fathomless"
	Fiend                         = "Fiend"
	Genie                         = "Genie"
	GreatOldOne                   = "Great Old One"
	Hexblade                      = "Hexblade"
	Undying                       = "Undying"
)

func (t *WarlockPatronType) UnmarshalJSON(b []byte) error {
	type T WarlockPatronType
	var r *T = (*T)(t)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	switch *t {
	case
		Archfey,
		Celestial,
		Fathomless,
		Fiend,
		Genie,
		GreatOldOne,
		Hexblade,
		Undying:
		return nil
	}
	return ErrInvalidWarlockPatronType
}

type WarlockClass struct {
	CharacterID int               `json:"character_id" db:"character_id"`
	Patron      WarlockPatronType `json:"patron" db:"patron"`
}

type WizardTraditionType string

const (
	WAbjuration    WizardTraditionType = "Abjuration"
	Bladesinging                       = "Bladesinging"
	Chronurgy                          = "Chronurgy"
	WConjuration                       = "Conjuration"
	WDivination                        = "Divination"
	WEnchantment                       = "Enchantment"
	WEvocation                         = "Evocation"
	Graviturgy                         = "Graviturgy"
	WIllusion                          = "Illusion"
	WNecromancy                        = "Necromancy"
	OrderOfScribes                     = "Order of Scribes"
	Transmutation                      = "Transmutation"
	WarMagic                           = "War Magic"
)

func (t *WizardTraditionType) UnmarshalJSON(b []byte) error {
	type T WizardTraditionType
	var r *T = (*T)(t)
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	switch *t {
	case
		WAbjuration,
		Bladesinging,
		Chronurgy,
		WConjuration,
		WDivination,
		WEnchantment,
		WEvocation,
		Graviturgy,
		WIllusion,
		WNecromancy,
		OrderOfScribes,
		Transmutation,
		WarMagic:
		return nil
	}
	return ErrInvalidWizardTraditionType
}

type WizardClass struct {
	CharacterID     int                 `json:"character_id" db:"character_id"`
	ArcaneTradition WizardTraditionType `json:"arcane_tradition" db:"arcane_tradition"`
}

/* Model for global statistics. */
type Stats struct {
	NumPlayersCreated    int `json:"num_player_account" db:"num_player_account"`
	NumCharactersCreated int `json:"num_character_created" db:"num_character_created"`
	NumCampaignsCreated  int `json:"num_campaign_created" db:"num_campaign_created"`
	NumSpellsCreated     int `json:"num_spells_created" db:"num_spells_created"`
	NumItemsCreated      int `json:"num_items_created" db:"num_items_created"`
}
