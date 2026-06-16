-- name: CreateSnippet :one
INSERT INTO snippets(title, content, created, expires)
VALUES(
        $1,
        $2,
        NOW(),
        NOW() + ($3 * INTERVAL '1 days')
    )
RETURNING id;
-- name: GetSnippet :one
SELECT *
FROM snippets
WHERE(id = $1);