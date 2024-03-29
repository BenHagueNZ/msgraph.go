// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// EventRequestBuilder is request builder for Event
type EventRequestBuilder struct{ BaseRequestBuilder }

// Request returns EventRequest
func (b *EventRequestBuilder) Request() *EventRequest {
	return &EventRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EventRequest is request for Event
type EventRequest struct{ BaseRequest }

// Get performs GET request for Event
func (r *EventRequest) Get(ctx context.Context) (resObj *Event, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Event
func (r *EventRequest) Update(ctx context.Context, reqObj *Event) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Event
func (r *EventRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// EventMessageRequestBuilder is request builder for EventMessage
type EventMessageRequestBuilder struct{ BaseRequestBuilder }

// Request returns EventMessageRequest
func (b *EventMessageRequestBuilder) Request() *EventMessageRequest {
	return &EventMessageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EventMessageRequest is request for EventMessage
type EventMessageRequest struct{ BaseRequest }

// Get performs GET request for EventMessage
func (r *EventMessageRequest) Get(ctx context.Context) (resObj *EventMessage, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EventMessage
func (r *EventMessageRequest) Update(ctx context.Context, reqObj *EventMessage) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EventMessage
func (r *EventMessageRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// EventMessageDetailRequestBuilder is request builder for EventMessageDetail
type EventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns EventMessageDetailRequest
func (b *EventMessageDetailRequestBuilder) Request() *EventMessageDetailRequest {
	return &EventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EventMessageDetailRequest is request for EventMessageDetail
type EventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for EventMessageDetail
func (r *EventMessageDetailRequest) Get(ctx context.Context) (resObj *EventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EventMessageDetail
func (r *EventMessageDetailRequest) Update(ctx context.Context, reqObj *EventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EventMessageDetail
func (r *EventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// EventMessageRequestObjectRequestBuilder is request builder for EventMessageRequestObject
type EventMessageRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns EventMessageRequestObjectRequest
func (b *EventMessageRequestObjectRequestBuilder) Request() *EventMessageRequestObjectRequest {
	return &EventMessageRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EventMessageRequestObjectRequest is request for EventMessageRequestObject
type EventMessageRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for EventMessageRequestObject
func (r *EventMessageRequestObjectRequest) Get(ctx context.Context) (resObj *EventMessageRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EventMessageRequestObject
func (r *EventMessageRequestObjectRequest) Update(ctx context.Context, reqObj *EventMessageRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EventMessageRequestObject
func (r *EventMessageRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// EventMessageResponseRequestBuilder is request builder for EventMessageResponse
type EventMessageResponseRequestBuilder struct{ BaseRequestBuilder }

// Request returns EventMessageResponseRequest
func (b *EventMessageResponseRequestBuilder) Request() *EventMessageResponseRequest {
	return &EventMessageResponseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EventMessageResponseRequest is request for EventMessageResponse
type EventMessageResponseRequest struct{ BaseRequest }

// Get performs GET request for EventMessageResponse
func (r *EventMessageResponseRequest) Get(ctx context.Context) (resObj *EventMessageResponse, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EventMessageResponse
func (r *EventMessageResponseRequest) Update(ctx context.Context, reqObj *EventMessageResponse) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EventMessageResponse
func (r *EventMessageResponseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type EventCancelRequestBuilder struct{ BaseRequestBuilder }

// Cancel action undocumentedrav
func (b *EventRequestBuilder) Cancel(reqObj *EventCancelRequestParameter) *EventCancelRequestBuilder {
	bb := &EventCancelRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Cancel"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type EventCancelRequest struct{ BaseRequest }

func (b *EventCancelRequestBuilder) Request() *EventCancelRequest {
	return &EventCancelRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *EventCancelRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type EventAcceptRequestBuilder struct{ BaseRequestBuilder }

// Accept action undocumentedrav
func (b *EventRequestBuilder) Accept(reqObj *EventAcceptRequestParameter) *EventAcceptRequestBuilder {
	bb := &EventAcceptRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Accept"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type EventAcceptRequest struct{ BaseRequest }

func (b *EventAcceptRequestBuilder) Request() *EventAcceptRequest {
	return &EventAcceptRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *EventAcceptRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type EventDeclineRequestBuilder struct{ BaseRequestBuilder }

// Decline action undocumentedrav
func (b *EventRequestBuilder) Decline(reqObj *EventDeclineRequestParameter) *EventDeclineRequestBuilder {
	bb := &EventDeclineRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Decline"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type EventDeclineRequest struct{ BaseRequest }

func (b *EventDeclineRequestBuilder) Request() *EventDeclineRequest {
	return &EventDeclineRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *EventDeclineRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type EventDismissReminderRequestBuilder struct{ BaseRequestBuilder }

// DismissReminder action undocumentedrav
func (b *EventRequestBuilder) DismissReminder(reqObj *EventDismissReminderRequestParameter) *EventDismissReminderRequestBuilder {
	bb := &EventDismissReminderRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/DismissReminder"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type EventDismissReminderRequest struct{ BaseRequest }

func (b *EventDismissReminderRequestBuilder) Request() *EventDismissReminderRequest {
	return &EventDismissReminderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *EventDismissReminderRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type EventForwardRequestBuilder struct{ BaseRequestBuilder }

// Forward action undocumentedrav
func (b *EventRequestBuilder) Forward(reqObj *EventForwardRequestParameter) *EventForwardRequestBuilder {
	bb := &EventForwardRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Forward"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type EventForwardRequest struct{ BaseRequest }

func (b *EventForwardRequestBuilder) Request() *EventForwardRequest {
	return &EventForwardRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *EventForwardRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type EventSnoozeReminderRequestBuilder struct{ BaseRequestBuilder }

// SnoozeReminder action undocumentedrav
func (b *EventRequestBuilder) SnoozeReminder(reqObj *EventSnoozeReminderRequestParameter) *EventSnoozeReminderRequestBuilder {
	bb := &EventSnoozeReminderRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/SnoozeReminder"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type EventSnoozeReminderRequest struct{ BaseRequest }

func (b *EventSnoozeReminderRequestBuilder) Request() *EventSnoozeReminderRequest {
	return &EventSnoozeReminderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *EventSnoozeReminderRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type EventTentativelyAcceptRequestBuilder struct{ BaseRequestBuilder }

// TentativelyAccept action undocumentedrav
func (b *EventRequestBuilder) TentativelyAccept(reqObj *EventTentativelyAcceptRequestParameter) *EventTentativelyAcceptRequestBuilder {
	bb := &EventTentativelyAcceptRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/TentativelyAccept"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type EventTentativelyAcceptRequest struct{ BaseRequest }

func (b *EventTentativelyAcceptRequestBuilder) Request() *EventTentativelyAcceptRequest {
	return &EventTentativelyAcceptRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *EventTentativelyAcceptRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
