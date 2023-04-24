// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// RecommendationCompleteRequestParameter undocumented
type RecommendationCompleteRequestParameter struct {
}

// RecommendationDismissRequestParameter undocumented
type RecommendationDismissRequestParameter struct {
	// DismissReason undocumented
	DismissReason *string `json:"dismissReason,omitempty"`
}

// RecommendationPostponeRequestParameter undocumented
type RecommendationPostponeRequestParameter struct {
	// PostponeUntilDateTime undocumented
	PostponeUntilDateTime *time.Time `json:"postponeUntilDateTime,omitempty"`
}

// RecommendationReactivateRequestParameter undocumented
type RecommendationReactivateRequestParameter struct {
}

// ImpactedResources returns request builder for ImpactedResource collection
func (b *RecommendationBaseRequestBuilder) ImpactedResources() *RecommendationBaseImpactedResourcesCollectionRequestBuilder {
	bb := &RecommendationBaseImpactedResourcesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/impactedResources"
	return bb
}

// RecommendationBaseImpactedResourcesCollectionRequestBuilder is request builder for ImpactedResource collection rcn
type RecommendationBaseImpactedResourcesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ImpactedResource collection
func (b *RecommendationBaseImpactedResourcesCollectionRequestBuilder) Request() *RecommendationBaseImpactedResourcesCollectionRequest {
	return &RecommendationBaseImpactedResourcesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ImpactedResource item
func (b *RecommendationBaseImpactedResourcesCollectionRequestBuilder) ID(id string) *ImpactedResourceRequestBuilder {
	bb := &ImpactedResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// RecommendationBaseImpactedResourcesCollectionRequest is request for ImpactedResource collection
type RecommendationBaseImpactedResourcesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ImpactedResource collection
func (r *RecommendationBaseImpactedResourcesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ImpactedResource, error) {
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
	var values []ImpactedResource
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
			value  []ImpactedResource
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

// GetN performs GET request for ImpactedResource collection, max N pages
func (r *RecommendationBaseImpactedResourcesCollectionRequest) GetN(ctx context.Context, n int) ([]ImpactedResource, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ImpactedResource collection
func (r *RecommendationBaseImpactedResourcesCollectionRequest) Get(ctx context.Context) ([]ImpactedResource, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ImpactedResource collection
func (r *RecommendationBaseImpactedResourcesCollectionRequest) Add(ctx context.Context, reqObj *ImpactedResource) (resObj *ImpactedResource, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *RecommendationBaseRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}