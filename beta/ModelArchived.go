
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// ArchivedPrintJob undocumented
type ArchivedPrintJob struct {
    // Object is the base model of ArchivedPrintJob
    Object
    // AcquiredByPrinter undocumented
    AcquiredByPrinter *bool `json:"acquiredByPrinter,omitempty"`
    // AcquiredDateTime undocumented
    AcquiredDateTime *time.Time `json:"acquiredDateTime,omitempty"`
    // BlackAndWhitePageCount undocumented
    BlackAndWhitePageCount *int `json:"blackAndWhitePageCount,omitempty"`
    // ColorPageCount undocumented
    ColorPageCount *int `json:"colorPageCount,omitempty"`
    // CompletionDateTime undocumented
    CompletionDateTime *time.Time `json:"completionDateTime,omitempty"`
    // CopiesPrinted undocumented
    CopiesPrinted *int `json:"copiesPrinted,omitempty"`
    // CreatedBy undocumented
    CreatedBy *UserIdentity `json:"createdBy,omitempty"`
    // CreatedDateTime undocumented
    CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
    // DuplexPageCount undocumented
    DuplexPageCount *int `json:"duplexPageCount,omitempty"`
    // ID undocumented
    ID *string `json:"id,omitempty"`
    // PageCount undocumented
    PageCount *int `json:"pageCount,omitempty"`
    // PrinterID undocumented
    PrinterID *string `json:"printerId,omitempty"`
    // ProcessingState undocumented
    ProcessingState *PrintJobProcessingState `json:"processingState,omitempty"`
    // SimplexPageCount undocumented
    SimplexPageCount *int `json:"simplexPageCount,omitempty"`
}
