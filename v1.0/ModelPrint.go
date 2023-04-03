// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Print undocumented
type Print struct {
	// Object is the base model of Print
	Object

	ODataType string `json:"@odata.type"`
	// Settings undocumented
	Settings *PrintSettings `json:"settings,omitempty"`
	// Connectors undocumented
	Connectors []PrintConnector `json:"connectors,omitempty"`
	// Operations undocumented
	Operations []PrintOperation `json:"operations,omitempty"`
	// Printers undocumented
	Printers []Printer `json:"printers,omitempty"`
	// Services undocumented
	Services []PrintService `json:"services,omitempty"`
	// Shares undocumented
	Shares []PrinterShare `json:"shares,omitempty"`
	// TaskDefinitions undocumented
	TaskDefinitions []PrintTaskDefinition `json:"taskDefinitions,omitempty"`
}

func NewPrint() (*Print, error) {
	newPrint := &Print{
		ODataType: "#microsoft.graph.Print",
	}
	return newPrint, nil
}

// PrintCertificateSigningRequestObject undocumented
type PrintCertificateSigningRequestObject struct {
	// Object is the base model of PrintCertificateSigningRequestObject
	Object

	ODataType string `json:"@odata.type"`
	// Content undocumented
	Content *string `json:"content,omitempty"`
	// TransportKey undocumented
	TransportKey *string `json:"transportKey,omitempty"`
}

func NewPrintCertificateSigningRequestObject() (*PrintCertificateSigningRequestObject, error) {
	newPrintCertificateSigningRequestObject := &PrintCertificateSigningRequestObject{
		ODataType: "#microsoft.graph.PrintCertificateSigningRequestObject",
	}
	return newPrintCertificateSigningRequestObject, nil
}

// PrintConnector undocumented
type PrintConnector struct {
	// Entity is the base model of PrintConnector
	Entity

	ODataType string `json:"@odata.type"`
	// AppVersion undocumented
	AppVersion *string `json:"appVersion,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// FullyQualifiedDomainName undocumented
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty"`
	// Location undocumented
	Location *PrinterLocation `json:"location,omitempty"`
	// OperatingSystem undocumented
	OperatingSystem *string `json:"operatingSystem,omitempty"`
	// RegisteredDateTime undocumented
	RegisteredDateTime *time.Time `json:"registeredDateTime,omitempty"`
}

func NewPrintConnector() (*PrintConnector, error) {
	newPrintConnector := &PrintConnector{
		ODataType: "#microsoft.graph.PrintConnector",
	}
	return newPrintConnector, nil
}

// PrintDocument undocumented
type PrintDocument struct {
	// Entity is the base model of PrintDocument
	Entity

	ODataType string `json:"@odata.type"`
	// ContentType undocumented
	ContentType *string `json:"contentType,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Size undocumented
	Size *int `json:"size,omitempty"`
}

func NewPrintDocument() (*PrintDocument, error) {
	newPrintDocument := &PrintDocument{
		ODataType: "#microsoft.graph.PrintDocument",
	}
	return newPrintDocument, nil
}

// PrintDocumentUploadProperties undocumented
type PrintDocumentUploadProperties struct {
	// Object is the base model of PrintDocumentUploadProperties
	Object

	ODataType string `json:"@odata.type"`
	// ContentType undocumented
	ContentType *string `json:"contentType,omitempty"`
	// DocumentName undocumented
	DocumentName *string `json:"documentName,omitempty"`
	// Size undocumented
	Size *int `json:"size,omitempty"`
}

func NewPrintDocumentUploadProperties() (*PrintDocumentUploadProperties, error) {
	newPrintDocumentUploadProperties := &PrintDocumentUploadProperties{
		ODataType: "#microsoft.graph.PrintDocumentUploadProperties",
	}
	return newPrintDocumentUploadProperties, nil
}

// PrintJob undocumented
type PrintJob struct {
	// Entity is the base model of PrintJob
	Entity

	ODataType string `json:"@odata.type"`
	// Configuration undocumented
	Configuration *PrintJobConfiguration `json:"configuration,omitempty"`
	// CreatedBy undocumented
	CreatedBy *UserIdentity `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// IsFetchable undocumented
	IsFetchable *bool `json:"isFetchable,omitempty"`
	// RedirectedFrom undocumented
	RedirectedFrom *string `json:"redirectedFrom,omitempty"`
	// RedirectedTo undocumented
	RedirectedTo *string `json:"redirectedTo,omitempty"`
	// Status undocumented
	Status *PrintJobStatus `json:"status,omitempty"`
	// Documents undocumented
	Documents []PrintDocument `json:"documents,omitempty"`
	// Tasks undocumented
	Tasks []PrintTask `json:"tasks,omitempty"`
}

func NewPrintJob() (*PrintJob, error) {
	newPrintJob := &PrintJob{
		ODataType: "#microsoft.graph.PrintJob",
	}
	return newPrintJob, nil
}

// PrintJobConfiguration undocumented
type PrintJobConfiguration struct {
	// Object is the base model of PrintJobConfiguration
	Object

	ODataType string `json:"@odata.type"`
	// Collate undocumented
	Collate *bool `json:"collate,omitempty"`
	// ColorMode undocumented
	ColorMode *PrintColorMode `json:"colorMode,omitempty"`
	// Copies undocumented
	Copies *int `json:"copies,omitempty"`
	// Dpi undocumented
	Dpi *int `json:"dpi,omitempty"`
	// DuplexMode undocumented
	DuplexMode *PrintDuplexMode `json:"duplexMode,omitempty"`
	// FeedOrientation undocumented
	FeedOrientation *PrinterFeedOrientation `json:"feedOrientation,omitempty"`
	// Finishings undocumented
	Finishings []PrintFinishing `json:"finishings,omitempty"`
	// FitPdfToPage undocumented
	FitPdfToPage *bool `json:"fitPdfToPage,omitempty"`
	// InputBin undocumented
	InputBin *string `json:"inputBin,omitempty"`
	// Margin undocumented
	Margin *PrintMargin `json:"margin,omitempty"`
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
	// PageRanges undocumented
	PageRanges []IntegerRange `json:"pageRanges,omitempty"`
	// PagesPerSheet undocumented
	PagesPerSheet *int `json:"pagesPerSheet,omitempty"`
	// Quality undocumented
	Quality *PrintQuality `json:"quality,omitempty"`
	// Scaling undocumented
	Scaling *PrintScaling `json:"scaling,omitempty"`
}

func NewPrintJobConfiguration() (*PrintJobConfiguration, error) {
	newPrintJobConfiguration := &PrintJobConfiguration{
		ODataType: "#microsoft.graph.PrintJobConfiguration",
	}
	return newPrintJobConfiguration, nil
}

// PrintJobStatus undocumented
type PrintJobStatus struct {
	// Object is the base model of PrintJobStatus
	Object

	ODataType string `json:"@odata.type"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// Details undocumented
	Details []PrintJobStateDetail `json:"details,omitempty"`
	// IsAcquiredByPrinter undocumented
	IsAcquiredByPrinter *bool `json:"isAcquiredByPrinter,omitempty"`
	// State undocumented
	State *PrintJobProcessingState `json:"state,omitempty"`
}

func NewPrintJobStatus() (*PrintJobStatus, error) {
	newPrintJobStatus := &PrintJobStatus{
		ODataType: "#microsoft.graph.PrintJobStatus",
	}
	return newPrintJobStatus, nil
}

// PrintMargin undocumented
type PrintMargin struct {
	// Object is the base model of PrintMargin
	Object

	ODataType string `json:"@odata.type"`
	// Bottom undocumented
	Bottom *int `json:"bottom,omitempty"`
	// Left undocumented
	Left *int `json:"left,omitempty"`
	// Right undocumented
	Right *int `json:"right,omitempty"`
	// Top undocumented
	Top *int `json:"top,omitempty"`
}

func NewPrintMargin() (*PrintMargin, error) {
	newPrintMargin := &PrintMargin{
		ODataType: "#microsoft.graph.PrintMargin",
	}
	return newPrintMargin, nil
}

// PrintOperation undocumented
type PrintOperation struct {
	// Entity is the base model of PrintOperation
	Entity

	ODataType string `json:"@odata.type"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Status undocumented
	Status *PrintOperationStatus `json:"status,omitempty"`
}

func NewPrintOperation() (*PrintOperation, error) {
	newPrintOperation := &PrintOperation{
		ODataType: "#microsoft.graph.PrintOperation",
	}
	return newPrintOperation, nil
}

// PrintOperationStatus undocumented
type PrintOperationStatus struct {
	// Object is the base model of PrintOperationStatus
	Object

	ODataType string `json:"@odata.type"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// State undocumented
	State *PrintOperationProcessingState `json:"state,omitempty"`
}

func NewPrintOperationStatus() (*PrintOperationStatus, error) {
	newPrintOperationStatus := &PrintOperationStatus{
		ODataType: "#microsoft.graph.PrintOperationStatus",
	}
	return newPrintOperationStatus, nil
}

// PrintService undocumented
type PrintService struct {
	// Entity is the base model of PrintService
	Entity

	ODataType string `json:"@odata.type"`
	// Endpoints undocumented
	Endpoints []PrintServiceEndpoint `json:"endpoints,omitempty"`
}

func NewPrintService() (*PrintService, error) {
	newPrintService := &PrintService{
		ODataType: "#microsoft.graph.PrintService",
	}
	return newPrintService, nil
}

// PrintServiceEndpoint undocumented
type PrintServiceEndpoint struct {
	// Entity is the base model of PrintServiceEndpoint
	Entity

	ODataType string `json:"@odata.type"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// URI undocumented
	URI *string `json:"uri,omitempty"`
}

func NewPrintServiceEndpoint() (*PrintServiceEndpoint, error) {
	newPrintServiceEndpoint := &PrintServiceEndpoint{
		ODataType: "#microsoft.graph.PrintServiceEndpoint",
	}
	return newPrintServiceEndpoint, nil
}

// PrintSettings undocumented
type PrintSettings struct {
	// Object is the base model of PrintSettings
	Object

	ODataType string `json:"@odata.type"`
	// DocumentConversionEnabled undocumented
	DocumentConversionEnabled *bool `json:"documentConversionEnabled,omitempty"`
}

func NewPrintSettings() (*PrintSettings, error) {
	newPrintSettings := &PrintSettings{
		ODataType: "#microsoft.graph.PrintSettings",
	}
	return newPrintSettings, nil
}

// PrintTask undocumented
type PrintTask struct {
	// Entity is the base model of PrintTask
	Entity

	ODataType string `json:"@odata.type"`
	// ParentURL undocumented
	ParentURL *string `json:"parentUrl,omitempty"`
	// Status undocumented
	Status *PrintTaskStatus `json:"status,omitempty"`
	// Definition undocumented
	Definition *PrintTaskDefinition `json:"definition,omitempty"`
	// Trigger undocumented
	Trigger *PrintTaskTrigger `json:"trigger,omitempty"`
}

func NewPrintTask() (*PrintTask, error) {
	newPrintTask := &PrintTask{
		ODataType: "#microsoft.graph.PrintTask",
	}
	return newPrintTask, nil
}

// PrintTaskDefinition undocumented
type PrintTaskDefinition struct {
	// Entity is the base model of PrintTaskDefinition
	Entity

	ODataType string `json:"@odata.type"`
	// CreatedBy undocumented
	CreatedBy *AppIdentity `json:"createdBy,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Tasks undocumented
	Tasks []PrintTask `json:"tasks,omitempty"`
}

func NewPrintTaskDefinition() (*PrintTaskDefinition, error) {
	newPrintTaskDefinition := &PrintTaskDefinition{
		ODataType: "#microsoft.graph.PrintTaskDefinition",
	}
	return newPrintTaskDefinition, nil
}

// PrintTaskStatus undocumented
type PrintTaskStatus struct {
	// Object is the base model of PrintTaskStatus
	Object

	ODataType string `json:"@odata.type"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// State undocumented
	State *PrintTaskProcessingState `json:"state,omitempty"`
}

func NewPrintTaskStatus() (*PrintTaskStatus, error) {
	newPrintTaskStatus := &PrintTaskStatus{
		ODataType: "#microsoft.graph.PrintTaskStatus",
	}
	return newPrintTaskStatus, nil
}

// PrintTaskTrigger undocumented
type PrintTaskTrigger struct {
	// Entity is the base model of PrintTaskTrigger
	Entity

	ODataType string `json:"@odata.type"`
	// Event undocumented
	Event *PrintEvent `json:"event,omitempty"`
	// Definition undocumented
	Definition *PrintTaskDefinition `json:"definition,omitempty"`
}

func NewPrintTaskTrigger() (*PrintTaskTrigger, error) {
	newPrintTaskTrigger := &PrintTaskTrigger{
		ODataType: "#microsoft.graph.PrintTaskTrigger",
	}
	return newPrintTaskTrigger, nil
}

// PrintUsage undocumented
type PrintUsage struct {
	// Entity is the base model of PrintUsage
	Entity

	ODataType string `json:"@odata.type"`
	// CompletedBlackAndWhiteJobCount undocumented
	CompletedBlackAndWhiteJobCount *int `json:"completedBlackAndWhiteJobCount,omitempty"`
	// CompletedColorJobCount undocumented
	CompletedColorJobCount *int `json:"completedColorJobCount,omitempty"`
	// IncompleteJobCount undocumented
	IncompleteJobCount *int `json:"incompleteJobCount,omitempty"`
	// UsageDate undocumented
	UsageDate *Date `json:"usageDate,omitempty"`
}

func NewPrintUsage() (*PrintUsage, error) {
	newPrintUsage := &PrintUsage{
		ODataType: "#microsoft.graph.PrintUsage",
	}
	return newPrintUsage, nil
}

// PrintUsageByPrinter undocumented
type PrintUsageByPrinter struct {
	// PrintUsage is the base model of PrintUsageByPrinter
	PrintUsage

	ODataType string `json:"@odata.type"`
	// PrinterID undocumented
	PrinterID *string `json:"printerId,omitempty"`
}

func NewPrintUsageByPrinter() (*PrintUsageByPrinter, error) {
	newPrintUsageByPrinter := &PrintUsageByPrinter{
		ODataType: "#microsoft.graph.PrintUsageByPrinter",
	}
	return newPrintUsageByPrinter, nil
}

// PrintUsageByUser undocumented
type PrintUsageByUser struct {
	// PrintUsage is the base model of PrintUsageByUser
	PrintUsage

	ODataType string `json:"@odata.type"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

func NewPrintUsageByUser() (*PrintUsageByUser, error) {
	newPrintUsageByUser := &PrintUsageByUser{
		ODataType: "#microsoft.graph.PrintUsageByUser",
	}
	return newPrintUsageByUser, nil
}
