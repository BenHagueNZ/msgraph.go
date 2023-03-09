// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// CloudCommunicationsGetPresencesByUserIDRequestParameter undocumented
type CloudCommunicationsGetPresencesByUserIDRequestParameter struct {
	// IDs undocumented
	IDs []string `json:"ids,omitempty"`
}

// CallRecords returns request builder for CallRecord collection
func (b *CloudCommunicationsRequestBuilder) CallRecords() *CloudCommunicationsCallRecordsCollectionRequestBuilder {
	bb := &CloudCommunicationsCallRecordsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/callRecords"
	return bb
}

// CloudCommunicationsCallRecordsCollectionRequestBuilder is request builder for CallRecord collection
type CloudCommunicationsCallRecordsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for CallRecord collection
func (b *CloudCommunicationsCallRecordsCollectionRequestBuilder) Request() *CloudCommunicationsCallRecordsCollectionRequest {
	return &CloudCommunicationsCallRecordsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for CallRecord item
func (b *CloudCommunicationsCallRecordsCollectionRequestBuilder) ID(id string) *CallRecordRequestBuilder {
	bb := &CallRecordRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CloudCommunicationsCallRecordsCollectionRequest is request for CallRecord collection
type CloudCommunicationsCallRecordsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for CallRecord collection
func (r *CloudCommunicationsCallRecordsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]CallRecord, error) {
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
	var values []CallRecord
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
			value  []CallRecord
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

// GetN performs GET request for CallRecord collection, max N pages
func (r *CloudCommunicationsCallRecordsCollectionRequest) GetN(ctx context.Context, n int) ([]CallRecord, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for CallRecord collection
func (r *CloudCommunicationsCallRecordsCollectionRequest) Get(ctx context.Context) ([]CallRecord, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for CallRecord collection
func (r *CloudCommunicationsCallRecordsCollectionRequest) Add(ctx context.Context, reqObj *CallRecord) (resObj *CallRecord, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Calls returns request builder for Call collection
func (b *CloudCommunicationsRequestBuilder) Calls() *CloudCommunicationsCallsCollectionRequestBuilder {
	bb := &CloudCommunicationsCallsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/calls"
	return bb
}

// CloudCommunicationsCallsCollectionRequestBuilder is request builder for Call collection
type CloudCommunicationsCallsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Call collection
func (b *CloudCommunicationsCallsCollectionRequestBuilder) Request() *CloudCommunicationsCallsCollectionRequest {
	return &CloudCommunicationsCallsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Call item
func (b *CloudCommunicationsCallsCollectionRequestBuilder) ID(id string) *CallRequestBuilder {
	bb := &CallRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CloudCommunicationsCallsCollectionRequest is request for Call collection
type CloudCommunicationsCallsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Call collection
func (r *CloudCommunicationsCallsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Call, error) {
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
func (r *CloudCommunicationsCallsCollectionRequest) GetN(ctx context.Context, n int) ([]Call, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Call collection
func (r *CloudCommunicationsCallsCollectionRequest) Get(ctx context.Context) ([]Call, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Call collection
func (r *CloudCommunicationsCallsCollectionRequest) Add(ctx context.Context, reqObj *Call) (resObj *Call, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// OnlineMeetings returns request builder for OnlineMeeting collection
func (b *CloudCommunicationsRequestBuilder) OnlineMeetings() *CloudCommunicationsOnlineMeetingsCollectionRequestBuilder {
	bb := &CloudCommunicationsOnlineMeetingsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/onlineMeetings"
	return bb
}

// CloudCommunicationsOnlineMeetingsCollectionRequestBuilder is request builder for OnlineMeeting collection
type CloudCommunicationsOnlineMeetingsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnlineMeeting collection
func (b *CloudCommunicationsOnlineMeetingsCollectionRequestBuilder) Request() *CloudCommunicationsOnlineMeetingsCollectionRequest {
	return &CloudCommunicationsOnlineMeetingsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnlineMeeting item
func (b *CloudCommunicationsOnlineMeetingsCollectionRequestBuilder) ID(id string) *OnlineMeetingRequestBuilder {
	bb := &OnlineMeetingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CloudCommunicationsOnlineMeetingsCollectionRequest is request for OnlineMeeting collection
type CloudCommunicationsOnlineMeetingsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnlineMeeting collection
func (r *CloudCommunicationsOnlineMeetingsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]OnlineMeeting, error) {
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
func (r *CloudCommunicationsOnlineMeetingsCollectionRequest) GetN(ctx context.Context, n int) ([]OnlineMeeting, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for OnlineMeeting collection
func (r *CloudCommunicationsOnlineMeetingsCollectionRequest) Get(ctx context.Context) ([]OnlineMeeting, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for OnlineMeeting collection
func (r *CloudCommunicationsOnlineMeetingsCollectionRequest) Add(ctx context.Context, reqObj *OnlineMeeting) (resObj *OnlineMeeting, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Presences returns request builder for Presence collection
func (b *CloudCommunicationsRequestBuilder) Presences() *CloudCommunicationsPresencesCollectionRequestBuilder {
	bb := &CloudCommunicationsPresencesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/presences"
	return bb
}

// CloudCommunicationsPresencesCollectionRequestBuilder is request builder for Presence collection
type CloudCommunicationsPresencesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Presence collection
func (b *CloudCommunicationsPresencesCollectionRequestBuilder) Request() *CloudCommunicationsPresencesCollectionRequest {
	return &CloudCommunicationsPresencesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Presence item
func (b *CloudCommunicationsPresencesCollectionRequestBuilder) ID(id string) *PresenceRequestBuilder {
	bb := &PresenceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CloudCommunicationsPresencesCollectionRequest is request for Presence collection
type CloudCommunicationsPresencesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Presence collection
func (r *CloudCommunicationsPresencesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Presence, error) {
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
	var values []Presence
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
			value  []Presence
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

// GetN performs GET request for Presence collection, max N pages
func (r *CloudCommunicationsPresencesCollectionRequest) GetN(ctx context.Context, n int) ([]Presence, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Presence collection
func (r *CloudCommunicationsPresencesCollectionRequest) Get(ctx context.Context) ([]Presence, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Presence collection
func (r *CloudCommunicationsPresencesCollectionRequest) Add(ctx context.Context, reqObj *Presence) (resObj *Presence, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
