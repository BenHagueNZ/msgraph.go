// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// CalendarRequestBuilder is request builder for Calendar
type CalendarRequestBuilder struct{ BaseRequestBuilder }

// Request returns CalendarRequest
func (b *CalendarRequestBuilder) Request() *CalendarRequest {
	return &CalendarRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CalendarRequest is request for Calendar
type CalendarRequest struct{ BaseRequest }

// Get performs GET request for Calendar
func (r *CalendarRequest) Get(ctx context.Context) (resObj *Calendar, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Calendar
func (r *CalendarRequest) Update(ctx context.Context, reqObj *Calendar) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Calendar
func (r *CalendarRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CalendarGroupRequestBuilder is request builder for CalendarGroup
type CalendarGroupRequestBuilder struct{ BaseRequestBuilder }

// Request returns CalendarGroupRequest
func (b *CalendarGroupRequestBuilder) Request() *CalendarGroupRequest {
	return &CalendarGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CalendarGroupRequest is request for CalendarGroup
type CalendarGroupRequest struct{ BaseRequest }

// Get performs GET request for CalendarGroup
func (r *CalendarGroupRequest) Get(ctx context.Context) (resObj *CalendarGroup, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CalendarGroup
func (r *CalendarGroupRequest) Update(ctx context.Context, reqObj *CalendarGroup) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CalendarGroup
func (r *CalendarGroupRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CalendarPermissionRequestBuilder is request builder for CalendarPermission
type CalendarPermissionRequestBuilder struct{ BaseRequestBuilder }

// Request returns CalendarPermissionRequest
func (b *CalendarPermissionRequestBuilder) Request() *CalendarPermissionRequest {
	return &CalendarPermissionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CalendarPermissionRequest is request for CalendarPermission
type CalendarPermissionRequest struct{ BaseRequest }

// Get performs GET request for CalendarPermission
func (r *CalendarPermissionRequest) Get(ctx context.Context) (resObj *CalendarPermission, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CalendarPermission
func (r *CalendarPermissionRequest) Update(ctx context.Context, reqObj *CalendarPermission) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CalendarPermission
func (r *CalendarPermissionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CalendarSharingMessageRequestBuilder is request builder for CalendarSharingMessage
type CalendarSharingMessageRequestBuilder struct{ BaseRequestBuilder }

// Request returns CalendarSharingMessageRequest
func (b *CalendarSharingMessageRequestBuilder) Request() *CalendarSharingMessageRequest {
	return &CalendarSharingMessageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CalendarSharingMessageRequest is request for CalendarSharingMessage
type CalendarSharingMessageRequest struct{ BaseRequest }

// Get performs GET request for CalendarSharingMessage
func (r *CalendarSharingMessageRequest) Get(ctx context.Context) (resObj *CalendarSharingMessage, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CalendarSharingMessage
func (r *CalendarSharingMessageRequest) Update(ctx context.Context, reqObj *CalendarSharingMessage) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CalendarSharingMessage
func (r *CalendarSharingMessageRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CalendarSharingMessageActionRequestBuilder is request builder for CalendarSharingMessageAction
type CalendarSharingMessageActionRequestBuilder struct{ BaseRequestBuilder }

// Request returns CalendarSharingMessageActionRequest
func (b *CalendarSharingMessageActionRequestBuilder) Request() *CalendarSharingMessageActionRequest {
	return &CalendarSharingMessageActionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CalendarSharingMessageActionRequest is request for CalendarSharingMessageAction
type CalendarSharingMessageActionRequest struct{ BaseRequest }

// Get performs GET request for CalendarSharingMessageAction
func (r *CalendarSharingMessageActionRequest) Get(ctx context.Context) (resObj *CalendarSharingMessageAction, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CalendarSharingMessageAction
func (r *CalendarSharingMessageActionRequest) Update(ctx context.Context, reqObj *CalendarSharingMessageAction) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CalendarSharingMessageAction
func (r *CalendarSharingMessageActionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type CalendarGetScheduleRequestBuilder struct{ BaseRequestBuilder }

// GetSchedule action undocumentedrac
func (b *CalendarRequestBuilder) GetSchedule(reqObj *CalendarGetScheduleRequestParameter) *CalendarGetScheduleRequestBuilder {
	bb := &CalendarGetScheduleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/GetSchedule"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type CalendarGetScheduleRequest struct{ BaseRequest }

func (b *CalendarGetScheduleRequestBuilder) Request() *CalendarGetScheduleRequest {
	return &CalendarGetScheduleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *CalendarGetScheduleRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ScheduleInformation, error) {
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
	var values []ScheduleInformation
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
			value  []ScheduleInformation
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
		req, _ = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

func (r *CalendarGetScheduleRequest) PostN(ctx context.Context, n int) ([]ScheduleInformation, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

func (r *CalendarGetScheduleRequest) Post(ctx context.Context) ([]ScheduleInformation, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

type CalendarSharingMessageAcceptRequestBuilder struct{ BaseRequestBuilder }

// Accept action undocumentedras
func (b *CalendarSharingMessageRequestBuilder) Accept(reqObj *CalendarSharingMessageAcceptRequestParameter) *CalendarSharingMessageAcceptRequestBuilder {
	bb := &CalendarSharingMessageAcceptRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Accept"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type CalendarSharingMessageAcceptRequest struct{ BaseRequest }

func (b *CalendarSharingMessageAcceptRequestBuilder) Request() *CalendarSharingMessageAcceptRequest {
	return &CalendarSharingMessageAcceptRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *CalendarSharingMessageAcceptRequest) Post(ctx context.Context) (resObj *Calendar, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}