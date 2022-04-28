package events

// AssetDelivered is the event originating from MEDIABANKEN notifying the system
// That an update for a VOD episide is availble on the storage.
// Note that this may trigger a creation of a new Episode as required
type AssetDelivered struct {
	Prefix string `json:"prefix"`
}