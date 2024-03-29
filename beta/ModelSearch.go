// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// SearchAggregation undocumented
type SearchAggregation struct {
	// Object is the base model of SearchAggregation
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Buckets undocumented
	Buckets []SearchBucket `json:"buckets,omitempty"`
	// Field undocumented
	Field *string `json:"field,omitempty"`
}

func NewSearchAggregation() (*SearchAggregation, error) {
	newSearchAggregation := &SearchAggregation{
		ODataType: "#microsoft.graph.SearchAggregation",
	}
	return newSearchAggregation, nil
}

// SearchAlteration undocumented
type SearchAlteration struct {
	// Object is the base model of SearchAlteration
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AlteredHighlightedQueryString undocumented
	AlteredHighlightedQueryString *string `json:"alteredHighlightedQueryString,omitempty"`
	// AlteredQueryString undocumented
	AlteredQueryString *string `json:"alteredQueryString,omitempty"`
	// AlteredQueryTokens undocumented
	AlteredQueryTokens []AlteredQueryToken `json:"alteredQueryTokens,omitempty"`
}

func NewSearchAlteration() (*SearchAlteration, error) {
	newSearchAlteration := &SearchAlteration{
		ODataType: "#microsoft.graph.SearchAlteration",
	}
	return newSearchAlteration, nil
}

// SearchAlterationOptions undocumented
type SearchAlterationOptions struct {
	// Object is the base model of SearchAlterationOptions
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// EnableModification undocumented
	EnableModification *bool `json:"enableModification,omitempty"`
	// EnableSuggestion undocumented
	EnableSuggestion *bool `json:"enableSuggestion,omitempty"`
}

func NewSearchAlterationOptions() (*SearchAlterationOptions, error) {
	newSearchAlterationOptions := &SearchAlterationOptions{
		ODataType: "#microsoft.graph.SearchAlterationOptions",
	}
	return newSearchAlterationOptions, nil
}

// SearchBucket undocumented
type SearchBucket struct {
	// Object is the base model of SearchBucket
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AggregationFilterToken undocumented
	AggregationFilterToken *string `json:"aggregationFilterToken,omitempty"`
	// Count undocumented
	Count *int `json:"count,omitempty"`
	// Key undocumented
	Key *string `json:"key,omitempty"`
}

func NewSearchBucket() (*SearchBucket, error) {
	newSearchBucket := &SearchBucket{
		ODataType: "#microsoft.graph.SearchBucket",
	}
	return newSearchBucket, nil
}

// SearchEntity undocumented
type SearchEntity struct {
	// Entity is the base model of SearchEntity
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Acronyms undocumented
	Acronyms []SearchAcronym `json:"acronyms,omitempty"`
	// Bookmarks undocumented
	Bookmarks []SearchBookmark `json:"bookmarks,omitempty"`
	// Qnas undocumented
	Qnas []SearchQna `json:"qnas,omitempty"`
}

func NewSearchEntity() (*SearchEntity, error) {
	newSearchEntity := &SearchEntity{
		ODataType: "#microsoft.graph.SearchEntity",
	}
	return newSearchEntity, nil
}

// SearchHit undocumented
type SearchHit struct {
	// Object is the base model of SearchHit
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ContentSource undocumented
	ContentSource *string `json:"contentSource,omitempty"`
	// HitID undocumented
	HitID *string `json:"hitId,omitempty"`
	// IsCollapsed undocumented
	IsCollapsed *bool `json:"isCollapsed,omitempty"`
	// Rank undocumented
	Rank *int `json:"rank,omitempty"`
	// ResultTemplateID undocumented
	ResultTemplateID *string `json:"resultTemplateId,omitempty"`
	// Summary undocumented
	Summary *string `json:"summary,omitempty"`
	// Underscoreid undocumented
	Underscoreid *string `json:"_id,omitempty"`
	// Underscorescore undocumented
	Underscorescore *int `json:"_score,omitempty"`
	// Underscoresummary undocumented
	Underscoresummary *string `json:"_summary,omitempty"`
	// Resource undocumented
	Resource *Entity `json:"resource,omitempty"`
	// Underscoresource undocumented
	Underscoresource *Entity `json:"_source,omitempty"`
}

func NewSearchHit() (*SearchHit, error) {
	newSearchHit := &SearchHit{
		ODataType: "#microsoft.graph.SearchHit",
	}
	return newSearchHit, nil
}

// SearchHitsContainer undocumented
type SearchHitsContainer struct {
	// Object is the base model of SearchHitsContainer
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Aggregations undocumented
	Aggregations []SearchAggregation `json:"aggregations,omitempty"`
	// Hits undocumented
	Hits []SearchHit `json:"hits,omitempty"`
	// MoreResultsAvailable undocumented
	MoreResultsAvailable *bool `json:"moreResultsAvailable,omitempty"`
	// Total undocumented
	Total *int `json:"total,omitempty"`
}

func NewSearchHitsContainer() (*SearchHitsContainer, error) {
	newSearchHitsContainer := &SearchHitsContainer{
		ODataType: "#microsoft.graph.SearchHitsContainer",
	}
	return newSearchHitsContainer, nil
}

// SearchQuery undocumented
type SearchQuery struct {
	// Object is the base model of SearchQuery
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// QueryString undocumented
	QueryString *string `json:"queryString,omitempty"`
	// QueryTemplate undocumented
	QueryTemplate *string `json:"queryTemplate,omitempty"`
	// QueryUnderscorestring undocumented
	QueryUnderscorestring *SearchQueryString `json:"query_string,omitempty"`
}

func NewSearchQuery() (*SearchQuery, error) {
	newSearchQuery := &SearchQuery{
		ODataType: "#microsoft.graph.SearchQuery",
	}
	return newSearchQuery, nil
}

// SearchQueryString undocumented
type SearchQueryString struct {
	// Object is the base model of SearchQueryString
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Query undocumented
	Query *string `json:"query,omitempty"`
}

func NewSearchQueryString() (*SearchQueryString, error) {
	newSearchQueryString := &SearchQueryString{
		ODataType: "#microsoft.graph.SearchQueryString",
	}
	return newSearchQueryString, nil
}

// SearchRequestObject undocumented
type SearchRequestObject struct {
	// Object is the base model of SearchRequestObject
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AggregationFilters undocumented
	AggregationFilters []string `json:"aggregationFilters,omitempty"`
	// Aggregations undocumented
	Aggregations []AggregationOption `json:"aggregations,omitempty"`
	// CollapseProperties undocumented
	CollapseProperties []CollapseProperty `json:"collapseProperties,omitempty"`
	// ContentSources undocumented
	ContentSources []string `json:"contentSources,omitempty"`
	// EnableTopResults undocumented
	EnableTopResults *bool `json:"enableTopResults,omitempty"`
	// EntityTypes undocumented
	EntityTypes []EntityType `json:"entityTypes,omitempty"`
	// Fields undocumented
	Fields []string `json:"fields,omitempty"`
	// From undocumented
	From *int `json:"from,omitempty"`
	// Query undocumented
	Query *SearchQuery `json:"query,omitempty"`
	// QueryAlterationOptions undocumented
	QueryAlterationOptions *SearchAlterationOptions `json:"queryAlterationOptions,omitempty"`
	// Region undocumented
	Region *string `json:"region,omitempty"`
	// ResultTemplateOptions undocumented
	ResultTemplateOptions *ResultTemplateOption `json:"resultTemplateOptions,omitempty"`
	// SharePointOneDriveOptions undocumented
	SharePointOneDriveOptions *SharePointOneDriveOptions `json:"sharePointOneDriveOptions,omitempty"`
	// Size undocumented
	Size *int `json:"size,omitempty"`
	// SortProperties undocumented
	SortProperties []SortProperty `json:"sortProperties,omitempty"`
	// StoredUnderscorefields undocumented
	StoredUnderscorefields []string `json:"stored_fields,omitempty"`
	// TrimDuplicates undocumented
	TrimDuplicates *bool `json:"trimDuplicates,omitempty"`
}

func NewSearchRequestObject() (*SearchRequestObject, error) {
	newSearchRequestObject := &SearchRequestObject{
		ODataType: "#microsoft.graph.SearchRequestObject",
	}
	return newSearchRequestObject, nil
}

// SearchResponse undocumented
type SearchResponse struct {
	// Object is the base model of SearchResponse
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// HitsContainers undocumented
	HitsContainers []SearchHitsContainer `json:"hitsContainers,omitempty"`
	// QueryAlterationResponse undocumented
	QueryAlterationResponse *AlterationResponse `json:"queryAlterationResponse,omitempty"`
	// ResultTemplates undocumented
	ResultTemplates *ResultTemplateDictionary `json:"resultTemplates,omitempty"`
	// SearchTerms undocumented
	SearchTerms []string `json:"searchTerms,omitempty"`
}

func NewSearchResponse() (*SearchResponse, error) {
	newSearchResponse := &SearchResponse{
		ODataType: "#microsoft.graph.SearchResponse",
	}
	return newSearchResponse, nil
}

// SearchResult undocumented
type SearchResult struct {
	// Object is the base model of SearchResult
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// OnClickTelemetryURL undocumented
	OnClickTelemetryURL *string `json:"onClickTelemetryUrl,omitempty"`
}

func NewSearchResult() (*SearchResult, error) {
	newSearchResult := &SearchResult{
		ODataType: "#microsoft.graph.SearchResult",
	}
	return newSearchResult, nil
}

// SearchAcronym undocumented
type SearchAcronym struct {
	// SearchSearchAnswer is the base model of SearchAcronym
	SearchSearchAnswer

	ODataType string `json:"@odata.type,omitempty"`
	// StandsFor undocumented
	StandsFor *string `json:"standsFor,omitempty"`
	// State undocumented
	State *SearchAnswerState `json:"state,omitempty"`
}

func NewSearchAcronym() (*SearchAcronym, error) {
	newSearchAcronym := &SearchAcronym{
		ODataType: "#microsoft.graph.SearchAcronym",
	}
	return newSearchAcronym, nil
}

// SearchAnswerKeyword undocumented
type SearchAnswerKeyword struct {
	// Object is the base model of SearchAnswerKeyword
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Keywords undocumented
	Keywords []string `json:"keywords,omitempty"`
	// MatchSimilarKeywords undocumented
	MatchSimilarKeywords *bool `json:"matchSimilarKeywords,omitempty"`
	// ReservedKeywords undocumented
	ReservedKeywords []string `json:"reservedKeywords,omitempty"`
}

func NewSearchAnswerKeyword() (*SearchAnswerKeyword, error) {
	newSearchAnswerKeyword := &SearchAnswerKeyword{
		ODataType: "#microsoft.graph.SearchAnswerKeyword",
	}
	return newSearchAnswerKeyword, nil
}

// SearchAnswerVariant undocumented
type SearchAnswerVariant struct {
	// Object is the base model of SearchAnswerVariant
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LanguageTag undocumented
	LanguageTag *string `json:"languageTag,omitempty"`
	// Platform undocumented
	Platform *DevicePlatformType `json:"platform,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
}

func NewSearchAnswerVariant() (*SearchAnswerVariant, error) {
	newSearchAnswerVariant := &SearchAnswerVariant{
		ODataType: "#microsoft.graph.SearchAnswerVariant",
	}
	return newSearchAnswerVariant, nil
}

// SearchBookmark undocumented
type SearchBookmark struct {
	// SearchSearchAnswer is the base model of SearchBookmark
	SearchSearchAnswer

	ODataType string `json:"@odata.type,omitempty"`
	// AvailabilityEndDateTime undocumented
	AvailabilityEndDateTime *time.Time `json:"availabilityEndDateTime,omitempty"`
	// AvailabilityStartDateTime undocumented
	AvailabilityStartDateTime *time.Time `json:"availabilityStartDateTime,omitempty"`
	// Categories undocumented
	Categories []string `json:"categories,omitempty"`
	// GroupIDs undocumented
	GroupIDs []string `json:"groupIds,omitempty"`
	// IsSuggested undocumented
	IsSuggested *bool `json:"isSuggested,omitempty"`
	// Keywords undocumented
	Keywords *SearchAnswerKeyword `json:"keywords,omitempty"`
	// LanguageTags undocumented
	LanguageTags []string `json:"languageTags,omitempty"`
	// Platforms undocumented
	Platforms []DevicePlatformType `json:"platforms,omitempty"`
	// PowerAppIDs undocumented
	PowerAppIDs []string `json:"powerAppIds,omitempty"`
	// State undocumented
	State *SearchAnswerState `json:"state,omitempty"`
	// TargetedVariations undocumented
	TargetedVariations []SearchAnswerVariant `json:"targetedVariations,omitempty"`
}

func NewSearchBookmark() (*SearchBookmark, error) {
	newSearchBookmark := &SearchBookmark{
		ODataType: "#microsoft.graph.SearchBookmark",
	}
	return newSearchBookmark, nil
}

// SearchIdentity undocumented
type SearchIdentity struct {
	// Object is the base model of SearchIdentity
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
}

func NewSearchIdentity() (*SearchIdentity, error) {
	newSearchIdentity := &SearchIdentity{
		ODataType: "#microsoft.graph.SearchIdentity",
	}
	return newSearchIdentity, nil
}

// SearchIdentitySet undocumented
type SearchIdentitySet struct {
	// Object is the base model of SearchIdentitySet
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Application undocumented
	Application *SearchIdentity `json:"application,omitempty"`
	// Device undocumented
	Device *SearchIdentity `json:"device,omitempty"`
	// User undocumented
	User *SearchIdentity `json:"user,omitempty"`
}

func NewSearchIdentitySet() (*SearchIdentitySet, error) {
	newSearchIdentitySet := &SearchIdentitySet{
		ODataType: "#microsoft.graph.SearchIdentitySet",
	}
	return newSearchIdentitySet, nil
}

// SearchQna undocumented
type SearchQna struct {
	// SearchSearchAnswer is the base model of SearchQna
	SearchSearchAnswer

	ODataType string `json:"@odata.type,omitempty"`
	// AvailabilityEndDateTime undocumented
	AvailabilityEndDateTime *time.Time `json:"availabilityEndDateTime,omitempty"`
	// AvailabilityStartDateTime undocumented
	AvailabilityStartDateTime *time.Time `json:"availabilityStartDateTime,omitempty"`
	// GroupIDs undocumented
	GroupIDs []string `json:"groupIds,omitempty"`
	// IsSuggested undocumented
	IsSuggested *bool `json:"isSuggested,omitempty"`
	// Keywords undocumented
	Keywords *SearchAnswerKeyword `json:"keywords,omitempty"`
	// LanguageTags undocumented
	LanguageTags []string `json:"languageTags,omitempty"`
	// Platforms undocumented
	Platforms []DevicePlatformType `json:"platforms,omitempty"`
	// State undocumented
	State *SearchAnswerState `json:"state,omitempty"`
	// TargetedVariations undocumented
	TargetedVariations []SearchAnswerVariant `json:"targetedVariations,omitempty"`
}

func NewSearchQna() (*SearchQna, error) {
	newSearchQna := &SearchQna{
		ODataType: "#microsoft.graph.SearchQna",
	}
	return newSearchQna, nil
}

// SearchSearchAnswer undocumented
type SearchSearchAnswer struct {
	// Entity is the base model of SearchSearchAnswer
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *SearchIdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
}

func NewSearchSearchAnswer() (*SearchSearchAnswer, error) {
	newSearchSearchAnswer := &SearchSearchAnswer{
		ODataType: "#microsoft.graph.SearchSearchAnswer",
	}
	return newSearchSearchAnswer, nil
}
