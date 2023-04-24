// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SpecifiedCaptiveNetworkPluginsRequestBuilder is request builder for SpecifiedCaptiveNetworkPlugins
type SpecifiedCaptiveNetworkPluginsRequestBuilder struct{ BaseRequestBuilder }

// Request returns SpecifiedCaptiveNetworkPluginsRequest
func (b *SpecifiedCaptiveNetworkPluginsRequestBuilder) Request() *SpecifiedCaptiveNetworkPluginsRequest {
	return &SpecifiedCaptiveNetworkPluginsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SpecifiedCaptiveNetworkPluginsRequest is request for SpecifiedCaptiveNetworkPlugins
type SpecifiedCaptiveNetworkPluginsRequest struct{ BaseRequest }

// Get performs GET request for SpecifiedCaptiveNetworkPlugins
func (r *SpecifiedCaptiveNetworkPluginsRequest) Get(ctx context.Context) (resObj *SpecifiedCaptiveNetworkPlugins, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SpecifiedCaptiveNetworkPlugins
func (r *SpecifiedCaptiveNetworkPluginsRequest) Update(ctx context.Context, reqObj *SpecifiedCaptiveNetworkPlugins) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SpecifiedCaptiveNetworkPlugins
func (r *SpecifiedCaptiveNetworkPluginsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
