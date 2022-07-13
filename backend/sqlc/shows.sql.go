// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: shows.sql

package sqlc

import (
	"context"
	"time"

	"github.com/lib/pq"
	null_v4 "gopkg.in/guregu/null.v4"
)

const getAccessForShows = `-- name: GetAccessForShows :many
SELECT id, published, available_from, available_to, usergroups, usergroups_downloads, usergroups_earlyaccess FROM shows_access WHERE id = ANY($1::int[])
`

func (q *Queries) GetAccessForShows(ctx context.Context, dollar_1 []int32) ([]ShowsAccess, error) {
	rows, err := q.db.QueryContext(ctx, getAccessForShows, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ShowsAccess
	for rows.Next() {
		var i ShowsAccess
		if err := rows.Scan(
			&i.ID,
			&i.Published,
			&i.AvailableFrom,
			&i.AvailableTo,
			&i.Usergroups,
			&i.UsergroupsDownloads,
			&i.UsergroupsEarlyaccess,
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

const getRolesForShow = `-- name: GetRolesForShow :many
SELECT DISTINCT usergroups_code FROM public.episodes_usergroups WHERE episodes_id IN
    (SELECT id FROM episodes WHERE season_id IN
    (SELECT id FROM seasons WHERE show_id = $1))
`

func (q *Queries) GetRolesForShow(ctx context.Context, showID int32) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getRolesForShow, showID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var usergroups_code string
		if err := rows.Scan(&usergroups_code); err != nil {
			return nil, err
		}
		items = append(items, usergroups_code)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getShow = `-- name: GetShow :one
SELECT agerating_code, available_from, available_to, date_created, date_updated, id, image_file_id, legacy_description_id, legacy_id, legacy_title_id, publish_date, status, type, user_created, user_updated FROM shows WHERE id = $1
`

func (q *Queries) GetShow(ctx context.Context, id int32) (Show, error) {
	row := q.db.QueryRowContext(ctx, getShow, id)
	var i Show
	err := row.Scan(
		&i.AgeratingCode,
		&i.AvailableFrom,
		&i.AvailableTo,
		&i.DateCreated,
		&i.DateUpdated,
		&i.ID,
		&i.ImageFileID,
		&i.LegacyDescriptionID,
		&i.LegacyID,
		&i.LegacyTitleID,
		&i.PublishDate,
		&i.Status,
		&i.Type,
		&i.UserCreated,
		&i.UserUpdated,
	)
	return i, err
}

const getShowTranslations = `-- name: GetShowTranslations :many
SELECT description, id, is_primary, languages_code, legacy_description_id, legacy_tags, legacy_tags_id, legacy_title_id, shows_id, title FROM public.shows_translations
`

func (q *Queries) GetShowTranslations(ctx context.Context) ([]ShowsTranslation, error) {
	rows, err := q.db.QueryContext(ctx, getShowTranslations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ShowsTranslation
	for rows.Next() {
		var i ShowsTranslation
		if err := rows.Scan(
			&i.Description,
			&i.ID,
			&i.IsPrimary,
			&i.LanguagesCode,
			&i.LegacyDescriptionID,
			&i.LegacyTags,
			&i.LegacyTagsID,
			&i.LegacyTitleID,
			&i.ShowsID,
			&i.Title,
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

const getShows = `-- name: GetShows :many
SELECT agerating_code, available_from, available_to, date_created, date_updated, id, image_file_id, legacy_description_id, legacy_id, legacy_title_id, publish_date, status, type, user_created, user_updated FROM public.shows
`

func (q *Queries) GetShows(ctx context.Context) ([]Show, error) {
	rows, err := q.db.QueryContext(ctx, getShows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Show
	for rows.Next() {
		var i Show
		if err := rows.Scan(
			&i.AgeratingCode,
			&i.AvailableFrom,
			&i.AvailableTo,
			&i.DateCreated,
			&i.DateUpdated,
			&i.ID,
			&i.ImageFileID,
			&i.LegacyDescriptionID,
			&i.LegacyID,
			&i.LegacyTitleID,
			&i.PublishDate,
			&i.Status,
			&i.Type,
			&i.UserCreated,
			&i.UserUpdated,
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

const getShowsWithTranslationsByID = `-- name: GetShowsWithTranslationsByID :many
SELECT id, image_file_id, title, description, published, available_from, available_to, usergroups, download_groups, early_access_groups FROM shows_expanded WHERE id = ANY($1::int[])
`

func (q *Queries) GetShowsWithTranslationsByID(ctx context.Context, dollar_1 []int32) ([]ShowsExpanded, error) {
	rows, err := q.db.QueryContext(ctx, getShowsWithTranslationsByID, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ShowsExpanded
	for rows.Next() {
		var i ShowsExpanded
		if err := rows.Scan(
			&i.ID,
			&i.ImageFileID,
			&i.Title,
			&i.Description,
			&i.Published,
			&i.AvailableFrom,
			&i.AvailableTo,
			pq.Array(&i.Usergroups),
			pq.Array(&i.DownloadGroups),
			pq.Array(&i.EarlyAccessGroups),
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

const getTranslationsForShow = `-- name: GetTranslationsForShow :many
SELECT description, id, is_primary, languages_code, legacy_description_id, legacy_tags, legacy_tags_id, legacy_title_id, shows_id, title FROM public.shows_translations WHERE shows_id = $1
`

func (q *Queries) GetTranslationsForShow(ctx context.Context, showsID int32) ([]ShowsTranslation, error) {
	rows, err := q.db.QueryContext(ctx, getTranslationsForShow, showsID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ShowsTranslation
	for rows.Next() {
		var i ShowsTranslation
		if err := rows.Scan(
			&i.Description,
			&i.ID,
			&i.IsPrimary,
			&i.LanguagesCode,
			&i.LegacyDescriptionID,
			&i.LegacyTags,
			&i.LegacyTagsID,
			&i.LegacyTitleID,
			&i.ShowsID,
			&i.Title,
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

const getVisibilityForShow = `-- name: GetVisibilityForShow :one
SELECT id, status, publish_date, available_from, available_to FROM public.shows WHERE id = $1
`

type GetVisibilityForShowRow struct {
	ID            int32        `db:"id" json:"id"`
	Status        string       `db:"status" json:"status"`
	PublishDate   time.Time    `db:"publish_date" json:"publishDate"`
	AvailableFrom null_v4.Time `db:"available_from" json:"availableFrom"`
	AvailableTo   null_v4.Time `db:"available_to" json:"availableTo"`
}

func (q *Queries) GetVisibilityForShow(ctx context.Context, id int32) (GetVisibilityForShowRow, error) {
	row := q.db.QueryRowContext(ctx, getVisibilityForShow, id)
	var i GetVisibilityForShowRow
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.PublishDate,
		&i.AvailableFrom,
		&i.AvailableTo,
	)
	return i, err
}

const getVisibilityForShows = `-- name: GetVisibilityForShows :many
SELECT id, status, publish_date, available_from, available_to FROM public.shows
`

type GetVisibilityForShowsRow struct {
	ID            int32        `db:"id" json:"id"`
	Status        string       `db:"status" json:"status"`
	PublishDate   time.Time    `db:"publish_date" json:"publishDate"`
	AvailableFrom null_v4.Time `db:"available_from" json:"availableFrom"`
	AvailableTo   null_v4.Time `db:"available_to" json:"availableTo"`
}

func (q *Queries) GetVisibilityForShows(ctx context.Context) ([]GetVisibilityForShowsRow, error) {
	rows, err := q.db.QueryContext(ctx, getVisibilityForShows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetVisibilityForShowsRow
	for rows.Next() {
		var i GetVisibilityForShowsRow
		if err := rows.Scan(
			&i.ID,
			&i.Status,
			&i.PublishDate,
			&i.AvailableFrom,
			&i.AvailableTo,
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

const refreshShowAccessView = `-- name: RefreshShowAccessView :one
SELECT update_access('shows_access')
`

func (q *Queries) RefreshShowAccessView(ctx context.Context) (bool, error) {
	row := q.db.QueryRowContext(ctx, refreshShowAccessView)
	var update_access bool
	err := row.Scan(&update_access)
	return update_access, err
}
