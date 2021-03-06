package section

import (
	"github.com/bcc-code/brunstadtv/backend/common"
	"github.com/bcc-code/brunstadtv/backend/sqlc"
	"github.com/graph-gophers/dataloader/v7"
)

// NewBatchLoader returns a configured batch loader for GQL Section
func NewBatchLoader(queries sqlc.Queries) *dataloader.Loader[int, *sqlc.SectionExpanded] {
	return common.NewBatchLoader(queries.GetSections, func(row sqlc.SectionExpanded) int {
		return int(row.ID)
	}, func(id int) int32 {
		return int32(id)
	})
}

// NewListBatchLoader returns related data for a page
func NewListBatchLoader(queries sqlc.Queries) *dataloader.Loader[int, []*sqlc.SectionExpanded] {
	return common.NewListBatchLoader(queries.GetSectionsForPageIDs, func(i sqlc.SectionExpanded) int {
		return int(i.PageID.ValueOrZero())
	}, func(id int) int32 {
		return int32(id)
	})
}
