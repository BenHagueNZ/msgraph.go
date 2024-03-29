// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// MetaDataKeyStringPairRequestBuilder is request builder for MetaDataKeyStringPair
type MetaDataKeyStringPairRequestBuilder struct{ BaseRequestBuilder }

// Request returns MetaDataKeyStringPairRequest
func (b *MetaDataKeyStringPairRequestBuilder) Request() *MetaDataKeyStringPairRequest {
	return &MetaDataKeyStringPairRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MetaDataKeyStringPairRequest is request for MetaDataKeyStringPair
type MetaDataKeyStringPairRequest struct{ BaseRequest }

// Get performs GET request for MetaDataKeyStringPair
func (r *MetaDataKeyStringPairRequest) Get(ctx context.Context) (resObj *MetaDataKeyStringPair, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MetaDataKeyStringPair
func (r *MetaDataKeyStringPairRequest) Update(ctx context.Context, reqObj *MetaDataKeyStringPair) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MetaDataKeyStringPair
func (r *MetaDataKeyStringPairRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MetaDataKeyValuePairRequestBuilder is request builder for MetaDataKeyValuePair
type MetaDataKeyValuePairRequestBuilder struct{ BaseRequestBuilder }

// Request returns MetaDataKeyValuePairRequest
func (b *MetaDataKeyValuePairRequestBuilder) Request() *MetaDataKeyValuePairRequest {
	return &MetaDataKeyValuePairRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MetaDataKeyValuePairRequest is request for MetaDataKeyValuePair
type MetaDataKeyValuePairRequest struct{ BaseRequest }

// Get performs GET request for MetaDataKeyValuePair
func (r *MetaDataKeyValuePairRequest) Get(ctx context.Context) (resObj *MetaDataKeyValuePair, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MetaDataKeyValuePair
func (r *MetaDataKeyValuePairRequest) Update(ctx context.Context, reqObj *MetaDataKeyValuePair) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MetaDataKeyValuePair
func (r *MetaDataKeyValuePairRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
