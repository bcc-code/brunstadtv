// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/tabbed/pqtype"
)

type Agerating struct {
	Code      string         `json:"code"`
	Title     sql.NullString `json:"title"`
	SortOrder int16          `json:"sortOrder"`
}

type AgeratingT struct {
	ID            int32          `json:"id"`
	AgeratingCode sql.NullString `json:"ageratingCode"`
	LanguageCode  string         `json:"languageCode"`
}

type Asset struct {
	ID                 int64          `json:"id"`
	SourceID           sql.NullString `json:"sourceID"`
	PublishedVersionID sql.NullInt64  `json:"publishedVersionID"`
	Name               int16          `json:"name"`
}

type AssetFormat struct {
	ID      int64          `json:"id"`
	Name    sql.NullString `json:"name"`
	Url     sql.NullString `json:"url"`
	Bitrate sql.NullInt64  `json:"bitrate"`
}

type AssetVersion struct {
	ID      int64         `json:"id"`
	AssetID sql.NullInt64 `json:"assetID"`
}

type AssetVersionFormat struct {
	ID             int64         `json:"id"`
	AssetversionID sql.NullInt64 `json:"assetversionID"`
	AssetformatID  sql.NullInt64 `json:"assetformatID"`
}

type CalendarEvent struct {
	ID               int64          `json:"id"`
	EventID          sql.NullString `json:"eventID"`
	Start            sql.NullTime   `json:"start"`
	End              sql.NullTime   `json:"end"`
	Locked           bool           `json:"locked"`
	UseImageFromBcco bool           `json:"useImageFromBcco"`
	ImageID          sql.NullInt64  `json:"imageID"`
	AudienceID       sql.NullInt64  `json:"audienceID"`
}

type CalendareventT struct {
	ID              int64          `json:"id"`
	CalendareventID int64          `json:"calendareventID"`
	LanguageCode    string         `json:"languageCode"`
	Name            sql.NullString `json:"name"`
	ImageID         sql.NullInt64  `json:"imageID"`
}

type Collectable struct {
	ID            int64          `json:"id"`
	Type          sql.NullString `json:"type"`
	AvailableFrom sql.NullTime   `json:"availableFrom"`
	AvailableTo   sql.NullTime   `json:"availableTo"`
	Status        int16          `json:"status"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
}

type Collection struct {
	ID          int64                 `json:"id"`
	Type        sql.NullString        `json:"type"`
	QueryID     sql.NullInt64         `json:"queryID"`
	QueryParams pqtype.NullRawMessage `json:"queryParams"`
	PageID      sql.NullInt64         `json:"pageID"`
}

type CollectionCollectable struct {
	ID            int64         `json:"id"`
	CollectionID  sql.NullInt64 `json:"collectionID"`
	CollectableID sql.NullInt64 `json:"collectableID"`
	SortOrder     int16         `json:"sortOrder"`
}

type Faq struct {
	ID       int16 `json:"id"`
	Category int16 `json:"category"`
}

type FaqT struct {
	ID           int16          `json:"id"`
	FaqID        int16          `json:"faqID"`
	LanguageCode string         `json:"languageCode"`
	Question     sql.NullString `json:"question"`
	Answer       sql.NullString `json:"answer"`
}

type Faqcategory struct {
	ID        int16 `json:"id"`
	SortOrder int16 `json:"sortOrder"`
	Status    int16 `json:"status"`
}

type FaqcategoryT struct {
	ID            int16  `json:"id"`
	FaqcategoryID int16  `json:"faqcategoryID"`
	LanguageCode  string `json:"languageCode"`
	Title         int16  `json:"title"`
}

type Image struct {
	ID  int64  `json:"id"`
	Url string `json:"url"`
}

type Language struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type MediaT struct {
	ID              int64          `json:"id"`
	MediaID         sql.NullInt64  `json:"mediaID"`
	LanguageCode    string         `json:"languageCode"`
	Title           sql.NullString `json:"title"`
	Description     sql.NullString `json:"description"`
	LongDescription sql.NullString `json:"longDescription"`
	ImageID         sql.NullInt64  `json:"imageID"`
}

type Medium struct {
	ID                int64          `json:"id"`
	CollectableType   sql.NullString `json:"collectableType"`
	MediaType         sql.NullString `json:"mediaType"`
	PrimaryGroupID    sql.NullInt64  `json:"primaryGroupID"`
	SubclippedMediaID sql.NullInt64  `json:"subclippedMediaID"`
	ReferenceMediaID  sql.NullInt64  `json:"referenceMediaID"`
	SequenceNumber    int16          `json:"sequenceNumber"`
	// for subclips and markers
	StartTime sql.NullFloat64 `json:"startTime"`
	// for subclips
	EndTime   sql.NullFloat64 `json:"endTime"`
	AssetID   sql.NullInt64   `json:"assetID"`
	Agerating sql.NullString  `json:"agerating"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
}

type Notification struct {
	ID            int64          `json:"id"`
	CohortID      string         `json:"cohortID"`
	Action        sql.NullString `json:"action"`
	InternalUrl   sql.NullString `json:"internalUrl"`
	ExternalUrl   sql.NullString `json:"externalUrl"`
	Badge         int16          `json:"badge"`
	ScheduledTime sql.NullTime   `json:"scheduledTime"`
}

type NotificationAction struct {
	Code string         `json:"code"`
	Name sql.NullString `json:"name"`
}

type NotificationT struct {
	ID             int64          `json:"id"`
	NotificationID sql.NullInt64  `json:"notificationID"`
	LanguageCode   string         `json:"languageCode"`
	Title          sql.NullString `json:"title"`
	Message        sql.NullString `json:"message"`
}

type Page struct {
	ID              int64          `json:"id"`
	CollectableType sql.NullString `json:"collectableType"`
}

type PageSection struct {
	ID        int64         `json:"id"`
	PageID    sql.NullInt64 `json:"pageID"`
	SectionID sql.NullInt64 `json:"sectionID"`
}

type PageT struct {
	ID           int64          `json:"id"`
	LanguageCode string         `json:"languageCode"`
	PageID       sql.NullInt64  `json:"pageID"`
	Title        sql.NullString `json:"title"`
	Description  sql.NullString `json:"description"`
}

type Query struct {
	ID            int64           `json:"id"`
	Template      json.RawMessage `json:"template"`
	Parameters    json.RawMessage `json:"parameters"`
	SystemDefined bool            `json:"systemDefined"`
}

type Section struct {
	ID              int64                 `json:"id"`
	Type            string                `json:"type"`
	CollectionID    sql.NullInt64         `json:"collectionID"`
	VisibilityRules pqtype.NullRawMessage `json:"visibilityRules"`
}

type SectionItem struct {
	ID            int64          `json:"id"`
	SectionID     int64          `json:"sectionID"`
	Type          string         `json:"type"`
	CollectableID sql.NullInt64  `json:"collectableID"`
	Url           sql.NullString `json:"url"`
	CollectionID  sql.NullInt64  `json:"collectionID"`
	SortOrder     int16          `json:"sortOrder"`
}

type SectionItemT struct {
	ID            int64          `json:"id"`
	LanguageCode  string         `json:"languageCode"`
	SectionItemID sql.NullInt64  `json:"sectionItemID"`
	Title         sql.NullString `json:"title"`
}

type SectionT struct {
	ID           int64          `json:"id"`
	SectionID    sql.NullInt64  `json:"sectionID"`
	LanguageCode string         `json:"languageCode"`
	Title        sql.NullString `json:"title"`
}

type SectionitemUsergroup struct {
	ID            int64          `json:"id"`
	SectionitemID sql.NullInt64  `json:"sectionitemID"`
	UsergroupID   sql.NullString `json:"usergroupID"`
}

type Systemdatum struct {
	ID                  int16         `json:"id"`
	Live                bool          `json:"live"`
	FullMaintenanceMode bool          `json:"fullMaintenanceMode"`
	MetaImageID         sql.NullInt64 `json:"metaImageID"`
}

type Tag struct {
	ID int64 `json:"id"`
	// speaker, composer, etc.
	Type sql.NullString `json:"type"`
}

type TagCollectable struct {
	ID            int64         `json:"id"`
	CollectableID sql.NullInt64 `json:"collectableID"`
	TagID         sql.NullInt64 `json:"tagID"`
}

type TagT struct {
	ID           int64         `json:"id"`
	TagID        sql.NullInt64 `json:"tagID"`
	LanguageCode string        `json:"languageCode"`
	Title        int16         `json:"title"`
}

type Tvguideentry struct {
	ID                int64         `json:"id"`
	Start             sql.NullTime  `json:"start"`
	End               sql.NullTime  `json:"end"`
	EventID           sql.NullInt64 `json:"eventID"`
	MediaID           sql.NullInt64 `json:"mediaID"`
	ImageID           sql.NullInt64 `json:"imageID"`
	UseImageFromMedia bool          `json:"useImageFromMedia"`
}

type TvguideentryT struct {
	ID             int64          `json:"id"`
	TvguideentryID sql.NullInt64  `json:"tvguideentryID"`
	LanguageCode   string         `json:"languageCode"`
	ImageID        sql.NullInt64  `json:"imageID"`
	Title          sql.NullString `json:"title"`
	Description    sql.NullString `json:"description"`
}

type UserDataDevice struct {
	ID        int64          `json:"id"`
	UserID    string         `json:"userID"`
	PushToken sql.NullString `json:"pushToken"`
	App       string         `json:"app"`
}

type UserDataList struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"userID"`
}

type UserDataListMedium struct {
	ID        int64         `json:"id"`
	MediaID   int64         `json:"mediaID"`
	ListID    sql.NullInt64 `json:"listID"`
	SortOrder int16         `json:"sortOrder"`
}

type UserDataMediaProgress struct {
	ID          int64          `json:"id"`
	MediaID     sql.NullInt64  `json:"mediaID"`
	UserID      sql.NullString `json:"userID"`
	Seconds     sql.NullInt32  `json:"seconds"`
	LastUpdated sql.NullTime   `json:"lastUpdated"`
}

type UserDataUserList struct {
	ID        int64         `json:"id"`
	UserID    string        `json:"userID"`
	ListID    sql.NullInt64 `json:"listID"`
	SortOrder int16         `json:"sortOrder"`
}

type UserDataUsergroupUser struct {
	ID          int64  `json:"id"`
	UsergroupID string `json:"usergroupID"`
	UserID      string `json:"userID"`
}

type Usergroup struct {
	ID string `json:"id"`
}

type UsergroupCollectable struct {
	ID            int64          `json:"id"`
	UsergroupID   sql.NullString `json:"usergroupID"`
	CollectableID sql.NullInt64  `json:"collectableID"`
}
