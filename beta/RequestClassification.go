// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ClassificationAttributeRequestBuilder is request builder for ClassificationAttribute
type ClassificationAttributeRequestBuilder struct{ BaseRequestBuilder }

// Request returns ClassificationAttributeRequest
func (b *ClassificationAttributeRequestBuilder) Request() *ClassificationAttributeRequest {
	return &ClassificationAttributeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ClassificationAttributeRequest is request for ClassificationAttribute
type ClassificationAttributeRequest struct{ BaseRequest }

// Get performs GET request for ClassificationAttribute
func (r *ClassificationAttributeRequest) Get(ctx context.Context) (resObj *ClassificationAttribute, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ClassificationAttribute
func (r *ClassificationAttributeRequest) Update(ctx context.Context, reqObj *ClassificationAttribute) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ClassificationAttribute
func (r *ClassificationAttributeRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ClassificationErrorRequestBuilder is request builder for ClassificationError
type ClassificationErrorRequestBuilder struct{ BaseRequestBuilder }

// Request returns ClassificationErrorRequest
func (b *ClassificationErrorRequestBuilder) Request() *ClassificationErrorRequest {
	return &ClassificationErrorRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ClassificationErrorRequest is request for ClassificationError
type ClassificationErrorRequest struct{ BaseRequest }

// Get performs GET request for ClassificationError
func (r *ClassificationErrorRequest) Get(ctx context.Context) (resObj *ClassificationError, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ClassificationError
func (r *ClassificationErrorRequest) Update(ctx context.Context, reqObj *ClassificationError) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ClassificationError
func (r *ClassificationErrorRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ClassificationInnerErrorRequestBuilder is request builder for ClassificationInnerError
type ClassificationInnerErrorRequestBuilder struct{ BaseRequestBuilder }

// Request returns ClassificationInnerErrorRequest
func (b *ClassificationInnerErrorRequestBuilder) Request() *ClassificationInnerErrorRequest {
	return &ClassificationInnerErrorRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ClassificationInnerErrorRequest is request for ClassificationInnerError
type ClassificationInnerErrorRequest struct{ BaseRequest }

// Get performs GET request for ClassificationInnerError
func (r *ClassificationInnerErrorRequest) Get(ctx context.Context) (resObj *ClassificationInnerError, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ClassificationInnerError
func (r *ClassificationInnerErrorRequest) Update(ctx context.Context, reqObj *ClassificationInnerError) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ClassificationInnerError
func (r *ClassificationInnerErrorRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ClassificationJobResponseRequestBuilder is request builder for ClassificationJobResponse
type ClassificationJobResponseRequestBuilder struct{ BaseRequestBuilder }

// Request returns ClassificationJobResponseRequest
func (b *ClassificationJobResponseRequestBuilder) Request() *ClassificationJobResponseRequest {
	return &ClassificationJobResponseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ClassificationJobResponseRequest is request for ClassificationJobResponse
type ClassificationJobResponseRequest struct{ BaseRequest }

// Get performs GET request for ClassificationJobResponse
func (r *ClassificationJobResponseRequest) Get(ctx context.Context) (resObj *ClassificationJobResponse, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ClassificationJobResponse
func (r *ClassificationJobResponseRequest) Update(ctx context.Context, reqObj *ClassificationJobResponse) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ClassificationJobResponse
func (r *ClassificationJobResponseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ClassificationRequestContentMetaDataRequestBuilder is request builder for ClassificationRequestContentMetaData
type ClassificationRequestContentMetaDataRequestBuilder struct{ BaseRequestBuilder }

// Request returns ClassificationRequestContentMetaDataRequest
func (b *ClassificationRequestContentMetaDataRequestBuilder) Request() *ClassificationRequestContentMetaDataRequest {
	return &ClassificationRequestContentMetaDataRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ClassificationRequestContentMetaDataRequest is request for ClassificationRequestContentMetaData
type ClassificationRequestContentMetaDataRequest struct{ BaseRequest }

// Get performs GET request for ClassificationRequestContentMetaData
func (r *ClassificationRequestContentMetaDataRequest) Get(ctx context.Context) (resObj *ClassificationRequestContentMetaData, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ClassificationRequestContentMetaData
func (r *ClassificationRequestContentMetaDataRequest) Update(ctx context.Context, reqObj *ClassificationRequestContentMetaData) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ClassificationRequestContentMetaData
func (r *ClassificationRequestContentMetaDataRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ClassificationResultRequestBuilder is request builder for ClassificationResult
type ClassificationResultRequestBuilder struct{ BaseRequestBuilder }

// Request returns ClassificationResultRequest
func (b *ClassificationResultRequestBuilder) Request() *ClassificationResultRequest {
	return &ClassificationResultRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ClassificationResultRequest is request for ClassificationResult
type ClassificationResultRequest struct{ BaseRequest }

// Get performs GET request for ClassificationResult
func (r *ClassificationResultRequest) Get(ctx context.Context) (resObj *ClassificationResult, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ClassificationResult
func (r *ClassificationResultRequest) Update(ctx context.Context, reqObj *ClassificationResult) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ClassificationResult
func (r *ClassificationResultRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
