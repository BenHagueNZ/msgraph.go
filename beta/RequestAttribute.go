// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// AttributeMappingFunctionSchemaRequestBuilder is request builder for AttributeMappingFunctionSchema
type AttributeMappingFunctionSchemaRequestBuilder struct{ BaseRequestBuilder }

// Request returns AttributeMappingFunctionSchemaRequest
func (b *AttributeMappingFunctionSchemaRequestBuilder) Request() *AttributeMappingFunctionSchemaRequest {
	return &AttributeMappingFunctionSchemaRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AttributeMappingFunctionSchemaRequest is request for AttributeMappingFunctionSchema
type AttributeMappingFunctionSchemaRequest struct{ BaseRequest }

// Get performs GET request for AttributeMappingFunctionSchema
func (r *AttributeMappingFunctionSchemaRequest) Get(ctx context.Context) (resObj *AttributeMappingFunctionSchema, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AttributeMappingFunctionSchema
func (r *AttributeMappingFunctionSchemaRequest) Update(ctx context.Context, reqObj *AttributeMappingFunctionSchema) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AttributeMappingFunctionSchema
func (r *AttributeMappingFunctionSchemaRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
