// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// CountryNamedLocation returns request builder for CountryNamedLocation collection
func (b *ConditionalAccessRootNamedLocationsCollectionRequestBuilder) CountryNamedLocation() *ConditionalAccessRootNamedLocationsCollectionCountryNamedLocationCollectionRequestBuilder {
	bb := &ConditionalAccessRootNamedLocationsCollectionCountryNamedLocationCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// ConditionalAccessRootNamedLocationsCollectionCountryNamedLocationCollectionRequestBuilder is request builder for CountryNamedLocation collection rcn
type ConditionalAccessRootNamedLocationsCollectionCountryNamedLocationCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for CountryNamedLocation collection
func (b *ConditionalAccessRootNamedLocationsCollectionCountryNamedLocationCollectionRequestBuilder) Request() *ConditionalAccessRootNamedLocationsCollectionCountryNamedLocationCollectionRequest {
	return &ConditionalAccessRootNamedLocationsCollectionCountryNamedLocationCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for CountryNamedLocation item
func (b *ConditionalAccessRootNamedLocationsCollectionCountryNamedLocationCollectionRequestBuilder) ID(id string) *CountryNamedLocationRequestBuilder {
	bb := &CountryNamedLocationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ConditionalAccessRootNamedLocationsCollectionCountryNamedLocationCollectionRequest is request for CountryNamedLocation collection
type ConditionalAccessRootNamedLocationsCollectionCountryNamedLocationCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for CountryNamedLocation collection
func (r *ConditionalAccessRootNamedLocationsCollectionCountryNamedLocationCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]CountryNamedLocation, error) {
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
	var values []CountryNamedLocation
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
			value  []CountryNamedLocation
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

// GetN performs GET request for CountryNamedLocation collection, max N pages
func (r *ConditionalAccessRootNamedLocationsCollectionCountryNamedLocationCollectionRequest) GetN(ctx context.Context, n int) ([]CountryNamedLocation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for CountryNamedLocation collection
func (r *ConditionalAccessRootNamedLocationsCollectionCountryNamedLocationCollectionRequest) Get(ctx context.Context) ([]CountryNamedLocation, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for CountryNamedLocation collection
func (r *ConditionalAccessRootNamedLocationsCollectionCountryNamedLocationCollectionRequest) Add(ctx context.Context, reqObj *CountryNamedLocation) (resObj *CountryNamedLocation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *CountryRegionRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
