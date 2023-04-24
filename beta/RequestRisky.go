// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// RiskyServicePrincipalRequestBuilder is request builder for RiskyServicePrincipal
type RiskyServicePrincipalRequestBuilder struct{ BaseRequestBuilder }

// Request returns RiskyServicePrincipalRequest
func (b *RiskyServicePrincipalRequestBuilder) Request() *RiskyServicePrincipalRequest {
	return &RiskyServicePrincipalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RiskyServicePrincipalRequest is request for RiskyServicePrincipal
type RiskyServicePrincipalRequest struct{ BaseRequest }

// Get performs GET request for RiskyServicePrincipal
func (r *RiskyServicePrincipalRequest) Get(ctx context.Context) (resObj *RiskyServicePrincipal, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RiskyServicePrincipal
func (r *RiskyServicePrincipalRequest) Update(ctx context.Context, reqObj *RiskyServicePrincipal) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RiskyServicePrincipal
func (r *RiskyServicePrincipalRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RiskyServicePrincipalHistoryItemRequestBuilder is request builder for RiskyServicePrincipalHistoryItem
type RiskyServicePrincipalHistoryItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns RiskyServicePrincipalHistoryItemRequest
func (b *RiskyServicePrincipalHistoryItemRequestBuilder) Request() *RiskyServicePrincipalHistoryItemRequest {
	return &RiskyServicePrincipalHistoryItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RiskyServicePrincipalHistoryItemRequest is request for RiskyServicePrincipalHistoryItem
type RiskyServicePrincipalHistoryItemRequest struct{ BaseRequest }

// Get performs GET request for RiskyServicePrincipalHistoryItem
func (r *RiskyServicePrincipalHistoryItemRequest) Get(ctx context.Context) (resObj *RiskyServicePrincipalHistoryItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RiskyServicePrincipalHistoryItem
func (r *RiskyServicePrincipalHistoryItemRequest) Update(ctx context.Context, reqObj *RiskyServicePrincipalHistoryItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RiskyServicePrincipalHistoryItem
func (r *RiskyServicePrincipalHistoryItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RiskyUserRequestBuilder is request builder for RiskyUser
type RiskyUserRequestBuilder struct{ BaseRequestBuilder }

// Request returns RiskyUserRequest
func (b *RiskyUserRequestBuilder) Request() *RiskyUserRequest {
	return &RiskyUserRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RiskyUserRequest is request for RiskyUser
type RiskyUserRequest struct{ BaseRequest }

// Get performs GET request for RiskyUser
func (r *RiskyUserRequest) Get(ctx context.Context) (resObj *RiskyUser, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RiskyUser
func (r *RiskyUserRequest) Update(ctx context.Context, reqObj *RiskyUser) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RiskyUser
func (r *RiskyUserRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RiskyUserHistoryItemRequestBuilder is request builder for RiskyUserHistoryItem
type RiskyUserHistoryItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns RiskyUserHistoryItemRequest
func (b *RiskyUserHistoryItemRequestBuilder) Request() *RiskyUserHistoryItemRequest {
	return &RiskyUserHistoryItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RiskyUserHistoryItemRequest is request for RiskyUserHistoryItem
type RiskyUserHistoryItemRequest struct{ BaseRequest }

// Get performs GET request for RiskyUserHistoryItem
func (r *RiskyUserHistoryItemRequest) Get(ctx context.Context) (resObj *RiskyUserHistoryItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RiskyUserHistoryItem
func (r *RiskyUserHistoryItemRequest) Update(ctx context.Context, reqObj *RiskyUserHistoryItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RiskyUserHistoryItem
func (r *RiskyUserHistoryItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type RiskyServicePrincipalCollectionDismissRequestBuilder struct{ BaseRequestBuilder }

// Dismiss action undocumentedrav
func (b *IdentityProtectionRootRiskyServicePrincipalsCollectionRequestBuilder) Dismiss(reqObj *RiskyServicePrincipalCollectionDismissRequestParameter) *RiskyServicePrincipalCollectionDismissRequestBuilder {
	bb := &RiskyServicePrincipalCollectionDismissRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Dismiss"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type RiskyServicePrincipalCollectionDismissRequest struct{ BaseRequest }

func (b *RiskyServicePrincipalCollectionDismissRequestBuilder) Request() *RiskyServicePrincipalCollectionDismissRequest {
	return &RiskyServicePrincipalCollectionDismissRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *RiskyServicePrincipalCollectionDismissRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type RiskyServicePrincipalCollectionConfirmCompromisedRequestBuilder struct{ BaseRequestBuilder }

// ConfirmCompromised action undocumentedrav
func (b *IdentityProtectionRootRiskyServicePrincipalsCollectionRequestBuilder) ConfirmCompromised(reqObj *RiskyServicePrincipalCollectionConfirmCompromisedRequestParameter) *RiskyServicePrincipalCollectionConfirmCompromisedRequestBuilder {
	bb := &RiskyServicePrincipalCollectionConfirmCompromisedRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ConfirmCompromised"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type RiskyServicePrincipalCollectionConfirmCompromisedRequest struct{ BaseRequest }

func (b *RiskyServicePrincipalCollectionConfirmCompromisedRequestBuilder) Request() *RiskyServicePrincipalCollectionConfirmCompromisedRequest {
	return &RiskyServicePrincipalCollectionConfirmCompromisedRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *RiskyServicePrincipalCollectionConfirmCompromisedRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type RiskyUserCollectionDismissRequestBuilder struct{ BaseRequestBuilder }

// Dismiss action undocumentedrav
func (b *IdentityProtectionRootRiskyUsersCollectionRequestBuilder) Dismiss(reqObj *RiskyUserCollectionDismissRequestParameter) *RiskyUserCollectionDismissRequestBuilder {
	bb := &RiskyUserCollectionDismissRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Dismiss"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type RiskyUserCollectionDismissRequest struct{ BaseRequest }

func (b *RiskyUserCollectionDismissRequestBuilder) Request() *RiskyUserCollectionDismissRequest {
	return &RiskyUserCollectionDismissRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *RiskyUserCollectionDismissRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type RiskyUserCollectionConfirmCompromisedRequestBuilder struct{ BaseRequestBuilder }

// ConfirmCompromised action undocumentedrav
func (b *IdentityProtectionRootRiskyUsersCollectionRequestBuilder) ConfirmCompromised(reqObj *RiskyUserCollectionConfirmCompromisedRequestParameter) *RiskyUserCollectionConfirmCompromisedRequestBuilder {
	bb := &RiskyUserCollectionConfirmCompromisedRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ConfirmCompromised"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type RiskyUserCollectionConfirmCompromisedRequest struct{ BaseRequest }

func (b *RiskyUserCollectionConfirmCompromisedRequestBuilder) Request() *RiskyUserCollectionConfirmCompromisedRequest {
	return &RiskyUserCollectionConfirmCompromisedRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *RiskyUserCollectionConfirmCompromisedRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
