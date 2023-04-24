// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// AccountRequestBuilder is request builder for Account
type AccountRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccountRequest
func (b *AccountRequestBuilder) Request() *AccountRequest {
	return &AccountRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AccountRequest is request for Account
type AccountRequest struct{ BaseRequest }

// Get performs GET request for Account
func (r *AccountRequest) Get(ctx context.Context) (resObj *Account, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Account
func (r *AccountRequest) Update(ctx context.Context, reqObj *Account) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Account
func (r *AccountRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AccountAliasRequestBuilder is request builder for AccountAlias
type AccountAliasRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccountAliasRequest
func (b *AccountAliasRequestBuilder) Request() *AccountAliasRequest {
	return &AccountAliasRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AccountAliasRequest is request for AccountAlias
type AccountAliasRequest struct{ BaseRequest }

// Get performs GET request for AccountAlias
func (r *AccountAliasRequest) Get(ctx context.Context) (resObj *AccountAlias, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AccountAlias
func (r *AccountAliasRequest) Update(ctx context.Context, reqObj *AccountAlias) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AccountAlias
func (r *AccountAliasRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AccountTargetContentRequestBuilder is request builder for AccountTargetContent
type AccountTargetContentRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccountTargetContentRequest
func (b *AccountTargetContentRequestBuilder) Request() *AccountTargetContentRequest {
	return &AccountTargetContentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AccountTargetContentRequest is request for AccountTargetContent
type AccountTargetContentRequest struct{ BaseRequest }

// Get performs GET request for AccountTargetContent
func (r *AccountTargetContentRequest) Get(ctx context.Context) (resObj *AccountTargetContent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AccountTargetContent
func (r *AccountTargetContentRequest) Update(ctx context.Context, reqObj *AccountTargetContent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AccountTargetContent
func (r *AccountTargetContentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
