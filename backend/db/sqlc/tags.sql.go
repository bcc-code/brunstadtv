// Code generated by sqlc. DO NOT EDIT.
// source: tags.sql

package db

import (
	"context"

	null_v4 "gopkg.in/guregu/null.v4"
)

const getTag = `-- name: GetTag :one
SELECT id, type FROM tag
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTag(ctx context.Context, id int64) (Tag, error) {
	row := q.db.QueryRowContext(ctx, getTag, id)
	var i Tag
	err := row.Scan(&i.ID, &i.Type)
	return i, err
}

const insertTag = `-- name: InsertTag :one
WITH t AS (
    INSERT INTO tag (
        type
    ) VALUES (
        $1
    ) RETURNING id, type
),
tt AS (
    INSERT INTO tag_t (
        tag_id,
        language_code,
        title
    ) SELECT
        t.id,
        $2,
        $3
    FROM t RETURNING id, tag_id, language_code, title
)
SELECT 
    t.id,
    t.type,
    tt.title
    FROM t,tt
`

type InsertTagParams struct {
	Type         null_v4.String `db:"type" json:"type"`
	LanguageCode string         `db:"language_code" json:"languageCode"`
	Title        int16          `db:"title" json:"title"`
}

type InsertTagRow struct {
	ID    int64          `db:"id" json:"id"`
	Type  null_v4.String `db:"type" json:"type"`
	Title int16          `db:"title" json:"title"`
}

func (q *Queries) InsertTag(ctx context.Context, arg InsertTagParams) (InsertTagRow, error) {
	row := q.db.QueryRowContext(ctx, insertTag, arg.Type, arg.LanguageCode, arg.Title)
	var i InsertTagRow
	err := row.Scan(&i.ID, &i.Type, &i.Title)
	return i, err
}
