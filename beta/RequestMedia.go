// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// MediaRequestBuilder is request builder for Media
type MediaRequestBuilder struct{ BaseRequestBuilder }

// Request returns MediaRequest
func (b *MediaRequestBuilder) Request() *MediaRequest {
	return &MediaRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MediaRequest is request for Media
type MediaRequest struct{ BaseRequest }

// Get performs GET request for Media
func (r *MediaRequest) Get(ctx context.Context) (resObj *Media, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Media
func (r *MediaRequest) Update(ctx context.Context, reqObj *Media) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Media
func (r *MediaRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MediaConfigRequestBuilder is request builder for MediaConfig
type MediaConfigRequestBuilder struct{ BaseRequestBuilder }

// Request returns MediaConfigRequest
func (b *MediaConfigRequestBuilder) Request() *MediaConfigRequest {
	return &MediaConfigRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MediaConfigRequest is request for MediaConfig
type MediaConfigRequest struct{ BaseRequest }

// Get performs GET request for MediaConfig
func (r *MediaConfigRequest) Get(ctx context.Context) (resObj *MediaConfig, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MediaConfig
func (r *MediaConfigRequest) Update(ctx context.Context, reqObj *MediaConfig) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MediaConfig
func (r *MediaConfigRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MediaContentRatingAustraliaRequestBuilder is request builder for MediaContentRatingAustralia
type MediaContentRatingAustraliaRequestBuilder struct{ BaseRequestBuilder }

// Request returns MediaContentRatingAustraliaRequest
func (b *MediaContentRatingAustraliaRequestBuilder) Request() *MediaContentRatingAustraliaRequest {
	return &MediaContentRatingAustraliaRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MediaContentRatingAustraliaRequest is request for MediaContentRatingAustralia
type MediaContentRatingAustraliaRequest struct{ BaseRequest }

// Get performs GET request for MediaContentRatingAustralia
func (r *MediaContentRatingAustraliaRequest) Get(ctx context.Context) (resObj *MediaContentRatingAustralia, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MediaContentRatingAustralia
func (r *MediaContentRatingAustraliaRequest) Update(ctx context.Context, reqObj *MediaContentRatingAustralia) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MediaContentRatingAustralia
func (r *MediaContentRatingAustraliaRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MediaContentRatingCanadaRequestBuilder is request builder for MediaContentRatingCanada
type MediaContentRatingCanadaRequestBuilder struct{ BaseRequestBuilder }

// Request returns MediaContentRatingCanadaRequest
func (b *MediaContentRatingCanadaRequestBuilder) Request() *MediaContentRatingCanadaRequest {
	return &MediaContentRatingCanadaRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MediaContentRatingCanadaRequest is request for MediaContentRatingCanada
type MediaContentRatingCanadaRequest struct{ BaseRequest }

// Get performs GET request for MediaContentRatingCanada
func (r *MediaContentRatingCanadaRequest) Get(ctx context.Context) (resObj *MediaContentRatingCanada, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MediaContentRatingCanada
func (r *MediaContentRatingCanadaRequest) Update(ctx context.Context, reqObj *MediaContentRatingCanada) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MediaContentRatingCanada
func (r *MediaContentRatingCanadaRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MediaContentRatingFranceRequestBuilder is request builder for MediaContentRatingFrance
type MediaContentRatingFranceRequestBuilder struct{ BaseRequestBuilder }

// Request returns MediaContentRatingFranceRequest
func (b *MediaContentRatingFranceRequestBuilder) Request() *MediaContentRatingFranceRequest {
	return &MediaContentRatingFranceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MediaContentRatingFranceRequest is request for MediaContentRatingFrance
type MediaContentRatingFranceRequest struct{ BaseRequest }

// Get performs GET request for MediaContentRatingFrance
func (r *MediaContentRatingFranceRequest) Get(ctx context.Context) (resObj *MediaContentRatingFrance, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MediaContentRatingFrance
func (r *MediaContentRatingFranceRequest) Update(ctx context.Context, reqObj *MediaContentRatingFrance) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MediaContentRatingFrance
func (r *MediaContentRatingFranceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MediaContentRatingGermanyRequestBuilder is request builder for MediaContentRatingGermany
type MediaContentRatingGermanyRequestBuilder struct{ BaseRequestBuilder }

// Request returns MediaContentRatingGermanyRequest
func (b *MediaContentRatingGermanyRequestBuilder) Request() *MediaContentRatingGermanyRequest {
	return &MediaContentRatingGermanyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MediaContentRatingGermanyRequest is request for MediaContentRatingGermany
type MediaContentRatingGermanyRequest struct{ BaseRequest }

// Get performs GET request for MediaContentRatingGermany
func (r *MediaContentRatingGermanyRequest) Get(ctx context.Context) (resObj *MediaContentRatingGermany, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MediaContentRatingGermany
func (r *MediaContentRatingGermanyRequest) Update(ctx context.Context, reqObj *MediaContentRatingGermany) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MediaContentRatingGermany
func (r *MediaContentRatingGermanyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MediaContentRatingIrelandRequestBuilder is request builder for MediaContentRatingIreland
type MediaContentRatingIrelandRequestBuilder struct{ BaseRequestBuilder }

// Request returns MediaContentRatingIrelandRequest
func (b *MediaContentRatingIrelandRequestBuilder) Request() *MediaContentRatingIrelandRequest {
	return &MediaContentRatingIrelandRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MediaContentRatingIrelandRequest is request for MediaContentRatingIreland
type MediaContentRatingIrelandRequest struct{ BaseRequest }

// Get performs GET request for MediaContentRatingIreland
func (r *MediaContentRatingIrelandRequest) Get(ctx context.Context) (resObj *MediaContentRatingIreland, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MediaContentRatingIreland
func (r *MediaContentRatingIrelandRequest) Update(ctx context.Context, reqObj *MediaContentRatingIreland) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MediaContentRatingIreland
func (r *MediaContentRatingIrelandRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MediaContentRatingJapanRequestBuilder is request builder for MediaContentRatingJapan
type MediaContentRatingJapanRequestBuilder struct{ BaseRequestBuilder }

// Request returns MediaContentRatingJapanRequest
func (b *MediaContentRatingJapanRequestBuilder) Request() *MediaContentRatingJapanRequest {
	return &MediaContentRatingJapanRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MediaContentRatingJapanRequest is request for MediaContentRatingJapan
type MediaContentRatingJapanRequest struct{ BaseRequest }

// Get performs GET request for MediaContentRatingJapan
func (r *MediaContentRatingJapanRequest) Get(ctx context.Context) (resObj *MediaContentRatingJapan, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MediaContentRatingJapan
func (r *MediaContentRatingJapanRequest) Update(ctx context.Context, reqObj *MediaContentRatingJapan) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MediaContentRatingJapan
func (r *MediaContentRatingJapanRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MediaContentRatingNewZealandRequestBuilder is request builder for MediaContentRatingNewZealand
type MediaContentRatingNewZealandRequestBuilder struct{ BaseRequestBuilder }

// Request returns MediaContentRatingNewZealandRequest
func (b *MediaContentRatingNewZealandRequestBuilder) Request() *MediaContentRatingNewZealandRequest {
	return &MediaContentRatingNewZealandRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MediaContentRatingNewZealandRequest is request for MediaContentRatingNewZealand
type MediaContentRatingNewZealandRequest struct{ BaseRequest }

// Get performs GET request for MediaContentRatingNewZealand
func (r *MediaContentRatingNewZealandRequest) Get(ctx context.Context) (resObj *MediaContentRatingNewZealand, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MediaContentRatingNewZealand
func (r *MediaContentRatingNewZealandRequest) Update(ctx context.Context, reqObj *MediaContentRatingNewZealand) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MediaContentRatingNewZealand
func (r *MediaContentRatingNewZealandRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MediaContentRatingUnitedKingdomRequestBuilder is request builder for MediaContentRatingUnitedKingdom
type MediaContentRatingUnitedKingdomRequestBuilder struct{ BaseRequestBuilder }

// Request returns MediaContentRatingUnitedKingdomRequest
func (b *MediaContentRatingUnitedKingdomRequestBuilder) Request() *MediaContentRatingUnitedKingdomRequest {
	return &MediaContentRatingUnitedKingdomRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MediaContentRatingUnitedKingdomRequest is request for MediaContentRatingUnitedKingdom
type MediaContentRatingUnitedKingdomRequest struct{ BaseRequest }

// Get performs GET request for MediaContentRatingUnitedKingdom
func (r *MediaContentRatingUnitedKingdomRequest) Get(ctx context.Context) (resObj *MediaContentRatingUnitedKingdom, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MediaContentRatingUnitedKingdom
func (r *MediaContentRatingUnitedKingdomRequest) Update(ctx context.Context, reqObj *MediaContentRatingUnitedKingdom) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MediaContentRatingUnitedKingdom
func (r *MediaContentRatingUnitedKingdomRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MediaContentRatingUnitedStatesRequestBuilder is request builder for MediaContentRatingUnitedStates
type MediaContentRatingUnitedStatesRequestBuilder struct{ BaseRequestBuilder }

// Request returns MediaContentRatingUnitedStatesRequest
func (b *MediaContentRatingUnitedStatesRequestBuilder) Request() *MediaContentRatingUnitedStatesRequest {
	return &MediaContentRatingUnitedStatesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MediaContentRatingUnitedStatesRequest is request for MediaContentRatingUnitedStates
type MediaContentRatingUnitedStatesRequest struct{ BaseRequest }

// Get performs GET request for MediaContentRatingUnitedStates
func (r *MediaContentRatingUnitedStatesRequest) Get(ctx context.Context) (resObj *MediaContentRatingUnitedStates, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MediaContentRatingUnitedStates
func (r *MediaContentRatingUnitedStatesRequest) Update(ctx context.Context, reqObj *MediaContentRatingUnitedStates) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MediaContentRatingUnitedStates
func (r *MediaContentRatingUnitedStatesRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MediaInfoRequestBuilder is request builder for MediaInfo
type MediaInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns MediaInfoRequest
func (b *MediaInfoRequestBuilder) Request() *MediaInfoRequest {
	return &MediaInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MediaInfoRequest is request for MediaInfo
type MediaInfoRequest struct{ BaseRequest }

// Get performs GET request for MediaInfo
func (r *MediaInfoRequest) Get(ctx context.Context) (resObj *MediaInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MediaInfo
func (r *MediaInfoRequest) Update(ctx context.Context, reqObj *MediaInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MediaInfo
func (r *MediaInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MediaPromptRequestBuilder is request builder for MediaPrompt
type MediaPromptRequestBuilder struct{ BaseRequestBuilder }

// Request returns MediaPromptRequest
func (b *MediaPromptRequestBuilder) Request() *MediaPromptRequest {
	return &MediaPromptRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MediaPromptRequest is request for MediaPrompt
type MediaPromptRequest struct{ BaseRequest }

// Get performs GET request for MediaPrompt
func (r *MediaPromptRequest) Get(ctx context.Context) (resObj *MediaPrompt, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MediaPrompt
func (r *MediaPromptRequest) Update(ctx context.Context, reqObj *MediaPrompt) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MediaPrompt
func (r *MediaPromptRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MediaSourceRequestBuilder is request builder for MediaSource
type MediaSourceRequestBuilder struct{ BaseRequestBuilder }

// Request returns MediaSourceRequest
func (b *MediaSourceRequestBuilder) Request() *MediaSourceRequest {
	return &MediaSourceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MediaSourceRequest is request for MediaSource
type MediaSourceRequest struct{ BaseRequest }

// Get performs GET request for MediaSource
func (r *MediaSourceRequest) Get(ctx context.Context) (resObj *MediaSource, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MediaSource
func (r *MediaSourceRequest) Update(ctx context.Context, reqObj *MediaSource) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MediaSource
func (r *MediaSourceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MediaStreamRequestBuilder is request builder for MediaStream
type MediaStreamRequestBuilder struct{ BaseRequestBuilder }

// Request returns MediaStreamRequest
func (b *MediaStreamRequestBuilder) Request() *MediaStreamRequest {
	return &MediaStreamRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MediaStreamRequest is request for MediaStream
type MediaStreamRequest struct{ BaseRequest }

// Get performs GET request for MediaStream
func (r *MediaStreamRequest) Get(ctx context.Context) (resObj *MediaStream, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MediaStream
func (r *MediaStreamRequest) Update(ctx context.Context, reqObj *MediaStream) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MediaStream
func (r *MediaStreamRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
