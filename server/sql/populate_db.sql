-- Populate DB with some preset values for testing
-- Since this is just used for testing, don't expect many of these numbers to make sense in a real game

-- This won't work if run directly, it must be run through the go test script (for proper password hashing)

INSERT INTO Character VALUES
(
    DEFAULT,
    'Korgoth',
    135,
    215,
    'Chaotic Evil',
    'Male',
    'Destroyer of Worlds. Harbinger of Doom.',
    'Half-Orc',
    30,
    18,
    8,
    8,
    8,
    10,
    14,
    14,
    0,
    0,
    'Barbarian',
    'mc1'
),
(
    DEFAULT,
    'Vlad',
    60,
    170,
    'Lawful Neutral',
    'Male',
    'A powerful and feared wizard',
    'Half-Elf',
    30,
    8,
    8,
    18,
    18,
    10,
    9,
    6,
    0,
    300,
    'Wizard',
    'alove'
),
(
    DEFAULT,
    'Jane',
    50,
    160,
    'Lawful Good',
    'Female',
    'A natural and just leader.',
    'Human',
    30,
    10,
    10,
    12,
    12,
    18,
    10,
    10,
    0,
    300,
    'Fighter',
    'inew'
),
(
    DEFAULT,
    'Nilys',
    30,
    90,
    'Chaotic Neutral',
    'Female',
    'A thief, loyal only to her friends and no one else',
    'Gnome',
    20,
    9,
    16,
    14,
    14,
    10,
    8,
    7,
    0,
    400,
    'Rogue',
    'at1'
);

INSERT INTO Barbarians VALUES (1, 'Berserker');
INSERT INTO Wizards VALUES (2, 'Evocation');
INSERT INTO Fighters VALUES (3, 'Battle Master', 'Great Weapon Fighting');
INSERT INTO Rogues VALUES (4, 'Thief');

INSERT INTO CharacterAbilities VALUES
(
    1,
    'Rage'
),
(
    3,
    'Second Wind'
),
(
    3,
    'Action Surge'
),
(
    3,
    'Sneak Attack'
),
(
    3,
    'Thieves'' Cant'
),
(
    4,
    'Cunning Action'
);

INSERT INTO Items VALUES
(
    1,
    'Warhammer',
    'Weapon',
    'Common',
    2,
    15,
    1,
    'Martial Melee Weapon'
),
(
    1,
    'Chainmail',
    'Armor',
    'Common',
    55,
    75,
    1,
    'Heavy Armor'
),
(
    2,
    'Wand',
    'Wand',
    'Common',
    1,
    10,
    1,
    'Arcane Focus'
),
(
    2,
    'Pride Silk Outfit',
    'Armor',
    'Common',
    4,
    500,
    1,
    'Light Armor'
),
(
    2,
    'Blight Ichor',
    'Potion',
    'Common',
    0,
    200,
    10,
    'Vulnerability to Psychic Damage for 1 hour'
),
(
    3,
    'Greatsword',
    'Weapon',
    'Common',
    6,
    50,
    1,
    'Martial Melee Weapon'
),
(
    3,
    'Breastplate',
    'Armor',
    'Common',
    20,
    400,
    1,
    'Medium Armor'
),
(
    4,
    'Dagger',
    'Weapon',
    'Common',
    1,
    2,
    2,
    'Simple Melee Weapon'
),
(
    4,
    'Leather',
    'Armor',
    'Common',
    10,
    10,
    1,
    'Light Armor'
);

INSERT INTO Spells VALUES
(
    2,
    'Fire Bolt',
    0,
    'Evocation',
    FALSE,
    'Hurl a mote of fire at a creature',
    1,
    120,
    0
),
(
    2,
    'Frostbite',
    0,
    'Evocation',
    FALSE,
    'Cause numbing frost to form on one creature',
    1,
    60,
    0
),
(
    2,
    'Burning Hands',
    1,
    'Evocation',
    FALSE,
    'Cause a thin sheet of flames to shoot forth from your fingertips',
    1,
    15,
    0
);

INSERT INTO SpellComponents VALUES
(
    2,
    'Fire Bolt',
    'Verbal'
),
(
    2,
    'Fire Bolt',
    'Somatic'
),
(
    2,
    'Frostbite',
    'Verbal'
),
(
    2,
    'Frostbite',
    'Somatic'
),
(
    2,
    'Burning Hands',
    'Verbal'
),
(
    2,
    'Burning Hands',
    'Somatic'
);

INSERT INTO Campaign VALUES
(
    1,
    'Campaign 1',
    'Baldur''s Gate',
    'Doing nothing of substance',
    'ae1'
),
(
    2,
    'Campaign 2',
    'Icewind Dale',
    'Fighting a troll',
    'ae1'
);

INSERT INTO CampaignMilestones VALUES
(
    1,
    'Started your wonderful journey'
),
(
    1,
    'Defeated the big bad Dragon'
),
(
    2,
    'Started your wonderful journey'
);

INSERT INTO BelongsTo VALUES
(
    1,
    1
),
(
    2,
    1
),
(
    3,
    1
),
(
    4,
    1
),
(
    1,
    2
),
(
    2,
    2
),
(
    3,
    2
),
(
    4,
    2
);
