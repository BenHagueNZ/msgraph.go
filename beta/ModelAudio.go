
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// Audio undocumented
type Audio struct {
    // Object is the base model of Audio
    Object
    // Album undocumented
    Album *string `json:"album,omitempty"`
    // AlbumArtist undocumented
    AlbumArtist *string `json:"albumArtist,omitempty"`
    // Artist undocumented
    Artist *string `json:"artist,omitempty"`
    // Bitrate undocumented
    Bitrate *int `json:"bitrate,omitempty"`
    // Composers undocumented
    Composers *string `json:"composers,omitempty"`
    // Copyright undocumented
    Copyright *string `json:"copyright,omitempty"`
    // Disc undocumented
    Disc *int `json:"disc,omitempty"`
    // DiscCount undocumented
    DiscCount *int `json:"discCount,omitempty"`
    // Duration undocumented
    Duration *int `json:"duration,omitempty"`
    // Genre undocumented
    Genre *string `json:"genre,omitempty"`
    // HasDrm undocumented
    HasDrm *bool `json:"hasDrm,omitempty"`
    // IsVariableBitrate undocumented
    IsVariableBitrate *bool `json:"isVariableBitrate,omitempty"`
    // Title undocumented
    Title *string `json:"title,omitempty"`
    // Track undocumented
    Track *int `json:"track,omitempty"`
    // TrackCount undocumented
    TrackCount *int `json:"trackCount,omitempty"`
    // Year undocumented
    Year *int `json:"year,omitempty"`
}

// AudioConferencing undocumented
type AudioConferencing struct {
    // Object is the base model of AudioConferencing
    Object
    // ConferenceID undocumented
    ConferenceID *string `json:"conferenceId,omitempty"`
    // DialinURL undocumented
    DialinURL *string `json:"dialinUrl,omitempty"`
    // TollFreeNumber undocumented
    TollFreeNumber *string `json:"tollFreeNumber,omitempty"`
    // TollFreeNumbers undocumented
    TollFreeNumbers []string `json:"tollFreeNumbers,omitempty"`
    // TollNumber undocumented
    TollNumber *string `json:"tollNumber,omitempty"`
    // TollNumbers undocumented
    TollNumbers []string `json:"tollNumbers,omitempty"`
}

// AudioRoutingGroup undocumented
type AudioRoutingGroup struct {
    // Entity is the base model of AudioRoutingGroup
    Entity
    // Receivers undocumented
    Receivers []string `json:"receivers,omitempty"`
    // RoutingMode undocumented
    RoutingMode *RoutingMode `json:"routingMode,omitempty"`
    // Sources undocumented
    Sources []string `json:"sources,omitempty"`
}
