// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// Calls returns request builder for Call collection
func (b *CommsApplicationRequestBuilder) Calls() *CommsApplicationCallsCollectionRequestBuilder {
	bb := &CommsApplicationCallsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/calls"
	return bb
}

// CommsApplicationCallsCollectionRequestBuilder is request builder for Call collection
type CommsApplicationCallsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Call collection
func (b *CommsApplicationCallsCollectionRequestBuilder) Request() *CommsApplicationCallsCollectionRequest {
	return &CommsApplicationCallsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Call item
func (b *CommsApplicationCallsCollectionRequestBuilder) ID(id string) *CallRequestBuilder {
	bb := &CallRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CommsApplicationCallsCollectionRequest is request for Call collection
type CommsApplicationCallsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Call collection
func (r *CommsApplicationCallsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Call, error) {
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
	var values []Call
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
			value  []Call
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

// GetN performs GET request for Call collection, max N pages
func (r *CommsApplicationCallsCollectionRequest) GetN(ctx context.Context, n int) ([]Call, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Call collection
func (r *CommsApplicationCallsCollectionRequest) Get(ctx context.Context) ([]Call, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Call collection
func (r *CommsApplicationCallsCollectionRequest) Add(ctx context.Context, reqObj *Call) (resObj *Call, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// OnlineMeetings returns request builder for OnlineMeeting collection
func (b *CommsApplicationRequestBuilder) OnlineMeetings() *CommsApplicationOnlineMeetingsCollectionRequestBuilder {
	bb := &CommsApplicationOnlineMeetingsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/onlineMeetings"
	return bb
}

// CommsApplicationOnlineMeetingsCollectionRequestBuilder is request builder for OnlineMeeting collection
type CommsApplicationOnlineMeetingsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnlineMeeting collection
func (b *CommsApplicationOnlineMeetingsCollectionRequestBuilder) Request() *CommsApplicationOnlineMeetingsCollectionRequest {
	return &CommsApplicationOnlineMeetingsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnlineMeeting item
func (b *CommsApplicationOnlineMeetingsCollectionRequestBuilder) ID(id string) *OnlineMeetingRequestBuilder {
	bb := &OnlineMeetingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CommsApplicationOnlineMeetingsCollectionRequest is request for OnlineMeeting collection
type CommsApplicationOnlineMeetingsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnlineMeeting collection
func (r *CommsApplicationOnlineMeetingsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]OnlineMeeting, error) {
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
	var values []OnlineMeeting
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
			value  []OnlineMeeting
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

// GetN performs GET request for OnlineMeeting collection, max N pages
func (r *CommsApplicationOnlineMeetingsCollectionRequest) GetN(ctx context.Context, n int) ([]OnlineMeeting, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for OnlineMeeting collection
func (r *CommsApplicationOnlineMeetingsCollectionRequest) Get(ctx context.Context) ([]OnlineMeeting, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for OnlineMeeting collection
func (r *CommsApplicationOnlineMeetingsCollectionRequest) Add(ctx context.Context, reqObj *OnlineMeeting) (resObj *OnlineMeeting, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
