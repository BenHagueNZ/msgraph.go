// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// EdiscoveryAddToReviewSetOperation undocumented
type EdiscoveryAddToReviewSetOperation struct {
	// EdiscoveryCaseOperation is the base model of EdiscoveryAddToReviewSetOperation
	EdiscoveryCaseOperation

	ODataType string `json:"@odata.type,omitempty"`
	// ReviewSet undocumented
	ReviewSet *EdiscoveryReviewSet `json:"reviewSet,omitempty"`
	// SourceCollection undocumented
	SourceCollection *EdiscoverySourceCollection `json:"sourceCollection,omitempty"`
}

func NewEdiscoveryAddToReviewSetOperation() (*EdiscoveryAddToReviewSetOperation, error) {
	newEdiscoveryAddToReviewSetOperation := &EdiscoveryAddToReviewSetOperation{
		ODataType: "#microsoft.graph.EdiscoveryAddToReviewSetOperation",
	}
	return newEdiscoveryAddToReviewSetOperation, nil
}

// EdiscoveryCase undocumented
type EdiscoveryCase struct {
	// Entity is the base model of EdiscoveryCase
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ClosedBy undocumented
	ClosedBy *IdentitySet `json:"closedBy,omitempty"`
	// ClosedDateTime undocumented
	ClosedDateTime *time.Time `json:"closedDateTime,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Status undocumented
	Status *EdiscoveryCaseStatus `json:"status,omitempty"`
	// Custodians undocumented
	Custodians []EdiscoveryCustodian `json:"custodians,omitempty"`
	// LegalHolds undocumented
	LegalHolds []EdiscoveryLegalHold `json:"legalHolds,omitempty"`
	// NoncustodialDataSources undocumented
	NoncustodialDataSources []EdiscoveryNoncustodialDataSource `json:"noncustodialDataSources,omitempty"`
	// Operations undocumented
	Operations []EdiscoveryCaseOperation `json:"operations,omitempty"`
	// ReviewSets undocumented
	ReviewSets []EdiscoveryReviewSet `json:"reviewSets,omitempty"`
	// Settings undocumented
	Settings *EdiscoveryCaseSettings `json:"settings,omitempty"`
	// SourceCollections undocumented
	SourceCollections []EdiscoverySourceCollection `json:"sourceCollections,omitempty"`
	// Tags undocumented
	Tags []EdiscoveryTag `json:"tags,omitempty"`
}

func NewEdiscoveryCase() (*EdiscoveryCase, error) {
	newEdiscoveryCase := &EdiscoveryCase{
		ODataType: "#microsoft.graph.EdiscoveryCase",
	}
	return newEdiscoveryCase, nil
}

// EdiscoveryCaseExportOperation undocumented
type EdiscoveryCaseExportOperation struct {
	// EdiscoveryCaseOperation is the base model of EdiscoveryCaseExportOperation
	EdiscoveryCaseOperation

	ODataType string `json:"@odata.type,omitempty"`
	// AzureBlobContainer undocumented
	AzureBlobContainer *string `json:"azureBlobContainer,omitempty"`
	// AzureBlobToken undocumented
	AzureBlobToken *string `json:"azureBlobToken,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// ExportOptions undocumented
	ExportOptions *EdiscoveryExportOptions `json:"exportOptions,omitempty"`
	// ExportStructure undocumented
	ExportStructure *EdiscoveryExportFileStructure `json:"exportStructure,omitempty"`
	// OutputFolderID undocumented
	OutputFolderID *string `json:"outputFolderId,omitempty"`
	// OutputName undocumented
	OutputName *string `json:"outputName,omitempty"`
	// ReviewSet undocumented
	ReviewSet *EdiscoveryReviewSet `json:"reviewSet,omitempty"`
}

func NewEdiscoveryCaseExportOperation() (*EdiscoveryCaseExportOperation, error) {
	newEdiscoveryCaseExportOperation := &EdiscoveryCaseExportOperation{
		ODataType: "#microsoft.graph.EdiscoveryCaseExportOperation",
	}
	return newEdiscoveryCaseExportOperation, nil
}

// EdiscoveryCaseHoldOperation undocumented
type EdiscoveryCaseHoldOperation struct {
	// EdiscoveryCaseOperation is the base model of EdiscoveryCaseHoldOperation
	EdiscoveryCaseOperation

	ODataType string `json:"@odata.type,omitempty"`
}

func NewEdiscoveryCaseHoldOperation() (*EdiscoveryCaseHoldOperation, error) {
	newEdiscoveryCaseHoldOperation := &EdiscoveryCaseHoldOperation{
		ODataType: "#microsoft.graph.EdiscoveryCaseHoldOperation",
	}
	return newEdiscoveryCaseHoldOperation, nil
}

// EdiscoveryCaseIndexOperation undocumented
type EdiscoveryCaseIndexOperation struct {
	// EdiscoveryCaseOperation is the base model of EdiscoveryCaseIndexOperation
	EdiscoveryCaseOperation

	ODataType string `json:"@odata.type,omitempty"`
}

func NewEdiscoveryCaseIndexOperation() (*EdiscoveryCaseIndexOperation, error) {
	newEdiscoveryCaseIndexOperation := &EdiscoveryCaseIndexOperation{
		ODataType: "#microsoft.graph.EdiscoveryCaseIndexOperation",
	}
	return newEdiscoveryCaseIndexOperation, nil
}

// EdiscoveryCaseOperation undocumented
type EdiscoveryCaseOperation struct {
	// Entity is the base model of EdiscoveryCaseOperation
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Action undocumented
	Action *EdiscoveryCaseAction `json:"action,omitempty"`
	// CompletedDateTime undocumented
	CompletedDateTime *time.Time `json:"completedDateTime,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// PercentProgress undocumented
	PercentProgress *int `json:"percentProgress,omitempty"`
	// ResultInfo undocumented
	ResultInfo *ResultInfo `json:"resultInfo,omitempty"`
	// Status undocumented
	Status *EdiscoveryCaseOperationStatus `json:"status,omitempty"`
}

func NewEdiscoveryCaseOperation() (*EdiscoveryCaseOperation, error) {
	newEdiscoveryCaseOperation := &EdiscoveryCaseOperation{
		ODataType: "#microsoft.graph.EdiscoveryCaseOperation",
	}
	return newEdiscoveryCaseOperation, nil
}

// EdiscoveryCaseSettings undocumented
type EdiscoveryCaseSettings struct {
	// Entity is the base model of EdiscoveryCaseSettings
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Ocr undocumented
	Ocr *EdiscoveryOcrSettings `json:"ocr,omitempty"`
	// RedundancyDetection undocumented
	RedundancyDetection *EdiscoveryRedundancyDetectionSettings `json:"redundancyDetection,omitempty"`
	// TopicModeling undocumented
	TopicModeling *EdiscoveryTopicModelingSettings `json:"topicModeling,omitempty"`
}

func NewEdiscoveryCaseSettings() (*EdiscoveryCaseSettings, error) {
	newEdiscoveryCaseSettings := &EdiscoveryCaseSettings{
		ODataType: "#microsoft.graph.EdiscoveryCaseSettings",
	}
	return newEdiscoveryCaseSettings, nil
}

// EdiscoveryCustodian undocumented
type EdiscoveryCustodian struct {
	// EdiscoveryDataSourceContainer is the base model of EdiscoveryCustodian
	EdiscoveryDataSourceContainer

	ODataType string `json:"@odata.type,omitempty"`
	// AcknowledgedDateTime undocumented
	AcknowledgedDateTime *time.Time `json:"acknowledgedDateTime,omitempty"`
	// ApplyHoldToSources undocumented
	ApplyHoldToSources *bool `json:"applyHoldToSources,omitempty"`
	// Email undocumented
	Email *string `json:"email,omitempty"`
	// SiteSources undocumented
	SiteSources []EdiscoverySiteSource `json:"siteSources,omitempty"`
	// UnifiedGroupSources undocumented
	UnifiedGroupSources []EdiscoveryUnifiedGroupSource `json:"unifiedGroupSources,omitempty"`
	// UserSources undocumented
	UserSources []EdiscoveryUserSource `json:"userSources,omitempty"`
}

func NewEdiscoveryCustodian() (*EdiscoveryCustodian, error) {
	newEdiscoveryCustodian := &EdiscoveryCustodian{
		ODataType: "#microsoft.graph.EdiscoveryCustodian",
	}
	return newEdiscoveryCustodian, nil
}

// EdiscoveryDataSource undocumented
type EdiscoveryDataSource struct {
	// Entity is the base model of EdiscoveryDataSource
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// HoldStatus undocumented
	HoldStatus *EdiscoveryDataSourceHoldStatus `json:"holdStatus,omitempty"`
}

func NewEdiscoveryDataSource() (*EdiscoveryDataSource, error) {
	newEdiscoveryDataSource := &EdiscoveryDataSource{
		ODataType: "#microsoft.graph.EdiscoveryDataSource",
	}
	return newEdiscoveryDataSource, nil
}

// EdiscoveryDataSourceContainer undocumented
type EdiscoveryDataSourceContainer struct {
	// Entity is the base model of EdiscoveryDataSourceContainer
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// HoldStatus undocumented
	HoldStatus *EdiscoveryDataSourceHoldStatus `json:"holdStatus,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// ReleasedDateTime undocumented
	ReleasedDateTime *time.Time `json:"releasedDateTime,omitempty"`
	// Status undocumented
	Status *EdiscoveryDataSourceContainerStatus `json:"status,omitempty"`
	// LastIndexOperation undocumented
	LastIndexOperation *EdiscoveryCaseIndexOperation `json:"lastIndexOperation,omitempty"`
}

func NewEdiscoveryDataSourceContainer() (*EdiscoveryDataSourceContainer, error) {
	newEdiscoveryDataSourceContainer := &EdiscoveryDataSourceContainer{
		ODataType: "#microsoft.graph.EdiscoveryDataSourceContainer",
	}
	return newEdiscoveryDataSourceContainer, nil
}

// EdiscoveryEdiscoveryroot undocumented
type EdiscoveryEdiscoveryroot struct {
	// Entity is the base model of EdiscoveryEdiscoveryroot
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Cases undocumented
	Cases []EdiscoveryCase `json:"cases,omitempty"`
}

func NewEdiscoveryEdiscoveryroot() (*EdiscoveryEdiscoveryroot, error) {
	newEdiscoveryEdiscoveryroot := &EdiscoveryEdiscoveryroot{
		ODataType: "#microsoft.graph.EdiscoveryEdiscoveryroot",
	}
	return newEdiscoveryEdiscoveryroot, nil
}

// EdiscoveryEstimateStatisticsOperation undocumented
type EdiscoveryEstimateStatisticsOperation struct {
	// EdiscoveryCaseOperation is the base model of EdiscoveryEstimateStatisticsOperation
	EdiscoveryCaseOperation

	ODataType string `json:"@odata.type,omitempty"`
	// IndexedItemCount undocumented
	IndexedItemCount *int `json:"indexedItemCount,omitempty"`
	// IndexedItemsSize undocumented
	IndexedItemsSize *int `json:"indexedItemsSize,omitempty"`
	// MailboxCount undocumented
	MailboxCount *int `json:"mailboxCount,omitempty"`
	// SiteCount undocumented
	SiteCount *int `json:"siteCount,omitempty"`
	// UnindexedItemCount undocumented
	UnindexedItemCount *int `json:"unindexedItemCount,omitempty"`
	// UnindexedItemsSize undocumented
	UnindexedItemsSize *int `json:"unindexedItemsSize,omitempty"`
	// SourceCollection undocumented
	SourceCollection *EdiscoverySourceCollection `json:"sourceCollection,omitempty"`
}

func NewEdiscoveryEstimateStatisticsOperation() (*EdiscoveryEstimateStatisticsOperation, error) {
	newEdiscoveryEstimateStatisticsOperation := &EdiscoveryEstimateStatisticsOperation{
		ODataType: "#microsoft.graph.EdiscoveryEstimateStatisticsOperation",
	}
	return newEdiscoveryEstimateStatisticsOperation, nil
}

// EdiscoveryLegalHold undocumented
type EdiscoveryLegalHold struct {
	// Entity is the base model of EdiscoveryLegalHold
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ContentQuery undocumented
	ContentQuery *string `json:"contentQuery,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Errors undocumented
	Errors []string `json:"errors,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Status undocumented
	Status *EdiscoveryLegalHoldStatus `json:"status,omitempty"`
	// SiteSources undocumented
	SiteSources []EdiscoverySiteSource `json:"siteSources,omitempty"`
	// UnifiedGroupSources undocumented
	UnifiedGroupSources []EdiscoveryUnifiedGroupSource `json:"unifiedGroupSources,omitempty"`
	// UserSources undocumented
	UserSources []EdiscoveryUserSource `json:"userSources,omitempty"`
}

func NewEdiscoveryLegalHold() (*EdiscoveryLegalHold, error) {
	newEdiscoveryLegalHold := &EdiscoveryLegalHold{
		ODataType: "#microsoft.graph.EdiscoveryLegalHold",
	}
	return newEdiscoveryLegalHold, nil
}

// EdiscoveryNoncustodialDataSource undocumented
type EdiscoveryNoncustodialDataSource struct {
	// EdiscoveryDataSourceContainer is the base model of EdiscoveryNoncustodialDataSource
	EdiscoveryDataSourceContainer

	ODataType string `json:"@odata.type,omitempty"`
	// ApplyHoldToSource undocumented
	ApplyHoldToSource *bool `json:"applyHoldToSource,omitempty"`
	// DataSource undocumented
	DataSource *EdiscoveryDataSource `json:"dataSource,omitempty"`
}

func NewEdiscoveryNoncustodialDataSource() (*EdiscoveryNoncustodialDataSource, error) {
	newEdiscoveryNoncustodialDataSource := &EdiscoveryNoncustodialDataSource{
		ODataType: "#microsoft.graph.EdiscoveryNoncustodialDataSource",
	}
	return newEdiscoveryNoncustodialDataSource, nil
}

// EdiscoveryOcrSettings undocumented
type EdiscoveryOcrSettings struct {
	// Object is the base model of EdiscoveryOcrSettings
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// MaxImageSize undocumented
	MaxImageSize *int `json:"maxImageSize,omitempty"`
	// Timeout undocumented
	Timeout *Duration `json:"timeout,omitempty"`
}

func NewEdiscoveryOcrSettings() (*EdiscoveryOcrSettings, error) {
	newEdiscoveryOcrSettings := &EdiscoveryOcrSettings{
		ODataType: "#microsoft.graph.EdiscoveryOcrSettings",
	}
	return newEdiscoveryOcrSettings, nil
}

// EdiscoveryPurgeDataOperation undocumented
type EdiscoveryPurgeDataOperation struct {
	// EdiscoveryCaseOperation is the base model of EdiscoveryPurgeDataOperation
	EdiscoveryCaseOperation

	ODataType string `json:"@odata.type,omitempty"`
}

func NewEdiscoveryPurgeDataOperation() (*EdiscoveryPurgeDataOperation, error) {
	newEdiscoveryPurgeDataOperation := &EdiscoveryPurgeDataOperation{
		ODataType: "#microsoft.graph.EdiscoveryPurgeDataOperation",
	}
	return newEdiscoveryPurgeDataOperation, nil
}

// EdiscoveryRedundancyDetectionSettings undocumented
type EdiscoveryRedundancyDetectionSettings struct {
	// Object is the base model of EdiscoveryRedundancyDetectionSettings
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// MaxWords undocumented
	MaxWords *int `json:"maxWords,omitempty"`
	// MinWords undocumented
	MinWords *int `json:"minWords,omitempty"`
	// SimilarityThreshold undocumented
	SimilarityThreshold *int `json:"similarityThreshold,omitempty"`
}

func NewEdiscoveryRedundancyDetectionSettings() (*EdiscoveryRedundancyDetectionSettings, error) {
	newEdiscoveryRedundancyDetectionSettings := &EdiscoveryRedundancyDetectionSettings{
		ODataType: "#microsoft.graph.EdiscoveryRedundancyDetectionSettings",
	}
	return newEdiscoveryRedundancyDetectionSettings, nil
}

// EdiscoveryReviewSet undocumented
type EdiscoveryReviewSet struct {
	// Entity is the base model of EdiscoveryReviewSet
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Queries undocumented
	Queries []EdiscoveryReviewSetQuery `json:"queries,omitempty"`
}

func NewEdiscoveryReviewSet() (*EdiscoveryReviewSet, error) {
	newEdiscoveryReviewSet := &EdiscoveryReviewSet{
		ODataType: "#microsoft.graph.EdiscoveryReviewSet",
	}
	return newEdiscoveryReviewSet, nil
}

// EdiscoveryReviewSetQuery undocumented
type EdiscoveryReviewSetQuery struct {
	// Entity is the base model of EdiscoveryReviewSetQuery
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Query undocumented
	Query *string `json:"query,omitempty"`
}

func NewEdiscoveryReviewSetQuery() (*EdiscoveryReviewSetQuery, error) {
	newEdiscoveryReviewSetQuery := &EdiscoveryReviewSetQuery{
		ODataType: "#microsoft.graph.EdiscoveryReviewSetQuery",
	}
	return newEdiscoveryReviewSetQuery, nil
}

// EdiscoverySiteSource undocumented
type EdiscoverySiteSource struct {
	// EdiscoveryDataSource is the base model of EdiscoverySiteSource
	EdiscoveryDataSource

	ODataType string `json:"@odata.type,omitempty"`
	// Site undocumented
	Site *Site `json:"site,omitempty"`
}

func NewEdiscoverySiteSource() (*EdiscoverySiteSource, error) {
	newEdiscoverySiteSource := &EdiscoverySiteSource{
		ODataType: "#microsoft.graph.EdiscoverySiteSource",
	}
	return newEdiscoverySiteSource, nil
}

// EdiscoverySourceCollection undocumented
type EdiscoverySourceCollection struct {
	// Entity is the base model of EdiscoverySourceCollection
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ContentQuery undocumented
	ContentQuery *string `json:"contentQuery,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DataSourceScopes undocumented
	DataSourceScopes *EdiscoveryDataSourceScopes `json:"dataSourceScopes,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// AdditionalSources undocumented
	AdditionalSources []EdiscoveryDataSource `json:"additionalSources,omitempty"`
	// AddToReviewSetOperation undocumented
	AddToReviewSetOperation *EdiscoveryAddToReviewSetOperation `json:"addToReviewSetOperation,omitempty"`
	// CustodianSources undocumented
	CustodianSources []EdiscoveryDataSource `json:"custodianSources,omitempty"`
	// LastEstimateStatisticsOperation undocumented
	LastEstimateStatisticsOperation *EdiscoveryEstimateStatisticsOperation `json:"lastEstimateStatisticsOperation,omitempty"`
	// NoncustodialSources undocumented
	NoncustodialSources []EdiscoveryNoncustodialDataSource `json:"noncustodialSources,omitempty"`
}

func NewEdiscoverySourceCollection() (*EdiscoverySourceCollection, error) {
	newEdiscoverySourceCollection := &EdiscoverySourceCollection{
		ODataType: "#microsoft.graph.EdiscoverySourceCollection",
	}
	return newEdiscoverySourceCollection, nil
}

// EdiscoveryTag undocumented
type EdiscoveryTag struct {
	// Entity is the base model of EdiscoveryTag
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ChildSelectability undocumented
	ChildSelectability *EdiscoveryChildSelectability `json:"childSelectability,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// ChildTags undocumented
	ChildTags []EdiscoveryTag `json:"childTags,omitempty"`
	// Parent undocumented
	Parent *EdiscoveryTag `json:"parent,omitempty"`
}

func NewEdiscoveryTag() (*EdiscoveryTag, error) {
	newEdiscoveryTag := &EdiscoveryTag{
		ODataType: "#microsoft.graph.EdiscoveryTag",
	}
	return newEdiscoveryTag, nil
}

// EdiscoveryTagOperation undocumented
type EdiscoveryTagOperation struct {
	// EdiscoveryCaseOperation is the base model of EdiscoveryTagOperation
	EdiscoveryCaseOperation

	ODataType string `json:"@odata.type,omitempty"`
}

func NewEdiscoveryTagOperation() (*EdiscoveryTagOperation, error) {
	newEdiscoveryTagOperation := &EdiscoveryTagOperation{
		ODataType: "#microsoft.graph.EdiscoveryTagOperation",
	}
	return newEdiscoveryTagOperation, nil
}

// EdiscoveryTopicModelingSettings undocumented
type EdiscoveryTopicModelingSettings struct {
	// Object is the base model of EdiscoveryTopicModelingSettings
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// DynamicallyAdjustTopicCount undocumented
	DynamicallyAdjustTopicCount *bool `json:"dynamicallyAdjustTopicCount,omitempty"`
	// IgnoreNumbers undocumented
	IgnoreNumbers *bool `json:"ignoreNumbers,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// TopicCount undocumented
	TopicCount *int `json:"topicCount,omitempty"`
}

func NewEdiscoveryTopicModelingSettings() (*EdiscoveryTopicModelingSettings, error) {
	newEdiscoveryTopicModelingSettings := &EdiscoveryTopicModelingSettings{
		ODataType: "#microsoft.graph.EdiscoveryTopicModelingSettings",
	}
	return newEdiscoveryTopicModelingSettings, nil
}

// EdiscoveryUnifiedGroupSource undocumented
type EdiscoveryUnifiedGroupSource struct {
	// EdiscoveryDataSource is the base model of EdiscoveryUnifiedGroupSource
	EdiscoveryDataSource

	ODataType string `json:"@odata.type,omitempty"`
	// IncludedSources undocumented
	IncludedSources *EdiscoverySourceType `json:"includedSources,omitempty"`
	// Group undocumented
	Group *Group `json:"group,omitempty"`
}

func NewEdiscoveryUnifiedGroupSource() (*EdiscoveryUnifiedGroupSource, error) {
	newEdiscoveryUnifiedGroupSource := &EdiscoveryUnifiedGroupSource{
		ODataType: "#microsoft.graph.EdiscoveryUnifiedGroupSource",
	}
	return newEdiscoveryUnifiedGroupSource, nil
}

// EdiscoveryUserSource undocumented
type EdiscoveryUserSource struct {
	// EdiscoveryDataSource is the base model of EdiscoveryUserSource
	EdiscoveryDataSource

	ODataType string `json:"@odata.type,omitempty"`
	// Email undocumented
	Email *string `json:"email,omitempty"`
	// IncludedSources undocumented
	IncludedSources *EdiscoverySourceType `json:"includedSources,omitempty"`
	// SiteWebURL undocumented
	SiteWebURL *string `json:"siteWebUrl,omitempty"`
}

func NewEdiscoveryUserSource() (*EdiscoveryUserSource, error) {
	newEdiscoveryUserSource := &EdiscoveryUserSource{
		ODataType: "#microsoft.graph.EdiscoveryUserSource",
	}
	return newEdiscoveryUserSource, nil
}
