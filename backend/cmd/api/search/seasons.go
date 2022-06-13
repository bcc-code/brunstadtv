package search

import (
	"context"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/bcc-code/brunstadtv/backend/sqlc"
	"github.com/bcc-code/mediabank-bridge/log"
	"strconv"
)

func indexSeasons(queries *sqlc.Queries, ctx context.Context, index *search.Index) {
	items, err := queries.GetSeasons(ctx)
	if err != nil {
		return
	}
	itemTranslations, err := queries.GetSeasonTranslations(ctx)
	if err != nil {
		return
	}
	translationDictionary := mapToRelatedId(itemTranslations)
	objects := mapToSearchObjects(items, func(item sqlc.Season) map[string]any {
		values := map[string]any{}
		itemId := int(item.ID)
		values["objectID"] = "season-" + strconv.Itoa(itemId)
		translations := translationDictionary[itemId]
		for _, translation := range translations {
			translatedValues := map[string]string{
				description: translation.Description.ValueOrZero(),
				title:       translation.Title.ValueOrZero(),
			}
			for field, value := range translatedValues {
				if value != "" {
					values[field+"_"+translation.LanguagesCode] = value
				}
			}
		}
		return values
	})

	err = indexObjects(index, objects)

	if err != nil {
		log.L.Error().Err(err).Msg("Failed to index objects")
		return
	}
}
