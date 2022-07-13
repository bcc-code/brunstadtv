package asset

import (
	"context"
	"github.com/bcc-code/brunstadtv/backend/common"
	"strconv"

	"github.com/bcc-code/brunstadtv/backend/sqlc"
	"github.com/graph-gophers/dataloader/v7"
)

// NewBatchFileLoader returns a configured batch loader for GQL File
func NewBatchFileLoader(queries sqlc.Queries) *dataloader.Loader[int, []*sqlc.GetFilesForEpisodesRow] {
	return common.NewListBatchLoader(queries.GetFilesForEpisodes, func(row sqlc.GetFilesForEpisodesRow) int {
		return int(row.EpisodesID)
	}, func(id int) int32 {
		return int32(id)
	})
}

// GetFilesForEpisode retrieves file assets currently associated with the specified asset
//
// It uses the dataloader to efficiently load data from DB or cache (as available)
func GetFilesForEpisode(ctx context.Context, loader *dataloader.Loader[int, []*sqlc.GetFilesForEpisodesRow], id string) ([]*sqlc.GetFilesForEpisodesRow, error) {
	intID, _ := strconv.ParseInt(id, 10, 32)
	thunk := loader.Load(ctx, int(intID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// NewBatchStreamLoader returns a configured batch loader for GQL Stream
func NewBatchStreamLoader(queries sqlc.Queries) *dataloader.Loader[int, []*sqlc.GetStreamsForEpisodesRow] {
	return common.NewListBatchLoader(queries.GetStreamsForEpisodes, func(row sqlc.GetStreamsForEpisodesRow) int {
		return int(row.EpisodesID)
	}, func(id int) int32 {
		return int32(id)
	})
}

// GetStreamsForEpisode retrieves file assets currently associated with the specified asset
//
// It uses the dataloader to efficiently load data from DB or cache (as available)
func GetStreamsForEpisode(ctx context.Context, loader *dataloader.Loader[int, []*sqlc.GetStreamsForEpisodesRow], id string) ([]*sqlc.GetStreamsForEpisodesRow, error) {
	intID, _ := strconv.ParseInt(id, 10, 32)
	thunk := loader.Load(ctx, int(intID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result, nil
}
