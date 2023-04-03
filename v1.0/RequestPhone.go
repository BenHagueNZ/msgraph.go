// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// PhoneRequestBuilder is request builder for Phone
type PhoneRequestBuilder struct{ BaseRequestBuilder }

// Request returns PhoneRequest
func (b *PhoneRequestBuilder) Request() *PhoneRequest {
	return &PhoneRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PhoneRequest is request for Phone
type PhoneRequest struct{ BaseRequest }

// Get performs GET request for Phone
func (r *PhoneRequest) Get(ctx context.Context) (resObj *Phone, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Phone
func (r *PhoneRequest) Update(ctx context.Context, reqObj *Phone) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Phone
func (r *PhoneRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PhoneAuthenticationMethodRequestBuilder is request builder for PhoneAuthenticationMethod
type PhoneAuthenticationMethodRequestBuilder struct{ BaseRequestBuilder }

// Request returns PhoneAuthenticationMethodRequest
func (b *PhoneAuthenticationMethodRequestBuilder) Request() *PhoneAuthenticationMethodRequest {
	return &PhoneAuthenticationMethodRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PhoneAuthenticationMethodRequest is request for PhoneAuthenticationMethod
type PhoneAuthenticationMethodRequest struct{ BaseRequest }

// Get performs GET request for PhoneAuthenticationMethod
func (r *PhoneAuthenticationMethodRequest) Get(ctx context.Context) (resObj *PhoneAuthenticationMethod, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PhoneAuthenticationMethod
func (r *PhoneAuthenticationMethodRequest) Update(ctx context.Context, reqObj *PhoneAuthenticationMethod) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PhoneAuthenticationMethod
func (r *PhoneAuthenticationMethodRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type PhoneAuthenticationMethodDisableSmsSignInRequestBuilder struct{ BaseRequestBuilder }

// DisableSmsSignIn action undocumentedrav
func (b *PhoneAuthenticationMethodRequestBuilder) DisableSmsSignIn(reqObj *PhoneAuthenticationMethodDisableSmsSignInRequestParameter) *PhoneAuthenticationMethodDisableSmsSignInRequestBuilder {
	bb := &PhoneAuthenticationMethodDisableSmsSignInRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/DisableSmsSignIn"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type PhoneAuthenticationMethodDisableSmsSignInRequest struct{ BaseRequest }

func (b *PhoneAuthenticationMethodDisableSmsSignInRequestBuilder) Request() *PhoneAuthenticationMethodDisableSmsSignInRequest {
	return &PhoneAuthenticationMethodDisableSmsSignInRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *PhoneAuthenticationMethodDisableSmsSignInRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type PhoneAuthenticationMethodEnableSmsSignInRequestBuilder struct{ BaseRequestBuilder }

// EnableSmsSignIn action undocumentedrav
func (b *PhoneAuthenticationMethodRequestBuilder) EnableSmsSignIn(reqObj *PhoneAuthenticationMethodEnableSmsSignInRequestParameter) *PhoneAuthenticationMethodEnableSmsSignInRequestBuilder {
	bb := &PhoneAuthenticationMethodEnableSmsSignInRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/EnableSmsSignIn"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type PhoneAuthenticationMethodEnableSmsSignInRequest struct{ BaseRequest }

func (b *PhoneAuthenticationMethodEnableSmsSignInRequestBuilder) Request() *PhoneAuthenticationMethodEnableSmsSignInRequest {
	return &PhoneAuthenticationMethodEnableSmsSignInRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *PhoneAuthenticationMethodEnableSmsSignInRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
