// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// ItemActionStat undocumented
type ItemActionStat struct {
	// Object is the base model of ItemActionStat
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ActionCount undocumented
	ActionCount *int `json:"actionCount,omitempty"`
	// ActorCount undocumented
	ActorCount *int `json:"actorCount,omitempty"`
}

func NewItemActionStat() (*ItemActionStat, error) {
	newItemActionStat := &ItemActionStat{
		ODataType: "#microsoft.graph.ItemActionStat",
	}
	return newItemActionStat, nil
}

// ItemActivity undocumented
type ItemActivity struct {
	// Entity is the base model of ItemActivity
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Access undocumented
	Access *AccessAction `json:"access,omitempty"`
	// ActivityDateTime undocumented
	ActivityDateTime *time.Time `json:"activityDateTime,omitempty"`
	// Actor undocumented
	Actor *IdentitySet `json:"actor,omitempty"`
	// DriveItem undocumented
	DriveItem *DriveItem `json:"driveItem,omitempty"`
}

func NewItemActivity() (*ItemActivity, error) {
	newItemActivity := &ItemActivity{
		ODataType: "#microsoft.graph.ItemActivity",
	}
	return newItemActivity, nil
}

// ItemActivityStat undocumented
type ItemActivityStat struct {
	// Entity is the base model of ItemActivityStat
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Access undocumented
	Access *ItemActionStat `json:"access,omitempty"`
	// Create undocumented
	Create *ItemActionStat `json:"create,omitempty"`
	// Delete undocumented
	Delete *ItemActionStat `json:"delete,omitempty"`
	// Edit undocumented
	Edit *ItemActionStat `json:"edit,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// IncompleteData undocumented
	IncompleteData *IncompleteData `json:"incompleteData,omitempty"`
	// IsTrending undocumented
	IsTrending *bool `json:"isTrending,omitempty"`
	// Move undocumented
	Move *ItemActionStat `json:"move,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// Activities undocumented
	Activities []ItemActivity `json:"activities,omitempty"`
}

func NewItemActivityStat() (*ItemActivityStat, error) {
	newItemActivityStat := &ItemActivityStat{
		ODataType: "#microsoft.graph.ItemActivityStat",
	}
	return newItemActivityStat, nil
}

// ItemAnalytics undocumented
type ItemAnalytics struct {
	// Entity is the base model of ItemAnalytics
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AllTime undocumented
	AllTime *ItemActivityStat `json:"allTime,omitempty"`
	// ItemActivityStats undocumented
	ItemActivityStats []ItemActivityStat `json:"itemActivityStats,omitempty"`
	// LastSevenDays undocumented
	LastSevenDays *ItemActivityStat `json:"lastSevenDays,omitempty"`
}

func NewItemAnalytics() (*ItemAnalytics, error) {
	newItemAnalytics := &ItemAnalytics{
		ODataType: "#microsoft.graph.ItemAnalytics",
	}
	return newItemAnalytics, nil
}

// ItemAttachment undocumented
type ItemAttachment struct {
	// Attachment is the base model of ItemAttachment
	Attachment

	ODataType string `json:"@odata.type,omitempty"`
	// Item undocumented
	Item *OutlookItem `json:"item,omitempty"`
}

func NewItemAttachment() (*ItemAttachment, error) {
	newItemAttachment := &ItemAttachment{
		ODataType: "#microsoft.graph.ItemAttachment",
	}
	return newItemAttachment, nil
}

// ItemBody undocumented
type ItemBody struct {
	// Object is the base model of ItemBody
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Content undocumented
	Content *string `json:"content,omitempty"`
	// ContentType undocumented
	ContentType *BodyType `json:"contentType,omitempty"`
}

func NewItemBody() (*ItemBody, error) {
	newItemBody := &ItemBody{
		ODataType: "#microsoft.graph.ItemBody",
	}
	return newItemBody, nil
}

// ItemPreviewInfo undocumented
type ItemPreviewInfo struct {
	// Object is the base model of ItemPreviewInfo
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// GetURL undocumented
	GetURL *string `json:"getUrl,omitempty"`
	// PostParameters undocumented
	PostParameters *string `json:"postParameters,omitempty"`
	// PostURL undocumented
	PostURL *string `json:"postUrl,omitempty"`
}

func NewItemPreviewInfo() (*ItemPreviewInfo, error) {
	newItemPreviewInfo := &ItemPreviewInfo{
		ODataType: "#microsoft.graph.ItemPreviewInfo",
	}
	return newItemPreviewInfo, nil
}

// ItemReference undocumented
type ItemReference struct {
	// Object is the base model of ItemReference
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// DriveID undocumented
	DriveID *string `json:"driveId,omitempty"`
	// DriveType undocumented
	DriveType *string `json:"driveType,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Path undocumented
	Path *string `json:"path,omitempty"`
	// ShareID undocumented
	ShareID *string `json:"shareId,omitempty"`
	// SharepointIDs undocumented
	SharepointIDs *SharepointIDs `json:"sharepointIds,omitempty"`
	// SiteID undocumented
	SiteID *string `json:"siteId,omitempty"`
}

func NewItemReference() (*ItemReference, error) {
	newItemReference := &ItemReference{
		ODataType: "#microsoft.graph.ItemReference",
	}
	return newItemReference, nil
}
