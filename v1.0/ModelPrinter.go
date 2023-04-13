// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Printer undocumented
type Printer struct {
	// PrinterBase is the base model of Printer
	PrinterBase

	ODataType string `json:"@odata.type,omitempty"`
	// HasPhysicalDevice undocumented
	HasPhysicalDevice *bool `json:"hasPhysicalDevice,omitempty"`
	// IsShared undocumented
	IsShared *bool `json:"isShared,omitempty"`
	// LastSeenDateTime undocumented
	LastSeenDateTime *time.Time `json:"lastSeenDateTime,omitempty"`
	// RegisteredDateTime undocumented
	RegisteredDateTime *time.Time `json:"registeredDateTime,omitempty"`
	// Connectors undocumented
	Connectors []PrintConnector `json:"connectors,omitempty"`
	// Shares undocumented
	Shares []PrinterShare `json:"shares,omitempty"`
	// TaskTriggers undocumented
	TaskTriggers []PrintTaskTrigger `json:"taskTriggers,omitempty"`
}

func NewPrinter() (*Printer, error) {
	newPrinter := &Printer{
		ODataType: "#microsoft.graph.Printer",
	}
	return newPrinter, nil
}

// PrinterBase undocumented
type PrinterBase struct {
	// Entity is the base model of PrinterBase
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Capabilities undocumented
	Capabilities *PrinterCapabilities `json:"capabilities,omitempty"`
	// Defaults undocumented
	Defaults *PrinterDefaults `json:"defaults,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsAcceptingJobs undocumented
	IsAcceptingJobs *bool `json:"isAcceptingJobs,omitempty"`
	// Location undocumented
	Location *PrinterLocation `json:"location,omitempty"`
	// Manufacturer undocumented
	Manufacturer *string `json:"manufacturer,omitempty"`
	// Model undocumented
	Model *string `json:"model,omitempty"`
	// Status undocumented
	Status *PrinterStatus `json:"status,omitempty"`
	// Jobs undocumented
	Jobs []PrintJob `json:"jobs,omitempty"`
}

func NewPrinterBase() (*PrinterBase, error) {
	newPrinterBase := &PrinterBase{
		ODataType: "#microsoft.graph.PrinterBase",
	}
	return newPrinterBase, nil
}

// PrinterCapabilities undocumented
type PrinterCapabilities struct {
	// Object is the base model of PrinterCapabilities
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// BottomMargins undocumented
	BottomMargins []int `json:"bottomMargins,omitempty"`
	// Collation undocumented
	Collation *bool `json:"collation,omitempty"`
	// ColorModes undocumented
	ColorModes []PrintColorMode `json:"colorModes,omitempty"`
	// ContentTypes undocumented
	ContentTypes []string `json:"contentTypes,omitempty"`
	// CopiesPerJob undocumented
	CopiesPerJob *IntegerRange `json:"copiesPerJob,omitempty"`
	// Dpis undocumented
	Dpis []int `json:"dpis,omitempty"`
	// DuplexModes undocumented
	DuplexModes []PrintDuplexMode `json:"duplexModes,omitempty"`
	// FeedOrientations undocumented
	FeedOrientations []PrinterFeedOrientation `json:"feedOrientations,omitempty"`
	// Finishings undocumented
	Finishings []PrintFinishing `json:"finishings,omitempty"`
	// InputBins undocumented
	InputBins []string `json:"inputBins,omitempty"`
	// IsColorPrintingSupported undocumented
	IsColorPrintingSupported *bool `json:"isColorPrintingSupported,omitempty"`
	// IsPageRangeSupported undocumented
	IsPageRangeSupported *bool `json:"isPageRangeSupported,omitempty"`
	// LeftMargins undocumented
	LeftMargins []int `json:"leftMargins,omitempty"`
	// MediaColors undocumented
	MediaColors []string `json:"mediaColors,omitempty"`
	// MediaSizes undocumented
	MediaSizes []string `json:"mediaSizes,omitempty"`
	// MediaTypes undocumented
	MediaTypes []string `json:"mediaTypes,omitempty"`
	// MultipageLayouts undocumented
	MultipageLayouts []PrintMultipageLayout `json:"multipageLayouts,omitempty"`
	// Orientations undocumented
	Orientations []PrintOrientation `json:"orientations,omitempty"`
	// OutputBins undocumented
	OutputBins []string `json:"outputBins,omitempty"`
	// PagesPerSheet undocumented
	PagesPerSheet []int `json:"pagesPerSheet,omitempty"`
	// Qualities undocumented
	Qualities []PrintQuality `json:"qualities,omitempty"`
	// RightMargins undocumented
	RightMargins []int `json:"rightMargins,omitempty"`
	// Scalings undocumented
	Scalings []PrintScaling `json:"scalings,omitempty"`
	// SupportsFitPdfToPage undocumented
	SupportsFitPdfToPage *bool `json:"supportsFitPdfToPage,omitempty"`
	// TopMargins undocumented
	TopMargins []int `json:"topMargins,omitempty"`
}

func NewPrinterCapabilities() (*PrinterCapabilities, error) {
	newPrinterCapabilities := &PrinterCapabilities{
		ODataType: "#microsoft.graph.PrinterCapabilities",
	}
	return newPrinterCapabilities, nil
}

// PrinterCreateOperation undocumented
type PrinterCreateOperation struct {
	// PrintOperation is the base model of PrinterCreateOperation
	PrintOperation

	ODataType string `json:"@odata.type,omitempty"`
	// Certificate undocumented
	Certificate *string `json:"certificate,omitempty"`
	// Printer undocumented
	Printer *Printer `json:"printer,omitempty"`
}

func NewPrinterCreateOperation() (*PrinterCreateOperation, error) {
	newPrinterCreateOperation := &PrinterCreateOperation{
		ODataType: "#microsoft.graph.PrinterCreateOperation",
	}
	return newPrinterCreateOperation, nil
}

// PrinterDefaults undocumented
type PrinterDefaults struct {
	// Object is the base model of PrinterDefaults
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ColorMode undocumented
	ColorMode *PrintColorMode `json:"colorMode,omitempty"`
	// ContentType undocumented
	ContentType *string `json:"contentType,omitempty"`
	// CopiesPerJob undocumented
	CopiesPerJob *int `json:"copiesPerJob,omitempty"`
	// Dpi undocumented
	Dpi *int `json:"dpi,omitempty"`
	// DuplexMode undocumented
	DuplexMode *PrintDuplexMode `json:"duplexMode,omitempty"`
	// Finishings undocumented
	Finishings []PrintFinishing `json:"finishings,omitempty"`
	// FitPdfToPage undocumented
	FitPdfToPage *bool `json:"fitPdfToPage,omitempty"`
	// InputBin undocumented
	InputBin *string `json:"inputBin,omitempty"`
	// MediaColor undocumented
	MediaColor *string `json:"mediaColor,omitempty"`
	// MediaSize undocumented
	MediaSize *string `json:"mediaSize,omitempty"`
	// MediaType undocumented
	MediaType *string `json:"mediaType,omitempty"`
	// MultipageLayout undocumented
	MultipageLayout *PrintMultipageLayout `json:"multipageLayout,omitempty"`
	// Orientation undocumented
	Orientation *PrintOrientation `json:"orientation,omitempty"`
	// OutputBin undocumented
	OutputBin *string `json:"outputBin,omitempty"`
	// PagesPerSheet undocumented
	PagesPerSheet *int `json:"pagesPerSheet,omitempty"`
	// Quality undocumented
	Quality *PrintQuality `json:"quality,omitempty"`
	// Scaling undocumented
	Scaling *PrintScaling `json:"scaling,omitempty"`
}

func NewPrinterDefaults() (*PrinterDefaults, error) {
	newPrinterDefaults := &PrinterDefaults{
		ODataType: "#microsoft.graph.PrinterDefaults",
	}
	return newPrinterDefaults, nil
}

// PrinterLocation undocumented
type PrinterLocation struct {
	// Object is the base model of PrinterLocation
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AltitudeInMeters undocumented
	AltitudeInMeters *int `json:"altitudeInMeters,omitempty"`
	// Building undocumented
	Building *string `json:"building,omitempty"`
	// City undocumented
	City *string `json:"city,omitempty"`
	// CountryOrRegion undocumented
	CountryOrRegion *string `json:"countryOrRegion,omitempty"`
	// Floor undocumented
	Floor *string `json:"floor,omitempty"`
	// FloorDescription undocumented
	FloorDescription *string `json:"floorDescription,omitempty"`
	// Latitude undocumented
	Latitude *float64 `json:"latitude,omitempty"`
	// Longitude undocumented
	Longitude *float64 `json:"longitude,omitempty"`
	// Organization undocumented
	Organization []string `json:"organization,omitempty"`
	// PostalCode undocumented
	PostalCode *string `json:"postalCode,omitempty"`
	// RoomDescription undocumented
	RoomDescription *string `json:"roomDescription,omitempty"`
	// RoomName undocumented
	RoomName *string `json:"roomName,omitempty"`
	// Site undocumented
	Site *string `json:"site,omitempty"`
	// StateOrProvince undocumented
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
	// StreetAddress undocumented
	StreetAddress *string `json:"streetAddress,omitempty"`
	// Subdivision undocumented
	Subdivision []string `json:"subdivision,omitempty"`
	// Subunit undocumented
	Subunit []string `json:"subunit,omitempty"`
}

func NewPrinterLocation() (*PrinterLocation, error) {
	newPrinterLocation := &PrinterLocation{
		ODataType: "#microsoft.graph.PrinterLocation",
	}
	return newPrinterLocation, nil
}

// PrinterShare undocumented
type PrinterShare struct {
	// PrinterBase is the base model of PrinterShare
	PrinterBase

	ODataType string `json:"@odata.type,omitempty"`
	// AllowAllUsers undocumented
	AllowAllUsers *bool `json:"allowAllUsers,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// AllowedGroups undocumented
	AllowedGroups []Group `json:"allowedGroups,omitempty"`
	// AllowedUsers undocumented
	AllowedUsers []User `json:"allowedUsers,omitempty"`
	// Printer undocumented
	Printer *Printer `json:"printer,omitempty"`
}

func NewPrinterShare() (*PrinterShare, error) {
	newPrinterShare := &PrinterShare{
		ODataType: "#microsoft.graph.PrinterShare",
	}
	return newPrinterShare, nil
}

// PrinterStatus undocumented
type PrinterStatus struct {
	// Object is the base model of PrinterStatus
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// Details undocumented
	Details []PrinterProcessingStateDetail `json:"details,omitempty"`
	// State undocumented
	State *PrinterProcessingState `json:"state,omitempty"`
}

func NewPrinterStatus() (*PrinterStatus, error) {
	newPrinterStatus := &PrinterStatus{
		ODataType: "#microsoft.graph.PrinterStatus",
	}
	return newPrinterStatus, nil
}
