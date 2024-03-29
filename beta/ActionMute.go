// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// MuteParticipantOperation returns request builder for MuteParticipantOperation collection
func (b *CallOperationsCollectionRequestBuilder) MuteParticipantOperation() *CallOperationsCollectionMuteParticipantOperationCollectionRequestBuilder {
	bb := &CallOperationsCollectionMuteParticipantOperationCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// CallOperationsCollectionMuteParticipantOperationCollectionRequestBuilder is request builder for MuteParticipantOperation collection rcn
type CallOperationsCollectionMuteParticipantOperationCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MuteParticipantOperation collection
func (b *CallOperationsCollectionMuteParticipantOperationCollectionRequestBuilder) Request() *CallOperationsCollectionMuteParticipantOperationCollectionRequest {
	return &CallOperationsCollectionMuteParticipantOperationCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MuteParticipantOperation item
func (b *CallOperationsCollectionMuteParticipantOperationCollectionRequestBuilder) ID(id string) *MuteParticipantOperationRequestBuilder {
	bb := &MuteParticipantOperationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CallOperationsCollectionMuteParticipantOperationCollectionRequest is request for MuteParticipantOperation collection
type CallOperationsCollectionMuteParticipantOperationCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MuteParticipantOperation collection
func (r *CallOperationsCollectionMuteParticipantOperationCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MuteParticipantOperation, error) {
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
	var values []MuteParticipantOperation
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
			value  []MuteParticipantOperation
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

// GetN performs GET request for MuteParticipantOperation collection, max N pages
func (r *CallOperationsCollectionMuteParticipantOperationCollectionRequest) GetN(ctx context.Context, n int) ([]MuteParticipantOperation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MuteParticipantOperation collection
func (r *CallOperationsCollectionMuteParticipantOperationCollectionRequest) Get(ctx context.Context) ([]MuteParticipantOperation, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MuteParticipantOperation collection
func (r *CallOperationsCollectionMuteParticipantOperationCollectionRequest) Add(ctx context.Context, reqObj *MuteParticipantOperation) (resObj *MuteParticipantOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// MuteParticipantsOperation returns request builder for MuteParticipantsOperation collection
func (b *CallOperationsCollectionRequestBuilder) MuteParticipantsOperation() *CallOperationsCollectionMuteParticipantsOperationCollectionRequestBuilder {
	bb := &CallOperationsCollectionMuteParticipantsOperationCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// CallOperationsCollectionMuteParticipantsOperationCollectionRequestBuilder is request builder for MuteParticipantsOperation collection rcn
type CallOperationsCollectionMuteParticipantsOperationCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MuteParticipantsOperation collection
func (b *CallOperationsCollectionMuteParticipantsOperationCollectionRequestBuilder) Request() *CallOperationsCollectionMuteParticipantsOperationCollectionRequest {
	return &CallOperationsCollectionMuteParticipantsOperationCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MuteParticipantsOperation item
func (b *CallOperationsCollectionMuteParticipantsOperationCollectionRequestBuilder) ID(id string) *MuteParticipantsOperationRequestBuilder {
	bb := &MuteParticipantsOperationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CallOperationsCollectionMuteParticipantsOperationCollectionRequest is request for MuteParticipantsOperation collection
type CallOperationsCollectionMuteParticipantsOperationCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MuteParticipantsOperation collection
func (r *CallOperationsCollectionMuteParticipantsOperationCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MuteParticipantsOperation, error) {
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
	var values []MuteParticipantsOperation
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
			value  []MuteParticipantsOperation
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

// GetN performs GET request for MuteParticipantsOperation collection, max N pages
func (r *CallOperationsCollectionMuteParticipantsOperationCollectionRequest) GetN(ctx context.Context, n int) ([]MuteParticipantsOperation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MuteParticipantsOperation collection
func (r *CallOperationsCollectionMuteParticipantsOperationCollectionRequest) Get(ctx context.Context) ([]MuteParticipantsOperation, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MuteParticipantsOperation collection
func (r *CallOperationsCollectionMuteParticipantsOperationCollectionRequest) Add(ctx context.Context, reqObj *MuteParticipantsOperation) (resObj *MuteParticipantsOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
