// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// ApplicationSegments returns request builder for IPApplicationSegment collection
func (b *IPSegmentConfigurationRequestBuilder) ApplicationSegments() *IPSegmentConfigurationApplicationSegmentsCollectionRequestBuilder {
	bb := &IPSegmentConfigurationApplicationSegmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/applicationSegments"
	return bb
}

// IPSegmentConfigurationApplicationSegmentsCollectionRequestBuilder is request builder for IPApplicationSegment collection rcn
type IPSegmentConfigurationApplicationSegmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for IPApplicationSegment collection
func (b *IPSegmentConfigurationApplicationSegmentsCollectionRequestBuilder) Request() *IPSegmentConfigurationApplicationSegmentsCollectionRequest {
	return &IPSegmentConfigurationApplicationSegmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for IPApplicationSegment item
func (b *IPSegmentConfigurationApplicationSegmentsCollectionRequestBuilder) ID(id string) *IPApplicationSegmentRequestBuilder {
	bb := &IPApplicationSegmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// IPSegmentConfigurationApplicationSegmentsCollectionRequest is request for IPApplicationSegment collection
type IPSegmentConfigurationApplicationSegmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for IPApplicationSegment collection
func (r *IPSegmentConfigurationApplicationSegmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]IPApplicationSegment, error) {
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
	var values []IPApplicationSegment
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
			value  []IPApplicationSegment
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

// GetN performs GET request for IPApplicationSegment collection, max N pages
func (r *IPSegmentConfigurationApplicationSegmentsCollectionRequest) GetN(ctx context.Context, n int) ([]IPApplicationSegment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for IPApplicationSegment collection
func (r *IPSegmentConfigurationApplicationSegmentsCollectionRequest) Get(ctx context.Context) ([]IPApplicationSegment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for IPApplicationSegment collection
func (r *IPSegmentConfigurationApplicationSegmentsCollectionRequest) Add(ctx context.Context, reqObj *IPApplicationSegment) (resObj *IPApplicationSegment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// IPNamedLocation returns request builder for IPNamedLocation collection
func (b *ConditionalAccessRootNamedLocationsCollectionRequestBuilder) IPNamedLocation() *ConditionalAccessRootNamedLocationsCollectionIPNamedLocationCollectionRequestBuilder {
	bb := &ConditionalAccessRootNamedLocationsCollectionIPNamedLocationCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// ConditionalAccessRootNamedLocationsCollectionIPNamedLocationCollectionRequestBuilder is request builder for IPNamedLocation collection rcn
type ConditionalAccessRootNamedLocationsCollectionIPNamedLocationCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for IPNamedLocation collection
func (b *ConditionalAccessRootNamedLocationsCollectionIPNamedLocationCollectionRequestBuilder) Request() *ConditionalAccessRootNamedLocationsCollectionIPNamedLocationCollectionRequest {
	return &ConditionalAccessRootNamedLocationsCollectionIPNamedLocationCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for IPNamedLocation item
func (b *ConditionalAccessRootNamedLocationsCollectionIPNamedLocationCollectionRequestBuilder) ID(id string) *IPNamedLocationRequestBuilder {
	bb := &IPNamedLocationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ConditionalAccessRootNamedLocationsCollectionIPNamedLocationCollectionRequest is request for IPNamedLocation collection
type ConditionalAccessRootNamedLocationsCollectionIPNamedLocationCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for IPNamedLocation collection
func (r *ConditionalAccessRootNamedLocationsCollectionIPNamedLocationCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]IPNamedLocation, error) {
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
	var values []IPNamedLocation
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
			value  []IPNamedLocation
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

// GetN performs GET request for IPNamedLocation collection, max N pages
func (r *ConditionalAccessRootNamedLocationsCollectionIPNamedLocationCollectionRequest) GetN(ctx context.Context, n int) ([]IPNamedLocation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for IPNamedLocation collection
func (r *ConditionalAccessRootNamedLocationsCollectionIPNamedLocationCollectionRequest) Get(ctx context.Context) ([]IPNamedLocation, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for IPNamedLocation collection
func (r *ConditionalAccessRootNamedLocationsCollectionIPNamedLocationCollectionRequest) Add(ctx context.Context, reqObj *IPNamedLocation) (resObj *IPNamedLocation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *IPSecurityProfileRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
