-- Initialize DB schema
-- Only needs to be run once per server

CREATE TABLE Player (
    username        varchar(25) PRIMARY KEY,
    password        text NOT NULL,
    name            varchar(50) NOT NULL
);

CREATE TYPE e_class AS ENUM (
    'Barbarian',
    'Bard',
    'Cleric',
    'Druid',
    'Fighter',
    'Monk',
    'Paladin',
    'Ranger',
    'Rogue',
    'Sorcerer',
    'Warlock',
    'Wizard',
    'None'
);

CREATE TYPE e_alignment AS ENUM (
    'Lawful Good',
    'Neutral Good',
    'Chaotic Good',
    'Lawful Neutral',
    'True Neutral',
    'Chaotic Neutral',
    'Lawful Evil',
    'Neutral Evil',
    'Chaotic Evil'
);

CREATE TYPE e_race AS ENUM (
    'Dragonborn',
    'Dwarf',
    'Elf',
    'Gnome',
    'Half-Elf',
    'Halfling',
    'Half-Orc',
    'Human',
    'Tiefling'
);

CREATE TYPE e_sex AS ENUM (
    'Male',
    'Female',
    'Other'
);

CREATE TABLE Character (
    id                  serial PRIMARY KEY,
    name                text NOT NULL,
    weight              int NOT NULL CHECK (weight > 0 AND weight <= 1000), -- Kilograms
    height              int NOT NULL CHECK (height > 0 AND height <= 1000), -- Centimetres
    alignment           e_alignment NOT NULL,
    sex                 e_sex NOT NULL,
    background          text,
    race                e_race NOT NULL,
    speed               int NOT NULL CHECK (speed > 0 AND speed <= 1280),
    strength            int NOT NULL CHECK (strength > 0 AND strength <= 30),
    dexterity           int NOT NULL CHECK (dexterity > 0 AND dexterity <= 30),
    intelligence        int NOT NULL CHECK (intelligence > 0 AND intelligence <= 30),
    wisdom              int NOT NULL CHECK (wisdom > 0 AND wisdom <= 30),
    charisma            int NOT NULL CHECK (charisma > 0 AND charisma <= 30),
    constitution        int NOT NULL CHECK (constitution > 0 AND constitution <= 30),
    hp_max              int NOT NULL CHECK (hp_max > 0 AND hp_max <= 440),
    ability_points      int NOT NULL CHECK (ability_points >= 0 AND ability_points <= 180), -- Unused ability points
    xp_points           int NOT NULL CHECK (xp_points >= 0 AND xp_points <= 355000),
    class               e_class NOT NULL,
    player_username     text NOT NULL,
    UNIQUE (name, player_username), -- No player should have multiple characters of the same name
    FOREIGN KEY (player_username) REFERENCES Player(username)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TYPE e_item_type AS ENUM (
    'Armor',
    'Potion',
    'Ring',
    'Rod',
    'Scroll',
    'Staff',
    'Wand',
    'Weapon',
    'Wondrous Item'
);

CREATE TYPE e_item_rarity AS ENUM (
    'Common',
    'Uncommon',
    'Rare',
    'Very Rare',
    'Legendary',
    'Artifact'
);

CREATE TABLE Items (
    character_id        int,
    item_name           text,
    type                e_item_type NOT NULL,
    rarity              e_item_rarity NOT NULL,
    weight              int CHECK (weight >= 0) NOT NULL, -- In pounds
    gold_value          int CHECK (gold_value >= 0) NOT NULL,
    quantity            int CHECK (quantity >= 0) NOT NULL,
    description         text,
    PRIMARY KEY (character_id, item_name),
    FOREIGN KEY (character_id) REFERENCES Character(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TYPE e_school AS ENUM (
    'Abjuration',
    'Conjuration',
    'Divination',
    'Enchantment',
    'Evocation',
    'Illusion',
    'Necromancy',
    'Transmutation'
);

CREATE TABLE Spells (
    character_id        int,
    spell_name          text,
    level               int NOT NULL CHECK (level >= 0 AND level <= 9),
    school              e_school NOT NULL,
    concentration       bool NOT NULL,
    description         text,
    casting_time        int NOT NULL CHECK (casting_time >= 0), -- # Actions
    range               int NOT NULL CHECK (range >= 0), -- in feet
    duration            int NOT NULL CHECK (duration >= 0),
    PRIMARY KEY (character_id, spell_name),
    FOREIGN KEY (character_id) REFERENCES Character(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE Campaign (
    id                  serial PRIMARY KEY,
    name                text NOT NULL,
    current_location    text NOT NULL,
    state               text,
    dungeon_master      varchar(25),
    FOREIGN KEY (dungeon_master) REFERENCES Player(username)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);

CREATE TABLE CampaignMilestones (
    campaign_id         int,
    milestone           text,
    PRIMARY KEY (campaign_id, milestone),
    FOREIGN KEY (campaign_id) REFERENCES Campaign(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE BelongsTo (
    character_id        int,
    campaign_id         int,
    PRIMARY KEY (character_id, campaign_id),
    FOREIGN KEY (character_id) REFERENCES Character(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (campaign_id) REFERENCES Campaign(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);


-- Classes listed below

CREATE TYPE e_barbarian_primal_path AS ENUM (
    'Ancestral Guardian',
    'Battlerager',
    'Beast',
    'Berserker',
    'Storm Herald',
    'Totem Warrior',
    'Wild Magic',
    'Zealot'
);

CREATE TABLE Barbarians (
    character_id        int,
    primal_path         e_barbarian_primal_path,
    PRIMARY KEY (character_id),
    FOREIGN KEY (character_id) REFERENCES Character(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TYPE e_bard_college AS ENUM (
    'Creation',
    'Eloquence',
    'Glamour',
    'Lore',
    'Swords',
    'Valor',
    'Whispers'
);

CREATE TABLE Bards (
    character_id        int,
    college             e_bard_college,
    PRIMARY KEY (character_id),
    FOREIGN KEY (character_id) REFERENCES Character(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TYPE e_cleric_domain AS ENUM (
    'Arcana',
    'Death',
    'Forge',
    'Grave',
    'Knowledge',
    'Life',
    'Light',
    'Nature',
    'Order',
    'Peace',
    'Tempest',
    'Trickery',
    'Twilight',
    'War'
);

CREATE TABLE Clerics (
    character_id        int,
    divine_domain       e_cleric_domain,
    PRIMARY KEY (character_id),
    FOREIGN KEY (character_id) REFERENCES Character(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TYPE e_druid_circle AS ENUM (
    'Dreams',
    'The Land',
    'The Moon',
    'The Shepherd',
    'The Spores',
    'The Stars',
    'Wildfire'
);

CREATE TABLE Druids (
    character_id        int,
    druid_circle        e_druid_circle,
    PRIMARY KEY (character_id),
    FOREIGN KEY (character_id) REFERENCES Character(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TYPE e_fighter_archetype AS ENUM (
    'Arcane Archer',
    'Banneret',
    'Battle Master',
    'Cavalier',
    'Champion',
    'Echo Knight',
    'Eldritch Knight',
    'Psi Warrior',
    'Rune Knight',
    'Samurai'
);

CREATE TYPE e_fighting_style AS ENUM (
    'Archery',
    'Blind Fighting',
    'Defense',
    'Druidic Warrior',
    'Dueling',
    'Great Weapon Fighting',
    'Interception',
    'Protection',
    'Superior Technique',
    'Thrown Weapon Fighting',
    'Two-Weapon Fighting',
    'Unarmed Fighting'
);

CREATE TABLE Fighters (
    character_id        int,
    archetype           e_fighter_archetype,
    fighting_style      e_fighting_style,
    PRIMARY KEY (character_id),
    FOREIGN KEY (character_id) REFERENCES Character(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TYPE e_monk_tradition AS ENUM (
    'Astral Self',
    'Drunken Master',
    'Four Elements',
    'Kensei',
    'Long Death',
    'Mercy',
    'Open Hand',
    'Shadow',
    'Sun Soul'
);

CREATE TABLE Monks (
    character_id        int,
    monastic_tradition  e_monk_tradition,
    PRIMARY KEY (character_id),
    FOREIGN KEY (character_id) REFERENCES Character(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TYPE e_paladin_oath AS ENUM (
    'The Ancients',
    'Conquest',
    'The Crown',
    'Devotion',
    'Glory',
    'Redemption',
    'Vengeance',
    'The Watchers',
    'Oathbreaker'
);

CREATE TABLE Paladins (
    character_id        int,
    oath                e_paladin_oath,
    fighting_style      e_fighting_style,
    PRIMARY KEY (character_id),
    FOREIGN KEY (character_id) REFERENCES Character(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TYPE e_ranger_conclave AS ENUM (
    'Beast Master',
    'Fey Wanderer',
    'Gloom Stalker',
    'Horizon Walker',
    'Hunter',
    'Monster Slayer',
    'Swarmkeeper'
);

CREATE TABLE Rangers (
    character_id        int,
    conclave            e_ranger_conclave,
    favoured_enemy      text,
    PRIMARY KEY (character_id),
    FOREIGN KEY (character_id) REFERENCES Character(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TYPE e_rogue_archetype AS ENUM (
    'Arcane Trickster',
    'Assassin',
    'Inquisitive',
    'Mastermind',
    'Phantom',
    'Scout',
    'Soulknife',
    'Swashbuckler',
    'Thief'
);

CREATE TABLE Rogues (
    character_id        int,
    archetype           e_rogue_archetype,
    PRIMARY KEY (character_id),
    FOREIGN KEY (character_id) REFERENCES Character(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TYPE e_sorcerer_origin AS ENUM (
    'Aberrant Mind',
    'Clockwork Soul',
    'Draconic Bloodline',
    'Divine Soul',
    'Shadow',
    'Storm',
    'Wild Magic'
);

CREATE TABLE Sorcerers (
    character_id        int,
    sorcerous_origin    e_sorcerer_origin,
    PRIMARY KEY (character_id),
    FOREIGN KEY (character_id) REFERENCES Character(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TYPE e_warlock_patron AS ENUM (
    'Archfey',
    'Celestial',
    'Fathomless',
    'Fiend',
    'Genie',
    'Great Old One',
    'Hexblade',
    'Undying'
);

CREATE TABLE Warlocks (
    character_id        int,
    patron              e_warlock_patron,
    PRIMARY KEY (character_id),
    FOREIGN KEY (character_id) REFERENCES Character(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TYPE e_wizard_tradition AS ENUM (
    'Abjuration',
    'Bladesinging',
    'Chronurgy',
    'Conjuration',
    'Divination',
    'Enchantment',
    'Evocation',
    'Graviturgy',
    'Illusion',
    'Necromancy',
    'Order of Scribes',
    'Transmutation',
    'War Magic'
);

CREATE TABLE Wizards (
    character_id        int,
    arcane_tradition    e_wizard_tradition,
    PRIMARY KEY (character_id),
    FOREIGN KEY (character_id) REFERENCES Character(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE Stats (
    num_player_account int DEFAULT 0,
    num_character_created int DEFAULT 0,
    num_campaign_created int DEFAULT 0,
    num_spells_created int DEFAULT 0,
    num_items_created int DEFAULT 0
);

INSERT INTO Stats VALUES (0, 0, 0, 0, 0);

-- Update Stats.num_player_account
CREATE FUNCTION increment_player_account() RETURNS trigger AS $_$
BEGIN
UPDATE Stats SET num_player_account = num_player_account + 1;
RETURN NEW;
END $_$ LANGUAGE 'plpgsql';

CREATE TRIGGER player_registration
AFTER INSERT ON Player
FOR EACH ROW
EXECUTE PROCEDURE increment_player_account();

-- Update Stats.num_character_created
CREATE FUNCTION increment_character_count() RETURNS trigger AS $_$
BEGIN
UPDATE Stats SET num_character_created = num_character_created + 1;
RETURN NEW;
END $_$ LANGUAGE 'plpgsql';

CREATE TRIGGER character_creation
AFTER INSERT ON Character
FOR EACH ROW
EXECUTE PROCEDURE increment_character_count();

-- Update Stats.num_campaign_created
CREATE FUNCTION increment_campaign_count() RETURNS trigger AS $_$
BEGIN
UPDATE Stats SET num_campaign_created = num_campaign_created + 1;
RETURN NEW;
END $_$ LANGUAGE 'plpgsql';

CREATE TRIGGER campaign_creation
AFTER INSERT ON Campaign
FOR EACH ROW
EXECUTE PROCEDURE increment_campaign_count();

-- Update Stats.num_spells_created
CREATE FUNCTION increment_spell_count() RETURNS trigger AS $_$
BEGIN
UPDATE Stats SET num_spells_created = num_spells_created + 1;
RETURN NEW;
END $_$ LANGUAGE 'plpgsql';

CREATE TRIGGER spell_creation
AFTER INSERT ON Spells
FOR EACH ROW
EXECUTE PROCEDURE increment_spell_count();

-- Update Stats.num_items_created
CREATE FUNCTION increment_item_count() RETURNS trigger AS $_$
BEGIN
UPDATE Stats SET num_items_created = num_items_created + 1;
RETURN NEW;
END $_$ LANGUAGE 'plpgsql';

CREATE TRIGGER item_creation
AFTER INSERT ON Items
FOR EACH ROW
EXECUTE PROCEDURE increment_item_count();
