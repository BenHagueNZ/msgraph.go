// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// TeleconferenceDeviceAudioQuality undocumented
type TeleconferenceDeviceAudioQuality struct {
	// TeleconferenceDeviceMediaQuality is the base model of TeleconferenceDeviceAudioQuality
	TeleconferenceDeviceMediaQuality

	ODataType string `json:"@odata.type"`
}

func NewTeleconferenceDeviceAudioQuality() (*TeleconferenceDeviceAudioQuality, error) {
	newTeleconferenceDeviceAudioQuality := &TeleconferenceDeviceAudioQuality{
		ODataType: "#microsoft.graph.TeleconferenceDeviceAudioQuality",
	}
	return newTeleconferenceDeviceAudioQuality, nil
}

// TeleconferenceDeviceMediaQuality undocumented
type TeleconferenceDeviceMediaQuality struct {
	// Object is the base model of TeleconferenceDeviceMediaQuality
	Object

	ODataType string `json:"@odata.type"`
	// AverageInboundJitter undocumented
	AverageInboundJitter *Duration `json:"averageInboundJitter,omitempty"`
	// AverageInboundPacketLossRateInPercentage undocumented
	AverageInboundPacketLossRateInPercentage *float64 `json:"averageInboundPacketLossRateInPercentage,omitempty"`
	// AverageInboundRoundTripDelay undocumented
	AverageInboundRoundTripDelay *Duration `json:"averageInboundRoundTripDelay,omitempty"`
	// AverageOutboundJitter undocumented
	AverageOutboundJitter *Duration `json:"averageOutboundJitter,omitempty"`
	// AverageOutboundPacketLossRateInPercentage undocumented
	AverageOutboundPacketLossRateInPercentage *float64 `json:"averageOutboundPacketLossRateInPercentage,omitempty"`
	// AverageOutboundRoundTripDelay undocumented
	AverageOutboundRoundTripDelay *Duration `json:"averageOutboundRoundTripDelay,omitempty"`
	// ChannelIndex undocumented
	ChannelIndex *int `json:"channelIndex,omitempty"`
	// InboundPackets undocumented
	InboundPackets *int `json:"inboundPackets,omitempty"`
	// LocalIPAddress undocumented
	LocalIPAddress *string `json:"localIPAddress,omitempty"`
	// LocalPort undocumented
	LocalPort *int `json:"localPort,omitempty"`
	// MaximumInboundJitter undocumented
	MaximumInboundJitter *Duration `json:"maximumInboundJitter,omitempty"`
	// MaximumInboundPacketLossRateInPercentage undocumented
	MaximumInboundPacketLossRateInPercentage *float64 `json:"maximumInboundPacketLossRateInPercentage,omitempty"`
	// MaximumInboundRoundTripDelay undocumented
	MaximumInboundRoundTripDelay *Duration `json:"maximumInboundRoundTripDelay,omitempty"`
	// MaximumOutboundJitter undocumented
	MaximumOutboundJitter *Duration `json:"maximumOutboundJitter,omitempty"`
	// MaximumOutboundPacketLossRateInPercentage undocumented
	MaximumOutboundPacketLossRateInPercentage *float64 `json:"maximumOutboundPacketLossRateInPercentage,omitempty"`
	// MaximumOutboundRoundTripDelay undocumented
	MaximumOutboundRoundTripDelay *Duration `json:"maximumOutboundRoundTripDelay,omitempty"`
	// MediaDuration undocumented
	MediaDuration *Duration `json:"mediaDuration,omitempty"`
	// NetworkLinkSpeedInBytes undocumented
	NetworkLinkSpeedInBytes *int `json:"networkLinkSpeedInBytes,omitempty"`
	// OutboundPackets undocumented
	OutboundPackets *int `json:"outboundPackets,omitempty"`
	// RemoteIPAddress undocumented
	RemoteIPAddress *string `json:"remoteIPAddress,omitempty"`
	// RemotePort undocumented
	RemotePort *int `json:"remotePort,omitempty"`
}

func NewTeleconferenceDeviceMediaQuality() (*TeleconferenceDeviceMediaQuality, error) {
	newTeleconferenceDeviceMediaQuality := &TeleconferenceDeviceMediaQuality{
		ODataType: "#microsoft.graph.TeleconferenceDeviceMediaQuality",
	}
	return newTeleconferenceDeviceMediaQuality, nil
}

// TeleconferenceDeviceQuality undocumented
type TeleconferenceDeviceQuality struct {
	// Object is the base model of TeleconferenceDeviceQuality
	Object

	ODataType string `json:"@odata.type"`
	// CallChainID undocumented
	CallChainID *UUID `json:"callChainId,omitempty"`
	// CloudServiceDeploymentEnvironment undocumented
	CloudServiceDeploymentEnvironment *string `json:"cloudServiceDeploymentEnvironment,omitempty"`
	// CloudServiceDeploymentID undocumented
	CloudServiceDeploymentID *string `json:"cloudServiceDeploymentId,omitempty"`
	// CloudServiceInstanceName undocumented
	CloudServiceInstanceName *string `json:"cloudServiceInstanceName,omitempty"`
	// CloudServiceName undocumented
	CloudServiceName *string `json:"cloudServiceName,omitempty"`
	// DeviceDescription undocumented
	DeviceDescription *string `json:"deviceDescription,omitempty"`
	// DeviceName undocumented
	DeviceName *string `json:"deviceName,omitempty"`
	// MediaLegID undocumented
	MediaLegID *UUID `json:"mediaLegId,omitempty"`
	// MediaQualityList undocumented
	MediaQualityList []TeleconferenceDeviceMediaQuality `json:"mediaQualityList,omitempty"`
	// ParticipantID undocumented
	ParticipantID *UUID `json:"participantId,omitempty"`
}

func NewTeleconferenceDeviceQuality() (*TeleconferenceDeviceQuality, error) {
	newTeleconferenceDeviceQuality := &TeleconferenceDeviceQuality{
		ODataType: "#microsoft.graph.TeleconferenceDeviceQuality",
	}
	return newTeleconferenceDeviceQuality, nil
}

// TeleconferenceDeviceScreenSharingQuality undocumented
type TeleconferenceDeviceScreenSharingQuality struct {
	// TeleconferenceDeviceVideoQuality is the base model of TeleconferenceDeviceScreenSharingQuality
	TeleconferenceDeviceVideoQuality

	ODataType string `json:"@odata.type"`
}

func NewTeleconferenceDeviceScreenSharingQuality() (*TeleconferenceDeviceScreenSharingQuality, error) {
	newTeleconferenceDeviceScreenSharingQuality := &TeleconferenceDeviceScreenSharingQuality{
		ODataType: "#microsoft.graph.TeleconferenceDeviceScreenSharingQuality",
	}
	return newTeleconferenceDeviceScreenSharingQuality, nil
}

// TeleconferenceDeviceVideoQuality undocumented
type TeleconferenceDeviceVideoQuality struct {
	// TeleconferenceDeviceMediaQuality is the base model of TeleconferenceDeviceVideoQuality
	TeleconferenceDeviceMediaQuality

	ODataType string `json:"@odata.type"`
	// AverageInboundBitRate undocumented
	AverageInboundBitRate *float64 `json:"averageInboundBitRate,omitempty"`
	// AverageInboundFrameRate undocumented
	AverageInboundFrameRate *float64 `json:"averageInboundFrameRate,omitempty"`
	// AverageOutboundBitRate undocumented
	AverageOutboundBitRate *float64 `json:"averageOutboundBitRate,omitempty"`
	// AverageOutboundFrameRate undocumented
	AverageOutboundFrameRate *float64 `json:"averageOutboundFrameRate,omitempty"`
}

func NewTeleconferenceDeviceVideoQuality() (*TeleconferenceDeviceVideoQuality, error) {
	newTeleconferenceDeviceVideoQuality := &TeleconferenceDeviceVideoQuality{
		ODataType: "#microsoft.graph.TeleconferenceDeviceVideoQuality",
	}
	return newTeleconferenceDeviceVideoQuality, nil
}
