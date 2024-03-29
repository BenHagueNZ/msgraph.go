// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// RequestObjectRequestBuilder is request builder for RequestObject
type RequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns RequestObjectRequest
func (b *RequestObjectRequestBuilder) Request() *RequestObjectRequest {
	return &RequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RequestObjectRequest is request for RequestObject
type RequestObjectRequest struct{ BaseRequest }

// Get performs GET request for RequestObject
func (r *RequestObjectRequest) Get(ctx context.Context) (resObj *RequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RequestObject
func (r *RequestObjectRequest) Update(ctx context.Context, reqObj *RequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RequestObject
func (r *RequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RequestRemoteHelpSessionAccessResponseRequestBuilder is request builder for RequestRemoteHelpSessionAccessResponse
type RequestRemoteHelpSessionAccessResponseRequestBuilder struct{ BaseRequestBuilder }

// Request returns RequestRemoteHelpSessionAccessResponseRequest
func (b *RequestRemoteHelpSessionAccessResponseRequestBuilder) Request() *RequestRemoteHelpSessionAccessResponseRequest {
	return &RequestRemoteHelpSessionAccessResponseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RequestRemoteHelpSessionAccessResponseRequest is request for RequestRemoteHelpSessionAccessResponse
type RequestRemoteHelpSessionAccessResponseRequest struct{ BaseRequest }

// Get performs GET request for RequestRemoteHelpSessionAccessResponse
func (r *RequestRemoteHelpSessionAccessResponseRequest) Get(ctx context.Context) (resObj *RequestRemoteHelpSessionAccessResponse, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RequestRemoteHelpSessionAccessResponse
func (r *RequestRemoteHelpSessionAccessResponseRequest) Update(ctx context.Context, reqObj *RequestRemoteHelpSessionAccessResponse) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RequestRemoteHelpSessionAccessResponse
func (r *RequestRemoteHelpSessionAccessResponseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RequestScheduleRequestBuilder is request builder for RequestSchedule
type RequestScheduleRequestBuilder struct{ BaseRequestBuilder }

// Request returns RequestScheduleRequest
func (b *RequestScheduleRequestBuilder) Request() *RequestScheduleRequest {
	return &RequestScheduleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RequestScheduleRequest is request for RequestSchedule
type RequestScheduleRequest struct{ BaseRequest }

// Get performs GET request for RequestSchedule
func (r *RequestScheduleRequest) Get(ctx context.Context) (resObj *RequestSchedule, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RequestSchedule
func (r *RequestScheduleRequest) Update(ctx context.Context, reqObj *RequestSchedule) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RequestSchedule
func (r *RequestScheduleRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RequestSignatureVerificationRequestBuilder is request builder for RequestSignatureVerification
type RequestSignatureVerificationRequestBuilder struct{ BaseRequestBuilder }

// Request returns RequestSignatureVerificationRequest
func (b *RequestSignatureVerificationRequestBuilder) Request() *RequestSignatureVerificationRequest {
	return &RequestSignatureVerificationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RequestSignatureVerificationRequest is request for RequestSignatureVerification
type RequestSignatureVerificationRequest struct{ BaseRequest }

// Get performs GET request for RequestSignatureVerification
func (r *RequestSignatureVerificationRequest) Get(ctx context.Context) (resObj *RequestSignatureVerification, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RequestSignatureVerification
func (r *RequestSignatureVerificationRequest) Update(ctx context.Context, reqObj *RequestSignatureVerification) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RequestSignatureVerification
func (r *RequestSignatureVerificationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type RequestObjectStopRequestBuilder struct{ BaseRequestBuilder }

// Stop action undocumentedrav
func (b *RequestObjectRequestBuilder) Stop(reqObj *RequestObjectStopRequestParameter) *RequestObjectStopRequestBuilder {
	bb := &RequestObjectStopRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Stop"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type RequestObjectStopRequest struct{ BaseRequest }

func (b *RequestObjectStopRequestBuilder) Request() *RequestObjectStopRequest {
	return &RequestObjectStopRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *RequestObjectStopRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type RequestObjectRecordDecisionsRequestBuilder struct{ BaseRequestBuilder }

// RecordDecisions action undocumentedrav
func (b *RequestObjectRequestBuilder) RecordDecisions(reqObj *RequestObjectRecordDecisionsRequestParameter) *RequestObjectRecordDecisionsRequestBuilder {
	bb := &RequestObjectRecordDecisionsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/RecordDecisions"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type RequestObjectRecordDecisionsRequest struct{ BaseRequest }

func (b *RequestObjectRecordDecisionsRequestBuilder) Request() *RequestObjectRecordDecisionsRequest {
	return &RequestObjectRecordDecisionsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *RequestObjectRecordDecisionsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
