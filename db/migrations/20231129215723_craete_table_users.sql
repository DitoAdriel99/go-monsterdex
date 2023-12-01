-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users 
(
    id          BIGSERIAL          PRIMARY KEY,
    fullname    VARCHAR(225)    NOT NULL,
    email       VARCHAR(225)    NOT NULL,
    password    VARCHAR(225)    NOT NULL,
    role        VARCHAR(10)     NOT NULL,
    created_at  TIMESTAMP       NOT NULL,
    updated_at  TIMESTAMP       NOT NULL
);
INSERT INTO users (fullname, email, password, role, created_at, updated_at)
VALUES ('Admin', 'admin@gmail.com', '$2a$14$2xYJc4OuqcLnb0HhiPQ5l.ArlzBTdV8m8uniGGBEpGbqZa/R8yEJe', 'admin', NOW(), NOW()),
('Akuntes1', 'akuntes@gmail.com', '$2a$14$2xYJc4OuqcLnb0HhiPQ5l.ArlzBTdV8m8uniGGBEpGbqZa/R8yEJe','user',  NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users
-- +goose StatementEnd
