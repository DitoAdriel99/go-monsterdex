-- +goose Up
-- +goose StatementBegin
CREATE TABLE if NOT EXISTS monster_owner (
	id              BIGSERIAL       PRIMARY KEY,
	monster_id      INTEGER			NOT NULL,
	user_id         INTEGER			NOT NULL,
    is_catched      BOOLEAN         NOT NULL,
	created_at      TIMESTAMP       NOT NULL,
	updated_at      TIMESTAMP       NOT NULL
);
CREATE INDEX idx_monster_owner ON monster_owner (is_catched);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if EXISTS monster_owner;
-- +goose StatementEnd
