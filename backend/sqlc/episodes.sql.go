// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: episodes.sql

package sqlc

import (
	"context"
)

const getEpisode = `-- name: GetEpisode :one
SELECT agerating_code, asset_id, available_from, available_to, date_created, date_updated, episode_number, id, image_file_id, legacy_description_id, legacy_extra_description_id, legacy_id, legacy_program_id, legacy_tags_id, legacy_title_id, migration_data, publish_date, season_id, status, type, user_created, user_updated FROM public.episodes WHERE id = $1
`

func (q *Queries) GetEpisode(ctx context.Context, id int32) (Episode, error) {
	row := q.db.QueryRowContext(ctx, getEpisode, id)
	var i Episode
	err := row.Scan(
		&i.AgeratingCode,
		&i.AssetID,
		&i.AvailableFrom,
		&i.AvailableTo,
		&i.DateCreated,
		&i.DateUpdated,
		&i.EpisodeNumber,
		&i.ID,
		&i.ImageFileID,
		&i.LegacyDescriptionID,
		&i.LegacyExtraDescriptionID,
		&i.LegacyID,
		&i.LegacyProgramID,
		&i.LegacyTagsID,
		&i.LegacyTitleID,
		&i.MigrationData,
		&i.PublishDate,
		&i.SeasonID,
		&i.Status,
		&i.Type,
		&i.UserCreated,
		&i.UserUpdated,
	)
	return i, err
}

const getEpisodeTranslations = `-- name: GetEpisodeTranslations :many
SELECT description, episodes_id, extra_description, id, is_primary, languages_code, title FROM public.episodes_translations
`

func (q *Queries) GetEpisodeTranslations(ctx context.Context) ([]EpisodesTranslation, error) {
	rows, err := q.db.QueryContext(ctx, getEpisodeTranslations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EpisodesTranslation
	for rows.Next() {
		var i EpisodesTranslation
		if err := rows.Scan(
			&i.Description,
			&i.EpisodesID,
			&i.ExtraDescription,
			&i.ID,
			&i.IsPrimary,
			&i.LanguagesCode,
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

const getEpisodes = `-- name: GetEpisodes :many
SELECT agerating_code, asset_id, available_from, available_to, date_created, date_updated, episode_number, id, image_file_id, legacy_description_id, legacy_extra_description_id, legacy_id, legacy_program_id, legacy_tags_id, legacy_title_id, migration_data, publish_date, season_id, status, type, user_created, user_updated FROM public.episodes
`

func (q *Queries) GetEpisodes(ctx context.Context) ([]Episode, error) {
	rows, err := q.db.QueryContext(ctx, getEpisodes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Episode
	for rows.Next() {
		var i Episode
		if err := rows.Scan(
			&i.AgeratingCode,
			&i.AssetID,
			&i.AvailableFrom,
			&i.AvailableTo,
			&i.DateCreated,
			&i.DateUpdated,
			&i.EpisodeNumber,
			&i.ID,
			&i.ImageFileID,
			&i.LegacyDescriptionID,
			&i.LegacyExtraDescriptionID,
			&i.LegacyID,
			&i.LegacyProgramID,
			&i.LegacyTagsID,
			&i.LegacyTitleID,
			&i.MigrationData,
			&i.PublishDate,
			&i.SeasonID,
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

const getTranslationsForEpisode = `-- name: GetTranslationsForEpisode :many
SELECT description, episodes_id, extra_description, id, is_primary, languages_code, title FROM public.episodes_translations WHERE episodes_id = $1
`

func (q *Queries) GetTranslationsForEpisode(ctx context.Context, episodesID int32) ([]EpisodesTranslation, error) {
	rows, err := q.db.QueryContext(ctx, getTranslationsForEpisode, episodesID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EpisodesTranslation
	for rows.Next() {
		var i EpisodesTranslation
		if err := rows.Scan(
			&i.Description,
			&i.EpisodesID,
			&i.ExtraDescription,
			&i.ID,
			&i.IsPrimary,
			&i.LanguagesCode,
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
