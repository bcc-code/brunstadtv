package search

import (
	"context"
	"strconv"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/bcc-code/brunstadtv/backend/sqlc"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

func (service *Service) mapSeasonToSearchObject(
	ctx context.Context,
	item sqlc.Season,
	image *sqlc.DirectusFile,
) searchObject {
	object := searchObject{}
	itemId := int(item.ID)
	object[rolesField] = service.getRolesForSeason(ctx, item.ID)
	object[idField] = "seasons-" + strconv.Itoa(itemId)
	object[createdAtField] = item.DateCreated.UTC().Unix()
	object[updatedAtField] = item.DateUpdated.UTC().Unix()
	object[publishedAtField] = item.PublishDate.UTC().Unix()
	if image != nil && image.FilenameDisk.Valid {
		object[imageField] = image.GetImageUrl()
	}

	object.assignVisibility(service.getVisibilityForSeason(ctx, item.ID))
	title, description := toLocaleStrings(service.getTranslationsForSeason(ctx, item.ID))
	object.mapFromLocaleString(titleField, title)
	object.mapFromLocaleString(descriptionField, description)
	showTitle, _ := toLocaleStrings(service.getTranslationsForShow(ctx, item.ShowID))
	object.mapFromLocaleString(showTitleField, showTitle)
	return object
}

func (service *Service) indexSeasons(
	ctx context.Context,
	items []sqlc.Season,
	imageDict map[uuid.UUID]sqlc.DirectusFile,
	index *search.Index,
) error {
	objects := lo.Map(items, func(item sqlc.Season, _ int) searchObject {
		var image *sqlc.DirectusFile
		if item.ImageFileID.Valid {
			thumbnailResult := imageDict[item.ImageFileID.UUID]
			image = &thumbnailResult
		}
		return service.mapSeasonToSearchObject(ctx, item, image)
	})

	return indexObjects(index, objects)
}

func (service *Service) indexSeason(ctx context.Context, item sqlc.Season) error {
	var image *sqlc.DirectusFile
	if item.ImageFileID.Valid {
		thumbnailResult, _ := service.queries.GetFile(ctx, item.ImageFileID.UUID)
		image = &thumbnailResult
	}

	object := service.mapSeasonToSearchObject(ctx, item, image)
	_, err := service.index.SaveObject(object)
	return err
}
