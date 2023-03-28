-- +migrate Up 
-- +migrate StatementBegin

CREATE TABLE signature_requests(
    id SERIAL PRIMARY KEY, 
    status VARCHAR(32),
    requester_name VARCHAR(256) NOT NULL,
    requester_email VARCHAR(256) NOT NULL,
    description VARCHAR(256) NOT NULL, 
    approver_id INTEGER  REFERENCES users(id), 
    created_at TIMESTAMP DEFAULT NOW(), 
    updated_at TIMESTAMP DEFAULT NOW()
)

-- +migrate StatementEnd