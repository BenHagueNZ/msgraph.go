// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// WorkflowRequestBuilder is request builder for Workflow
type WorkflowRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkflowRequest
func (b *WorkflowRequestBuilder) Request() *WorkflowRequest {
	return &WorkflowRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkflowRequest is request for Workflow
type WorkflowRequest struct{ BaseRequest }

// Get performs GET request for Workflow
func (r *WorkflowRequest) Get(ctx context.Context) (resObj *Workflow, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Workflow
func (r *WorkflowRequest) Update(ctx context.Context, reqObj *Workflow) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Workflow
func (r *WorkflowRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// WorkflowBaseRequestBuilder is request builder for WorkflowBase
type WorkflowBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkflowBaseRequest
func (b *WorkflowBaseRequestBuilder) Request() *WorkflowBaseRequest {
	return &WorkflowBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkflowBaseRequest is request for WorkflowBase
type WorkflowBaseRequest struct{ BaseRequest }

// Get performs GET request for WorkflowBase
func (r *WorkflowBaseRequest) Get(ctx context.Context) (resObj *WorkflowBase, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkflowBase
func (r *WorkflowBaseRequest) Update(ctx context.Context, reqObj *WorkflowBase) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkflowBase
func (r *WorkflowBaseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// WorkflowExecutionConditionsRequestBuilder is request builder for WorkflowExecutionConditions
type WorkflowExecutionConditionsRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkflowExecutionConditionsRequest
func (b *WorkflowExecutionConditionsRequestBuilder) Request() *WorkflowExecutionConditionsRequest {
	return &WorkflowExecutionConditionsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkflowExecutionConditionsRequest is request for WorkflowExecutionConditions
type WorkflowExecutionConditionsRequest struct{ BaseRequest }

// Get performs GET request for WorkflowExecutionConditions
func (r *WorkflowExecutionConditionsRequest) Get(ctx context.Context) (resObj *WorkflowExecutionConditions, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkflowExecutionConditions
func (r *WorkflowExecutionConditionsRequest) Update(ctx context.Context, reqObj *WorkflowExecutionConditions) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkflowExecutionConditions
func (r *WorkflowExecutionConditionsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// WorkflowExecutionTriggerRequestBuilder is request builder for WorkflowExecutionTrigger
type WorkflowExecutionTriggerRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkflowExecutionTriggerRequest
func (b *WorkflowExecutionTriggerRequestBuilder) Request() *WorkflowExecutionTriggerRequest {
	return &WorkflowExecutionTriggerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkflowExecutionTriggerRequest is request for WorkflowExecutionTrigger
type WorkflowExecutionTriggerRequest struct{ BaseRequest }

// Get performs GET request for WorkflowExecutionTrigger
func (r *WorkflowExecutionTriggerRequest) Get(ctx context.Context) (resObj *WorkflowExecutionTrigger, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkflowExecutionTrigger
func (r *WorkflowExecutionTriggerRequest) Update(ctx context.Context, reqObj *WorkflowExecutionTrigger) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkflowExecutionTrigger
func (r *WorkflowExecutionTriggerRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// WorkflowTemplateRequestBuilder is request builder for WorkflowTemplate
type WorkflowTemplateRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkflowTemplateRequest
func (b *WorkflowTemplateRequestBuilder) Request() *WorkflowTemplateRequest {
	return &WorkflowTemplateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkflowTemplateRequest is request for WorkflowTemplate
type WorkflowTemplateRequest struct{ BaseRequest }

// Get performs GET request for WorkflowTemplate
func (r *WorkflowTemplateRequest) Get(ctx context.Context) (resObj *WorkflowTemplate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkflowTemplate
func (r *WorkflowTemplateRequest) Update(ctx context.Context, reqObj *WorkflowTemplate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkflowTemplate
func (r *WorkflowTemplateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// WorkflowVersionRequestBuilder is request builder for WorkflowVersion
type WorkflowVersionRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkflowVersionRequest
func (b *WorkflowVersionRequestBuilder) Request() *WorkflowVersionRequest {
	return &WorkflowVersionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkflowVersionRequest is request for WorkflowVersion
type WorkflowVersionRequest struct{ BaseRequest }

// Get performs GET request for WorkflowVersion
func (r *WorkflowVersionRequest) Get(ctx context.Context) (resObj *WorkflowVersion, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkflowVersion
func (r *WorkflowVersionRequest) Update(ctx context.Context, reqObj *WorkflowVersion) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkflowVersion
func (r *WorkflowVersionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
