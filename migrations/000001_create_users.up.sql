CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    tg_id BIGINT NOT NULL UNIQUE,
    first_name TEXT NOT NULL,
    username TEXT UNIQUE
);
