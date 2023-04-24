// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ApplicationRequestBuilder is request builder for Application
type ApplicationRequestBuilder struct{ BaseRequestBuilder }

// Request returns ApplicationRequest
func (b *ApplicationRequestBuilder) Request() *ApplicationRequest {
	return &ApplicationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ApplicationRequest is request for Application
type ApplicationRequest struct{ BaseRequest }

// Get performs GET request for Application
func (r *ApplicationRequest) Get(ctx context.Context) (resObj *Application, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Application
func (r *ApplicationRequest) Update(ctx context.Context, reqObj *Application) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Application
func (r *ApplicationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ApplicationEnforcedRestrictionsSessionControlRequestBuilder is request builder for ApplicationEnforcedRestrictionsSessionControl
type ApplicationEnforcedRestrictionsSessionControlRequestBuilder struct{ BaseRequestBuilder }

// Request returns ApplicationEnforcedRestrictionsSessionControlRequest
func (b *ApplicationEnforcedRestrictionsSessionControlRequestBuilder) Request() *ApplicationEnforcedRestrictionsSessionControlRequest {
	return &ApplicationEnforcedRestrictionsSessionControlRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ApplicationEnforcedRestrictionsSessionControlRequest is request for ApplicationEnforcedRestrictionsSessionControl
type ApplicationEnforcedRestrictionsSessionControlRequest struct{ BaseRequest }

// Get performs GET request for ApplicationEnforcedRestrictionsSessionControl
func (r *ApplicationEnforcedRestrictionsSessionControlRequest) Get(ctx context.Context) (resObj *ApplicationEnforcedRestrictionsSessionControl, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ApplicationEnforcedRestrictionsSessionControl
func (r *ApplicationEnforcedRestrictionsSessionControlRequest) Update(ctx context.Context, reqObj *ApplicationEnforcedRestrictionsSessionControl) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ApplicationEnforcedRestrictionsSessionControl
func (r *ApplicationEnforcedRestrictionsSessionControlRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ApplicationSegmentRequestBuilder is request builder for ApplicationSegment
type ApplicationSegmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns ApplicationSegmentRequest
func (b *ApplicationSegmentRequestBuilder) Request() *ApplicationSegmentRequest {
	return &ApplicationSegmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ApplicationSegmentRequest is request for ApplicationSegment
type ApplicationSegmentRequest struct{ BaseRequest }

// Get performs GET request for ApplicationSegment
func (r *ApplicationSegmentRequest) Get(ctx context.Context) (resObj *ApplicationSegment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ApplicationSegment
func (r *ApplicationSegmentRequest) Update(ctx context.Context, reqObj *ApplicationSegment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ApplicationSegment
func (r *ApplicationSegmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ApplicationServicePrincipalRequestBuilder is request builder for ApplicationServicePrincipal
type ApplicationServicePrincipalRequestBuilder struct{ BaseRequestBuilder }

// Request returns ApplicationServicePrincipalRequest
func (b *ApplicationServicePrincipalRequestBuilder) Request() *ApplicationServicePrincipalRequest {
	return &ApplicationServicePrincipalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ApplicationServicePrincipalRequest is request for ApplicationServicePrincipal
type ApplicationServicePrincipalRequest struct{ BaseRequest }

// Get performs GET request for ApplicationServicePrincipal
func (r *ApplicationServicePrincipalRequest) Get(ctx context.Context) (resObj *ApplicationServicePrincipal, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ApplicationServicePrincipal
func (r *ApplicationServicePrincipalRequest) Update(ctx context.Context, reqObj *ApplicationServicePrincipal) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ApplicationServicePrincipal
func (r *ApplicationServicePrincipalRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ApplicationSignInDetailedSummaryRequestBuilder is request builder for ApplicationSignInDetailedSummary
type ApplicationSignInDetailedSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns ApplicationSignInDetailedSummaryRequest
func (b *ApplicationSignInDetailedSummaryRequestBuilder) Request() *ApplicationSignInDetailedSummaryRequest {
	return &ApplicationSignInDetailedSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ApplicationSignInDetailedSummaryRequest is request for ApplicationSignInDetailedSummary
type ApplicationSignInDetailedSummaryRequest struct{ BaseRequest }

// Get performs GET request for ApplicationSignInDetailedSummary
func (r *ApplicationSignInDetailedSummaryRequest) Get(ctx context.Context) (resObj *ApplicationSignInDetailedSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ApplicationSignInDetailedSummary
func (r *ApplicationSignInDetailedSummaryRequest) Update(ctx context.Context, reqObj *ApplicationSignInDetailedSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ApplicationSignInDetailedSummary
func (r *ApplicationSignInDetailedSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ApplicationSignInSummaryRequestBuilder is request builder for ApplicationSignInSummary
type ApplicationSignInSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns ApplicationSignInSummaryRequest
func (b *ApplicationSignInSummaryRequestBuilder) Request() *ApplicationSignInSummaryRequest {
	return &ApplicationSignInSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ApplicationSignInSummaryRequest is request for ApplicationSignInSummary
type ApplicationSignInSummaryRequest struct{ BaseRequest }

// Get performs GET request for ApplicationSignInSummary
func (r *ApplicationSignInSummaryRequest) Get(ctx context.Context) (resObj *ApplicationSignInSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ApplicationSignInSummary
func (r *ApplicationSignInSummaryRequest) Update(ctx context.Context, reqObj *ApplicationSignInSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ApplicationSignInSummary
func (r *ApplicationSignInSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ApplicationTemplateRequestBuilder is request builder for ApplicationTemplate
type ApplicationTemplateRequestBuilder struct{ BaseRequestBuilder }

// Request returns ApplicationTemplateRequest
func (b *ApplicationTemplateRequestBuilder) Request() *ApplicationTemplateRequest {
	return &ApplicationTemplateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ApplicationTemplateRequest is request for ApplicationTemplate
type ApplicationTemplateRequest struct{ BaseRequest }

// Get performs GET request for ApplicationTemplate
func (r *ApplicationTemplateRequest) Get(ctx context.Context) (resObj *ApplicationTemplate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ApplicationTemplate
func (r *ApplicationTemplateRequest) Update(ctx context.Context, reqObj *ApplicationTemplate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ApplicationTemplate
func (r *ApplicationTemplateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type ApplicationSetVerifiedPublisherRequestBuilder struct{ BaseRequestBuilder }

// SetVerifiedPublisher action undocumentedrav
func (b *ApplicationRequestBuilder) SetVerifiedPublisher(reqObj *ApplicationSetVerifiedPublisherRequestParameter) *ApplicationSetVerifiedPublisherRequestBuilder {
	bb := &ApplicationSetVerifiedPublisherRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/SetVerifiedPublisher"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ApplicationSetVerifiedPublisherRequest struct{ BaseRequest }

func (b *ApplicationSetVerifiedPublisherRequestBuilder) Request() *ApplicationSetVerifiedPublisherRequest {
	return &ApplicationSetVerifiedPublisherRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ApplicationSetVerifiedPublisherRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type ApplicationUnsetVerifiedPublisherRequestBuilder struct{ BaseRequestBuilder }

// UnsetVerifiedPublisher action undocumentedrav
func (b *ApplicationRequestBuilder) UnsetVerifiedPublisher(reqObj *ApplicationUnsetVerifiedPublisherRequestParameter) *ApplicationUnsetVerifiedPublisherRequestBuilder {
	bb := &ApplicationUnsetVerifiedPublisherRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/UnsetVerifiedPublisher"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ApplicationUnsetVerifiedPublisherRequest struct{ BaseRequest }

func (b *ApplicationUnsetVerifiedPublisherRequestBuilder) Request() *ApplicationUnsetVerifiedPublisherRequest {
	return &ApplicationUnsetVerifiedPublisherRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ApplicationUnsetVerifiedPublisherRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type ApplicationAddKeyRequestBuilder struct{ BaseRequestBuilder }

// AddKey action undocumentedras
func (b *ApplicationRequestBuilder) AddKey(reqObj *ApplicationAddKeyRequestParameter) *ApplicationAddKeyRequestBuilder {
	bb := &ApplicationAddKeyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/AddKey"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ApplicationAddKeyRequest struct{ BaseRequest }

func (b *ApplicationAddKeyRequestBuilder) Request() *ApplicationAddKeyRequest {
	return &ApplicationAddKeyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ApplicationAddKeyRequest) Post(ctx context.Context) (resObj *KeyCredential, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type ApplicationAddPasswordRequestBuilder struct{ BaseRequestBuilder }

// AddPassword action undocumentedras
func (b *ApplicationRequestBuilder) AddPassword(reqObj *ApplicationAddPasswordRequestParameter) *ApplicationAddPasswordRequestBuilder {
	bb := &ApplicationAddPasswordRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/AddPassword"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ApplicationAddPasswordRequest struct{ BaseRequest }

func (b *ApplicationAddPasswordRequestBuilder) Request() *ApplicationAddPasswordRequest {
	return &ApplicationAddPasswordRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ApplicationAddPasswordRequest) Post(ctx context.Context) (resObj *PasswordCredential, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type ApplicationRemoveKeyRequestBuilder struct{ BaseRequestBuilder }

// RemoveKey action undocumentedrav
func (b *ApplicationRequestBuilder) RemoveKey(reqObj *ApplicationRemoveKeyRequestParameter) *ApplicationRemoveKeyRequestBuilder {
	bb := &ApplicationRemoveKeyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/RemoveKey"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ApplicationRemoveKeyRequest struct{ BaseRequest }

func (b *ApplicationRemoveKeyRequestBuilder) Request() *ApplicationRemoveKeyRequest {
	return &ApplicationRemoveKeyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ApplicationRemoveKeyRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type ApplicationRemovePasswordRequestBuilder struct{ BaseRequestBuilder }

// RemovePassword action undocumentedrav
func (b *ApplicationRequestBuilder) RemovePassword(reqObj *ApplicationRemovePasswordRequestParameter) *ApplicationRemovePasswordRequestBuilder {
	bb := &ApplicationRemovePasswordRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/RemovePassword"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ApplicationRemovePasswordRequest struct{ BaseRequest }

func (b *ApplicationRemovePasswordRequestBuilder) Request() *ApplicationRemovePasswordRequest {
	return &ApplicationRemovePasswordRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ApplicationRemovePasswordRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type ApplicationTemplateInstantiateRequestBuilder struct{ BaseRequestBuilder }

// Instantiate action undocumentedras
func (b *ApplicationTemplateRequestBuilder) Instantiate(reqObj *ApplicationTemplateInstantiateRequestParameter) *ApplicationTemplateInstantiateRequestBuilder {
	bb := &ApplicationTemplateInstantiateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Instantiate"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ApplicationTemplateInstantiateRequest struct{ BaseRequest }

func (b *ApplicationTemplateInstantiateRequestBuilder) Request() *ApplicationTemplateInstantiateRequest {
	return &ApplicationTemplateInstantiateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ApplicationTemplateInstantiateRequest) Post(ctx context.Context) (resObj *ApplicationServicePrincipal, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
