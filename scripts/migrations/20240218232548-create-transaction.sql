
-- +migrate Up
CREATE TABLE transaction (
    id SERIAL PRIMARY KEY,
    account_id INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    type CHAR(1) NOT NULL,
    description VARCHAR(10) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_transaction_account_id
        FOREIGN KEY (account_id) REFERENCES account(id)
);

-- +migrate Down
DROP TABLE IF EXISTS transaction;
