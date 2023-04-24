// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// PersonRequestBuilder is request builder for Person
type PersonRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonRequest
func (b *PersonRequestBuilder) Request() *PersonRequest {
	return &PersonRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PersonRequest is request for Person
type PersonRequest struct{ BaseRequest }

// Get performs GET request for Person
func (r *PersonRequest) Get(ctx context.Context) (resObj *Person, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Person
func (r *PersonRequest) Update(ctx context.Context, reqObj *Person) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Person
func (r *PersonRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PersonAnnotationRequestBuilder is request builder for PersonAnnotation
type PersonAnnotationRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonAnnotationRequest
func (b *PersonAnnotationRequestBuilder) Request() *PersonAnnotationRequest {
	return &PersonAnnotationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PersonAnnotationRequest is request for PersonAnnotation
type PersonAnnotationRequest struct{ BaseRequest }

// Get performs GET request for PersonAnnotation
func (r *PersonAnnotationRequest) Get(ctx context.Context) (resObj *PersonAnnotation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersonAnnotation
func (r *PersonAnnotationRequest) Update(ctx context.Context, reqObj *PersonAnnotation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersonAnnotation
func (r *PersonAnnotationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PersonAnnualEventRequestBuilder is request builder for PersonAnnualEvent
type PersonAnnualEventRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonAnnualEventRequest
func (b *PersonAnnualEventRequestBuilder) Request() *PersonAnnualEventRequest {
	return &PersonAnnualEventRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PersonAnnualEventRequest is request for PersonAnnualEvent
type PersonAnnualEventRequest struct{ BaseRequest }

// Get performs GET request for PersonAnnualEvent
func (r *PersonAnnualEventRequest) Get(ctx context.Context) (resObj *PersonAnnualEvent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersonAnnualEvent
func (r *PersonAnnualEventRequest) Update(ctx context.Context, reqObj *PersonAnnualEvent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersonAnnualEvent
func (r *PersonAnnualEventRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PersonAwardRequestBuilder is request builder for PersonAward
type PersonAwardRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonAwardRequest
func (b *PersonAwardRequestBuilder) Request() *PersonAwardRequest {
	return &PersonAwardRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PersonAwardRequest is request for PersonAward
type PersonAwardRequest struct{ BaseRequest }

// Get performs GET request for PersonAward
func (r *PersonAwardRequest) Get(ctx context.Context) (resObj *PersonAward, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersonAward
func (r *PersonAwardRequest) Update(ctx context.Context, reqObj *PersonAward) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersonAward
func (r *PersonAwardRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PersonCertificationRequestBuilder is request builder for PersonCertification
type PersonCertificationRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonCertificationRequest
func (b *PersonCertificationRequestBuilder) Request() *PersonCertificationRequest {
	return &PersonCertificationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PersonCertificationRequest is request for PersonCertification
type PersonCertificationRequest struct{ BaseRequest }

// Get performs GET request for PersonCertification
func (r *PersonCertificationRequest) Get(ctx context.Context) (resObj *PersonCertification, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersonCertification
func (r *PersonCertificationRequest) Update(ctx context.Context, reqObj *PersonCertification) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersonCertification
func (r *PersonCertificationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PersonDataSourceRequestBuilder is request builder for PersonDataSource
type PersonDataSourceRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonDataSourceRequest
func (b *PersonDataSourceRequestBuilder) Request() *PersonDataSourceRequest {
	return &PersonDataSourceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PersonDataSourceRequest is request for PersonDataSource
type PersonDataSourceRequest struct{ BaseRequest }

// Get performs GET request for PersonDataSource
func (r *PersonDataSourceRequest) Get(ctx context.Context) (resObj *PersonDataSource, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersonDataSource
func (r *PersonDataSourceRequest) Update(ctx context.Context, reqObj *PersonDataSource) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersonDataSource
func (r *PersonDataSourceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PersonDataSourcesRequestBuilder is request builder for PersonDataSources
type PersonDataSourcesRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonDataSourcesRequest
func (b *PersonDataSourcesRequestBuilder) Request() *PersonDataSourcesRequest {
	return &PersonDataSourcesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PersonDataSourcesRequest is request for PersonDataSources
type PersonDataSourcesRequest struct{ BaseRequest }

// Get performs GET request for PersonDataSources
func (r *PersonDataSourcesRequest) Get(ctx context.Context) (resObj *PersonDataSources, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersonDataSources
func (r *PersonDataSourcesRequest) Update(ctx context.Context, reqObj *PersonDataSources) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersonDataSources
func (r *PersonDataSourcesRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PersonExtensionRequestBuilder is request builder for PersonExtension
type PersonExtensionRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonExtensionRequest
func (b *PersonExtensionRequestBuilder) Request() *PersonExtensionRequest {
	return &PersonExtensionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PersonExtensionRequest is request for PersonExtension
type PersonExtensionRequest struct{ BaseRequest }

// Get performs GET request for PersonExtension
func (r *PersonExtensionRequest) Get(ctx context.Context) (resObj *PersonExtension, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersonExtension
func (r *PersonExtensionRequest) Update(ctx context.Context, reqObj *PersonExtension) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersonExtension
func (r *PersonExtensionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PersonInterestRequestBuilder is request builder for PersonInterest
type PersonInterestRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonInterestRequest
func (b *PersonInterestRequestBuilder) Request() *PersonInterestRequest {
	return &PersonInterestRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PersonInterestRequest is request for PersonInterest
type PersonInterestRequest struct{ BaseRequest }

// Get performs GET request for PersonInterest
func (r *PersonInterestRequest) Get(ctx context.Context) (resObj *PersonInterest, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersonInterest
func (r *PersonInterestRequest) Update(ctx context.Context, reqObj *PersonInterest) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersonInterest
func (r *PersonInterestRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PersonNameRequestBuilder is request builder for PersonName
type PersonNameRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonNameRequest
func (b *PersonNameRequestBuilder) Request() *PersonNameRequest {
	return &PersonNameRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PersonNameRequest is request for PersonName
type PersonNameRequest struct{ BaseRequest }

// Get performs GET request for PersonName
func (r *PersonNameRequest) Get(ctx context.Context) (resObj *PersonName, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersonName
func (r *PersonNameRequest) Update(ctx context.Context, reqObj *PersonName) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersonName
func (r *PersonNameRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PersonNamePronounciationRequestBuilder is request builder for PersonNamePronounciation
type PersonNamePronounciationRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonNamePronounciationRequest
func (b *PersonNamePronounciationRequestBuilder) Request() *PersonNamePronounciationRequest {
	return &PersonNamePronounciationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PersonNamePronounciationRequest is request for PersonNamePronounciation
type PersonNamePronounciationRequest struct{ BaseRequest }

// Get performs GET request for PersonNamePronounciation
func (r *PersonNamePronounciationRequest) Get(ctx context.Context) (resObj *PersonNamePronounciation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersonNamePronounciation
func (r *PersonNamePronounciationRequest) Update(ctx context.Context, reqObj *PersonNamePronounciation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersonNamePronounciation
func (r *PersonNamePronounciationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PersonOrGroupColumnRequestBuilder is request builder for PersonOrGroupColumn
type PersonOrGroupColumnRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonOrGroupColumnRequest
func (b *PersonOrGroupColumnRequestBuilder) Request() *PersonOrGroupColumnRequest {
	return &PersonOrGroupColumnRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PersonOrGroupColumnRequest is request for PersonOrGroupColumn
type PersonOrGroupColumnRequest struct{ BaseRequest }

// Get performs GET request for PersonOrGroupColumn
func (r *PersonOrGroupColumnRequest) Get(ctx context.Context) (resObj *PersonOrGroupColumn, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersonOrGroupColumn
func (r *PersonOrGroupColumnRequest) Update(ctx context.Context, reqObj *PersonOrGroupColumn) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersonOrGroupColumn
func (r *PersonOrGroupColumnRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PersonResponsibilityRequestBuilder is request builder for PersonResponsibility
type PersonResponsibilityRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonResponsibilityRequest
func (b *PersonResponsibilityRequestBuilder) Request() *PersonResponsibilityRequest {
	return &PersonResponsibilityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PersonResponsibilityRequest is request for PersonResponsibility
type PersonResponsibilityRequest struct{ BaseRequest }

// Get performs GET request for PersonResponsibility
func (r *PersonResponsibilityRequest) Get(ctx context.Context) (resObj *PersonResponsibility, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersonResponsibility
func (r *PersonResponsibilityRequest) Update(ctx context.Context, reqObj *PersonResponsibility) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersonResponsibility
func (r *PersonResponsibilityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PersonWebsiteRequestBuilder is request builder for PersonWebsite
type PersonWebsiteRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonWebsiteRequest
func (b *PersonWebsiteRequestBuilder) Request() *PersonWebsiteRequest {
	return &PersonWebsiteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PersonWebsiteRequest is request for PersonWebsite
type PersonWebsiteRequest struct{ BaseRequest }

// Get performs GET request for PersonWebsite
func (r *PersonWebsiteRequest) Get(ctx context.Context) (resObj *PersonWebsite, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersonWebsite
func (r *PersonWebsiteRequest) Update(ctx context.Context, reqObj *PersonWebsite) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersonWebsite
func (r *PersonWebsiteRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}