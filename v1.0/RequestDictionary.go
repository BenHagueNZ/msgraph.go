// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// DictionaryRequestBuilder is request builder for Dictionary
type DictionaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns DictionaryRequest
func (b *DictionaryRequestBuilder) Request() *DictionaryRequest {
	return &DictionaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DictionaryRequest is request for Dictionary
type DictionaryRequest struct{ BaseRequest }

// Get performs GET request for Dictionary
func (r *DictionaryRequest) Get(ctx context.Context) (resObj *Dictionary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Dictionary
func (r *DictionaryRequest) Update(ctx context.Context, reqObj *Dictionary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Dictionary
func (r *DictionaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
