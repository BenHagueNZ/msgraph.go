// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// BookingBusinesses returns request builder for BookingBusiness collection rcn
func (b *SolutionsRootRequestBuilder) BookingBusinesses() *SolutionsRootBookingBusinessesCollectionRequestBuilder {
	bb := &SolutionsRootBookingBusinessesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/bookingBusinesses"
	return bb
}

// SolutionsRootBookingBusinessesCollectionRequestBuilder is request builder for BookingBusiness collection
type SolutionsRootBookingBusinessesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for BookingBusiness collection
func (b *SolutionsRootBookingBusinessesCollectionRequestBuilder) Request() *SolutionsRootBookingBusinessesCollectionRequest {
	return &SolutionsRootBookingBusinessesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for BookingBusiness item
func (b *SolutionsRootBookingBusinessesCollectionRequestBuilder) ID(id string) *BookingBusinessRequestBuilder {
	bb := &BookingBusinessRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SolutionsRootBookingBusinessesCollectionRequest is request for BookingBusiness collection
type SolutionsRootBookingBusinessesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for BookingBusiness collection
func (r *SolutionsRootBookingBusinessesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]BookingBusiness, error) {
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
	var values []BookingBusiness
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
			value  []BookingBusiness
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

// GetN performs GET request for BookingBusiness collection, max N pages
func (r *SolutionsRootBookingBusinessesCollectionRequest) GetN(ctx context.Context, n int) ([]BookingBusiness, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for BookingBusiness collection
func (r *SolutionsRootBookingBusinessesCollectionRequest) Get(ctx context.Context) ([]BookingBusiness, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for BookingBusiness collection
func (r *SolutionsRootBookingBusinessesCollectionRequest) Add(ctx context.Context, reqObj *BookingBusiness) (resObj *BookingBusiness, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// BookingCurrencies returns request builder for BookingCurrency collection rcn
func (b *SolutionsRootRequestBuilder) BookingCurrencies() *SolutionsRootBookingCurrenciesCollectionRequestBuilder {
	bb := &SolutionsRootBookingCurrenciesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/bookingCurrencies"
	return bb
}

// SolutionsRootBookingCurrenciesCollectionRequestBuilder is request builder for BookingCurrency collection
type SolutionsRootBookingCurrenciesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for BookingCurrency collection
func (b *SolutionsRootBookingCurrenciesCollectionRequestBuilder) Request() *SolutionsRootBookingCurrenciesCollectionRequest {
	return &SolutionsRootBookingCurrenciesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for BookingCurrency item
func (b *SolutionsRootBookingCurrenciesCollectionRequestBuilder) ID(id string) *BookingCurrencyRequestBuilder {
	bb := &BookingCurrencyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SolutionsRootBookingCurrenciesCollectionRequest is request for BookingCurrency collection
type SolutionsRootBookingCurrenciesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for BookingCurrency collection
func (r *SolutionsRootBookingCurrenciesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]BookingCurrency, error) {
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
	var values []BookingCurrency
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
			value  []BookingCurrency
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

// GetN performs GET request for BookingCurrency collection, max N pages
func (r *SolutionsRootBookingCurrenciesCollectionRequest) GetN(ctx context.Context, n int) ([]BookingCurrency, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for BookingCurrency collection
func (r *SolutionsRootBookingCurrenciesCollectionRequest) Get(ctx context.Context) ([]BookingCurrency, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for BookingCurrency collection
func (r *SolutionsRootBookingCurrenciesCollectionRequest) Add(ctx context.Context, reqObj *BookingCurrency) (resObj *BookingCurrency, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
