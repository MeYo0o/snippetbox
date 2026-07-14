-- +goose up
CREATE TABLE IF NOT EXISTS snippets(
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created TIMESTAMPTZ NOT NULL,
    expires TIMESTAMPTZ NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_snippets_created ON snippets(created);
-- +goose down
DROP TABLE snippets;