// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: seasons.sql

package sqlc

import (
	"context"

	null_v4 "gopkg.in/guregu/null.v4"
)

const getRolesForSeason = `-- name: GetRolesForSeason :many
SELECT DISTINCT usergroups_code FROM public.episodes_usergroups WHERE episodes_id IN
    (SELECT id FROM public.episodes WHERE season_id = $1)
`

func (q *Queries) GetRolesForSeason(ctx context.Context, seasonID null_v4.Int) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getRolesForSeason, seasonID)
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

const getSeason = `-- name: GetSeason :one
SELECT agerating_code, available_from, available_to, date_created, date_updated, id, image_file_id, legacy_description_id, legacy_id, legacy_title_id, publish_date, season_number, show_id, status, user_created, user_updated FROM public.seasons WHERE id = $1
`

func (q *Queries) GetSeason(ctx context.Context, id int32) (Season, error) {
	row := q.db.QueryRowContext(ctx, getSeason, id)
	var i Season
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
		&i.SeasonNumber,
		&i.ShowID,
		&i.Status,
		&i.UserCreated,
		&i.UserUpdated,
	)
	return i, err
}

const getSeasonTranslations = `-- name: GetSeasonTranslations :many
SELECT description, id, is_primary, languages_code, legacy_description_id, legacy_title_id, seasons_id, title FROM public.seasons_translations
`

func (q *Queries) GetSeasonTranslations(ctx context.Context) ([]SeasonsTranslation, error) {
	rows, err := q.db.QueryContext(ctx, getSeasonTranslations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SeasonsTranslation
	for rows.Next() {
		var i SeasonsTranslation
		if err := rows.Scan(
			&i.Description,
			&i.ID,
			&i.IsPrimary,
			&i.LanguagesCode,
			&i.LegacyDescriptionID,
			&i.LegacyTitleID,
			&i.SeasonsID,
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

const getSeasons = `-- name: GetSeasons :many
SELECT agerating_code, available_from, available_to, date_created, date_updated, id, image_file_id, legacy_description_id, legacy_id, legacy_title_id, publish_date, season_number, show_id, status, user_created, user_updated FROM public.seasons
`

func (q *Queries) GetSeasons(ctx context.Context) ([]Season, error) {
	rows, err := q.db.QueryContext(ctx, getSeasons)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Season
	for rows.Next() {
		var i Season
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
			&i.SeasonNumber,
			&i.ShowID,
			&i.Status,
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

const getTranslationsForSeason = `-- name: GetTranslationsForSeason :many
SELECT description, id, is_primary, languages_code, legacy_description_id, legacy_title_id, seasons_id, title FROM public.seasons_translations WHERE seasons_id = $1
`

func (q *Queries) GetTranslationsForSeason(ctx context.Context, seasonsID int32) ([]SeasonsTranslation, error) {
	rows, err := q.db.QueryContext(ctx, getTranslationsForSeason, seasonsID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SeasonsTranslation
	for rows.Next() {
		var i SeasonsTranslation
		if err := rows.Scan(
			&i.Description,
			&i.ID,
			&i.IsPrimary,
			&i.LanguagesCode,
			&i.LegacyDescriptionID,
			&i.LegacyTitleID,
			&i.SeasonsID,
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