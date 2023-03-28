-- +migrate Up 
-- +migrate StatementBegin

CREATE TABLE reset_password_tokens(
    id SERIAL PRIMARY KEY, 
    token VARCHAR(256),
    email VARCHAR(256) NOT NULL UNIQUE,
    valid BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW(), 
    updated_at TIMESTAMP DEFAULT NOW()
)

-- +migrate StatementEnd