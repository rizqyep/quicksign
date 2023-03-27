-- +migrate Up 
-- +migrate StatementBegin

CREATE TABLE users(
    id SERIAL PRIMARY KEY, 
    first_name VARCHAR(256),
    last_name VARCHAR (256),
    email VARCHAR(256) NOT NULL UNIQUE,
    username VARCHAR(32) UNIQUE NOT NULL, 
    password VARCHAR(256) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(), 
    updated_at TIMESTAMP DEFAULT NOW()
)

-- +migrate StatementEnd