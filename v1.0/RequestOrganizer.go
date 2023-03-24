// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// OrganizerMeetingInfoRequestBuilder is request builder for OrganizerMeetingInfo
type OrganizerMeetingInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns OrganizerMeetingInfoRequest
func (b *OrganizerMeetingInfoRequestBuilder) Request() *OrganizerMeetingInfoRequest {
	return &OrganizerMeetingInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OrganizerMeetingInfoRequest is request for OrganizerMeetingInfo
type OrganizerMeetingInfoRequest struct{ BaseRequest }

// Get performs GET request for OrganizerMeetingInfo
func (r *OrganizerMeetingInfoRequest) Get(ctx context.Context) (resObj *OrganizerMeetingInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OrganizerMeetingInfo
func (r *OrganizerMeetingInfoRequest) Update(ctx context.Context, reqObj *OrganizerMeetingInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OrganizerMeetingInfo
func (r *OrganizerMeetingInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
