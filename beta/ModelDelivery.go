
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// DeliveryOptimizationBandwidth undocumented
type DeliveryOptimizationBandwidth struct {
    // Object is the base model of DeliveryOptimizationBandwidth
    Object
}

// DeliveryOptimizationBandwidthAbsolute undocumented
type DeliveryOptimizationBandwidthAbsolute struct {
    // DeliveryOptimizationBandwidth is the base model of DeliveryOptimizationBandwidthAbsolute
    DeliveryOptimizationBandwidth
    // MaximumDownloadBandwidthInKilobytesPerSecond undocumented
    MaximumDownloadBandwidthInKilobytesPerSecond *int `json:"maximumDownloadBandwidthInKilobytesPerSecond,omitempty"`
    // MaximumUploadBandwidthInKilobytesPerSecond undocumented
    MaximumUploadBandwidthInKilobytesPerSecond *int `json:"maximumUploadBandwidthInKilobytesPerSecond,omitempty"`
}

// DeliveryOptimizationBandwidthBusinessHoursLimit undocumented
type DeliveryOptimizationBandwidthBusinessHoursLimit struct {
    // Object is the base model of DeliveryOptimizationBandwidthBusinessHoursLimit
    Object
    // BandwidthBeginBusinessHours undocumented
    BandwidthBeginBusinessHours *int `json:"bandwidthBeginBusinessHours,omitempty"`
    // BandwidthEndBusinessHours undocumented
    BandwidthEndBusinessHours *int `json:"bandwidthEndBusinessHours,omitempty"`
    // BandwidthPercentageDuringBusinessHours undocumented
    BandwidthPercentageDuringBusinessHours *int `json:"bandwidthPercentageDuringBusinessHours,omitempty"`
    // BandwidthPercentageOutsideBusinessHours undocumented
    BandwidthPercentageOutsideBusinessHours *int `json:"bandwidthPercentageOutsideBusinessHours,omitempty"`
}

// DeliveryOptimizationBandwidthHoursWithPercentage undocumented
type DeliveryOptimizationBandwidthHoursWithPercentage struct {
    // DeliveryOptimizationBandwidth is the base model of DeliveryOptimizationBandwidthHoursWithPercentage
    DeliveryOptimizationBandwidth
    // BandwidthBackgroundPercentageHours undocumented
    BandwidthBackgroundPercentageHours *DeliveryOptimizationBandwidthBusinessHoursLimit `json:"bandwidthBackgroundPercentageHours,omitempty"`
    // BandwidthForegroundPercentageHours undocumented
    BandwidthForegroundPercentageHours *DeliveryOptimizationBandwidthBusinessHoursLimit `json:"bandwidthForegroundPercentageHours,omitempty"`
}

// DeliveryOptimizationBandwidthPercentage undocumented
type DeliveryOptimizationBandwidthPercentage struct {
    // DeliveryOptimizationBandwidth is the base model of DeliveryOptimizationBandwidthPercentage
    DeliveryOptimizationBandwidth
    // MaximumBackgroundBandwidthPercentage undocumented
    MaximumBackgroundBandwidthPercentage *int `json:"maximumBackgroundBandwidthPercentage,omitempty"`
    // MaximumForegroundBandwidthPercentage undocumented
    MaximumForegroundBandwidthPercentage *int `json:"maximumForegroundBandwidthPercentage,omitempty"`
}

// DeliveryOptimizationGroupIDCustom undocumented
type DeliveryOptimizationGroupIDCustom struct {
    // DeliveryOptimizationGroupIDSource is the base model of DeliveryOptimizationGroupIDCustom
    DeliveryOptimizationGroupIDSource
    // GroupIDCustom undocumented
    GroupIDCustom *string `json:"groupIdCustom,omitempty"`
}

// DeliveryOptimizationGroupIDSource undocumented
type DeliveryOptimizationGroupIDSource struct {
    // Object is the base model of DeliveryOptimizationGroupIDSource
    Object
}

// DeliveryOptimizationGroupIDSourceOptions undocumented
type DeliveryOptimizationGroupIDSourceOptions struct {
    // DeliveryOptimizationGroupIDSource is the base model of DeliveryOptimizationGroupIDSourceOptions
    DeliveryOptimizationGroupIDSource
    // GroupIDSourceOption undocumented
    GroupIDSourceOption *DeliveryOptimizationGroupIDOptionsType `json:"groupIdSourceOption,omitempty"`
}

// DeliveryOptimizationMaxCacheSize undocumented
type DeliveryOptimizationMaxCacheSize struct {
    // Object is the base model of DeliveryOptimizationMaxCacheSize
    Object
}

// DeliveryOptimizationMaxCacheSizeAbsolute undocumented
type DeliveryOptimizationMaxCacheSizeAbsolute struct {
    // DeliveryOptimizationMaxCacheSize is the base model of DeliveryOptimizationMaxCacheSizeAbsolute
    DeliveryOptimizationMaxCacheSize
    // MaximumCacheSizeInGigabytes undocumented
    MaximumCacheSizeInGigabytes *int `json:"maximumCacheSizeInGigabytes,omitempty"`
}

// DeliveryOptimizationMaxCacheSizePercentage undocumented
type DeliveryOptimizationMaxCacheSizePercentage struct {
    // DeliveryOptimizationMaxCacheSize is the base model of DeliveryOptimizationMaxCacheSizePercentage
    DeliveryOptimizationMaxCacheSize
    // MaximumCacheSizePercentage undocumented
    MaximumCacheSizePercentage *int `json:"maximumCacheSizePercentage,omitempty"`
}
