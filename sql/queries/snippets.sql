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
WHERE expires > NOW()
    AND id = $1;
-- name: GetLatestSnippets :many
SELECT *
FROM snippets
WHERE expires > NOW()
ORDER BY id DESC
LIMIT 10;