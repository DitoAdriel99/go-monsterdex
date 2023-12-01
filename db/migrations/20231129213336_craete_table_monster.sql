-- +goose Up
-- +goose StatementBegin
CREATE TABLE if NOT EXISTS monsters (
	id                    BIGSERIAL       PRIMARY KEY,
	name                  VARCHAR(255)    NOT NULL,
	monster_category_id   INTEGER         NOT NULL,
	description           TEXT,
	image     		        TEXT,
	types_id              TEXT[]          NOT NULL,
	height                FLOAT     		  NOT NULL,
	weight                FLOAT     		  NOT NULL,
	stats_hp              INTEGER         NOT NULL,
	stats_attack          INTEGER         NOT NULL,
	stats_defense         INTEGER         NOT NULL,
	stats_speed           INTEGER         NOT NULL,
	is_active             BOOLEAN         DEFAULT true,
	created_at            TIMESTAMP       NOT NULL,
	updated_at            TIMESTAMP       NOT NULL
);
CREATE INDEX idx_monsters ON monsters (id,name,types_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX if EXISTS idx_monsters;
DROP TABLE if EXISTS monsters;
-- +goose StatementEnd
