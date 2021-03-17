-- Populate DB with some preset values for testing
-- Since this is just used for testing, don't expect many of these numbers to make sense in a real game

INSERT INTO Player VALUES
('ae1', 'password1', 'Albert Einstein'),
('mc1', 'password2', 'Marie Curie'),
('alove', 'password3', 'Ada Lovelace'),
('inew', 'password4', 'Isaac Newton'),
('at1', 'password5', 'Alan Turing');

INSERT INTO Character VALUES
(
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

INSERT INTO Barbarians VALUES ('Korgoth', 'Berserker');
INSERT INTO Wizards VALUES ('Vlad', 'Evocation');
INSERT INTO Fighters VALUES ('Jane', 'Battle Master', 'Great Weapon Fighting');
INSERT INTO Rogues VALUES ('Nilys', 'Thief');

INSERT INTO CharacterAbilities VALUES
(
    'Korgoth',
    'Rage'
),
(
    'Jane',
    'Second Wind'
),
(
    'Jane',
    'Action Surge'
),
(
    'Nilys',
    'Sneak Attack'
),
(
    'Nilys',
    'Thieves'' Cant'
),
(
    'Nilys',
    'Cunning Action'
);

INSERT INTO Items VALUES
(
    'Korgoth',
    'Warhammer',
    'Weapon',
    'Common',
    2,
    15,
    1,
    'Martial Melee Weapon'
),
(
    'Korgoth',
    'Chainmail',
    'Armor',
    'Common',
    55,
    75,
    1,
    'Heavy Armor'
),
(
    'Vlad',
    'Wand',
    'Wand',
    'Common',
    1,
    10,
    1,
    'Arcane Focus'
),
(
    'Vlad',
    'Pride Silk Outfit',
    'Armor',
    'Common',
    4,
    500,
    1,
    'Light Armor'
),
(
    'Vlad',
    'Blight Ichor',
    'Potion',
    'Common',
    0,
    200,
    10,
    'Vulnerability to Psychic Damage for 1 hour'
),
(
    'Jane',
    'Greatsword',
    'Weapon',
    'Common',
    6,
    50,
    1,
    'Martial Melee Weapon'
),
(
    'Jane',
    'Breastplate',
    'Armor',
    'Common',
    20,
    400,
    1,
    'Medium Armor'
),
(
    'Nilys',
    'Dagger',
    'Weapon',
    'Common',
    1,
    2,
    2,
    'Simple Melee Weapon'
),
(
    'Nilys',
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
    'Vlad',
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
    'Vlad',
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
    'Vlad',
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
    'Vlad',
    'Fire Bolt',
    'Verbal'
),
(
    'Vlad',
    'Fire Bolt',
    'Somatic'
),
(
    'Vlad',
    'Frostbite',
    'Verbal'
),
(
    'Vlad',
    'Frostbite',
    'Somatic'
),
(
    'Vlad',
    'Burning Hands',
    'Verbal'
),
(
    'Vlad',
    'Burning Hands',
    'Somatic'
);

INSERT INTO Campaign VALUES
(
    'Campaign 1',
    'Baldur''s Gate',
    'Doing nothing of substance'
),
(
    'Campaign 2',
    'Icewind Dale',
    'Fighting a troll'
);

INSERT INTO CampaignMilestones VALUES
(
    'Campaign 1',
    'Started your wonderful journey'
),
(
    'Campaign 1',
    'Defeated the big bad Dragon'
),
(
    'Campaign 2',
    'Started your wonderful journey'
);

INSERT INTO BelongsTo VALUES
(
    'Korgoth',
    'Campaign 1'
),
(
    'Vlad',
    'Campaign 1'
),
(
    'Jane',
    'Campaign 1'
),
(
    'Nilys',
    'Campaign 1'
),
(
    'Korgoth',
    'Campaign 2'
),
(
    'Vlad',
    'Campaign 2'
),
(
    'Jane',
    'Campaign 2'
),
(
    'Nilys',
    'Campaign 2'
);
