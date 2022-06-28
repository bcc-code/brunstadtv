package events

// Collection of possible sources for events
const (
	SourceMediaBanken    = "mediabanken"
	SourceCloudScheduler = "cloudscheduler"
)

// Collection of possible event types
const (
	TypeAssetDelivered = "asset.delivered"
	TypeRefreshView    = "view.refresh"
	TypeSearchReindex  = "search.reindex"
	TypeSearchIndex    = "search.index"
)
