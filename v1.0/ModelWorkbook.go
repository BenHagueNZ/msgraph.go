// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "encoding/json"

// Workbook undocumented
type Workbook struct {
	// Entity is the base model of Workbook
	Entity

	ODataType string `json:"@odata.type"`
	// Application undocumented
	Application *WorkbookApplication `json:"application,omitempty"`
	// Comments undocumented
	Comments []WorkbookComment `json:"comments,omitempty"`
	// Functions undocumented
	Functions *WorkbookFunctions `json:"functions,omitempty"`
	// Names undocumented
	Names []WorkbookNamedItem `json:"names,omitempty"`
	// Operations undocumented
	Operations []WorkbookOperation `json:"operations,omitempty"`
	// Tables undocumented
	Tables []WorkbookTable `json:"tables,omitempty"`
	// Worksheets undocumented
	Worksheets []WorkbookWorksheet `json:"worksheets,omitempty"`
}

func NewWorkbook() (*Workbook, error) {
	newWorkbook := &Workbook{
		ODataType: "#microsoft.graph.Workbook",
	}
	return newWorkbook, nil
}

// WorkbookApplication undocumented
type WorkbookApplication struct {
	// Entity is the base model of WorkbookApplication
	Entity

	ODataType string `json:"@odata.type"`
	// CalculationMode undocumented
	CalculationMode *string `json:"calculationMode,omitempty"`
}

func NewWorkbookApplication() (*WorkbookApplication, error) {
	newWorkbookApplication := &WorkbookApplication{
		ODataType: "#microsoft.graph.WorkbookApplication",
	}
	return newWorkbookApplication, nil
}

// WorkbookChart undocumented
type WorkbookChart struct {
	// Entity is the base model of WorkbookChart
	Entity

	ODataType string `json:"@odata.type"`
	// Height undocumented
	Height *float64 `json:"height,omitempty"`
	// Left undocumented
	Left *float64 `json:"left,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Top undocumented
	Top *float64 `json:"top,omitempty"`
	// Width undocumented
	Width *float64 `json:"width,omitempty"`
	// Axes undocumented
	Axes *WorkbookChartAxes `json:"axes,omitempty"`
	// DataLabels undocumented
	DataLabels *WorkbookChartDataLabels `json:"dataLabels,omitempty"`
	// Format undocumented
	Format *WorkbookChartAreaFormat `json:"format,omitempty"`
	// Legend undocumented
	Legend *WorkbookChartLegend `json:"legend,omitempty"`
	// Series undocumented
	Series []WorkbookChartSeries `json:"series,omitempty"`
	// Title undocumented
	Title *WorkbookChartTitle `json:"title,omitempty"`
	// Worksheet undocumented
	Worksheet *WorkbookWorksheet `json:"worksheet,omitempty"`
}

func NewWorkbookChart() (*WorkbookChart, error) {
	newWorkbookChart := &WorkbookChart{
		ODataType: "#microsoft.graph.WorkbookChart",
	}
	return newWorkbookChart, nil
}

// WorkbookChartAreaFormat undocumented
type WorkbookChartAreaFormat struct {
	// Entity is the base model of WorkbookChartAreaFormat
	Entity

	ODataType string `json:"@odata.type"`
	// Fill undocumented
	Fill *WorkbookChartFill `json:"fill,omitempty"`
	// Font undocumented
	Font *WorkbookChartFont `json:"font,omitempty"`
}

func NewWorkbookChartAreaFormat() (*WorkbookChartAreaFormat, error) {
	newWorkbookChartAreaFormat := &WorkbookChartAreaFormat{
		ODataType: "#microsoft.graph.WorkbookChartAreaFormat",
	}
	return newWorkbookChartAreaFormat, nil
}

// WorkbookChartAxes undocumented
type WorkbookChartAxes struct {
	// Entity is the base model of WorkbookChartAxes
	Entity

	ODataType string `json:"@odata.type"`
	// CategoryAxis undocumented
	CategoryAxis *WorkbookChartAxis `json:"categoryAxis,omitempty"`
	// SeriesAxis undocumented
	SeriesAxis *WorkbookChartAxis `json:"seriesAxis,omitempty"`
	// ValueAxis undocumented
	ValueAxis *WorkbookChartAxis `json:"valueAxis,omitempty"`
}

func NewWorkbookChartAxes() (*WorkbookChartAxes, error) {
	newWorkbookChartAxes := &WorkbookChartAxes{
		ODataType: "#microsoft.graph.WorkbookChartAxes",
	}
	return newWorkbookChartAxes, nil
}

// WorkbookChartAxis undocumented
type WorkbookChartAxis struct {
	// Entity is the base model of WorkbookChartAxis
	Entity

	ODataType string `json:"@odata.type"`
	// MajorUnit undocumented
	MajorUnit json.RawMessage `json:"majorUnit,omitempty"`
	// Maximum undocumented
	Maximum json.RawMessage `json:"maximum,omitempty"`
	// Minimum undocumented
	Minimum json.RawMessage `json:"minimum,omitempty"`
	// MinorUnit undocumented
	MinorUnit json.RawMessage `json:"minorUnit,omitempty"`
	// Format undocumented
	Format *WorkbookChartAxisFormat `json:"format,omitempty"`
	// MajorGridlines undocumented
	MajorGridlines *WorkbookChartGridlines `json:"majorGridlines,omitempty"`
	// MinorGridlines undocumented
	MinorGridlines *WorkbookChartGridlines `json:"minorGridlines,omitempty"`
	// Title undocumented
	Title *WorkbookChartAxisTitle `json:"title,omitempty"`
}

func NewWorkbookChartAxis() (*WorkbookChartAxis, error) {
	newWorkbookChartAxis := &WorkbookChartAxis{
		ODataType: "#microsoft.graph.WorkbookChartAxis",
	}
	return newWorkbookChartAxis, nil
}

// WorkbookChartAxisFormat undocumented
type WorkbookChartAxisFormat struct {
	// Entity is the base model of WorkbookChartAxisFormat
	Entity

	ODataType string `json:"@odata.type"`
	// Font undocumented
	Font *WorkbookChartFont `json:"font,omitempty"`
	// Line undocumented
	Line *WorkbookChartLineFormat `json:"line,omitempty"`
}

func NewWorkbookChartAxisFormat() (*WorkbookChartAxisFormat, error) {
	newWorkbookChartAxisFormat := &WorkbookChartAxisFormat{
		ODataType: "#microsoft.graph.WorkbookChartAxisFormat",
	}
	return newWorkbookChartAxisFormat, nil
}

// WorkbookChartAxisTitle undocumented
type WorkbookChartAxisTitle struct {
	// Entity is the base model of WorkbookChartAxisTitle
	Entity

	ODataType string `json:"@odata.type"`
	// Text undocumented
	Text *string `json:"text,omitempty"`
	// Visible undocumented
	Visible *bool `json:"visible,omitempty"`
	// Format undocumented
	Format *WorkbookChartAxisTitleFormat `json:"format,omitempty"`
}

func NewWorkbookChartAxisTitle() (*WorkbookChartAxisTitle, error) {
	newWorkbookChartAxisTitle := &WorkbookChartAxisTitle{
		ODataType: "#microsoft.graph.WorkbookChartAxisTitle",
	}
	return newWorkbookChartAxisTitle, nil
}

// WorkbookChartAxisTitleFormat undocumented
type WorkbookChartAxisTitleFormat struct {
	// Entity is the base model of WorkbookChartAxisTitleFormat
	Entity

	ODataType string `json:"@odata.type"`
	// Font undocumented
	Font *WorkbookChartFont `json:"font,omitempty"`
}

func NewWorkbookChartAxisTitleFormat() (*WorkbookChartAxisTitleFormat, error) {
	newWorkbookChartAxisTitleFormat := &WorkbookChartAxisTitleFormat{
		ODataType: "#microsoft.graph.WorkbookChartAxisTitleFormat",
	}
	return newWorkbookChartAxisTitleFormat, nil
}

// WorkbookChartDataLabelFormat undocumented
type WorkbookChartDataLabelFormat struct {
	// Entity is the base model of WorkbookChartDataLabelFormat
	Entity

	ODataType string `json:"@odata.type"`
	// Fill undocumented
	Fill *WorkbookChartFill `json:"fill,omitempty"`
	// Font undocumented
	Font *WorkbookChartFont `json:"font,omitempty"`
}

func NewWorkbookChartDataLabelFormat() (*WorkbookChartDataLabelFormat, error) {
	newWorkbookChartDataLabelFormat := &WorkbookChartDataLabelFormat{
		ODataType: "#microsoft.graph.WorkbookChartDataLabelFormat",
	}
	return newWorkbookChartDataLabelFormat, nil
}

// WorkbookChartDataLabels undocumented
type WorkbookChartDataLabels struct {
	// Entity is the base model of WorkbookChartDataLabels
	Entity

	ODataType string `json:"@odata.type"`
	// Position undocumented
	Position *string `json:"position,omitempty"`
	// Separator undocumented
	Separator *string `json:"separator,omitempty"`
	// ShowBubbleSize undocumented
	ShowBubbleSize *bool `json:"showBubbleSize,omitempty"`
	// ShowCategoryName undocumented
	ShowCategoryName *bool `json:"showCategoryName,omitempty"`
	// ShowLegendKey undocumented
	ShowLegendKey *bool `json:"showLegendKey,omitempty"`
	// ShowPercentage undocumented
	ShowPercentage *bool `json:"showPercentage,omitempty"`
	// ShowSeriesName undocumented
	ShowSeriesName *bool `json:"showSeriesName,omitempty"`
	// ShowValue undocumented
	ShowValue *bool `json:"showValue,omitempty"`
	// Format undocumented
	Format *WorkbookChartDataLabelFormat `json:"format,omitempty"`
}

func NewWorkbookChartDataLabels() (*WorkbookChartDataLabels, error) {
	newWorkbookChartDataLabels := &WorkbookChartDataLabels{
		ODataType: "#microsoft.graph.WorkbookChartDataLabels",
	}
	return newWorkbookChartDataLabels, nil
}

// WorkbookChartFill undocumented
type WorkbookChartFill struct {
	// Entity is the base model of WorkbookChartFill
	Entity

	ODataType string `json:"@odata.type"`
}

func NewWorkbookChartFill() (*WorkbookChartFill, error) {
	newWorkbookChartFill := &WorkbookChartFill{
		ODataType: "#microsoft.graph.WorkbookChartFill",
	}
	return newWorkbookChartFill, nil
}

// WorkbookChartFont undocumented
type WorkbookChartFont struct {
	// Entity is the base model of WorkbookChartFont
	Entity

	ODataType string `json:"@odata.type"`
	// Bold undocumented
	Bold *bool `json:"bold,omitempty"`
	// Color undocumented
	Color *string `json:"color,omitempty"`
	// Italic undocumented
	Italic *bool `json:"italic,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Size undocumented
	Size *float64 `json:"size,omitempty"`
	// Underline undocumented
	Underline *string `json:"underline,omitempty"`
}

func NewWorkbookChartFont() (*WorkbookChartFont, error) {
	newWorkbookChartFont := &WorkbookChartFont{
		ODataType: "#microsoft.graph.WorkbookChartFont",
	}
	return newWorkbookChartFont, nil
}

// WorkbookChartGridlines undocumented
type WorkbookChartGridlines struct {
	// Entity is the base model of WorkbookChartGridlines
	Entity

	ODataType string `json:"@odata.type"`
	// Visible undocumented
	Visible *bool `json:"visible,omitempty"`
	// Format undocumented
	Format *WorkbookChartGridlinesFormat `json:"format,omitempty"`
}

func NewWorkbookChartGridlines() (*WorkbookChartGridlines, error) {
	newWorkbookChartGridlines := &WorkbookChartGridlines{
		ODataType: "#microsoft.graph.WorkbookChartGridlines",
	}
	return newWorkbookChartGridlines, nil
}

// WorkbookChartGridlinesFormat undocumented
type WorkbookChartGridlinesFormat struct {
	// Entity is the base model of WorkbookChartGridlinesFormat
	Entity

	ODataType string `json:"@odata.type"`
	// Line undocumented
	Line *WorkbookChartLineFormat `json:"line,omitempty"`
}

func NewWorkbookChartGridlinesFormat() (*WorkbookChartGridlinesFormat, error) {
	newWorkbookChartGridlinesFormat := &WorkbookChartGridlinesFormat{
		ODataType: "#microsoft.graph.WorkbookChartGridlinesFormat",
	}
	return newWorkbookChartGridlinesFormat, nil
}

// WorkbookChartLegend undocumented
type WorkbookChartLegend struct {
	// Entity is the base model of WorkbookChartLegend
	Entity

	ODataType string `json:"@odata.type"`
	// Overlay undocumented
	Overlay *bool `json:"overlay,omitempty"`
	// Position undocumented
	Position *string `json:"position,omitempty"`
	// Visible undocumented
	Visible *bool `json:"visible,omitempty"`
	// Format undocumented
	Format *WorkbookChartLegendFormat `json:"format,omitempty"`
}

func NewWorkbookChartLegend() (*WorkbookChartLegend, error) {
	newWorkbookChartLegend := &WorkbookChartLegend{
		ODataType: "#microsoft.graph.WorkbookChartLegend",
	}
	return newWorkbookChartLegend, nil
}

// WorkbookChartLegendFormat undocumented
type WorkbookChartLegendFormat struct {
	// Entity is the base model of WorkbookChartLegendFormat
	Entity

	ODataType string `json:"@odata.type"`
	// Fill undocumented
	Fill *WorkbookChartFill `json:"fill,omitempty"`
	// Font undocumented
	Font *WorkbookChartFont `json:"font,omitempty"`
}

func NewWorkbookChartLegendFormat() (*WorkbookChartLegendFormat, error) {
	newWorkbookChartLegendFormat := &WorkbookChartLegendFormat{
		ODataType: "#microsoft.graph.WorkbookChartLegendFormat",
	}
	return newWorkbookChartLegendFormat, nil
}

// WorkbookChartLineFormat undocumented
type WorkbookChartLineFormat struct {
	// Entity is the base model of WorkbookChartLineFormat
	Entity

	ODataType string `json:"@odata.type"`
	// Color undocumented
	Color *string `json:"color,omitempty"`
}

func NewWorkbookChartLineFormat() (*WorkbookChartLineFormat, error) {
	newWorkbookChartLineFormat := &WorkbookChartLineFormat{
		ODataType: "#microsoft.graph.WorkbookChartLineFormat",
	}
	return newWorkbookChartLineFormat, nil
}

// WorkbookChartPoint undocumented
type WorkbookChartPoint struct {
	// Entity is the base model of WorkbookChartPoint
	Entity

	ODataType string `json:"@odata.type"`
	// Value undocumented
	Value json.RawMessage `json:"value,omitempty"`
	// Format undocumented
	Format *WorkbookChartPointFormat `json:"format,omitempty"`
}

func NewWorkbookChartPoint() (*WorkbookChartPoint, error) {
	newWorkbookChartPoint := &WorkbookChartPoint{
		ODataType: "#microsoft.graph.WorkbookChartPoint",
	}
	return newWorkbookChartPoint, nil
}

// WorkbookChartPointFormat undocumented
type WorkbookChartPointFormat struct {
	// Entity is the base model of WorkbookChartPointFormat
	Entity

	ODataType string `json:"@odata.type"`
	// Fill undocumented
	Fill *WorkbookChartFill `json:"fill,omitempty"`
}

func NewWorkbookChartPointFormat() (*WorkbookChartPointFormat, error) {
	newWorkbookChartPointFormat := &WorkbookChartPointFormat{
		ODataType: "#microsoft.graph.WorkbookChartPointFormat",
	}
	return newWorkbookChartPointFormat, nil
}

// WorkbookChartSeries undocumented
type WorkbookChartSeries struct {
	// Entity is the base model of WorkbookChartSeries
	Entity

	ODataType string `json:"@odata.type"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Format undocumented
	Format *WorkbookChartSeriesFormat `json:"format,omitempty"`
	// Points undocumented
	Points []WorkbookChartPoint `json:"points,omitempty"`
}

func NewWorkbookChartSeries() (*WorkbookChartSeries, error) {
	newWorkbookChartSeries := &WorkbookChartSeries{
		ODataType: "#microsoft.graph.WorkbookChartSeries",
	}
	return newWorkbookChartSeries, nil
}

// WorkbookChartSeriesFormat undocumented
type WorkbookChartSeriesFormat struct {
	// Entity is the base model of WorkbookChartSeriesFormat
	Entity

	ODataType string `json:"@odata.type"`
	// Fill undocumented
	Fill *WorkbookChartFill `json:"fill,omitempty"`
	// Line undocumented
	Line *WorkbookChartLineFormat `json:"line,omitempty"`
}

func NewWorkbookChartSeriesFormat() (*WorkbookChartSeriesFormat, error) {
	newWorkbookChartSeriesFormat := &WorkbookChartSeriesFormat{
		ODataType: "#microsoft.graph.WorkbookChartSeriesFormat",
	}
	return newWorkbookChartSeriesFormat, nil
}

// WorkbookChartTitle undocumented
type WorkbookChartTitle struct {
	// Entity is the base model of WorkbookChartTitle
	Entity

	ODataType string `json:"@odata.type"`
	// Overlay undocumented
	Overlay *bool `json:"overlay,omitempty"`
	// Text undocumented
	Text *string `json:"text,omitempty"`
	// Visible undocumented
	Visible *bool `json:"visible,omitempty"`
	// Format undocumented
	Format *WorkbookChartTitleFormat `json:"format,omitempty"`
}

func NewWorkbookChartTitle() (*WorkbookChartTitle, error) {
	newWorkbookChartTitle := &WorkbookChartTitle{
		ODataType: "#microsoft.graph.WorkbookChartTitle",
	}
	return newWorkbookChartTitle, nil
}

// WorkbookChartTitleFormat undocumented
type WorkbookChartTitleFormat struct {
	// Entity is the base model of WorkbookChartTitleFormat
	Entity

	ODataType string `json:"@odata.type"`
	// Fill undocumented
	Fill *WorkbookChartFill `json:"fill,omitempty"`
	// Font undocumented
	Font *WorkbookChartFont `json:"font,omitempty"`
}

func NewWorkbookChartTitleFormat() (*WorkbookChartTitleFormat, error) {
	newWorkbookChartTitleFormat := &WorkbookChartTitleFormat{
		ODataType: "#microsoft.graph.WorkbookChartTitleFormat",
	}
	return newWorkbookChartTitleFormat, nil
}

// WorkbookComment undocumented
type WorkbookComment struct {
	// Entity is the base model of WorkbookComment
	Entity

	ODataType string `json:"@odata.type"`
	// Content undocumented
	Content *string `json:"content,omitempty"`
	// ContentType undocumented
	ContentType *string `json:"contentType,omitempty"`
	// Replies undocumented
	Replies []WorkbookCommentReply `json:"replies,omitempty"`
}

func NewWorkbookComment() (*WorkbookComment, error) {
	newWorkbookComment := &WorkbookComment{
		ODataType: "#microsoft.graph.WorkbookComment",
	}
	return newWorkbookComment, nil
}

// WorkbookCommentReply undocumented
type WorkbookCommentReply struct {
	// Entity is the base model of WorkbookCommentReply
	Entity

	ODataType string `json:"@odata.type"`
	// Content undocumented
	Content *string `json:"content,omitempty"`
	// ContentType undocumented
	ContentType *string `json:"contentType,omitempty"`
}

func NewWorkbookCommentReply() (*WorkbookCommentReply, error) {
	newWorkbookCommentReply := &WorkbookCommentReply{
		ODataType: "#microsoft.graph.WorkbookCommentReply",
	}
	return newWorkbookCommentReply, nil
}

// WorkbookFilter undocumented
type WorkbookFilter struct {
	// Entity is the base model of WorkbookFilter
	Entity

	ODataType string `json:"@odata.type"`
	// Criteria undocumented
	Criteria *WorkbookFilterCriteria `json:"criteria,omitempty"`
}

func NewWorkbookFilter() (*WorkbookFilter, error) {
	newWorkbookFilter := &WorkbookFilter{
		ODataType: "#microsoft.graph.WorkbookFilter",
	}
	return newWorkbookFilter, nil
}

// WorkbookFilterCriteria undocumented
type WorkbookFilterCriteria struct {
	// Object is the base model of WorkbookFilterCriteria
	Object

	ODataType string `json:"@odata.type"`
	// Color undocumented
	Color *string `json:"color,omitempty"`
	// Criterion1 undocumented
	Criterion1 *string `json:"criterion1,omitempty"`
	// Criterion2 undocumented
	Criterion2 *string `json:"criterion2,omitempty"`
	// DynamicCriteria undocumented
	DynamicCriteria *string `json:"dynamicCriteria,omitempty"`
	// FilterOn undocumented
	FilterOn *string `json:"filterOn,omitempty"`
	// Icon undocumented
	Icon *WorkbookIcon `json:"icon,omitempty"`
	// Operator undocumented
	Operator *string `json:"operator,omitempty"`
	// Values undocumented
	Values json.RawMessage `json:"values,omitempty"`
}

func NewWorkbookFilterCriteria() (*WorkbookFilterCriteria, error) {
	newWorkbookFilterCriteria := &WorkbookFilterCriteria{
		ODataType: "#microsoft.graph.WorkbookFilterCriteria",
	}
	return newWorkbookFilterCriteria, nil
}

// WorkbookFilterDatetime undocumented
type WorkbookFilterDatetime struct {
	// Object is the base model of WorkbookFilterDatetime
	Object

	ODataType string `json:"@odata.type"`
	// Date undocumented
	Date *string `json:"date,omitempty"`
	// Specificity undocumented
	Specificity *string `json:"specificity,omitempty"`
}

func NewWorkbookFilterDatetime() (*WorkbookFilterDatetime, error) {
	newWorkbookFilterDatetime := &WorkbookFilterDatetime{
		ODataType: "#microsoft.graph.WorkbookFilterDatetime",
	}
	return newWorkbookFilterDatetime, nil
}

// WorkbookFormatProtection undocumented
type WorkbookFormatProtection struct {
	// Entity is the base model of WorkbookFormatProtection
	Entity

	ODataType string `json:"@odata.type"`
	// FormulaHidden undocumented
	FormulaHidden *bool `json:"formulaHidden,omitempty"`
	// Locked undocumented
	Locked *bool `json:"locked,omitempty"`
}

func NewWorkbookFormatProtection() (*WorkbookFormatProtection, error) {
	newWorkbookFormatProtection := &WorkbookFormatProtection{
		ODataType: "#microsoft.graph.WorkbookFormatProtection",
	}
	return newWorkbookFormatProtection, nil
}

// WorkbookFunctionResult undocumented
type WorkbookFunctionResult struct {
	// Entity is the base model of WorkbookFunctionResult
	Entity

	ODataType string `json:"@odata.type"`
	// Error undocumented
	Error *string `json:"error,omitempty"`
	// Value undocumented
	Value json.RawMessage `json:"value,omitempty"`
}

func NewWorkbookFunctionResult() (*WorkbookFunctionResult, error) {
	newWorkbookFunctionResult := &WorkbookFunctionResult{
		ODataType: "#microsoft.graph.WorkbookFunctionResult",
	}
	return newWorkbookFunctionResult, nil
}

// WorkbookFunctions undocumented
type WorkbookFunctions struct {
	// Entity is the base model of WorkbookFunctions
	Entity

	ODataType string `json:"@odata.type"`
}

func NewWorkbookFunctions() (*WorkbookFunctions, error) {
	newWorkbookFunctions := &WorkbookFunctions{
		ODataType: "#microsoft.graph.WorkbookFunctions",
	}
	return newWorkbookFunctions, nil
}

// WorkbookIcon undocumented
type WorkbookIcon struct {
	// Object is the base model of WorkbookIcon
	Object

	ODataType string `json:"@odata.type"`
	// Index undocumented
	Index *int `json:"index,omitempty"`
	// Set undocumented
	Set *string `json:"set,omitempty"`
}

func NewWorkbookIcon() (*WorkbookIcon, error) {
	newWorkbookIcon := &WorkbookIcon{
		ODataType: "#microsoft.graph.WorkbookIcon",
	}
	return newWorkbookIcon, nil
}

// WorkbookNamedItem undocumented
type WorkbookNamedItem struct {
	// Entity is the base model of WorkbookNamedItem
	Entity

	ODataType string `json:"@odata.type"`
	// Comment undocumented
	Comment *string `json:"comment,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Scope undocumented
	Scope *string `json:"scope,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// Value undocumented
	Value json.RawMessage `json:"value,omitempty"`
	// Visible undocumented
	Visible *bool `json:"visible,omitempty"`
	// Worksheet undocumented
	Worksheet *WorkbookWorksheet `json:"worksheet,omitempty"`
}

func NewWorkbookNamedItem() (*WorkbookNamedItem, error) {
	newWorkbookNamedItem := &WorkbookNamedItem{
		ODataType: "#microsoft.graph.WorkbookNamedItem",
	}
	return newWorkbookNamedItem, nil
}

// WorkbookOperation undocumented
type WorkbookOperation struct {
	// Entity is the base model of WorkbookOperation
	Entity

	ODataType string `json:"@odata.type"`
	// Error undocumented
	Error *WorkbookOperationError `json:"error,omitempty"`
	// ResourceLocation undocumented
	ResourceLocation *string `json:"resourceLocation,omitempty"`
	// Status undocumented
	Status *WorkbookOperationStatus `json:"status,omitempty"`
}

func NewWorkbookOperation() (*WorkbookOperation, error) {
	newWorkbookOperation := &WorkbookOperation{
		ODataType: "#microsoft.graph.WorkbookOperation",
	}
	return newWorkbookOperation, nil
}

// WorkbookOperationError undocumented
type WorkbookOperationError struct {
	// Object is the base model of WorkbookOperationError
	Object

	ODataType string `json:"@odata.type"`
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// InnerError undocumented
	InnerError *WorkbookOperationError `json:"innerError,omitempty"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
}

func NewWorkbookOperationError() (*WorkbookOperationError, error) {
	newWorkbookOperationError := &WorkbookOperationError{
		ODataType: "#microsoft.graph.WorkbookOperationError",
	}
	return newWorkbookOperationError, nil
}

// WorkbookPivotTable undocumented
type WorkbookPivotTable struct {
	// Entity is the base model of WorkbookPivotTable
	Entity

	ODataType string `json:"@odata.type"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Worksheet undocumented
	Worksheet *WorkbookWorksheet `json:"worksheet,omitempty"`
}

func NewWorkbookPivotTable() (*WorkbookPivotTable, error) {
	newWorkbookPivotTable := &WorkbookPivotTable{
		ODataType: "#microsoft.graph.WorkbookPivotTable",
	}
	return newWorkbookPivotTable, nil
}

// WorkbookRange undocumented
type WorkbookRange struct {
	// Entity is the base model of WorkbookRange
	Entity

	ODataType string `json:"@odata.type"`
	// Address undocumented
	Address *string `json:"address,omitempty"`
	// AddressLocal undocumented
	AddressLocal *string `json:"addressLocal,omitempty"`
	// CellCount undocumented
	CellCount *int `json:"cellCount,omitempty"`
	// ColumnCount undocumented
	ColumnCount *int `json:"columnCount,omitempty"`
	// ColumnHidden undocumented
	ColumnHidden *bool `json:"columnHidden,omitempty"`
	// ColumnIndex undocumented
	ColumnIndex *int `json:"columnIndex,omitempty"`
	// Formulas undocumented
	Formulas json.RawMessage `json:"formulas,omitempty"`
	// FormulasLocal undocumented
	FormulasLocal json.RawMessage `json:"formulasLocal,omitempty"`
	// FormulasR1C1 undocumented
	FormulasR1C1 json.RawMessage `json:"formulasR1C1,omitempty"`
	// Hidden undocumented
	Hidden *bool `json:"hidden,omitempty"`
	// NumberFormat undocumented
	NumberFormat json.RawMessage `json:"numberFormat,omitempty"`
	// RowCount undocumented
	RowCount *int `json:"rowCount,omitempty"`
	// RowHidden undocumented
	RowHidden *bool `json:"rowHidden,omitempty"`
	// RowIndex undocumented
	RowIndex *int `json:"rowIndex,omitempty"`
	// Text undocumented
	Text json.RawMessage `json:"text,omitempty"`
	// Values undocumented
	Values json.RawMessage `json:"values,omitempty"`
	// ValueTypes undocumented
	ValueTypes json.RawMessage `json:"valueTypes,omitempty"`
	// Format undocumented
	Format *WorkbookRangeFormat `json:"format,omitempty"`
	// Sort undocumented
	Sort *WorkbookRangeSort `json:"sort,omitempty"`
	// Worksheet undocumented
	Worksheet *WorkbookWorksheet `json:"worksheet,omitempty"`
}

func NewWorkbookRange() (*WorkbookRange, error) {
	newWorkbookRange := &WorkbookRange{
		ODataType: "#microsoft.graph.WorkbookRange",
	}
	return newWorkbookRange, nil
}

// WorkbookRangeBorder undocumented
type WorkbookRangeBorder struct {
	// Entity is the base model of WorkbookRangeBorder
	Entity

	ODataType string `json:"@odata.type"`
	// Color undocumented
	Color *string `json:"color,omitempty"`
	// SideIndex undocumented
	SideIndex *string `json:"sideIndex,omitempty"`
	// Style undocumented
	Style *string `json:"style,omitempty"`
	// Weight undocumented
	Weight *string `json:"weight,omitempty"`
}

func NewWorkbookRangeBorder() (*WorkbookRangeBorder, error) {
	newWorkbookRangeBorder := &WorkbookRangeBorder{
		ODataType: "#microsoft.graph.WorkbookRangeBorder",
	}
	return newWorkbookRangeBorder, nil
}

// WorkbookRangeFill undocumented
type WorkbookRangeFill struct {
	// Entity is the base model of WorkbookRangeFill
	Entity

	ODataType string `json:"@odata.type"`
	// Color undocumented
	Color *string `json:"color,omitempty"`
}

func NewWorkbookRangeFill() (*WorkbookRangeFill, error) {
	newWorkbookRangeFill := &WorkbookRangeFill{
		ODataType: "#microsoft.graph.WorkbookRangeFill",
	}
	return newWorkbookRangeFill, nil
}

// WorkbookRangeFont undocumented
type WorkbookRangeFont struct {
	// Entity is the base model of WorkbookRangeFont
	Entity

	ODataType string `json:"@odata.type"`
	// Bold undocumented
	Bold *bool `json:"bold,omitempty"`
	// Color undocumented
	Color *string `json:"color,omitempty"`
	// Italic undocumented
	Italic *bool `json:"italic,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Size undocumented
	Size *float64 `json:"size,omitempty"`
	// Underline undocumented
	Underline *string `json:"underline,omitempty"`
}

func NewWorkbookRangeFont() (*WorkbookRangeFont, error) {
	newWorkbookRangeFont := &WorkbookRangeFont{
		ODataType: "#microsoft.graph.WorkbookRangeFont",
	}
	return newWorkbookRangeFont, nil
}

// WorkbookRangeFormat undocumented
type WorkbookRangeFormat struct {
	// Entity is the base model of WorkbookRangeFormat
	Entity

	ODataType string `json:"@odata.type"`
	// ColumnWidth undocumented
	ColumnWidth *float64 `json:"columnWidth,omitempty"`
	// HorizontalAlignment undocumented
	HorizontalAlignment *string `json:"horizontalAlignment,omitempty"`
	// RowHeight undocumented
	RowHeight *float64 `json:"rowHeight,omitempty"`
	// VerticalAlignment undocumented
	VerticalAlignment *string `json:"verticalAlignment,omitempty"`
	// WrapText undocumented
	WrapText *bool `json:"wrapText,omitempty"`
	// Borders undocumented
	Borders []WorkbookRangeBorder `json:"borders,omitempty"`
	// Fill undocumented
	Fill *WorkbookRangeFill `json:"fill,omitempty"`
	// Font undocumented
	Font *WorkbookRangeFont `json:"font,omitempty"`
	// Protection undocumented
	Protection *WorkbookFormatProtection `json:"protection,omitempty"`
}

func NewWorkbookRangeFormat() (*WorkbookRangeFormat, error) {
	newWorkbookRangeFormat := &WorkbookRangeFormat{
		ODataType: "#microsoft.graph.WorkbookRangeFormat",
	}
	return newWorkbookRangeFormat, nil
}

// WorkbookRangeReference undocumented
type WorkbookRangeReference struct {
	// Object is the base model of WorkbookRangeReference
	Object

	ODataType string `json:"@odata.type"`
	// Address undocumented
	Address *string `json:"address,omitempty"`
}

func NewWorkbookRangeReference() (*WorkbookRangeReference, error) {
	newWorkbookRangeReference := &WorkbookRangeReference{
		ODataType: "#microsoft.graph.WorkbookRangeReference",
	}
	return newWorkbookRangeReference, nil
}

// WorkbookRangeSort undocumented
type WorkbookRangeSort struct {
	// Entity is the base model of WorkbookRangeSort
	Entity

	ODataType string `json:"@odata.type"`
}

func NewWorkbookRangeSort() (*WorkbookRangeSort, error) {
	newWorkbookRangeSort := &WorkbookRangeSort{
		ODataType: "#microsoft.graph.WorkbookRangeSort",
	}
	return newWorkbookRangeSort, nil
}

// WorkbookRangeView undocumented
type WorkbookRangeView struct {
	// Entity is the base model of WorkbookRangeView
	Entity

	ODataType string `json:"@odata.type"`
	// CellAddresses undocumented
	CellAddresses json.RawMessage `json:"cellAddresses,omitempty"`
	// ColumnCount undocumented
	ColumnCount *int `json:"columnCount,omitempty"`
	// Formulas undocumented
	Formulas json.RawMessage `json:"formulas,omitempty"`
	// FormulasLocal undocumented
	FormulasLocal json.RawMessage `json:"formulasLocal,omitempty"`
	// FormulasR1C1 undocumented
	FormulasR1C1 json.RawMessage `json:"formulasR1C1,omitempty"`
	// Index undocumented
	Index *int `json:"index,omitempty"`
	// NumberFormat undocumented
	NumberFormat json.RawMessage `json:"numberFormat,omitempty"`
	// RowCount undocumented
	RowCount *int `json:"rowCount,omitempty"`
	// Text undocumented
	Text json.RawMessage `json:"text,omitempty"`
	// Values undocumented
	Values json.RawMessage `json:"values,omitempty"`
	// ValueTypes undocumented
	ValueTypes json.RawMessage `json:"valueTypes,omitempty"`
	// Rows undocumented
	Rows []WorkbookRangeView `json:"rows,omitempty"`
}

func NewWorkbookRangeView() (*WorkbookRangeView, error) {
	newWorkbookRangeView := &WorkbookRangeView{
		ODataType: "#microsoft.graph.WorkbookRangeView",
	}
	return newWorkbookRangeView, nil
}

// WorkbookSessionInfo undocumented
type WorkbookSessionInfo struct {
	// Object is the base model of WorkbookSessionInfo
	Object

	ODataType string `json:"@odata.type"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// PersistChanges undocumented
	PersistChanges *bool `json:"persistChanges,omitempty"`
}

func NewWorkbookSessionInfo() (*WorkbookSessionInfo, error) {
	newWorkbookSessionInfo := &WorkbookSessionInfo{
		ODataType: "#microsoft.graph.WorkbookSessionInfo",
	}
	return newWorkbookSessionInfo, nil
}

// WorkbookSortField undocumented
type WorkbookSortField struct {
	// Object is the base model of WorkbookSortField
	Object

	ODataType string `json:"@odata.type"`
	// Ascending undocumented
	Ascending *bool `json:"ascending,omitempty"`
	// Color undocumented
	Color *string `json:"color,omitempty"`
	// DataOption undocumented
	DataOption *string `json:"dataOption,omitempty"`
	// Icon undocumented
	Icon *WorkbookIcon `json:"icon,omitempty"`
	// Key undocumented
	Key *int `json:"key,omitempty"`
	// SortOn undocumented
	SortOn *string `json:"sortOn,omitempty"`
}

func NewWorkbookSortField() (*WorkbookSortField, error) {
	newWorkbookSortField := &WorkbookSortField{
		ODataType: "#microsoft.graph.WorkbookSortField",
	}
	return newWorkbookSortField, nil
}

// WorkbookTable undocumented
type WorkbookTable struct {
	// Entity is the base model of WorkbookTable
	Entity

	ODataType string `json:"@odata.type"`
	// HighlightFirstColumn undocumented
	HighlightFirstColumn *bool `json:"highlightFirstColumn,omitempty"`
	// HighlightLastColumn undocumented
	HighlightLastColumn *bool `json:"highlightLastColumn,omitempty"`
	// LegacyID undocumented
	LegacyID *string `json:"legacyId,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// ShowBandedColumns undocumented
	ShowBandedColumns *bool `json:"showBandedColumns,omitempty"`
	// ShowBandedRows undocumented
	ShowBandedRows *bool `json:"showBandedRows,omitempty"`
	// ShowFilterButton undocumented
	ShowFilterButton *bool `json:"showFilterButton,omitempty"`
	// ShowHeaders undocumented
	ShowHeaders *bool `json:"showHeaders,omitempty"`
	// ShowTotals undocumented
	ShowTotals *bool `json:"showTotals,omitempty"`
	// Style undocumented
	Style *string `json:"style,omitempty"`
	// Columns undocumented
	Columns []WorkbookTableColumn `json:"columns,omitempty"`
	// Rows undocumented
	Rows []WorkbookTableRow `json:"rows,omitempty"`
	// Sort undocumented
	Sort *WorkbookTableSort `json:"sort,omitempty"`
	// Worksheet undocumented
	Worksheet *WorkbookWorksheet `json:"worksheet,omitempty"`
}

func NewWorkbookTable() (*WorkbookTable, error) {
	newWorkbookTable := &WorkbookTable{
		ODataType: "#microsoft.graph.WorkbookTable",
	}
	return newWorkbookTable, nil
}

// WorkbookTableColumn undocumented
type WorkbookTableColumn struct {
	// Entity is the base model of WorkbookTableColumn
	Entity

	ODataType string `json:"@odata.type"`
	// Index undocumented
	Index *int `json:"index,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Values undocumented
	Values json.RawMessage `json:"values,omitempty"`
	// Filter undocumented
	Filter *WorkbookFilter `json:"filter,omitempty"`
}

func NewWorkbookTableColumn() (*WorkbookTableColumn, error) {
	newWorkbookTableColumn := &WorkbookTableColumn{
		ODataType: "#microsoft.graph.WorkbookTableColumn",
	}
	return newWorkbookTableColumn, nil
}

// WorkbookTableRow undocumented
type WorkbookTableRow struct {
	// Entity is the base model of WorkbookTableRow
	Entity

	ODataType string `json:"@odata.type"`
	// Index undocumented
	Index *int `json:"index,omitempty"`
	// Values undocumented
	Values json.RawMessage `json:"values,omitempty"`
}

func NewWorkbookTableRow() (*WorkbookTableRow, error) {
	newWorkbookTableRow := &WorkbookTableRow{
		ODataType: "#microsoft.graph.WorkbookTableRow",
	}
	return newWorkbookTableRow, nil
}

// WorkbookTableSort undocumented
type WorkbookTableSort struct {
	// Entity is the base model of WorkbookTableSort
	Entity

	ODataType string `json:"@odata.type"`
	// Fields undocumented
	Fields []WorkbookSortField `json:"fields,omitempty"`
	// MatchCase undocumented
	MatchCase *bool `json:"matchCase,omitempty"`
	// Method undocumented
	Method *string `json:"method,omitempty"`
}

func NewWorkbookTableSort() (*WorkbookTableSort, error) {
	newWorkbookTableSort := &WorkbookTableSort{
		ODataType: "#microsoft.graph.WorkbookTableSort",
	}
	return newWorkbookTableSort, nil
}

// WorkbookWorksheet undocumented
type WorkbookWorksheet struct {
	// Entity is the base model of WorkbookWorksheet
	Entity

	ODataType string `json:"@odata.type"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Position undocumented
	Position *int `json:"position,omitempty"`
	// Visibility undocumented
	Visibility *string `json:"visibility,omitempty"`
	// Charts undocumented
	Charts []WorkbookChart `json:"charts,omitempty"`
	// Names undocumented
	Names []WorkbookNamedItem `json:"names,omitempty"`
	// PivotTables undocumented
	PivotTables []WorkbookPivotTable `json:"pivotTables,omitempty"`
	// Protection undocumented
	Protection *WorkbookWorksheetProtection `json:"protection,omitempty"`
	// Tables undocumented
	Tables []WorkbookTable `json:"tables,omitempty"`
}

func NewWorkbookWorksheet() (*WorkbookWorksheet, error) {
	newWorkbookWorksheet := &WorkbookWorksheet{
		ODataType: "#microsoft.graph.WorkbookWorksheet",
	}
	return newWorkbookWorksheet, nil
}

// WorkbookWorksheetProtection undocumented
type WorkbookWorksheetProtection struct {
	// Entity is the base model of WorkbookWorksheetProtection
	Entity

	ODataType string `json:"@odata.type"`
	// Options undocumented
	Options *WorkbookWorksheetProtectionOptions `json:"options,omitempty"`
	// Protected undocumented
	Protected *bool `json:"protected,omitempty"`
}

func NewWorkbookWorksheetProtection() (*WorkbookWorksheetProtection, error) {
	newWorkbookWorksheetProtection := &WorkbookWorksheetProtection{
		ODataType: "#microsoft.graph.WorkbookWorksheetProtection",
	}
	return newWorkbookWorksheetProtection, nil
}

// WorkbookWorksheetProtectionOptions undocumented
type WorkbookWorksheetProtectionOptions struct {
	// Object is the base model of WorkbookWorksheetProtectionOptions
	Object

	ODataType string `json:"@odata.type"`
	// AllowAutoFilter undocumented
	AllowAutoFilter *bool `json:"allowAutoFilter,omitempty"`
	// AllowDeleteColumns undocumented
	AllowDeleteColumns *bool `json:"allowDeleteColumns,omitempty"`
	// AllowDeleteRows undocumented
	AllowDeleteRows *bool `json:"allowDeleteRows,omitempty"`
	// AllowFormatCells undocumented
	AllowFormatCells *bool `json:"allowFormatCells,omitempty"`
	// AllowFormatColumns undocumented
	AllowFormatColumns *bool `json:"allowFormatColumns,omitempty"`
	// AllowFormatRows undocumented
	AllowFormatRows *bool `json:"allowFormatRows,omitempty"`
	// AllowInsertColumns undocumented
	AllowInsertColumns *bool `json:"allowInsertColumns,omitempty"`
	// AllowInsertHyperlinks undocumented
	AllowInsertHyperlinks *bool `json:"allowInsertHyperlinks,omitempty"`
	// AllowInsertRows undocumented
	AllowInsertRows *bool `json:"allowInsertRows,omitempty"`
	// AllowPivotTables undocumented
	AllowPivotTables *bool `json:"allowPivotTables,omitempty"`
	// AllowSort undocumented
	AllowSort *bool `json:"allowSort,omitempty"`
}

func NewWorkbookWorksheetProtectionOptions() (*WorkbookWorksheetProtectionOptions, error) {
	newWorkbookWorksheetProtectionOptions := &WorkbookWorksheetProtectionOptions{
		ODataType: "#microsoft.graph.WorkbookWorksheetProtectionOptions",
	}
	return newWorkbookWorksheetProtectionOptions, nil
}
