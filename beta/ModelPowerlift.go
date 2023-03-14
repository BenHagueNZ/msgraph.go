
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// PowerliftDownloadRequestObject undocumented
type PowerliftDownloadRequestObject struct {
    // Object is the base model of PowerliftDownloadRequestObject
    Object
    // Files undocumented
    Files []string `json:"files,omitempty"`
    // PowerliftID undocumented
    PowerliftID *UUID `json:"powerliftId,omitempty"`
}

// PowerliftIncidentMetadata undocumented
type PowerliftIncidentMetadata struct {
    // Object is the base model of PowerliftIncidentMetadata
    Object
    // Application undocumented
    Application *string `json:"application,omitempty"`
    // ClientVersion undocumented
    ClientVersion *string `json:"clientVersion,omitempty"`
    // CreatedAtDateTime undocumented
    CreatedAtDateTime *time.Time `json:"createdAtDateTime,omitempty"`
    // EasyID undocumented
    EasyID *string `json:"easyId,omitempty"`
    // FileNames undocumented
    FileNames []string `json:"fileNames,omitempty"`
    // Locale undocumented
    Locale *string `json:"locale,omitempty"`
    // Platform undocumented
    Platform *string `json:"platform,omitempty"`
    // PowerliftID undocumented
    PowerliftID *UUID `json:"powerliftId,omitempty"`
}