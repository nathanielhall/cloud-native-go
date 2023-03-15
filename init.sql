CREATE DATABASE IF NOT EXISTS todos_db;

\connect todos_db;

CREATE TABLE IF NOT EXISTS todos(
    id SERIAL NOT NULL PRIMARY KEY,
    description VARCHAR(100),
    priority INT,
    status VARCHAR(100)
);