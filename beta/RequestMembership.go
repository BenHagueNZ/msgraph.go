// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// MembershipOutlierInsightRequestBuilder is request builder for MembershipOutlierInsight
type MembershipOutlierInsightRequestBuilder struct{ BaseRequestBuilder }

// Request returns MembershipOutlierInsightRequest
func (b *MembershipOutlierInsightRequestBuilder) Request() *MembershipOutlierInsightRequest {
	return &MembershipOutlierInsightRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MembershipOutlierInsightRequest is request for MembershipOutlierInsight
type MembershipOutlierInsightRequest struct{ BaseRequest }

// Get performs GET request for MembershipOutlierInsight
func (r *MembershipOutlierInsightRequest) Get(ctx context.Context) (resObj *MembershipOutlierInsight, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MembershipOutlierInsight
func (r *MembershipOutlierInsightRequest) Update(ctx context.Context, reqObj *MembershipOutlierInsight) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MembershipOutlierInsight
func (r *MembershipOutlierInsightRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MembershipRuleEvaluationDetailsRequestBuilder is request builder for MembershipRuleEvaluationDetails
type MembershipRuleEvaluationDetailsRequestBuilder struct{ BaseRequestBuilder }

// Request returns MembershipRuleEvaluationDetailsRequest
func (b *MembershipRuleEvaluationDetailsRequestBuilder) Request() *MembershipRuleEvaluationDetailsRequest {
	return &MembershipRuleEvaluationDetailsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MembershipRuleEvaluationDetailsRequest is request for MembershipRuleEvaluationDetails
type MembershipRuleEvaluationDetailsRequest struct{ BaseRequest }

// Get performs GET request for MembershipRuleEvaluationDetails
func (r *MembershipRuleEvaluationDetailsRequest) Get(ctx context.Context) (resObj *MembershipRuleEvaluationDetails, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MembershipRuleEvaluationDetails
func (r *MembershipRuleEvaluationDetailsRequest) Update(ctx context.Context, reqObj *MembershipRuleEvaluationDetails) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MembershipRuleEvaluationDetails
func (r *MembershipRuleEvaluationDetailsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MembershipRuleProcessingStatusRequestBuilder is request builder for MembershipRuleProcessingStatus
type MembershipRuleProcessingStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns MembershipRuleProcessingStatusRequest
func (b *MembershipRuleProcessingStatusRequestBuilder) Request() *MembershipRuleProcessingStatusRequest {
	return &MembershipRuleProcessingStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MembershipRuleProcessingStatusRequest is request for MembershipRuleProcessingStatus
type MembershipRuleProcessingStatusRequest struct{ BaseRequest }

// Get performs GET request for MembershipRuleProcessingStatus
func (r *MembershipRuleProcessingStatusRequest) Get(ctx context.Context) (resObj *MembershipRuleProcessingStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MembershipRuleProcessingStatus
func (r *MembershipRuleProcessingStatusRequest) Update(ctx context.Context, reqObj *MembershipRuleProcessingStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MembershipRuleProcessingStatus
func (r *MembershipRuleProcessingStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}