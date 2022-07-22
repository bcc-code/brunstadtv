// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlmodel

import (
	"fmt"
	"io"
	"strconv"
)

type Item interface {
	IsItem()
}

type Page interface {
	IsPage()
}

type SearchResultItem interface {
	IsSearchResultItem()
}

type Section interface {
	IsSection()
}

type SectionBody interface {
	IsSectionBody()
}

type BubblesItemsSection struct {
	ID          string  `json:"id"`
	Items       []Item  `json:"items"`
	BorderColor *string `json:"borderColor"`
}

func (BubblesItemsSection) IsSectionBody() {}

type Calendar struct {
	Period *CalendarPeriod `json:"period"`
	Day    *CalendarDay    `json:"day"`
}

type CalendarDay struct {
	ID             string          `json:"id"`
	Events         []*Event        `json:"events"`
	TvGuideEntries []*TvGuideEntry `json:"tvGuideEntries"`
}

type CalendarPeriod struct {
	ID         string   `json:"id"`
	ActiveDays []string `json:"activeDays"`
	Events     []*Event `json:"events"`
}

type Chapter struct {
	ID    string `json:"id"`
	Start int    `json:"start"`
	Title string `json:"title"`
}

type ContainerSection struct {
	ID       string    `json:"id"`
	Title    *string   `json:"title"`
	Sections []Section `json:"sections"`
}

func (ContainerSection) IsSection() {}

type CustomPage struct {
	ID          string             `json:"id"`
	Code        string             `json:"code"`
	Title       *string            `json:"title"`
	Description *string            `json:"description"`
	Sections    *SectionConnection `json:"sections"`
}

func (CustomPage) IsPage() {}

type DefaultPage struct {
	ID          string             `json:"id"`
	Code        string             `json:"code"`
	Collection  string             `json:"collection"`
	Title       *string            `json:"title"`
	Description *string            `json:"description"`
	Sections    *SectionConnection `json:"sections"`
}

func (DefaultPage) IsPage() {}

type Episode struct {
	ID                string     `json:"id"`
	Title             string     `json:"title"`
	Description       string     `json:"description"`
	ExtraDescription  string     `json:"extraDescription"`
	Streams           []*Stream  `json:"streams"`
	Files             []*File    `json:"files"`
	Chapters          []*Chapter `json:"chapters"`
	Season            *Season    `json:"season"`
	Duration          int        `json:"duration"`
	AudioLanguages    []Language `json:"audioLanguages"`
	SubtitleLanguages []Language `json:"subtitleLanguages"`
	EpisodeNumber     *int       `json:"episodeNumber"`
}

type EpisodeItem struct {
	ID       string   `json:"id"`
	Title    *string  `json:"title"`
	ImageURL string   `json:"imageUrl"`
	Episode  *Episode `json:"episode"`
}

func (EpisodeItem) IsItem() {}

type EpisodePage struct {
	ID          string             `json:"id"`
	Code        string             `json:"code"`
	Title       *string            `json:"title"`
	Description *string            `json:"description"`
	Sections    *SectionConnection `json:"sections"`
	Episode     *Episode           `json:"episode"`
}

func (EpisodePage) IsPage() {}

type EpisodeSearchItem struct {
	ID          string  `json:"id"`
	Collection  string  `json:"collection"`
	Title       string  `json:"title"`
	Header      *string `json:"header"`
	Description *string `json:"description"`
	Highlight   *string `json:"highlight"`
	Image       *string `json:"image"`
	URL         string  `json:"url"`
	ShowID      *string `json:"showId"`
	ShowTitle   *string `json:"showTitle"`
	Show        *Show   `json:"show"`
	SeasonID    *string `json:"seasonId"`
	SeasonTitle *string `json:"seasonTitle"`
	Season      *Season `json:"season"`
}

func (EpisodeSearchItem) IsSearchResultItem() {}

type Event struct {
	ID             string          `json:"id"`
	Start          string          `json:"start"`
	End            string          `json:"end"`
	TvGuideEntries []*TvGuideEntry `json:"tvGuideEntries"`
	BannerImageURL string          `json:"bannerImageURL"`
}

type Faq struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type FAQCategory struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Questions []*Faq `json:"questions"`
}

type File struct {
	ID               string    `json:"id"`
	URL              string    `json:"url"`
	AudioLanguage    Language  `json:"audioLanguage"`
	SubtitleLanguage *Language `json:"subtitleLanguage"`
	Size             *int      `json:"size"`
	FileName         string    `json:"fileName"`
	MimeType         string    `json:"mimeType"`
}

type ItemSection struct {
	ID     string      `json:"id"`
	Title  *string     `json:"title"`
	Body   SectionBody `json:"body"`
	PageID string      `json:"pageId"`
}

func (ItemSection) IsSection() {}

type PageItem struct {
	ID       string  `json:"id"`
	Title    *string `json:"title"`
	ImageURL string  `json:"imageUrl"`
	PageID   string  `json:"pageId"`
	Page     Page    `json:"page"`
}

func (PageItem) IsItem() {}

type PaginationInfo struct {
	EndCursor   string `json:"endCursor"`
	HasNextPage bool   `json:"hasNextPage"`
}

type SearchResult struct {
	Hits   int                `json:"hits"`
	Page   int                `json:"page"`
	Result []SearchResultItem `json:"result"`
}

type Season struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Number      int        `json:"number"`
	Show        *Show      `json:"show"`
	Episodes    []*Episode `json:"episodes"`
}

type SeasonPage struct {
	ID          string             `json:"id"`
	Code        string             `json:"code"`
	Title       *string            `json:"title"`
	Description *string            `json:"description"`
	Sections    *SectionConnection `json:"sections"`
	Season      *Season            `json:"season"`
}

func (SeasonPage) IsPage() {}

type SeasonSearchItem struct {
	ID          string  `json:"id"`
	Collection  string  `json:"collection"`
	Title       string  `json:"title"`
	Header      *string `json:"header"`
	Description *string `json:"description"`
	Highlight   *string `json:"highlight"`
	Image       *string `json:"image"`
	URL         string  `json:"url"`
	ShowID      string  `json:"showId"`
	ShowTitle   string  `json:"showTitle"`
	Show        *Show   `json:"show"`
}

func (SeasonSearchItem) IsSearchResultItem() {}

type SectionConnection struct {
	Edges    []*SectionEdge  `json:"edges"`
	PageInfo *PaginationInfo `json:"pageInfo"`
	Cursor   string          `json:"cursor"`
}

type SectionEdge struct {
	ID   string  `json:"id"`
	Node Section `json:"node"`
}

type Settings struct {
	AudioLanguages    []Language `json:"audioLanguages"`
	SubtitleLanguages []Language `json:"subtitleLanguages"`
}

type Show struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	EpisodeCount int       `json:"episodeCount"`
	SeasonCount  int       `json:"seasonCount"`
	Seasons      []*Season `json:"seasons"`
}

type ShowPage struct {
	ID          string             `json:"id"`
	Code        string             `json:"code"`
	Title       *string            `json:"title"`
	Description *string            `json:"description"`
	Sections    *SectionConnection `json:"sections"`
	Show        *Show              `json:"show"`
}

func (ShowPage) IsPage() {}

type ShowSearchItem struct {
	ID          string  `json:"id"`
	Collection  string  `json:"collection"`
	Title       string  `json:"title"`
	Header      *string `json:"header"`
	Description *string `json:"description"`
	Highlight   *string `json:"highlight"`
	Image       *string `json:"image"`
	URL         string  `json:"url"`
}

func (ShowSearchItem) IsSearchResultItem() {}

type SliderItemsSection struct {
	ID    string `json:"id"`
	Items []Item `json:"items"`
}

func (SliderItemsSection) IsSectionBody() {}

type Stream struct {
	ID                string     `json:"id"`
	URL               string     `json:"url"`
	AudioLanguages    []Language `json:"audioLanguages"`
	SubtitleLanguages []Language `json:"subtitleLanguages"`
	Type              StreamType `json:"type"`
}

type TvGuideEntry struct {
	ID      string   `json:"id"`
	Start   string   `json:"start"`
	End     string   `json:"end"`
	Episode *Episode `json:"episode"`
}

type URLItem struct {
	ID       string  `json:"id"`
	Title    *string `json:"title"`
	ImageURL string  `json:"imageUrl"`
	URL      string  `json:"url"`
}

func (URLItem) IsItem() {}

type User struct {
	ID        *string   `json:"id"`
	Anonymous bool      `json:"anonymous"`
	BccMember bool      `json:"bccMember"`
	Audience  *string   `json:"audience"`
	Email     *string   `json:"email"`
	Settings  *Settings `json:"settings"`
	Roles     []string  `json:"roles"`
}

type Language string

const (
	LanguageEn Language = "en"
	LanguageNo Language = "no"
	LanguageDe Language = "de"
)

var AllLanguage = []Language{
	LanguageEn,
	LanguageNo,
	LanguageDe,
}

func (e Language) IsValid() bool {
	switch e {
	case LanguageEn, LanguageNo, LanguageDe:
		return true
	}
	return false
}

func (e Language) String() string {
	return string(e)
}

func (e *Language) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Language(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Language", str)
	}
	return nil
}

func (e Language) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type StreamType string

const (
	StreamTypeHls  StreamType = "hls"
	StreamTypeCmaf StreamType = "cmaf"
	StreamTypeDash StreamType = "dash"
)

var AllStreamType = []StreamType{
	StreamTypeHls,
	StreamTypeCmaf,
	StreamTypeDash,
}

func (e StreamType) IsValid() bool {
	switch e {
	case StreamTypeHls, StreamTypeCmaf, StreamTypeDash:
		return true
	}
	return false
}

func (e StreamType) String() string {
	return string(e)
}

func (e *StreamType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StreamType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StreamType", str)
	}
	return nil
}

func (e StreamType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
