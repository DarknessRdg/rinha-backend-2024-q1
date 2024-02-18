
-- +migrate Up
INSERT INTO account (id, balance, "limit")
VALUES
    (1, 0, 100000),
    (2, 0, 80000),
    (3, 0, 1000000),
    (4, 0, 10000000),
    (5, 0, 500000);

-- +migrate Down
DELETE FROM account WHERE id <= 5;