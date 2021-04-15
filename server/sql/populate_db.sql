-- Populate DB with some preset values for testing
-- Since this is just used for testing, don't expect many of these numbers to make sense in a real game

-- Passwords are formatted as "test<username>". Obviously this not
-- secure, but we need credentials for testing.
INSERT INTO Player (username, password, name) VALUES
(
    'kjerome',
    '$2a$10$90TeDDAllwO5dWpi0IS0y.iU6uUDHiKsZoLq9Ez/mRZipS0UIpZYq',
    'Kevin Jerome'
),
(
    'drowsell',
    '$2a$10$QDtmmPVkT0nsBNAZi1NBwOfaQVdhgqAYej3Y4rYH9mUCjMizfmrsS',
    'Dylan Rowsell'
),
(
    'ksmolko',
    '$2a$10$9g28.ZhUpLJBUuYOuFq0RO2ipZpElgbtCaoeKQPR6HEPaLtKdt.3W',
    'Kyle Smolko'
),
(
    'agervacio',
    '$2a$10$ebt5.5mQg3zAdataDc4dHOZNwNoBkhTPVMYeIip5De0imIMBhI642',
    'Arsenen Gervacio'
),
(
    'newuser',
    '$2a$10$obetuWhYWFuJW6lNvyYb9uothIStuAgmPwhX4/r8wG2nbWZkDfGlK',
    'New User'
);

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
    'Berserker',
    'kjerome'
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
    'Abjuration',
    'drowsell'
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
    'Psi Warrior',
    'ksmolko'
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
    'Assassin',
    'agervacio'
),
(
    DEFAULT,
    'Ives',
    180,
    180,
    'Lawful Evil',
    'Male',
    'Just a dude',
    'Half-Orc',
    30,
    9,
    16,
    14,
    14,
    10,
    8,
    7,
    0,
    400,
    'Monk',
    'Astral Self',
    'newuser'
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
    2,
    'Chainmail',
    'Armor',
    'Common',
    55,
    75,
    1,
    'Heavy Armor'
),
(
    3,
    'Wand',
    'Wand',
    'Common',
    1,
    10,
    1,
    'Arcane Focus'
),
(
    4,
    'Pride Silk Outfit',
    'Armor',
    'Common',
    4,
    500,
    1,
    'Light Armor'
),
(
    5,
    'Blight Ichor',
    'Potion',
    'Common',
    0,
    200,
    10,
    'Vulnerability to Psychic Damage for 1 hour'
),
(
    1,
    'Greatsword',
    'Weapon',
    'Common',
    6,
    50,
    1,
    'Martial Melee Weapon'
),
(
    2,
    'Breastplate',
    'Armor',
    'Common',
    20,
    400,
    1,
    'Medium Armor'
),
(
    3,
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
    1,
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
    3,
    'Burning Hands',
    1,
    'Evocation',
    FALSE,
    'Cause a thin sheet of flames to shoot forth from your fingertips',
    1,
    15,
    0
),
(
    4,
    'Shield',
    1,
    'Abjuration',
    FALSE,
    'An invisible barrier of magical force appears and protects you.',
    0,
    0,
    0
),
(
    5,
    'False Life',
    1,
    'Necromancy',
    FALSE,
    'Bolstering yourself with a necromantic facsimile of life.',
    1,
    0,
    60
);

INSERT INTO Campaign VALUES
(
    1,
    'Campaign 1',
    'Baldur''s Gate',
    'Doing nothing of substance',
    'kjerome'
),
(
    2,
    'Campaign 2',
    'Icewind Dale',
    'Fighting a troll',
    'drowsell'
),
(
    3,
    'Campaign 3',
    'Nowhere',
    'Not started',
    'ksmolko'
),
(
    4,
    'Campaign 4',
    'Baldur''s Gate',
    'Following Campaign 1',
    'agervacio'
),
(
    5,
    'Campaign 5',
    'The Void',
    'Not started',
    'newuser'
);

INSERT INTO CampaignMilestones VALUES
(
    1,
    'Started your wonderful journey'
),
(
    2,
    'Defeated the big bad Dragon'
),
(
    3,
    'Started your wonderful journey'
),
(
    4,
    'Started your wonderful journey'
),
(
    5,
    'Found an orphan'
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


-- Manually reset primary key indeces so that "serial" values continue from correct location
SELECT setval(pg_get_serial_sequence('Character', 'id'), (SELECT MAX(id) FROM Character));
SELECT setval(pg_get_serial_sequence('Campaign', 'id'), (SELECT MAX(id) FROM Campaign));
