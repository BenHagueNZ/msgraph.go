// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// AgentGroups returns request builder for OnPremisesAgentGroup collection
func (b *OnPremisesAgentRequestBuilder) AgentGroups() *OnPremisesAgentAgentGroupsCollectionRequestBuilder {
	bb := &OnPremisesAgentAgentGroupsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/agentGroups"
	return bb
}

// OnPremisesAgentAgentGroupsCollectionRequestBuilder is request builder for OnPremisesAgentGroup collection rcn
type OnPremisesAgentAgentGroupsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnPremisesAgentGroup collection
func (b *OnPremisesAgentAgentGroupsCollectionRequestBuilder) Request() *OnPremisesAgentAgentGroupsCollectionRequest {
	return &OnPremisesAgentAgentGroupsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnPremisesAgentGroup item
func (b *OnPremisesAgentAgentGroupsCollectionRequestBuilder) ID(id string) *OnPremisesAgentGroupRequestBuilder {
	bb := &OnPremisesAgentGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnPremisesAgentAgentGroupsCollectionRequest is request for OnPremisesAgentGroup collection
type OnPremisesAgentAgentGroupsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnPremisesAgentGroup collection
func (r *OnPremisesAgentAgentGroupsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]OnPremisesAgentGroup, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []OnPremisesAgentGroup
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []OnPremisesAgentGroup
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for OnPremisesAgentGroup collection, max N pages
func (r *OnPremisesAgentAgentGroupsCollectionRequest) GetN(ctx context.Context, n int) ([]OnPremisesAgentGroup, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for OnPremisesAgentGroup collection
func (r *OnPremisesAgentAgentGroupsCollectionRequest) Get(ctx context.Context) ([]OnPremisesAgentGroup, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for OnPremisesAgentGroup collection
func (r *OnPremisesAgentAgentGroupsCollectionRequest) Add(ctx context.Context, reqObj *OnPremisesAgentGroup) (resObj *OnPremisesAgentGroup, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Agents returns request builder for OnPremisesAgent collection
func (b *OnPremisesAgentGroupRequestBuilder) Agents() *OnPremisesAgentGroupAgentsCollectionRequestBuilder {
	bb := &OnPremisesAgentGroupAgentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/agents"
	return bb
}

// OnPremisesAgentGroupAgentsCollectionRequestBuilder is request builder for OnPremisesAgent collection rcn
type OnPremisesAgentGroupAgentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnPremisesAgent collection
func (b *OnPremisesAgentGroupAgentsCollectionRequestBuilder) Request() *OnPremisesAgentGroupAgentsCollectionRequest {
	return &OnPremisesAgentGroupAgentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnPremisesAgent item
func (b *OnPremisesAgentGroupAgentsCollectionRequestBuilder) ID(id string) *OnPremisesAgentRequestBuilder {
	bb := &OnPremisesAgentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnPremisesAgentGroupAgentsCollectionRequest is request for OnPremisesAgent collection
type OnPremisesAgentGroupAgentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnPremisesAgent collection
func (r *OnPremisesAgentGroupAgentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]OnPremisesAgent, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []OnPremisesAgent
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []OnPremisesAgent
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for OnPremisesAgent collection, max N pages
func (r *OnPremisesAgentGroupAgentsCollectionRequest) GetN(ctx context.Context, n int) ([]OnPremisesAgent, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for OnPremisesAgent collection
func (r *OnPremisesAgentGroupAgentsCollectionRequest) Get(ctx context.Context) ([]OnPremisesAgent, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for OnPremisesAgent collection
func (r *OnPremisesAgentGroupAgentsCollectionRequest) Add(ctx context.Context, reqObj *OnPremisesAgent) (resObj *OnPremisesAgent, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// PublishedResources returns request builder for PublishedResource collection
func (b *OnPremisesAgentGroupRequestBuilder) PublishedResources() *OnPremisesAgentGroupPublishedResourcesCollectionRequestBuilder {
	bb := &OnPremisesAgentGroupPublishedResourcesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/publishedResources"
	return bb
}

// OnPremisesAgentGroupPublishedResourcesCollectionRequestBuilder is request builder for PublishedResource collection rcn
type OnPremisesAgentGroupPublishedResourcesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PublishedResource collection
func (b *OnPremisesAgentGroupPublishedResourcesCollectionRequestBuilder) Request() *OnPremisesAgentGroupPublishedResourcesCollectionRequest {
	return &OnPremisesAgentGroupPublishedResourcesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PublishedResource item
func (b *OnPremisesAgentGroupPublishedResourcesCollectionRequestBuilder) ID(id string) *PublishedResourceRequestBuilder {
	bb := &PublishedResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnPremisesAgentGroupPublishedResourcesCollectionRequest is request for PublishedResource collection
type OnPremisesAgentGroupPublishedResourcesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PublishedResource collection
func (r *OnPremisesAgentGroupPublishedResourcesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PublishedResource, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []PublishedResource
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []PublishedResource
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for PublishedResource collection, max N pages
func (r *OnPremisesAgentGroupPublishedResourcesCollectionRequest) GetN(ctx context.Context, n int) ([]PublishedResource, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PublishedResource collection
func (r *OnPremisesAgentGroupPublishedResourcesCollectionRequest) Get(ctx context.Context) ([]PublishedResource, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PublishedResource collection
func (r *OnPremisesAgentGroupPublishedResourcesCollectionRequest) Add(ctx context.Context, reqObj *PublishedResource) (resObj *PublishedResource, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AgentGroups returns request builder for OnPremisesAgentGroup collection
func (b *OnPremisesPublishingProfileRequestBuilder) AgentGroups() *OnPremisesPublishingProfileAgentGroupsCollectionRequestBuilder {
	bb := &OnPremisesPublishingProfileAgentGroupsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/agentGroups"
	return bb
}

// OnPremisesPublishingProfileAgentGroupsCollectionRequestBuilder is request builder for OnPremisesAgentGroup collection rcn
type OnPremisesPublishingProfileAgentGroupsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnPremisesAgentGroup collection
func (b *OnPremisesPublishingProfileAgentGroupsCollectionRequestBuilder) Request() *OnPremisesPublishingProfileAgentGroupsCollectionRequest {
	return &OnPremisesPublishingProfileAgentGroupsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnPremisesAgentGroup item
func (b *OnPremisesPublishingProfileAgentGroupsCollectionRequestBuilder) ID(id string) *OnPremisesAgentGroupRequestBuilder {
	bb := &OnPremisesAgentGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnPremisesPublishingProfileAgentGroupsCollectionRequest is request for OnPremisesAgentGroup collection
type OnPremisesPublishingProfileAgentGroupsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnPremisesAgentGroup collection
func (r *OnPremisesPublishingProfileAgentGroupsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]OnPremisesAgentGroup, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []OnPremisesAgentGroup
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []OnPremisesAgentGroup
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for OnPremisesAgentGroup collection, max N pages
func (r *OnPremisesPublishingProfileAgentGroupsCollectionRequest) GetN(ctx context.Context, n int) ([]OnPremisesAgentGroup, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for OnPremisesAgentGroup collection
func (r *OnPremisesPublishingProfileAgentGroupsCollectionRequest) Get(ctx context.Context) ([]OnPremisesAgentGroup, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for OnPremisesAgentGroup collection
func (r *OnPremisesPublishingProfileAgentGroupsCollectionRequest) Add(ctx context.Context, reqObj *OnPremisesAgentGroup) (resObj *OnPremisesAgentGroup, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Agents returns request builder for OnPremisesAgent collection
func (b *OnPremisesPublishingProfileRequestBuilder) Agents() *OnPremisesPublishingProfileAgentsCollectionRequestBuilder {
	bb := &OnPremisesPublishingProfileAgentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/agents"
	return bb
}

// OnPremisesPublishingProfileAgentsCollectionRequestBuilder is request builder for OnPremisesAgent collection rcn
type OnPremisesPublishingProfileAgentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnPremisesAgent collection
func (b *OnPremisesPublishingProfileAgentsCollectionRequestBuilder) Request() *OnPremisesPublishingProfileAgentsCollectionRequest {
	return &OnPremisesPublishingProfileAgentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnPremisesAgent item
func (b *OnPremisesPublishingProfileAgentsCollectionRequestBuilder) ID(id string) *OnPremisesAgentRequestBuilder {
	bb := &OnPremisesAgentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnPremisesPublishingProfileAgentsCollectionRequest is request for OnPremisesAgent collection
type OnPremisesPublishingProfileAgentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnPremisesAgent collection
func (r *OnPremisesPublishingProfileAgentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]OnPremisesAgent, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []OnPremisesAgent
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []OnPremisesAgent
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for OnPremisesAgent collection, max N pages
func (r *OnPremisesPublishingProfileAgentsCollectionRequest) GetN(ctx context.Context, n int) ([]OnPremisesAgent, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for OnPremisesAgent collection
func (r *OnPremisesPublishingProfileAgentsCollectionRequest) Get(ctx context.Context) ([]OnPremisesAgent, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for OnPremisesAgent collection
func (r *OnPremisesPublishingProfileAgentsCollectionRequest) Add(ctx context.Context, reqObj *OnPremisesAgent) (resObj *OnPremisesAgent, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ConnectorGroups returns request builder for ConnectorGroup collection
func (b *OnPremisesPublishingProfileRequestBuilder) ConnectorGroups() *OnPremisesPublishingProfileConnectorGroupsCollectionRequestBuilder {
	bb := &OnPremisesPublishingProfileConnectorGroupsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/connectorGroups"
	return bb
}

// OnPremisesPublishingProfileConnectorGroupsCollectionRequestBuilder is request builder for ConnectorGroup collection rcn
type OnPremisesPublishingProfileConnectorGroupsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ConnectorGroup collection
func (b *OnPremisesPublishingProfileConnectorGroupsCollectionRequestBuilder) Request() *OnPremisesPublishingProfileConnectorGroupsCollectionRequest {
	return &OnPremisesPublishingProfileConnectorGroupsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ConnectorGroup item
func (b *OnPremisesPublishingProfileConnectorGroupsCollectionRequestBuilder) ID(id string) *ConnectorGroupRequestBuilder {
	bb := &ConnectorGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnPremisesPublishingProfileConnectorGroupsCollectionRequest is request for ConnectorGroup collection
type OnPremisesPublishingProfileConnectorGroupsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ConnectorGroup collection
func (r *OnPremisesPublishingProfileConnectorGroupsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ConnectorGroup, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ConnectorGroup
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ConnectorGroup
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for ConnectorGroup collection, max N pages
func (r *OnPremisesPublishingProfileConnectorGroupsCollectionRequest) GetN(ctx context.Context, n int) ([]ConnectorGroup, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ConnectorGroup collection
func (r *OnPremisesPublishingProfileConnectorGroupsCollectionRequest) Get(ctx context.Context) ([]ConnectorGroup, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ConnectorGroup collection
func (r *OnPremisesPublishingProfileConnectorGroupsCollectionRequest) Add(ctx context.Context, reqObj *ConnectorGroup) (resObj *ConnectorGroup, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Connectors returns request builder for Connector collection
func (b *OnPremisesPublishingProfileRequestBuilder) Connectors() *OnPremisesPublishingProfileConnectorsCollectionRequestBuilder {
	bb := &OnPremisesPublishingProfileConnectorsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/connectors"
	return bb
}

// OnPremisesPublishingProfileConnectorsCollectionRequestBuilder is request builder for Connector collection rcn
type OnPremisesPublishingProfileConnectorsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Connector collection
func (b *OnPremisesPublishingProfileConnectorsCollectionRequestBuilder) Request() *OnPremisesPublishingProfileConnectorsCollectionRequest {
	return &OnPremisesPublishingProfileConnectorsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Connector item
func (b *OnPremisesPublishingProfileConnectorsCollectionRequestBuilder) ID(id string) *ConnectorRequestBuilder {
	bb := &ConnectorRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnPremisesPublishingProfileConnectorsCollectionRequest is request for Connector collection
type OnPremisesPublishingProfileConnectorsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Connector collection
func (r *OnPremisesPublishingProfileConnectorsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Connector, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Connector
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []Connector
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for Connector collection, max N pages
func (r *OnPremisesPublishingProfileConnectorsCollectionRequest) GetN(ctx context.Context, n int) ([]Connector, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Connector collection
func (r *OnPremisesPublishingProfileConnectorsCollectionRequest) Get(ctx context.Context) ([]Connector, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Connector collection
func (r *OnPremisesPublishingProfileConnectorsCollectionRequest) Add(ctx context.Context, reqObj *Connector) (resObj *Connector, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// PublishedResources returns request builder for PublishedResource collection
func (b *OnPremisesPublishingProfileRequestBuilder) PublishedResources() *OnPremisesPublishingProfilePublishedResourcesCollectionRequestBuilder {
	bb := &OnPremisesPublishingProfilePublishedResourcesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/publishedResources"
	return bb
}

// OnPremisesPublishingProfilePublishedResourcesCollectionRequestBuilder is request builder for PublishedResource collection rcn
type OnPremisesPublishingProfilePublishedResourcesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PublishedResource collection
func (b *OnPremisesPublishingProfilePublishedResourcesCollectionRequestBuilder) Request() *OnPremisesPublishingProfilePublishedResourcesCollectionRequest {
	return &OnPremisesPublishingProfilePublishedResourcesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PublishedResource item
func (b *OnPremisesPublishingProfilePublishedResourcesCollectionRequestBuilder) ID(id string) *PublishedResourceRequestBuilder {
	bb := &PublishedResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnPremisesPublishingProfilePublishedResourcesCollectionRequest is request for PublishedResource collection
type OnPremisesPublishingProfilePublishedResourcesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PublishedResource collection
func (r *OnPremisesPublishingProfilePublishedResourcesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PublishedResource, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []PublishedResource
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []PublishedResource
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for PublishedResource collection, max N pages
func (r *OnPremisesPublishingProfilePublishedResourcesCollectionRequest) GetN(ctx context.Context, n int) ([]PublishedResource, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PublishedResource collection
func (r *OnPremisesPublishingProfilePublishedResourcesCollectionRequest) Get(ctx context.Context) ([]PublishedResource, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PublishedResource collection
func (r *OnPremisesPublishingProfilePublishedResourcesCollectionRequest) Add(ctx context.Context, reqObj *PublishedResource) (resObj *PublishedResource, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// CustomExtension is navigation property rn
func (b *OnTokenIssuanceStartCustomExtensionHandlerRequestBuilder) CustomExtension() *OnTokenIssuanceStartCustomExtensionRequestBuilder {
	bb := &OnTokenIssuanceStartCustomExtensionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/customExtension"
	return bb
}

// Entity is navigation property rn
func (b *OnPremisesAgentRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *OnPremisesAgentGroupRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *OnPremisesConditionalAccessSettingsRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *OnPremisesDirectorySynchronizationRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *OnPremisesPublishingProfileRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// OnTokenIssuanceStartCustomExtension returns request builder for OnTokenIssuanceStartCustomExtension collection
func (b *IdentityContainerCustomAuthenticationExtensionsCollectionRequestBuilder) OnTokenIssuanceStartCustomExtension() *IdentityContainerCustomAuthenticationExtensionsCollectionOnTokenIssuanceStartCustomExtensionCollectionRequestBuilder {
	bb := &IdentityContainerCustomAuthenticationExtensionsCollectionOnTokenIssuanceStartCustomExtensionCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// IdentityContainerCustomAuthenticationExtensionsCollectionOnTokenIssuanceStartCustomExtensionCollectionRequestBuilder is request builder for OnTokenIssuanceStartCustomExtension collection rcn
type IdentityContainerCustomAuthenticationExtensionsCollectionOnTokenIssuanceStartCustomExtensionCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnTokenIssuanceStartCustomExtension collection
func (b *IdentityContainerCustomAuthenticationExtensionsCollectionOnTokenIssuanceStartCustomExtensionCollectionRequestBuilder) Request() *IdentityContainerCustomAuthenticationExtensionsCollectionOnTokenIssuanceStartCustomExtensionCollectionRequest {
	return &IdentityContainerCustomAuthenticationExtensionsCollectionOnTokenIssuanceStartCustomExtensionCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnTokenIssuanceStartCustomExtension item
func (b *IdentityContainerCustomAuthenticationExtensionsCollectionOnTokenIssuanceStartCustomExtensionCollectionRequestBuilder) ID(id string) *OnTokenIssuanceStartCustomExtensionRequestBuilder {
	bb := &OnTokenIssuanceStartCustomExtensionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// IdentityContainerCustomAuthenticationExtensionsCollectionOnTokenIssuanceStartCustomExtensionCollectionRequest is request for OnTokenIssuanceStartCustomExtension collection
type IdentityContainerCustomAuthenticationExtensionsCollectionOnTokenIssuanceStartCustomExtensionCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnTokenIssuanceStartCustomExtension collection
func (r *IdentityContainerCustomAuthenticationExtensionsCollectionOnTokenIssuanceStartCustomExtensionCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]OnTokenIssuanceStartCustomExtension, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []OnTokenIssuanceStartCustomExtension
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []OnTokenIssuanceStartCustomExtension
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for OnTokenIssuanceStartCustomExtension collection, max N pages
func (r *IdentityContainerCustomAuthenticationExtensionsCollectionOnTokenIssuanceStartCustomExtensionCollectionRequest) GetN(ctx context.Context, n int) ([]OnTokenIssuanceStartCustomExtension, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for OnTokenIssuanceStartCustomExtension collection
func (r *IdentityContainerCustomAuthenticationExtensionsCollectionOnTokenIssuanceStartCustomExtensionCollectionRequest) Get(ctx context.Context) ([]OnTokenIssuanceStartCustomExtension, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for OnTokenIssuanceStartCustomExtension collection
func (r *IdentityContainerCustomAuthenticationExtensionsCollectionOnTokenIssuanceStartCustomExtensionCollectionRequest) Add(ctx context.Context, reqObj *OnTokenIssuanceStartCustomExtension) (resObj *OnTokenIssuanceStartCustomExtension, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// OnTokenIssuanceStartListener returns request builder for OnTokenIssuanceStartListener collection
func (b *IdentityContainerAuthenticationEventListenersCollectionRequestBuilder) OnTokenIssuanceStartListener() *IdentityContainerAuthenticationEventListenersCollectionOnTokenIssuanceStartListenerCollectionRequestBuilder {
	bb := &IdentityContainerAuthenticationEventListenersCollectionOnTokenIssuanceStartListenerCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// IdentityContainerAuthenticationEventListenersCollectionOnTokenIssuanceStartListenerCollectionRequestBuilder is request builder for OnTokenIssuanceStartListener collection rcn
type IdentityContainerAuthenticationEventListenersCollectionOnTokenIssuanceStartListenerCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnTokenIssuanceStartListener collection
func (b *IdentityContainerAuthenticationEventListenersCollectionOnTokenIssuanceStartListenerCollectionRequestBuilder) Request() *IdentityContainerAuthenticationEventListenersCollectionOnTokenIssuanceStartListenerCollectionRequest {
	return &IdentityContainerAuthenticationEventListenersCollectionOnTokenIssuanceStartListenerCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnTokenIssuanceStartListener item
func (b *IdentityContainerAuthenticationEventListenersCollectionOnTokenIssuanceStartListenerCollectionRequestBuilder) ID(id string) *OnTokenIssuanceStartListenerRequestBuilder {
	bb := &OnTokenIssuanceStartListenerRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// IdentityContainerAuthenticationEventListenersCollectionOnTokenIssuanceStartListenerCollectionRequest is request for OnTokenIssuanceStartListener collection
type IdentityContainerAuthenticationEventListenersCollectionOnTokenIssuanceStartListenerCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnTokenIssuanceStartListener collection
func (r *IdentityContainerAuthenticationEventListenersCollectionOnTokenIssuanceStartListenerCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]OnTokenIssuanceStartListener, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []OnTokenIssuanceStartListener
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []OnTokenIssuanceStartListener
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for OnTokenIssuanceStartListener collection, max N pages
func (r *IdentityContainerAuthenticationEventListenersCollectionOnTokenIssuanceStartListenerCollectionRequest) GetN(ctx context.Context, n int) ([]OnTokenIssuanceStartListener, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for OnTokenIssuanceStartListener collection
func (r *IdentityContainerAuthenticationEventListenersCollectionOnTokenIssuanceStartListenerCollectionRequest) Get(ctx context.Context) ([]OnTokenIssuanceStartListener, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for OnTokenIssuanceStartListener collection
func (r *IdentityContainerAuthenticationEventListenersCollectionOnTokenIssuanceStartListenerCollectionRequest) Add(ctx context.Context, reqObj *OnTokenIssuanceStartListener) (resObj *OnTokenIssuanceStartListener, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
