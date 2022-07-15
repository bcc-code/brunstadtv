package graph

import (
	"context"
	"github.com/bcc-code/brunstadtv/backend/common"
	gqlmodel "github.com/bcc-code/brunstadtv/backend/graph/model"
	"github.com/bcc-code/brunstadtv/backend/utils"
	"strconv"
)

func preloadItems(ctx context.Context, r *queryRootResolver, items []common.SearchResultItem) {
	var keysByCollection = map[string][]int{}

	for _, i := range items {
		keysByCollection[i.Collection] = append(keysByCollection[i.Collection], i.ID)
	}

	if keys, ok := keysByCollection["shows"]; ok {
		r.Loaders.ShowLoader.LoadMany(ctx, keys)
	}
	if keys, ok := keysByCollection["seasons"]; ok {
		r.Loaders.SeasonLoader.LoadMany(ctx, keys)
	}
	if keys, ok := keysByCollection["episodes"]; ok {
		r.Loaders.EpisodeLoader.LoadMany(ctx, keys)
	}
}

func filterOrAppend[t gqlmodel.SearchResultItem](ctx context.Context, results []gqlmodel.SearchResultItem, id int, factory func(context.Context, string) (*t, error)) []gqlmodel.SearchResultItem {
	item, err := factory(ctx, strconv.Itoa(id))
	if err != nil {
		return results
	}
	return append(results, *item)
}

func filterAndConvertToGQL(ctx context.Context, r *queryRootResolver, items []common.SearchResultItem) []gqlmodel.SearchResultItem {
	// Preload/fill query with all item IDs
	preloadItems(ctx, r, items)

	var results []gqlmodel.SearchResultItem
	for _, i := range items {
		switch i.Collection {
		case "shows":
			results = filterOrAppend(ctx, results, i.ID, r.Show)
		case "seasons":
			results = filterOrAppend(ctx, results, i.ID, r.Season)
		case "episodes":
			results = filterOrAppend(ctx, results, i.ID, r.Episode)
		}
	}
	return results
}

func searchResolver(r *queryRootResolver, ctx context.Context, queryString string, first *int, offset *int) (*gqlmodel.SearchResult, error) {
	ginCtx, err := utils.GinCtx(ctx)
	if err != nil {
		return nil, err
	}
	searchResult, err := r.SearchService.Search(ginCtx, common.SearchQuery{
		Query:  queryString,
		Limit:  first,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	return &gqlmodel.SearchResult{
		Result: filterAndConvertToGQL(ctx, r, searchResult.Result),
		Page:   searchResult.Page,
		Hits:   searchResult.HitCount,
	}, nil
}
