-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS todos
(
    id SERIAL NOT NULL PRIMARY KEY,
    description VARCHAR(100),
    priority INT,
    status VARCHAR(100)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS todos;
