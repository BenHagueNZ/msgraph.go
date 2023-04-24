// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Site undocumented
type Site struct {
	// BaseItem is the base model of Site
	BaseItem

	ODataType string `json:"@odata.type,omitempty"`
	// Deleted undocumented
	Deleted *Deleted `json:"deleted,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Root undocumented
	Root *Root `json:"root,omitempty"`
	// Settings undocumented
	Settings *SiteSettings `json:"settings,omitempty"`
	// SharepointIDs undocumented
	SharepointIDs *SharepointIDs `json:"sharepointIds,omitempty"`
	// SiteCollection undocumented
	SiteCollection *SiteCollection `json:"siteCollection,omitempty"`
	// InformationProtection undocumented
	InformationProtection *InformationProtection `json:"informationProtection,omitempty"`
	// Analytics undocumented
	Analytics *ItemAnalytics `json:"analytics,omitempty"`
	// Columns undocumented
	Columns []ColumnDefinition `json:"columns,omitempty"`
	// ContentTypes undocumented
	ContentTypes []ContentType `json:"contentTypes,omitempty"`
	// Drive undocumented
	Drive *Drive `json:"drive,omitempty"`
	// Drives undocumented
	Drives []Drive `json:"drives,omitempty"`
	// ExternalColumns undocumented
	ExternalColumns []ColumnDefinition `json:"externalColumns,omitempty"`
	// Items undocumented
	Items []BaseItem `json:"items,omitempty"`
	// Lists undocumented
	Lists []List `json:"lists,omitempty"`
	// Operations undocumented
	Operations []RichLongRunningOperation `json:"operations,omitempty"`
	// Pages undocumented
	Pages []SitePage `json:"pages,omitempty"`
	// Permissions undocumented
	Permissions []Permission `json:"permissions,omitempty"`
	// Sites undocumented
	Sites []Site `json:"sites,omitempty"`
	// TermStore undocumented
	TermStore *TermStoreStore `json:"termStore,omitempty"`
	// Onenote undocumented
	Onenote *Onenote `json:"onenote,omitempty"`
}

func NewSite() (*Site, error) {
	newSite := &Site{
		ODataType: "#microsoft.graph.Site",
	}
	return newSite, nil
}

// SiteCollection undocumented
type SiteCollection struct {
	// Object is the base model of SiteCollection
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// DataLocationCode undocumented
	DataLocationCode *string `json:"dataLocationCode,omitempty"`
	// Hostname undocumented
	Hostname *string `json:"hostname,omitempty"`
	// Root undocumented
	Root *Root `json:"root,omitempty"`
}

func NewSiteCollection() (*SiteCollection, error) {
	newSiteCollection := &SiteCollection{
		ODataType: "#microsoft.graph.SiteCollection",
	}
	return newSiteCollection, nil
}

// SitePage undocumented
type SitePage struct {
	// BaseItem is the base model of SitePage
	BaseItem

	ODataType string `json:"@odata.type,omitempty"`
	// ContentType undocumented
	ContentType *ContentTypeInfo `json:"contentType,omitempty"`
	// PageLayout undocumented
	PageLayout *PageLayoutType `json:"pageLayout,omitempty"`
	// PromotionKind undocumented
	PromotionKind *PagePromotionType `json:"promotionKind,omitempty"`
	// PublishingState undocumented
	PublishingState *PublicationFacet `json:"publishingState,omitempty"`
	// Reactions undocumented
	Reactions *ReactionsFacet `json:"reactions,omitempty"`
	// ShowComments undocumented
	ShowComments *bool `json:"showComments,omitempty"`
	// ShowRecommendedPages undocumented
	ShowRecommendedPages *bool `json:"showRecommendedPages,omitempty"`
	// ThumbnailWebURL undocumented
	ThumbnailWebURL *string `json:"thumbnailWebUrl,omitempty"`
	// Title undocumented
	Title *string `json:"title,omitempty"`
	// TitleArea undocumented
	TitleArea *TitleArea `json:"titleArea,omitempty"`
	// CanvasLayout undocumented
	CanvasLayout *CanvasLayout `json:"canvasLayout,omitempty"`
	// WebParts undocumented
	WebParts []WebPart `json:"webParts,omitempty"`
}

func NewSitePage() (*SitePage, error) {
	newSitePage := &SitePage{
		ODataType: "#microsoft.graph.SitePage",
	}
	return newSitePage, nil
}

// SiteSettings undocumented
type SiteSettings struct {
	// Object is the base model of SiteSettings
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// LanguageTag undocumented
	LanguageTag *string `json:"languageTag,omitempty"`
	// TimeZone undocumented
	TimeZone *string `json:"timeZone,omitempty"`
}

func NewSiteSettings() (*SiteSettings, error) {
	newSiteSettings := &SiteSettings{
		ODataType: "#microsoft.graph.SiteSettings",
	}
	return newSiteSettings, nil
}