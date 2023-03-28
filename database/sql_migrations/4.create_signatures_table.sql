-- +migrate Up 
-- +migrate StatementBegin

CREATE TABLE signatures(
    id SERIAL PRIMARY KEY, 
    user_id INTEGER NOT NULL,
    signature_token VARCHAR(64) UNIQUE, 
    description VARCHAR(256) NOT NULL,
    qr_code_url VARCHAR(256) NOT NULL, 
    request_id INTEGER REFERENCES signature_requests(id), 
    created_at TIMESTAMP DEFAULT NOW(), 
    updated_at TIMESTAMP DEFAULT NOW()
)

-- +migrate StatementEnd