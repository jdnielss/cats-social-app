-- Create the 'users' table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

-- Create the 'user_token' table
CREATE TABLE IF NOT EXISTS user_token (
    accessToken VARCHAR(255) PRIMARY KEY,
    expiredDate TIMESTAMP NOT NULL,
    userId INT NOT NULL,
    FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE
);

-- Create the 'cats' table
CREATE TABLE IF NOT EXISTS cats (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    race VARCHAR(30) NOT NULL,
    sex VARCHAR(10) NOT NULL,
    ageInMonth INT NOT NULL,
    description VARCHAR(200) NOT NULL,
    imageUrls VARCHAR(256) NOT NULL,
    hasMatched BOOLEAN NOT NULL,
    createdAt TIMESTAMP NOT NULL,
    userId INT NOT NULL,
    FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE
);

-- Create the 'matches' table
CREATE TABLE IF NOT EXISTS matches (
    id SERIAL PRIMARY KEY,
    matchCatId INT NOT NULL,
    catUserId INT NOT NULL,
    message TEXT,
    status VARCHAR(50),
    FOREIGN KEY (matchCatId) REFERENCES cats(id) ON DELETE CASCADE,
    FOREIGN KEY (catUserId) REFERENCES cats(id) ON DELETE CASCADE
);

-- Create indexes for the 'matches' table
CREATE INDEX IF NOT EXISTS idxMatchCatId ON matches (matchCatId);
CREATE INDEX IF NOT EXISTS idxCatUserId ON matches (catUserId);
