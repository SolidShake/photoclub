CREATE TABLE IF NOT EXISTS user_profile(
    id SERIAL PRIMARY KEY,
    user_id INTEGER UNIQUE,
    type VARCHAR(255),
    logo VARCHAR(255),
    about TEXT,
    FOREIGN KEY (user_id) REFERENCES users (id)
);