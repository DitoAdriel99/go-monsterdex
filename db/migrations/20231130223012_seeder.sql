-- +goose Up
-- +goose StatementBegin
INSERT INTO monster_types (name,created_at, updated_at)
VALUES ('Normal', NOW(), NOW()),
('Fire', NOW(), NOW()),
('Water', NOW(), NOW()),
('Grass', NOW(), NOW()),
('Electric', NOW(), NOW()),
('Ice', NOW(), NOW()),
('Fighting', NOW(), NOW()),
('Poison', NOW(), NOW()),
('Ground', NOW(), NOW()),
('Flying', NOW(), NOW()),
('Psychic', NOW(), NOW()),
('Bug', NOW(), NOW()),
('Rock', NOW(), NOW()),
('Ghost', NOW(), NOW()),
('Dragon', NOW(), NOW()),
('Dark', NOW(), NOW()),
('Steel', NOW(), NOW()),
('Fairy', NOW(), NOW()),
('???' , NOW(), NOW()), -- For unknown or special types
('Shadow', NOW(), NOW());

INSERT INTO monster_category (name,created_at, updated_at)
VALUES ('Fire Lizard', NOW(), NOW()),
('Water Turtle', NOW(), NOW()),
('Electric Rodent', NOW(), NOW()),
('Grass Snake', NOW(), NOW()),
('Rock Golem', NOW(), NOW()),
('Ice Bird', NOW(), NOW()),
('Psychic Spoon', NOW(), NOW()),
('Dark Bat', NOW(), NOW()),
('Fairy Pixie', NOW(), NOW()),
('Poisonous Frog', NOW(), NOW()),
('Flying Insect', NOW(), NOW()),
('Steel Mech', NOW(), NOW()),
('Dragon Serpent', NOW(), NOW()),
('Ghost Spirit', NOW(), NOW()),
('Normal Bunny', NOW(), NOW()),
('Ground Mole', NOW(), NOW()),
('Fighting Panda', NOW(), NOW()),
('Bug Beetle', NOW(), NOW()),
('Legendary Unicorn', NOW(), NOW()),
('Mystic Sphinx', NOW(), NOW());

INSERT INTO monsters (name, monster_category_id, description, image, types_id, height, weight, stats_hp, stats_attack, stats_defense, stats_speed, is_active, created_at, updated_at)
VALUES ('Charizard', 1, 'A fiery dragon-like creature', 'charizard.png', ARRAY[1, 6], 1.7, 90.5, 78, 84, 78, 100, true, NOW(), NOW()),
('Blastoise', 2, 'A powerful water-based tortoise', 'blastoise.png', ARRAY[2], 1.6, 85.5, 79, 83, 100, 78, true, NOW(), NOW()),
('Pikachu', 3, 'An electric rodent Pokémon', 'pikachu.png', ARRAY[4], 0.4, 6.0, 35, 55, 40, 90, true, NOW(), NOW()),
('Venusaur', 1, 'A grass-based dinosaur', 'venusaur.png', ARRAY[5, 6], 2.0, 100.0, 80, 82, 83, 80, true, NOW(), NOW()),
('Gyarados', 2, 'A fearsome water/flying serpent', 'gyarados.png', ARRAY[2, 6], 6.5, 235.0, 95, 125, 79, 81, true, NOW(), NOW()),
('Alakazam', 7, 'A psychic Pokémon with psychic spoons', 'alakazam.png', ARRAY[12], 1.5, 48.0, 55, 50, 45, 120, true, NOW(), NOW()),
('Gengar', 13, 'A mischievous ghostly creature', 'gengar.png', ARRAY[13, 14], 1.5, 40.5, 60, 65, 60, 110, true, NOW(), NOW()),
('Dragonite', 12, 'A majestic dragon-type Pokémon', 'dragonite.png', ARRAY[16, 6], 2.2, 210.0, 91, 134, 95, 80, true, NOW(), NOW()),
('Mewtwo', 14, 'A genetically created psychic Pokémon', 'mewtwo.png', ARRAY[12], 2.0, 122.0, 106, 110, 90, 130, true, NOW(), NOW()),
('Arceus', 18, 'A legendary deity Pokémon', 'arceus.png', ARRAY[17], 3.2, 320.0, 120, 120, 120, 120, true, NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
