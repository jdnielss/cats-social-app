-- Create the 'users' table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

-- Create the 'user_token' table
CREATE TABLE IF NOT EXISTS user_token (
    access_token VARCHAR(255) PRIMARY KEY,
    expired_date TIMESTAMP NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create the 'cats' table
CREATE TABLE IF NOT EXISTS cats (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    race VARCHAR(30) NOT NULL,
    sex VARCHAR(10) NOT NULL,
    age_in_month INT NOT NULL,
    description VARCHAR(200) NOT NULL,
    image_urls VARCHAR(256) NOT NULL,
    has_matched BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create the 'matches' table
CREATE TABLE IF NOT EXISTS matches (
    id SERIAL PRIMARY KEY,
    match_cat_id INT NOT NULL,
    cat_user_id INT NOT NULL,
    message TEXT,
    status VARCHAR(50),
    FOREIGN KEY (match_cat_id) REFERENCES cats(id) ON DELETE CASCADE,
    FOREIGN KEY (cat_user_id) REFERENCES cats(id) ON DELETE CASCADE
);

-- Create indexes for the 'matches' table
CREATE INDEX IF NOT EXISTS idx_match_cat_id ON matches (match_cat_id);
CREATE INDEX IF NOT EXISTS idx_cat_user_id ON matches (cat_user_id);
