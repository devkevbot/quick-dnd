-- Initialize DB
-- Only needs to be run once per server

CREATE DATABASE dnd_db;

\connect dnd_db

CREATE TABLE Player (
    username        text PRIMARY KEY,
    password        text NOT NULL,
    email           text UNIQUE NOT NULL,
    name            text
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
    'Wizard'
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

-- Letting a lot of attributes here be nullable as users may not wish to track some stuff
CREATE TABLE Character (
    name                text PRIMARY KEY,
    weight              int CHECK (weight > 0),
    height              int CHECK (height > 0),
    alignment           e_alignment,
    sex                 text,
    background          text,
    race                e_race,
    speed               int CHECK (speed > 0),
    strength            int CHECK (strength > 0),
    dexterity           int CHECK (dexterity > 0),
    intelligence        int CHECK (intelligence > 0),
    wisdom              int CHECK (wisdom > 0),
    charisma            int CHECK (charisma > 0),
    constitution        int CHECK (constitution > 0),
    hp_max              int CHECK (hp_max > 0),
    ability_points      int CHECK (ability_points >= 0),
    xp_points           int CHECK (xp_points >= 0),
    class               e_class,
    player_username     text NOT NULL,
    FOREIGN KEY (player_username) REFERENCES Player(username)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE CharacterAbilities (
    character_name      text,
    ability             text,
    PRIMARY KEY (character_name, ability),
    FOREIGN KEY (character_name) REFERENCES Character(name)
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
    character_name      text,
    item_name           text,
    type                e_item_type,
    rarity              e_item_rarity,
    weight              int CHECK (weight >= 0),
    gold_value          int CHECK (gold_value >= 0),
    description         text,
    PRIMARY KEY (character_name, item_name),
    FOREIGN KEY (character_name) REFERENCES Character(name)
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
    character_name      text,
    spell_name          text,
    level               int,
        CHECK (level >= 0),
        CHECK (level <= 9),
    school              e_school,
    concentration       bool,
    description         text,
    casting_time        int CHECK (casting_time >= 0),
    range               int CHECK (range >= 0),
    duration            int CHECK (duration >= 0),
    PRIMARY KEY (character_name, spell_name),
    FOREIGN KEY (character_name) REFERENCES Character(name)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE SpellComponents (
    character_name      text,
    spell_name          text,
    component           text,
    PRIMARY KEY (character_name, spell_name, component),
    FOREIGN KEY (character_name, spell_name) REFERENCES Spells(character_name, spell_name)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE Campaign (
    name                text PRIMARY KEY,
    current_location    text,
    state               text
);

CREATE TABLE CampaignMilestones (
    campaign_name       text,
    milestone           text,
    PRIMARY KEY (campaign_name, milestone),
    FOREIGN KEY (campaign_name) REFERENCES Campaign(name)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE BelongsTo (
    character_name      text,
    campaign_name       text,
    PRIMARY KEY (character_name, campaign_name),
    FOREIGN KEY (character_name) REFERENCES Character(name)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (campaign_name) REFERENCES Campaign(name)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);


-- Classes listed below

CREATE TYPE e_barbarian_primal_path AS ENUM (
    'Path of the Ancestral Guardian',
    'Path of the Battlerager',
    'Path of the Beast',
    'Path of the Berserker',
    'Path of the Storm Herald',
    'Path of the Totem Warrior',
    'Path of the Wild Magic',
    'Path of the Zealot'
);

CREATE TABLE Barbarians (
    character_name      text,
    primal_path         e_barbarian_primal_path,
    PRIMARY KEY (character_name),
    FOREIGN KEY (character_name) REFERENCES Character(name)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TYPE e_bard_college AS ENUM (
    'College of Creation',
    'College of Eloquence',
    'College of Glamour',
    'College of Lore',
    'College of Swords',
    'College of Valor',
    'College of Whispers'
);

CREATE TABLE Bards (
    character_name      text,
    college             e_bard_college,
    PRIMARY KEY (character_name),
    FOREIGN KEY (character_name) REFERENCES Character(name)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TYPE e_cleric_domain AS ENUM (
    'Arcana Domain',
    'Death Domain',
    'Forge Domain',
    'Grave Domain',
    'Knowledge Domain',
    'Life Domain',
    'Light Domain',
    'Nature Domain',
    'Order Domain',
    'Peace Domain',
    'Tempest Domain',
    'Trickery Domain',
    'Twilight Domain',
    'War Domain'
);

CREATE TABLE Clerics (
    character_name      text,
    divine_domain       e_cleric_domain,
    PRIMARY KEY (character_name),
    FOREIGN KEY (character_name) REFERENCES Character(name)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TYPE e_druid_circle AS ENUM (
    'Circle of Dreams',
    'Circle of the Land',
    'Circle of the Moon',
    'Circle of the Shepherd',
    'Circle of the Spores',
    'Circle of the Stars',
    'Circle of Wildfire'
);

CREATE TABLE Druids (
    character_name      text,
    druid_circle        e_druid_circle,
    PRIMARY KEY (character_name),
    FOREIGN KEY (character_name) REFERENCES Character(name)
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
    'Druidic Warrior'
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
    character_name      text,
    archetype           e_fighter_archetype,
    fighting_style      e_fighting_style,
    PRIMARY KEY (character_name),
    FOREIGN KEY (character_name) REFERENCES Character(name)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TYPE e_monk_tradition AS ENUM (
    'Way of the Astral Self',
    'Way of the Drunken Master',
    'Way of the Four Elements',
    'Way of the Kensei',
    'Way of the Long Death',
    'Way of the Mercy',
    'Way of the Open Hand',
    'Way of the Shadow',
    'Way of the Sun Soul'
);

CREATE TABLE Monks (
    character_name      text,
    monastic_tradition  e_monk_tradition,
    PRIMARY KEY (character_name),
    FOREIGN KEY (character_name) REFERENCES Character(name)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TYPE e_paladin_oath AS ENUM (
    'Oath of the Ancients',
    'Oath of Conquest',
    'Oath of the Crown',
    'Oath of Devotion',
    'Oath of Glory',
    'Oath of Redemption',
    'Oath of Vengeance',
    'Oath of the Watchers',
    'Oathbreaker'
);

CREATE TABLE Paladins (
    character_name      text,
    oath                e_paladin_oath,
    fighting_style      e_fighting_style,
    PRIMARY KEY (character_name),
    FOREIGN KEY (character_name) REFERENCES Character(name)
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
    character_name      text,
    conclave            e_ranger_conclave,
    favoured_enemy      text,
    PRIMARY KEY (character_name),
    FOREIGN KEY (character_name) REFERENCES Character(name)
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
    character_name      text,
    archetype           e_rogue_archetype,
    PRIMARY KEY (character_name),
    FOREIGN KEY (character_name) REFERENCES Character(name)
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
    character_name      text,
    sorcerous_origin    e_sorcerer_origin,
    PRIMARY KEY (character_name),
    FOREIGN KEY (character_name) REFERENCES Character(name)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TYPE e_warlock_patron AS ENUM (
    'The Archfey',
    'The Celestial',
    'The Fathomless',
    'The Fiend',
    'The Genie',
    'The Great Old One',
    'The Hexblade',
    'The Undying'
);

CREATE TABLE Warlocks (
    character_name      text,
    patron              e_warlock_patron,
    PRIMARY KEY (character_name),
    FOREIGN KEY (character_name) REFERENCES Character(name)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TYPE e_wizard_tradition AS ENUM (
    'School of Abjuration',
    'Bladesinging',
    'Chronurgy',
    'School of Conjuration',
    'School of Divination',
    'School of Enchantment',
    'School of Evocation',
    'Graviturgy',
    'School of Illusion',
    'School of Necromancy',
    'Order of Scribes',
    'School of Transmutation',
    'School of War Magic'
);

CREATE TABLE Wizards (
    character_name      text,
    arcane_tradition    e_wizard_tradition,
    PRIMARY KEY (character_name),
    FOREIGN KEY (character_name) REFERENCES Character(name)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);
