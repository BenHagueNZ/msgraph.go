// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// TeamworkSendActivityNotificationToRecipientsRequestParameter undocumented
type TeamworkSendActivityNotificationToRecipientsRequestParameter struct {
	// Topic undocumented
	Topic *TeamworkActivityTopic `json:"topic,omitempty"`
	// ActivityType undocumented
	ActivityType *string `json:"activityType,omitempty"`
	// ChainID undocumented
	ChainID *int `json:"chainId,omitempty"`
	// PreviewText undocumented
	PreviewText *ItemBody `json:"previewText,omitempty"`
	// TeamsAppID undocumented
	TeamsAppID *string `json:"teamsAppId,omitempty"`
	// TemplateParameters undocumented
	TemplateParameters []KeyValuePair `json:"templateParameters,omitempty"`
	// Recipients undocumented
	Recipients []TeamworkNotificationRecipient `json:"recipients,omitempty"`
}

// TeamworkDeviceRestartRequestParameter undocumented
type TeamworkDeviceRestartRequestParameter struct {
}

// TeamworkDeviceRunDiagnosticsRequestParameter undocumented
type TeamworkDeviceRunDiagnosticsRequestParameter struct {
}

// TeamworkDeviceUpdateSoftwareRequestParameter undocumented
type TeamworkDeviceUpdateSoftwareRequestParameter struct {
	// SoftwareType undocumented
	SoftwareType *TeamworkSoftwareType `json:"softwareType,omitempty"`
	// SoftwareVersion undocumented
	SoftwareVersion *string `json:"softwareVersion,omitempty"`
}

// DeletedTeams returns request builder for DeletedTeam collection
func (b *TeamworkRequestBuilder) DeletedTeams() *TeamworkDeletedTeamsCollectionRequestBuilder {
	bb := &TeamworkDeletedTeamsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deletedTeams"
	return bb
}

// TeamworkDeletedTeamsCollectionRequestBuilder is request builder for DeletedTeam collection rcn
type TeamworkDeletedTeamsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeletedTeam collection
func (b *TeamworkDeletedTeamsCollectionRequestBuilder) Request() *TeamworkDeletedTeamsCollectionRequest {
	return &TeamworkDeletedTeamsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeletedTeam item
func (b *TeamworkDeletedTeamsCollectionRequestBuilder) ID(id string) *DeletedTeamRequestBuilder {
	bb := &DeletedTeamRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamworkDeletedTeamsCollectionRequest is request for DeletedTeam collection
type TeamworkDeletedTeamsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DeletedTeam collection
func (r *TeamworkDeletedTeamsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DeletedTeam, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DeletedTeam
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []DeletedTeam
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for DeletedTeam collection, max N pages
func (r *TeamworkDeletedTeamsCollectionRequest) GetN(ctx context.Context, n int) ([]DeletedTeam, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DeletedTeam collection
func (r *TeamworkDeletedTeamsCollectionRequest) Get(ctx context.Context) ([]DeletedTeam, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DeletedTeam collection
func (r *TeamworkDeletedTeamsCollectionRequest) Add(ctx context.Context, reqObj *DeletedTeam) (resObj *DeletedTeam, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Devices returns request builder for TeamworkDevice collection
func (b *TeamworkRequestBuilder) Devices() *TeamworkDevicesCollectionRequestBuilder {
	bb := &TeamworkDevicesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/devices"
	return bb
}

// TeamworkDevicesCollectionRequestBuilder is request builder for TeamworkDevice collection rcn
type TeamworkDevicesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamworkDevice collection
func (b *TeamworkDevicesCollectionRequestBuilder) Request() *TeamworkDevicesCollectionRequest {
	return &TeamworkDevicesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamworkDevice item
func (b *TeamworkDevicesCollectionRequestBuilder) ID(id string) *TeamworkDeviceRequestBuilder {
	bb := &TeamworkDeviceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamworkDevicesCollectionRequest is request for TeamworkDevice collection
type TeamworkDevicesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamworkDevice collection
func (r *TeamworkDevicesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TeamworkDevice, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []TeamworkDevice
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []TeamworkDevice
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for TeamworkDevice collection, max N pages
func (r *TeamworkDevicesCollectionRequest) GetN(ctx context.Context, n int) ([]TeamworkDevice, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TeamworkDevice collection
func (r *TeamworkDevicesCollectionRequest) Get(ctx context.Context) ([]TeamworkDevice, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TeamworkDevice collection
func (r *TeamworkDevicesCollectionRequest) Add(ctx context.Context, reqObj *TeamworkDevice) (resObj *TeamworkDevice, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// TeamTemplates returns request builder for TeamTemplate collection
func (b *TeamworkRequestBuilder) TeamTemplates() *TeamworkTeamTemplatesCollectionRequestBuilder {
	bb := &TeamworkTeamTemplatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/teamTemplates"
	return bb
}

// TeamworkTeamTemplatesCollectionRequestBuilder is request builder for TeamTemplate collection rcn
type TeamworkTeamTemplatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamTemplate collection
func (b *TeamworkTeamTemplatesCollectionRequestBuilder) Request() *TeamworkTeamTemplatesCollectionRequest {
	return &TeamworkTeamTemplatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamTemplate item
func (b *TeamworkTeamTemplatesCollectionRequestBuilder) ID(id string) *TeamTemplateRequestBuilder {
	bb := &TeamTemplateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamworkTeamTemplatesCollectionRequest is request for TeamTemplate collection
type TeamworkTeamTemplatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamTemplate collection
func (r *TeamworkTeamTemplatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TeamTemplate, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []TeamTemplate
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []TeamTemplate
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for TeamTemplate collection, max N pages
func (r *TeamworkTeamTemplatesCollectionRequest) GetN(ctx context.Context, n int) ([]TeamTemplate, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TeamTemplate collection
func (r *TeamworkTeamTemplatesCollectionRequest) Get(ctx context.Context) ([]TeamTemplate, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TeamTemplate collection
func (r *TeamworkTeamTemplatesCollectionRequest) Add(ctx context.Context, reqObj *TeamTemplate) (resObj *TeamTemplate, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// TeamsAppSettings is navigation property rn
func (b *TeamworkRequestBuilder) TeamsAppSettings() *TeamsAppSettingsRequestBuilder {
	bb := &TeamsAppSettingsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/teamsAppSettings"
	return bb
}

// WorkforceIntegrations returns request builder for WorkforceIntegration collection
func (b *TeamworkRequestBuilder) WorkforceIntegrations() *TeamworkWorkforceIntegrationsCollectionRequestBuilder {
	bb := &TeamworkWorkforceIntegrationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/workforceIntegrations"
	return bb
}

// TeamworkWorkforceIntegrationsCollectionRequestBuilder is request builder for WorkforceIntegration collection rcn
type TeamworkWorkforceIntegrationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WorkforceIntegration collection
func (b *TeamworkWorkforceIntegrationsCollectionRequestBuilder) Request() *TeamworkWorkforceIntegrationsCollectionRequest {
	return &TeamworkWorkforceIntegrationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WorkforceIntegration item
func (b *TeamworkWorkforceIntegrationsCollectionRequestBuilder) ID(id string) *WorkforceIntegrationRequestBuilder {
	bb := &WorkforceIntegrationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamworkWorkforceIntegrationsCollectionRequest is request for WorkforceIntegration collection
type TeamworkWorkforceIntegrationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WorkforceIntegration collection
func (r *TeamworkWorkforceIntegrationsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]WorkforceIntegration, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []WorkforceIntegration
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []WorkforceIntegration
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for WorkforceIntegration collection, max N pages
func (r *TeamworkWorkforceIntegrationsCollectionRequest) GetN(ctx context.Context, n int) ([]WorkforceIntegration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for WorkforceIntegration collection
func (r *TeamworkWorkforceIntegrationsCollectionRequest) Get(ctx context.Context) ([]WorkforceIntegration, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for WorkforceIntegration collection
func (r *TeamworkWorkforceIntegrationsCollectionRequest) Add(ctx context.Context, reqObj *WorkforceIntegration) (resObj *WorkforceIntegration, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// CommunicationSpeaker is navigation property rn
func (b *TeamworkActivePeripheralsRequestBuilder) CommunicationSpeaker() *TeamworkPeripheralRequestBuilder {
	bb := &TeamworkPeripheralRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/communicationSpeaker"
	return bb
}

// ContentCamera is navigation property rn
func (b *TeamworkActivePeripheralsRequestBuilder) ContentCamera() *TeamworkPeripheralRequestBuilder {
	bb := &TeamworkPeripheralRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/contentCamera"
	return bb
}

// Microphone is navigation property rn
func (b *TeamworkActivePeripheralsRequestBuilder) Microphone() *TeamworkPeripheralRequestBuilder {
	bb := &TeamworkPeripheralRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/microphone"
	return bb
}

// RoomCamera is navigation property rn
func (b *TeamworkActivePeripheralsRequestBuilder) RoomCamera() *TeamworkPeripheralRequestBuilder {
	bb := &TeamworkPeripheralRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roomCamera"
	return bb
}

// Speaker is navigation property rn
func (b *TeamworkActivePeripheralsRequestBuilder) Speaker() *TeamworkPeripheralRequestBuilder {
	bb := &TeamworkPeripheralRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/speaker"
	return bb
}

// Cameras returns request builder for TeamworkPeripheral collection
func (b *TeamworkCameraConfigurationRequestBuilder) Cameras() *TeamworkCameraConfigurationCamerasCollectionRequestBuilder {
	bb := &TeamworkCameraConfigurationCamerasCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/cameras"
	return bb
}

// TeamworkCameraConfigurationCamerasCollectionRequestBuilder is request builder for TeamworkPeripheral collection rcn
type TeamworkCameraConfigurationCamerasCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamworkPeripheral collection
func (b *TeamworkCameraConfigurationCamerasCollectionRequestBuilder) Request() *TeamworkCameraConfigurationCamerasCollectionRequest {
	return &TeamworkCameraConfigurationCamerasCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamworkPeripheral item
func (b *TeamworkCameraConfigurationCamerasCollectionRequestBuilder) ID(id string) *TeamworkPeripheralRequestBuilder {
	bb := &TeamworkPeripheralRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamworkCameraConfigurationCamerasCollectionRequest is request for TeamworkPeripheral collection
type TeamworkCameraConfigurationCamerasCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamworkPeripheral collection
func (r *TeamworkCameraConfigurationCamerasCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TeamworkPeripheral, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []TeamworkPeripheral
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []TeamworkPeripheral
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for TeamworkPeripheral collection, max N pages
func (r *TeamworkCameraConfigurationCamerasCollectionRequest) GetN(ctx context.Context, n int) ([]TeamworkPeripheral, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TeamworkPeripheral collection
func (r *TeamworkCameraConfigurationCamerasCollectionRequest) Get(ctx context.Context) ([]TeamworkPeripheral, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TeamworkPeripheral collection
func (r *TeamworkCameraConfigurationCamerasCollectionRequest) Add(ctx context.Context, reqObj *TeamworkPeripheral) (resObj *TeamworkPeripheral, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DefaultContentCamera is navigation property rn
func (b *TeamworkCameraConfigurationRequestBuilder) DefaultContentCamera() *TeamworkPeripheralRequestBuilder {
	bb := &TeamworkPeripheralRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/defaultContentCamera"
	return bb
}

// Peripheral is navigation property rn
func (b *TeamworkConfiguredPeripheralRequestBuilder) Peripheral() *TeamworkPeripheralRequestBuilder {
	bb := &TeamworkPeripheralRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/peripheral"
	return bb
}

// Activity is navigation property rn
func (b *TeamworkDeviceRequestBuilder) Activity() *TeamworkDeviceActivityRequestBuilder {
	bb := &TeamworkDeviceActivityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/activity"
	return bb
}

// Configuration is navigation property rn
func (b *TeamworkDeviceRequestBuilder) Configuration() *TeamworkDeviceConfigurationRequestBuilder {
	bb := &TeamworkDeviceConfigurationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/configuration"
	return bb
}

// Health is navigation property rn
func (b *TeamworkDeviceRequestBuilder) Health() *TeamworkDeviceHealthRequestBuilder {
	bb := &TeamworkDeviceHealthRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/health"
	return bb
}

// Operations returns request builder for TeamworkDeviceOperation collection
func (b *TeamworkDeviceRequestBuilder) Operations() *TeamworkDeviceOperationsCollectionRequestBuilder {
	bb := &TeamworkDeviceOperationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/operations"
	return bb
}

// TeamworkDeviceOperationsCollectionRequestBuilder is request builder for TeamworkDeviceOperation collection rcn
type TeamworkDeviceOperationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamworkDeviceOperation collection
func (b *TeamworkDeviceOperationsCollectionRequestBuilder) Request() *TeamworkDeviceOperationsCollectionRequest {
	return &TeamworkDeviceOperationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamworkDeviceOperation item
func (b *TeamworkDeviceOperationsCollectionRequestBuilder) ID(id string) *TeamworkDeviceOperationRequestBuilder {
	bb := &TeamworkDeviceOperationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamworkDeviceOperationsCollectionRequest is request for TeamworkDeviceOperation collection
type TeamworkDeviceOperationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamworkDeviceOperation collection
func (r *TeamworkDeviceOperationsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TeamworkDeviceOperation, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []TeamworkDeviceOperation
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []TeamworkDeviceOperation
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for TeamworkDeviceOperation collection, max N pages
func (r *TeamworkDeviceOperationsCollectionRequest) GetN(ctx context.Context, n int) ([]TeamworkDeviceOperation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TeamworkDeviceOperation collection
func (r *TeamworkDeviceOperationsCollectionRequest) Get(ctx context.Context) ([]TeamworkDeviceOperation, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TeamworkDeviceOperation collection
func (r *TeamworkDeviceOperationsCollectionRequest) Add(ctx context.Context, reqObj *TeamworkDeviceOperation) (resObj *TeamworkDeviceOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Compute is navigation property rn
func (b *TeamworkHardwareConfigurationRequestBuilder) Compute() *TeamworkPeripheralRequestBuilder {
	bb := &TeamworkPeripheralRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/compute"
	return bb
}

// HdmiIngest is navigation property rn
func (b *TeamworkHardwareConfigurationRequestBuilder) HdmiIngest() *TeamworkPeripheralRequestBuilder {
	bb := &TeamworkPeripheralRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/hdmiIngest"
	return bb
}

// DefaultMicrophone is navigation property rn
func (b *TeamworkMicrophoneConfigurationRequestBuilder) DefaultMicrophone() *TeamworkPeripheralRequestBuilder {
	bb := &TeamworkPeripheralRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/defaultMicrophone"
	return bb
}

// Microphones returns request builder for TeamworkPeripheral collection
func (b *TeamworkMicrophoneConfigurationRequestBuilder) Microphones() *TeamworkMicrophoneConfigurationMicrophonesCollectionRequestBuilder {
	bb := &TeamworkMicrophoneConfigurationMicrophonesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/microphones"
	return bb
}

// TeamworkMicrophoneConfigurationMicrophonesCollectionRequestBuilder is request builder for TeamworkPeripheral collection rcn
type TeamworkMicrophoneConfigurationMicrophonesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamworkPeripheral collection
func (b *TeamworkMicrophoneConfigurationMicrophonesCollectionRequestBuilder) Request() *TeamworkMicrophoneConfigurationMicrophonesCollectionRequest {
	return &TeamworkMicrophoneConfigurationMicrophonesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamworkPeripheral item
func (b *TeamworkMicrophoneConfigurationMicrophonesCollectionRequestBuilder) ID(id string) *TeamworkPeripheralRequestBuilder {
	bb := &TeamworkPeripheralRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamworkMicrophoneConfigurationMicrophonesCollectionRequest is request for TeamworkPeripheral collection
type TeamworkMicrophoneConfigurationMicrophonesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamworkPeripheral collection
func (r *TeamworkMicrophoneConfigurationMicrophonesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TeamworkPeripheral, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []TeamworkPeripheral
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []TeamworkPeripheral
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for TeamworkPeripheral collection, max N pages
func (r *TeamworkMicrophoneConfigurationMicrophonesCollectionRequest) GetN(ctx context.Context, n int) ([]TeamworkPeripheral, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TeamworkPeripheral collection
func (r *TeamworkMicrophoneConfigurationMicrophonesCollectionRequest) Get(ctx context.Context) ([]TeamworkPeripheral, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TeamworkPeripheral collection
func (r *TeamworkMicrophoneConfigurationMicrophonesCollectionRequest) Add(ctx context.Context, reqObj *TeamworkPeripheral) (resObj *TeamworkPeripheral, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Peripheral is navigation property rn
func (b *TeamworkPeripheralHealthRequestBuilder) Peripheral() *TeamworkPeripheralRequestBuilder {
	bb := &TeamworkPeripheralRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/peripheral"
	return bb
}

// DefaultCommunicationSpeaker is navigation property rn
func (b *TeamworkSpeakerConfigurationRequestBuilder) DefaultCommunicationSpeaker() *TeamworkPeripheralRequestBuilder {
	bb := &TeamworkPeripheralRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/defaultCommunicationSpeaker"
	return bb
}

// DefaultSpeaker is navigation property rn
func (b *TeamworkSpeakerConfigurationRequestBuilder) DefaultSpeaker() *TeamworkPeripheralRequestBuilder {
	bb := &TeamworkPeripheralRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/defaultSpeaker"
	return bb
}

// Speakers returns request builder for TeamworkPeripheral collection
func (b *TeamworkSpeakerConfigurationRequestBuilder) Speakers() *TeamworkSpeakerConfigurationSpeakersCollectionRequestBuilder {
	bb := &TeamworkSpeakerConfigurationSpeakersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/speakers"
	return bb
}

// TeamworkSpeakerConfigurationSpeakersCollectionRequestBuilder is request builder for TeamworkPeripheral collection rcn
type TeamworkSpeakerConfigurationSpeakersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamworkPeripheral collection
func (b *TeamworkSpeakerConfigurationSpeakersCollectionRequestBuilder) Request() *TeamworkSpeakerConfigurationSpeakersCollectionRequest {
	return &TeamworkSpeakerConfigurationSpeakersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamworkPeripheral item
func (b *TeamworkSpeakerConfigurationSpeakersCollectionRequestBuilder) ID(id string) *TeamworkPeripheralRequestBuilder {
	bb := &TeamworkPeripheralRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamworkSpeakerConfigurationSpeakersCollectionRequest is request for TeamworkPeripheral collection
type TeamworkSpeakerConfigurationSpeakersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamworkPeripheral collection
func (r *TeamworkSpeakerConfigurationSpeakersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TeamworkPeripheral, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []TeamworkPeripheral
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []TeamworkPeripheral
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for TeamworkPeripheral collection, max N pages
func (r *TeamworkSpeakerConfigurationSpeakersCollectionRequest) GetN(ctx context.Context, n int) ([]TeamworkPeripheral, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TeamworkPeripheral collection
func (r *TeamworkSpeakerConfigurationSpeakersCollectionRequest) Get(ctx context.Context) ([]TeamworkPeripheral, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TeamworkPeripheral collection
func (r *TeamworkSpeakerConfigurationSpeakersCollectionRequest) Add(ctx context.Context, reqObj *TeamworkPeripheral) (resObj *TeamworkPeripheral, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Members returns request builder for TeamworkTagMember collection
func (b *TeamworkTagRequestBuilder) Members() *TeamworkTagMembersCollectionRequestBuilder {
	bb := &TeamworkTagMembersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/members"
	return bb
}

// TeamworkTagMembersCollectionRequestBuilder is request builder for TeamworkTagMember collection rcn
type TeamworkTagMembersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamworkTagMember collection
func (b *TeamworkTagMembersCollectionRequestBuilder) Request() *TeamworkTagMembersCollectionRequest {
	return &TeamworkTagMembersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamworkTagMember item
func (b *TeamworkTagMembersCollectionRequestBuilder) ID(id string) *TeamworkTagMemberRequestBuilder {
	bb := &TeamworkTagMemberRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamworkTagMembersCollectionRequest is request for TeamworkTagMember collection
type TeamworkTagMembersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamworkTagMember collection
func (r *TeamworkTagMembersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TeamworkTagMember, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []TeamworkTagMember
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []TeamworkTagMember
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for TeamworkTagMember collection, max N pages
func (r *TeamworkTagMembersCollectionRequest) GetN(ctx context.Context, n int) ([]TeamworkTagMember, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TeamworkTagMember collection
func (r *TeamworkTagMembersCollectionRequest) Get(ctx context.Context) ([]TeamworkTagMember, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TeamworkTagMember collection
func (r *TeamworkTagMembersCollectionRequest) Add(ctx context.Context, reqObj *TeamworkTagMember) (resObj *TeamworkTagMember, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *TeamworkRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *TeamworkBotRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *TeamworkDeviceRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *TeamworkDeviceActivityRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *TeamworkDeviceConfigurationRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *TeamworkDeviceHealthRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *TeamworkDeviceOperationRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *TeamworkPeripheralRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *TeamworkTagRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *TeamworkTagMemberRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
