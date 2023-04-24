// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// DecisionItemPrincipalResourceMembershipRequestBuilder is request builder for DecisionItemPrincipalResourceMembership
type DecisionItemPrincipalResourceMembershipRequestBuilder struct{ BaseRequestBuilder }

// Request returns DecisionItemPrincipalResourceMembershipRequest
func (b *DecisionItemPrincipalResourceMembershipRequestBuilder) Request() *DecisionItemPrincipalResourceMembershipRequest {
	return &DecisionItemPrincipalResourceMembershipRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DecisionItemPrincipalResourceMembershipRequest is request for DecisionItemPrincipalResourceMembership
type DecisionItemPrincipalResourceMembershipRequest struct{ BaseRequest }

// Get performs GET request for DecisionItemPrincipalResourceMembership
func (r *DecisionItemPrincipalResourceMembershipRequest) Get(ctx context.Context) (resObj *DecisionItemPrincipalResourceMembership, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DecisionItemPrincipalResourceMembership
func (r *DecisionItemPrincipalResourceMembershipRequest) Update(ctx context.Context, reqObj *DecisionItemPrincipalResourceMembership) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DecisionItemPrincipalResourceMembership
func (r *DecisionItemPrincipalResourceMembershipRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}