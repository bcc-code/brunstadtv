// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: pages.sql

package sqlc

import (
	"context"

	"github.com/lib/pq"
	"github.com/tabbed/pqtype"
	null_v4 "gopkg.in/guregu/null.v4"
)

const getPages = `-- name: getPages :many
WITH t AS (SELECT ts.pages_id,
                  json_object_agg(ts.languages_code, ts.title)       AS title,
                  json_object_agg(ts.languages_code, ts.description) AS description
           FROM pages_translations ts
           GROUP BY ts.pages_id),
     r AS (SELECT p_1.id                                                   AS page_id,
                  (SELECT array_agg(DISTINCT eu.usergroups_code) AS array_agg
                   FROM sections_usergroups eu
                   WHERE (eu.sections_id IN (SELECT e.id
                                             FROM episodes e
                                             WHERE e.season_id = p_1.id))) AS roles
           FROM pages p_1)
SELECT p.id,
       p.code,
       p.type,
       p.status::text = 'published'::text AS published,
       p.show_id,
       p.season_id,
       p.episode_id,
       p.collection,
       t.title,
       t.description,
       r.roles
FROM pages p
         LEFT JOIN t ON t.pages_id = p.id
         LEFT JOIN r ON r.page_id = p.id
WHERE id = ANY($1::int[])
`

type getPagesRow struct {
	ID          int32                 `db:"id" json:"id"`
	Code        null_v4.String        `db:"code" json:"code"`
	Type        null_v4.String        `db:"type" json:"type"`
	Published   bool                  `db:"published" json:"published"`
	ShowID      null_v4.Int           `db:"show_id" json:"showID"`
	SeasonID    null_v4.Int           `db:"season_id" json:"seasonID"`
	EpisodeID   null_v4.Int           `db:"episode_id" json:"episodeID"`
	Collection  null_v4.String        `db:"collection" json:"collection"`
	Title       pqtype.NullRawMessage `db:"title" json:"title"`
	Description pqtype.NullRawMessage `db:"description" json:"description"`
	Roles       interface{}           `db:"roles" json:"roles"`
}

func (q *Queries) getPages(ctx context.Context, dollar_1 []int32) ([]getPagesRow, error) {
	rows, err := q.db.QueryContext(ctx, getPages, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getPagesRow
	for rows.Next() {
		var i getPagesRow
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Type,
			&i.Published,
			&i.ShowID,
			&i.SeasonID,
			&i.EpisodeID,
			&i.Collection,
			&i.Title,
			&i.Description,
			&i.Roles,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPagesByCode = `-- name: getPagesByCode :many
WITH t AS (SELECT ts.pages_id,
                  json_object_agg(ts.languages_code, ts.title)       AS title,
                  json_object_agg(ts.languages_code, ts.description) AS description
           FROM pages_translations ts
           GROUP BY ts.pages_id),
     r AS (SELECT p_1.id                                                   AS page_id,
                  (SELECT array_agg(DISTINCT eu.usergroups_code) AS array_agg
                   FROM sections_usergroups eu
                   WHERE (eu.sections_id IN (SELECT e.id
                                             FROM episodes e
                                             WHERE e.season_id = p_1.id))) AS roles
           FROM pages p_1)
SELECT p.id,
       p.code,
       p.type,
       p.status::text = 'published'::text AS published,
       p.show_id,
       p.season_id,
       p.episode_id,
       p.collection,
       t.title,
       t.description,
       r.roles
FROM pages p
         LEFT JOIN t ON t.pages_id = p.id
         LEFT JOIN r ON r.page_id = p.id
WHERE code = ANY($1::varchar[])
`

type getPagesByCodeRow struct {
	ID          int32                 `db:"id" json:"id"`
	Code        null_v4.String        `db:"code" json:"code"`
	Type        null_v4.String        `db:"type" json:"type"`
	Published   bool                  `db:"published" json:"published"`
	ShowID      null_v4.Int           `db:"show_id" json:"showID"`
	SeasonID    null_v4.Int           `db:"season_id" json:"seasonID"`
	EpisodeID   null_v4.Int           `db:"episode_id" json:"episodeID"`
	Collection  null_v4.String        `db:"collection" json:"collection"`
	Title       pqtype.NullRawMessage `db:"title" json:"title"`
	Description pqtype.NullRawMessage `db:"description" json:"description"`
	Roles       interface{}           `db:"roles" json:"roles"`
}

func (q *Queries) getPagesByCode(ctx context.Context, dollar_1 []string) ([]getPagesByCodeRow, error) {
	rows, err := q.db.QueryContext(ctx, getPagesByCode, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []getPagesByCodeRow
	for rows.Next() {
		var i getPagesByCodeRow
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Type,
			&i.Published,
			&i.ShowID,
			&i.SeasonID,
			&i.EpisodeID,
			&i.Collection,
			&i.Title,
			&i.Description,
			&i.Roles,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPages = `-- name: listPages :many
WITH t AS (SELECT ts.pages_id,
                  json_object_agg(ts.languages_code, ts.title)       AS title,
                  json_object_agg(ts.languages_code, ts.description) AS description
           FROM pages_translations ts
           GROUP BY ts.pages_id),
     r AS (SELECT p_1.id                                                   AS page_id,
                  (SELECT array_agg(DISTINCT eu.usergroups_code) AS array_agg
                   FROM sections_usergroups eu
                   WHERE (eu.sections_id IN (SELECT e.id
                                             FROM episodes e
                                             WHERE e.season_id = p_1.id))) AS roles
           FROM pages p_1)
SELECT p.id,
       p.code,
       p.type,
       p.status::text = 'published'::text AS published,
       p.show_id,
       p.season_id,
       p.episode_id,
       p.collection,
       t.title,
       t.description,
       r.roles
FROM pages p
         LEFT JOIN t ON t.pages_id = p.id
         LEFT JOIN r ON r.page_id = p.id
`

type listPagesRow struct {
	ID          int32                 `db:"id" json:"id"`
	Code        null_v4.String        `db:"code" json:"code"`
	Type        null_v4.String        `db:"type" json:"type"`
	Published   bool                  `db:"published" json:"published"`
	ShowID      null_v4.Int           `db:"show_id" json:"showID"`
	SeasonID    null_v4.Int           `db:"season_id" json:"seasonID"`
	EpisodeID   null_v4.Int           `db:"episode_id" json:"episodeID"`
	Collection  null_v4.String        `db:"collection" json:"collection"`
	Title       pqtype.NullRawMessage `db:"title" json:"title"`
	Description pqtype.NullRawMessage `db:"description" json:"description"`
	Roles       interface{}           `db:"roles" json:"roles"`
}

func (q *Queries) listPages(ctx context.Context) ([]listPagesRow, error) {
	rows, err := q.db.QueryContext(ctx, listPages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []listPagesRow
	for rows.Next() {
		var i listPagesRow
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Type,
			&i.Published,
			&i.ShowID,
			&i.SeasonID,
			&i.EpisodeID,
			&i.Collection,
			&i.Title,
			&i.Description,
			&i.Roles,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
