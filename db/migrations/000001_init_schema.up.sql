CREATE TABLE IF NOT EXISTS cats (
    id UUID PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    race VARCHAR(30) NOT NULL,
    sex VARCHAR(10) NOT NULL,
    ageInMonth INT NOT NULL,
    description VARCHAR(200) NOT NULL,
    imageUrls VARCHAR(256) NOT NULL,
    hasMatched BOOLEAN NOT NULL,
    createdAt TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS "cats"


INSERT INTO cats (id, name, race, sex, ageInMonth, description, imageUrls, hasMatched, createdAt)
VALUES 
    ('a6c24f27-8c4f-4e5f-8c77-2f71a6c6a3a1', 'Fluffy', 'Persian', 'male', 12, 'A lovely Persian cat', 'http://example.com/cat1.jpg,http://example.com/cat2.jpg', true, NOW()),
    ('a6c24f27-8c4f-4e5f-8c77-2f71a6c6a3a2', 'Whiskers', 'Maine Coon', 'female', 24, 'A beautiful Maine Coon', 'http://example.com/cat3.jpg,http://example.com/cat4.jpg', false, NOW()),
    ('a6c24f27-8c4f-4e5f-8c77-2f71a6c6a3a3', 'Mittens', 'Siamese', 'male', 6, 'A playful Siamese kitten', 'http://example.com/cat5.jpg,http://example.com/cat6.jpg', false, NOW());
