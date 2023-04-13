// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// CallCollectionLogTeleconferenceDeviceQualityRequestParameter undocumented
type CallCollectionLogTeleconferenceDeviceQualityRequestParameter struct {
	// Quality undocumented
	Quality *TeleconferenceDeviceQuality `json:"quality,omitempty"`
}

// CallRedirectRequestParameter undocumented
type CallRedirectRequestParameter struct {
	// Targets undocumented
	Targets []InvitationParticipantInfo `json:"targets,omitempty"`
	// Timeout undocumented
	Timeout *int `json:"timeout,omitempty"`
	// CallbackURI undocumented
	CallbackURI *string `json:"callbackUri,omitempty"`
}

// CallAddLargeGalleryViewRequestParameter undocumented
type CallAddLargeGalleryViewRequestParameter struct {
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

// CallAnswerRequestParameter undocumented
type CallAnswerRequestParameter struct {
	// CallbackURI undocumented
	CallbackURI *string `json:"callbackUri,omitempty"`
	// MediaConfig undocumented
	MediaConfig *MediaConfig `json:"mediaConfig,omitempty"`
	// AcceptedModalities undocumented
	AcceptedModalities []Modality `json:"acceptedModalities,omitempty"`
	// ParticipantCapacity undocumented
	ParticipantCapacity *int `json:"participantCapacity,omitempty"`
	// CallOptions undocumented
	CallOptions *IncomingCallOptions `json:"callOptions,omitempty"`
}

// CallCancelMediaProcessingRequestParameter undocumented
type CallCancelMediaProcessingRequestParameter struct {
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

// CallChangeScreenSharingRoleRequestParameter undocumented
type CallChangeScreenSharingRoleRequestParameter struct {
	// Role undocumented
	Role *ScreenSharingRole `json:"role,omitempty"`
}

// CallKeepAliveRequestParameter undocumented
type CallKeepAliveRequestParameter struct {
}

// CallMuteRequestParameter undocumented
type CallMuteRequestParameter struct {
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

// CallPlayPromptRequestParameter undocumented
type CallPlayPromptRequestParameter struct {
	// Prompts undocumented
	Prompts []Prompt `json:"prompts,omitempty"`
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

// CallRecordResponseRequestParameter undocumented
type CallRecordResponseRequestParameter struct {
	// Prompts undocumented
	Prompts []Prompt `json:"prompts,omitempty"`
	// BargeInAllowed undocumented
	BargeInAllowed *bool `json:"bargeInAllowed,omitempty"`
	// InitialSilenceTimeoutInSeconds undocumented
	InitialSilenceTimeoutInSeconds *int `json:"initialSilenceTimeoutInSeconds,omitempty"`
	// MaxSilenceTimeoutInSeconds undocumented
	MaxSilenceTimeoutInSeconds *int `json:"maxSilenceTimeoutInSeconds,omitempty"`
	// MaxRecordDurationInSeconds undocumented
	MaxRecordDurationInSeconds *int `json:"maxRecordDurationInSeconds,omitempty"`
	// PlayBeep undocumented
	PlayBeep *bool `json:"playBeep,omitempty"`
	// StopTones undocumented
	StopTones []string `json:"stopTones,omitempty"`
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

// CallRejectRequestParameter undocumented
type CallRejectRequestParameter struct {
	// Reason undocumented
	Reason *RejectReason `json:"reason,omitempty"`
	// CallbackURI undocumented
	CallbackURI *string `json:"callbackUri,omitempty"`
}

// CallSubscribeToToneRequestParameter undocumented
type CallSubscribeToToneRequestParameter struct {
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

// CallTransferRequestParameter undocumented
type CallTransferRequestParameter struct {
	// TransferTarget undocumented
	TransferTarget *InvitationParticipantInfo `json:"transferTarget,omitempty"`
	// Transferee undocumented
	Transferee *ParticipantInfo `json:"transferee,omitempty"`
}

// CallUnmuteRequestParameter undocumented
type CallUnmuteRequestParameter struct {
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

// CallUpdateRecordingStatusRequestParameter undocumented
type CallUpdateRecordingStatusRequestParameter struct {
	// Status undocumented
	Status *RecordingStatus `json:"status,omitempty"`
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

// AudioRoutingGroups returns request builder for AudioRoutingGroup collection
func (b *CallRequestBuilder) AudioRoutingGroups() *CallAudioRoutingGroupsCollectionRequestBuilder {
	bb := &CallAudioRoutingGroupsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/audioRoutingGroups"
	return bb
}

// CallAudioRoutingGroupsCollectionRequestBuilder is request builder for AudioRoutingGroup collection
type CallAudioRoutingGroupsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AudioRoutingGroup collection
func (b *CallAudioRoutingGroupsCollectionRequestBuilder) Request() *CallAudioRoutingGroupsCollectionRequest {
	return &CallAudioRoutingGroupsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AudioRoutingGroup item
func (b *CallAudioRoutingGroupsCollectionRequestBuilder) ID(id string) *AudioRoutingGroupRequestBuilder {
	bb := &AudioRoutingGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CallAudioRoutingGroupsCollectionRequest is request for AudioRoutingGroup collection
type CallAudioRoutingGroupsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AudioRoutingGroup collection
func (r *CallAudioRoutingGroupsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AudioRoutingGroup, error) {
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
	var values []AudioRoutingGroup
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
			value  []AudioRoutingGroup
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

// GetN performs GET request for AudioRoutingGroup collection, max N pages
func (r *CallAudioRoutingGroupsCollectionRequest) GetN(ctx context.Context, n int) ([]AudioRoutingGroup, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AudioRoutingGroup collection
func (r *CallAudioRoutingGroupsCollectionRequest) Get(ctx context.Context) ([]AudioRoutingGroup, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AudioRoutingGroup collection
func (r *CallAudioRoutingGroupsCollectionRequest) Add(ctx context.Context, reqObj *AudioRoutingGroup) (resObj *AudioRoutingGroup, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ContentSharingSessions returns request builder for ContentSharingSession collection
func (b *CallRequestBuilder) ContentSharingSessions() *CallContentSharingSessionsCollectionRequestBuilder {
	bb := &CallContentSharingSessionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/contentSharingSessions"
	return bb
}

// CallContentSharingSessionsCollectionRequestBuilder is request builder for ContentSharingSession collection
type CallContentSharingSessionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ContentSharingSession collection
func (b *CallContentSharingSessionsCollectionRequestBuilder) Request() *CallContentSharingSessionsCollectionRequest {
	return &CallContentSharingSessionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ContentSharingSession item
func (b *CallContentSharingSessionsCollectionRequestBuilder) ID(id string) *ContentSharingSessionRequestBuilder {
	bb := &ContentSharingSessionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CallContentSharingSessionsCollectionRequest is request for ContentSharingSession collection
type CallContentSharingSessionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ContentSharingSession collection
func (r *CallContentSharingSessionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ContentSharingSession, error) {
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
	var values []ContentSharingSession
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
			value  []ContentSharingSession
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

// GetN performs GET request for ContentSharingSession collection, max N pages
func (r *CallContentSharingSessionsCollectionRequest) GetN(ctx context.Context, n int) ([]ContentSharingSession, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ContentSharingSession collection
func (r *CallContentSharingSessionsCollectionRequest) Get(ctx context.Context) ([]ContentSharingSession, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ContentSharingSession collection
func (r *CallContentSharingSessionsCollectionRequest) Add(ctx context.Context, reqObj *ContentSharingSession) (resObj *ContentSharingSession, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Operations returns request builder for CommsOperation collection
func (b *CallRequestBuilder) Operations() *CallOperationsCollectionRequestBuilder {
	bb := &CallOperationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/operations"
	return bb
}

// CallOperationsCollectionRequestBuilder is request builder for CommsOperation collection
type CallOperationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for CommsOperation collection
func (b *CallOperationsCollectionRequestBuilder) Request() *CallOperationsCollectionRequest {
	return &CallOperationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for CommsOperation item
func (b *CallOperationsCollectionRequestBuilder) ID(id string) *CommsOperationRequestBuilder {
	bb := &CommsOperationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CallOperationsCollectionRequest is request for CommsOperation collection
type CallOperationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for CommsOperation collection
func (r *CallOperationsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]CommsOperation, error) {
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
	var values []CommsOperation
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
			value  []CommsOperation
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

// GetN performs GET request for CommsOperation collection, max N pages
func (r *CallOperationsCollectionRequest) GetN(ctx context.Context, n int) ([]CommsOperation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for CommsOperation collection
func (r *CallOperationsCollectionRequest) Get(ctx context.Context) ([]CommsOperation, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for CommsOperation collection
func (r *CallOperationsCollectionRequest) Add(ctx context.Context, reqObj *CommsOperation) (resObj *CommsOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Participants returns request builder for Participant collection
func (b *CallRequestBuilder) Participants() *CallParticipantsCollectionRequestBuilder {
	bb := &CallParticipantsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/participants"
	return bb
}

// CallParticipantsCollectionRequestBuilder is request builder for Participant collection
type CallParticipantsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Participant collection
func (b *CallParticipantsCollectionRequestBuilder) Request() *CallParticipantsCollectionRequest {
	return &CallParticipantsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Participant item
func (b *CallParticipantsCollectionRequestBuilder) ID(id string) *ParticipantRequestBuilder {
	bb := &ParticipantRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CallParticipantsCollectionRequest is request for Participant collection
type CallParticipantsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Participant collection
func (r *CallParticipantsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Participant, error) {
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
	var values []Participant
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
			value  []Participant
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

// GetN performs GET request for Participant collection, max N pages
func (r *CallParticipantsCollectionRequest) GetN(ctx context.Context, n int) ([]Participant, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Participant collection
func (r *CallParticipantsCollectionRequest) Get(ctx context.Context) ([]Participant, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Participant collection
func (r *CallParticipantsCollectionRequest) Add(ctx context.Context, reqObj *Participant) (resObj *Participant, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *CallRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Sessions returns request builder for CallRecordsSession collection
func (b *CallRecordsCallRecordRequestBuilder) Sessions() *CallRecordsCallRecordSessionsCollectionRequestBuilder {
	bb := &CallRecordsCallRecordSessionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sessions"
	return bb
}

// CallRecordsCallRecordSessionsCollectionRequestBuilder is request builder for CallRecordsSession collection
type CallRecordsCallRecordSessionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for CallRecordsSession collection
func (b *CallRecordsCallRecordSessionsCollectionRequestBuilder) Request() *CallRecordsCallRecordSessionsCollectionRequest {
	return &CallRecordsCallRecordSessionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for CallRecordsSession item
func (b *CallRecordsCallRecordSessionsCollectionRequestBuilder) ID(id string) *CallRecordsSessionRequestBuilder {
	bb := &CallRecordsSessionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CallRecordsCallRecordSessionsCollectionRequest is request for CallRecordsSession collection
type CallRecordsCallRecordSessionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for CallRecordsSession collection
func (r *CallRecordsCallRecordSessionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]CallRecordsSession, error) {
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
	var values []CallRecordsSession
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
			value  []CallRecordsSession
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

// GetN performs GET request for CallRecordsSession collection, max N pages
func (r *CallRecordsCallRecordSessionsCollectionRequest) GetN(ctx context.Context, n int) ([]CallRecordsSession, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for CallRecordsSession collection
func (r *CallRecordsCallRecordSessionsCollectionRequest) Get(ctx context.Context) ([]CallRecordsSession, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for CallRecordsSession collection
func (r *CallRecordsCallRecordSessionsCollectionRequest) Add(ctx context.Context, reqObj *CallRecordsSession) (resObj *CallRecordsSession, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Segments returns request builder for CallRecordsSegment collection
func (b *CallRecordsSessionRequestBuilder) Segments() *CallRecordsSessionSegmentsCollectionRequestBuilder {
	bb := &CallRecordsSessionSegmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/segments"
	return bb
}

// CallRecordsSessionSegmentsCollectionRequestBuilder is request builder for CallRecordsSegment collection
type CallRecordsSessionSegmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for CallRecordsSegment collection
func (b *CallRecordsSessionSegmentsCollectionRequestBuilder) Request() *CallRecordsSessionSegmentsCollectionRequest {
	return &CallRecordsSessionSegmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for CallRecordsSegment item
func (b *CallRecordsSessionSegmentsCollectionRequestBuilder) ID(id string) *CallRecordsSegmentRequestBuilder {
	bb := &CallRecordsSegmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CallRecordsSessionSegmentsCollectionRequest is request for CallRecordsSegment collection
type CallRecordsSessionSegmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for CallRecordsSegment collection
func (r *CallRecordsSessionSegmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]CallRecordsSegment, error) {
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
	var values []CallRecordsSegment
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
			value  []CallRecordsSegment
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

// GetN performs GET request for CallRecordsSegment collection, max N pages
func (r *CallRecordsSessionSegmentsCollectionRequest) GetN(ctx context.Context, n int) ([]CallRecordsSegment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for CallRecordsSegment collection
func (r *CallRecordsSessionSegmentsCollectionRequest) Get(ctx context.Context) ([]CallRecordsSegment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for CallRecordsSegment collection
func (r *CallRecordsSessionSegmentsCollectionRequest) Add(ctx context.Context, reqObj *CallRecordsSegment) (resObj *CallRecordsSegment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
