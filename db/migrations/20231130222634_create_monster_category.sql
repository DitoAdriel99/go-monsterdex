-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS monster_category 
(
    id          BIGSERIAL       PRIMARY KEY,
    name        VARCHAR(225)    NOT NULL,
    created_at  TIMESTAMP       NOT NULL,
    updated_at  TIMESTAMP       NOT NULL
);
CREATE INDEX idx_monster_category ON monster_category (name);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS monster_category
-- +goose StatementEnd
