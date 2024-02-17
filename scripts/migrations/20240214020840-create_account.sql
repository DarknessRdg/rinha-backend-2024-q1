-- +migrate Up
CREATE TABLE account (
    id SERIAL PRIMARY KEY,
    balance INT NOT NULL,
    "limit" INT NOT NULL
);


-- +migrate Down
DROP TABLE IF EXISTS account;
