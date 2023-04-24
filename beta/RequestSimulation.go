// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SimulationRequestBuilder is request builder for Simulation
type SimulationRequestBuilder struct{ BaseRequestBuilder }

// Request returns SimulationRequest
func (b *SimulationRequestBuilder) Request() *SimulationRequest {
	return &SimulationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SimulationRequest is request for Simulation
type SimulationRequest struct{ BaseRequest }

// Get performs GET request for Simulation
func (r *SimulationRequest) Get(ctx context.Context) (resObj *Simulation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Simulation
func (r *SimulationRequest) Update(ctx context.Context, reqObj *Simulation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Simulation
func (r *SimulationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SimulationAutomationRequestBuilder is request builder for SimulationAutomation
type SimulationAutomationRequestBuilder struct{ BaseRequestBuilder }

// Request returns SimulationAutomationRequest
func (b *SimulationAutomationRequestBuilder) Request() *SimulationAutomationRequest {
	return &SimulationAutomationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SimulationAutomationRequest is request for SimulationAutomation
type SimulationAutomationRequest struct{ BaseRequest }

// Get performs GET request for SimulationAutomation
func (r *SimulationAutomationRequest) Get(ctx context.Context) (resObj *SimulationAutomation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SimulationAutomation
func (r *SimulationAutomationRequest) Update(ctx context.Context, reqObj *SimulationAutomation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SimulationAutomation
func (r *SimulationAutomationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SimulationAutomationRunRequestBuilder is request builder for SimulationAutomationRun
type SimulationAutomationRunRequestBuilder struct{ BaseRequestBuilder }

// Request returns SimulationAutomationRunRequest
func (b *SimulationAutomationRunRequestBuilder) Request() *SimulationAutomationRunRequest {
	return &SimulationAutomationRunRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SimulationAutomationRunRequest is request for SimulationAutomationRun
type SimulationAutomationRunRequest struct{ BaseRequest }

// Get performs GET request for SimulationAutomationRun
func (r *SimulationAutomationRunRequest) Get(ctx context.Context) (resObj *SimulationAutomationRun, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SimulationAutomationRun
func (r *SimulationAutomationRunRequest) Update(ctx context.Context, reqObj *SimulationAutomationRun) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SimulationAutomationRun
func (r *SimulationAutomationRunRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SimulationEventRequestBuilder is request builder for SimulationEvent
type SimulationEventRequestBuilder struct{ BaseRequestBuilder }

// Request returns SimulationEventRequest
func (b *SimulationEventRequestBuilder) Request() *SimulationEventRequest {
	return &SimulationEventRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SimulationEventRequest is request for SimulationEvent
type SimulationEventRequest struct{ BaseRequest }

// Get performs GET request for SimulationEvent
func (r *SimulationEventRequest) Get(ctx context.Context) (resObj *SimulationEvent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SimulationEvent
func (r *SimulationEventRequest) Update(ctx context.Context, reqObj *SimulationEvent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SimulationEvent
func (r *SimulationEventRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SimulationEventsContentRequestBuilder is request builder for SimulationEventsContent
type SimulationEventsContentRequestBuilder struct{ BaseRequestBuilder }

// Request returns SimulationEventsContentRequest
func (b *SimulationEventsContentRequestBuilder) Request() *SimulationEventsContentRequest {
	return &SimulationEventsContentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SimulationEventsContentRequest is request for SimulationEventsContent
type SimulationEventsContentRequest struct{ BaseRequest }

// Get performs GET request for SimulationEventsContent
func (r *SimulationEventsContentRequest) Get(ctx context.Context) (resObj *SimulationEventsContent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SimulationEventsContent
func (r *SimulationEventsContentRequest) Update(ctx context.Context, reqObj *SimulationEventsContent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SimulationEventsContent
func (r *SimulationEventsContentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SimulationReportRequestBuilder is request builder for SimulationReport
type SimulationReportRequestBuilder struct{ BaseRequestBuilder }

// Request returns SimulationReportRequest
func (b *SimulationReportRequestBuilder) Request() *SimulationReportRequest {
	return &SimulationReportRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SimulationReportRequest is request for SimulationReport
type SimulationReportRequest struct{ BaseRequest }

// Get performs GET request for SimulationReport
func (r *SimulationReportRequest) Get(ctx context.Context) (resObj *SimulationReport, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SimulationReport
func (r *SimulationReportRequest) Update(ctx context.Context, reqObj *SimulationReport) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SimulationReport
func (r *SimulationReportRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SimulationReportOverviewRequestBuilder is request builder for SimulationReportOverview
type SimulationReportOverviewRequestBuilder struct{ BaseRequestBuilder }

// Request returns SimulationReportOverviewRequest
func (b *SimulationReportOverviewRequestBuilder) Request() *SimulationReportOverviewRequest {
	return &SimulationReportOverviewRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SimulationReportOverviewRequest is request for SimulationReportOverview
type SimulationReportOverviewRequest struct{ BaseRequest }

// Get performs GET request for SimulationReportOverview
func (r *SimulationReportOverviewRequest) Get(ctx context.Context) (resObj *SimulationReportOverview, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SimulationReportOverview
func (r *SimulationReportOverviewRequest) Update(ctx context.Context, reqObj *SimulationReportOverview) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SimulationReportOverview
func (r *SimulationReportOverviewRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
