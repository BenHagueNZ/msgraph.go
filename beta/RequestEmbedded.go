// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// EmbeddedSIMActivationCodePoolRequestBuilder is request builder for EmbeddedSIMActivationCodePool
type EmbeddedSIMActivationCodePoolRequestBuilder struct{ BaseRequestBuilder }

// Request returns EmbeddedSIMActivationCodePoolRequest
func (b *EmbeddedSIMActivationCodePoolRequestBuilder) Request() *EmbeddedSIMActivationCodePoolRequest {
	return &EmbeddedSIMActivationCodePoolRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EmbeddedSIMActivationCodePoolRequest is request for EmbeddedSIMActivationCodePool
type EmbeddedSIMActivationCodePoolRequest struct{ BaseRequest }

// Get performs GET request for EmbeddedSIMActivationCodePool
func (r *EmbeddedSIMActivationCodePoolRequest) Get(ctx context.Context) (resObj *EmbeddedSIMActivationCodePool, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EmbeddedSIMActivationCodePool
func (r *EmbeddedSIMActivationCodePoolRequest) Update(ctx context.Context, reqObj *EmbeddedSIMActivationCodePool) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EmbeddedSIMActivationCodePool
func (r *EmbeddedSIMActivationCodePoolRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// EmbeddedSIMActivationCodePoolAssignmentRequestBuilder is request builder for EmbeddedSIMActivationCodePoolAssignment
type EmbeddedSIMActivationCodePoolAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns EmbeddedSIMActivationCodePoolAssignmentRequest
func (b *EmbeddedSIMActivationCodePoolAssignmentRequestBuilder) Request() *EmbeddedSIMActivationCodePoolAssignmentRequest {
	return &EmbeddedSIMActivationCodePoolAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EmbeddedSIMActivationCodePoolAssignmentRequest is request for EmbeddedSIMActivationCodePoolAssignment
type EmbeddedSIMActivationCodePoolAssignmentRequest struct{ BaseRequest }

// Get performs GET request for EmbeddedSIMActivationCodePoolAssignment
func (r *EmbeddedSIMActivationCodePoolAssignmentRequest) Get(ctx context.Context) (resObj *EmbeddedSIMActivationCodePoolAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EmbeddedSIMActivationCodePoolAssignment
func (r *EmbeddedSIMActivationCodePoolAssignmentRequest) Update(ctx context.Context, reqObj *EmbeddedSIMActivationCodePoolAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EmbeddedSIMActivationCodePoolAssignment
func (r *EmbeddedSIMActivationCodePoolAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// EmbeddedSIMDeviceStateRequestBuilder is request builder for EmbeddedSIMDeviceState
type EmbeddedSIMDeviceStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns EmbeddedSIMDeviceStateRequest
func (b *EmbeddedSIMDeviceStateRequestBuilder) Request() *EmbeddedSIMDeviceStateRequest {
	return &EmbeddedSIMDeviceStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EmbeddedSIMDeviceStateRequest is request for EmbeddedSIMDeviceState
type EmbeddedSIMDeviceStateRequest struct{ BaseRequest }

// Get performs GET request for EmbeddedSIMDeviceState
func (r *EmbeddedSIMDeviceStateRequest) Get(ctx context.Context) (resObj *EmbeddedSIMDeviceState, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EmbeddedSIMDeviceState
func (r *EmbeddedSIMDeviceStateRequest) Update(ctx context.Context, reqObj *EmbeddedSIMDeviceState) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EmbeddedSIMDeviceState
func (r *EmbeddedSIMDeviceStateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type EmbeddedSIMActivationCodePoolAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *EmbeddedSIMActivationCodePoolRequestBuilder) Assign(reqObj *EmbeddedSIMActivationCodePoolAssignRequestParameter) *EmbeddedSIMActivationCodePoolAssignRequestBuilder {
	bb := &EmbeddedSIMActivationCodePoolAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type EmbeddedSIMActivationCodePoolAssignRequest struct{ BaseRequest }

//
func (b *EmbeddedSIMActivationCodePoolAssignRequestBuilder) Request() *EmbeddedSIMActivationCodePoolAssignRequest {
	return &EmbeddedSIMActivationCodePoolAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *EmbeddedSIMActivationCodePoolAssignRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]EmbeddedSIMActivationCodePoolAssignment, error) {
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
	var values []EmbeddedSIMActivationCodePoolAssignment
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
			value  []EmbeddedSIMActivationCodePoolAssignment
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

//
func (r *EmbeddedSIMActivationCodePoolAssignRequest) PostN(ctx context.Context, n int) ([]EmbeddedSIMActivationCodePoolAssignment, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *EmbeddedSIMActivationCodePoolAssignRequest) Post(ctx context.Context) ([]EmbeddedSIMActivationCodePoolAssignment, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}
