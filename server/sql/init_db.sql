-- Initialize DB
-- Only needs to be run once per server

CREATE DATABASE dnd_db;

\connect dnd_db

CREATE TABLE Player (
    username        text PRIMARY KEY,
    password        text,
    email           text,
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

CREATE TABLE Character (
    name                text PRIMARY KEY,
    weight              int,
    height              int,
    alignment           e_alignment,
    sex                 text,
    background          text,
    race                e_race,
    speed               int,
    strength            int,
    dexterity           int,
    intelligence        int,
    wisdom              int,
    charisma            int,
    constitution        int,
    hp_max              int,
    ability_points      int,
    xp_points           int,
    class               e_class,
    player_username     text REFERENCES Player(username)
);

CREATE TABLE CharacterAbilities (
    character_name      text REFERENCES Character(name),
    ability             text,
    PRIMARY KEY (character_name, ability)
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
    character_name      text REFERENCES Character(name),
    item_name           text,
    type                e_item_type,
    rarity              e_item_rarity,
    weight              int,
    gold_value          int,
    description         text,
    PRIMARY KEY (character_name, item_name)
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
    character_name      text REFERENCES Character(name),
    spell_name          text,
    level               int,
    school              e_school,
    concentration       bool,
    description         text,
    casting_time        int,
    range               int,
    duration            int,
    PRIMARY KEY (character_name, spell_name)
);

CREATE TABLE SpellComponents (
    character_name      text,
    spell_name          text,
    component           text,
    FOREIGN KEY (character_name, spell_name) REFERENCES Spells(character_name, spell_name),
    PRIMARY KEY (character_name, spell_name, component)
);

CREATE TABLE Campaign (
    name                text PRIMARY KEY,
    current_location    text,
    state               text
);

CREATE TABLE CampaignMilestones (
    campaign_name       text REFERENCES Campaign(name),
    milestone           text,
    PRIMARY KEY (campaign_name, milestone)
);

CREATE TABLE BelongsTo (
    character_name      text REFERENCES Character(name),
    campaign_name       text REFERENCES Campaign(name),
    PRIMARY KEY (character_name, campaign_name)
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
    character_name      text REFERENCES Character(name),
    primal_path         e_barbarian_primal_path,
    PRIMARY KEY (character_name)
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
    character_name      text REFERENCES Character(name),
    college             e_bard_college,
    PRIMARY KEY (character_name)
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
    character_name      text REFERENCES Character(name),
    divine_domain       e_cleric_domain,
    PRIMARY KEY (character_name)
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
    character_name      text REFERENCES Character(name),
    druid_circle        e_druid_circle,
    PRIMARY KEY (character_name)
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
    character_name      text REFERENCES Character(name),
    archetype           e_fighter_archetype,
    fighting_style      e_fighting_style,
    PRIMARY KEY (character_name)
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
    character_name      text REFERENCES Character(name),
    monastic_tradition  e_monk_tradition,
    PRIMARY KEY (character_name)
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
    character_name      text REFERENCES Character(name),
    oath                e_paladin_oath,
    fighting_style      e_fighting_style,
    PRIMARY KEY (character_name)
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
    character_name      text REFERENCES Character(name),
    conclave            e_ranger_conclave,
    favoured_enemy      text,
    PRIMARY KEY (character_name)
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
    character_name      text REFERENCES Character(name),
    archetype           e_rogue_archetype,
    PRIMARY KEY (character_name)
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
    character_name      text REFERENCES Character(name),
    sorcerous_origin    e_sorcerer_origin,
    PRIMARY KEY (character_name)
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
    character_name      text REFERENCES Character(name),
    patron              e_warlock_patron,
    PRIMARY KEY (character_name)
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
    character_name      text REFERENCES Character(name),
    arcane_tradition    e_wizard_tradition,
    PRIMARY KEY (character_name)
);
