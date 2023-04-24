// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// OMASettingRequestBuilder is request builder for OMASetting
type OMASettingRequestBuilder struct{ BaseRequestBuilder }

// Request returns OMASettingRequest
func (b *OMASettingRequestBuilder) Request() *OMASettingRequest {
	return &OMASettingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OMASettingRequest is request for OMASetting
type OMASettingRequest struct{ BaseRequest }

// Get performs GET request for OMASetting
func (r *OMASettingRequest) Get(ctx context.Context) (resObj *OMASetting, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OMASetting
func (r *OMASettingRequest) Update(ctx context.Context, reqObj *OMASetting) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OMASetting
func (r *OMASettingRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OMASettingBase64RequestBuilder is request builder for OMASettingBase64
type OMASettingBase64RequestBuilder struct{ BaseRequestBuilder }

// Request returns OMASettingBase64Request
func (b *OMASettingBase64RequestBuilder) Request() *OMASettingBase64Request {
	return &OMASettingBase64Request{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OMASettingBase64Request is request for OMASettingBase64
type OMASettingBase64Request struct{ BaseRequest }

// Get performs GET request for OMASettingBase64
func (r *OMASettingBase64Request) Get(ctx context.Context) (resObj *OMASettingBase64, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OMASettingBase64
func (r *OMASettingBase64Request) Update(ctx context.Context, reqObj *OMASettingBase64) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OMASettingBase64
func (r *OMASettingBase64Request) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OMASettingBooleanRequestBuilder is request builder for OMASettingBoolean
type OMASettingBooleanRequestBuilder struct{ BaseRequestBuilder }

// Request returns OMASettingBooleanRequest
func (b *OMASettingBooleanRequestBuilder) Request() *OMASettingBooleanRequest {
	return &OMASettingBooleanRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OMASettingBooleanRequest is request for OMASettingBoolean
type OMASettingBooleanRequest struct{ BaseRequest }

// Get performs GET request for OMASettingBoolean
func (r *OMASettingBooleanRequest) Get(ctx context.Context) (resObj *OMASettingBoolean, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OMASettingBoolean
func (r *OMASettingBooleanRequest) Update(ctx context.Context, reqObj *OMASettingBoolean) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OMASettingBoolean
func (r *OMASettingBooleanRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OMASettingDateTimeRequestBuilder is request builder for OMASettingDateTime
type OMASettingDateTimeRequestBuilder struct{ BaseRequestBuilder }

// Request returns OMASettingDateTimeRequest
func (b *OMASettingDateTimeRequestBuilder) Request() *OMASettingDateTimeRequest {
	return &OMASettingDateTimeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OMASettingDateTimeRequest is request for OMASettingDateTime
type OMASettingDateTimeRequest struct{ BaseRequest }

// Get performs GET request for OMASettingDateTime
func (r *OMASettingDateTimeRequest) Get(ctx context.Context) (resObj *OMASettingDateTime, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OMASettingDateTime
func (r *OMASettingDateTimeRequest) Update(ctx context.Context, reqObj *OMASettingDateTime) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OMASettingDateTime
func (r *OMASettingDateTimeRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OMASettingFloatingPointRequestBuilder is request builder for OMASettingFloatingPoint
type OMASettingFloatingPointRequestBuilder struct{ BaseRequestBuilder }

// Request returns OMASettingFloatingPointRequest
func (b *OMASettingFloatingPointRequestBuilder) Request() *OMASettingFloatingPointRequest {
	return &OMASettingFloatingPointRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OMASettingFloatingPointRequest is request for OMASettingFloatingPoint
type OMASettingFloatingPointRequest struct{ BaseRequest }

// Get performs GET request for OMASettingFloatingPoint
func (r *OMASettingFloatingPointRequest) Get(ctx context.Context) (resObj *OMASettingFloatingPoint, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OMASettingFloatingPoint
func (r *OMASettingFloatingPointRequest) Update(ctx context.Context, reqObj *OMASettingFloatingPoint) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OMASettingFloatingPoint
func (r *OMASettingFloatingPointRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OMASettingIntegerRequestBuilder is request builder for OMASettingInteger
type OMASettingIntegerRequestBuilder struct{ BaseRequestBuilder }

// Request returns OMASettingIntegerRequest
func (b *OMASettingIntegerRequestBuilder) Request() *OMASettingIntegerRequest {
	return &OMASettingIntegerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OMASettingIntegerRequest is request for OMASettingInteger
type OMASettingIntegerRequest struct{ BaseRequest }

// Get performs GET request for OMASettingInteger
func (r *OMASettingIntegerRequest) Get(ctx context.Context) (resObj *OMASettingInteger, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OMASettingInteger
func (r *OMASettingIntegerRequest) Update(ctx context.Context, reqObj *OMASettingInteger) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OMASettingInteger
func (r *OMASettingIntegerRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OMASettingStringRequestBuilder is request builder for OMASettingString
type OMASettingStringRequestBuilder struct{ BaseRequestBuilder }

// Request returns OMASettingStringRequest
func (b *OMASettingStringRequestBuilder) Request() *OMASettingStringRequest {
	return &OMASettingStringRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OMASettingStringRequest is request for OMASettingString
type OMASettingStringRequest struct{ BaseRequest }

// Get performs GET request for OMASettingString
func (r *OMASettingStringRequest) Get(ctx context.Context) (resObj *OMASettingString, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OMASettingString
func (r *OMASettingStringRequest) Update(ctx context.Context, reqObj *OMASettingString) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OMASettingString
func (r *OMASettingStringRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OMASettingStringXMLRequestBuilder is request builder for OMASettingStringXML
type OMASettingStringXMLRequestBuilder struct{ BaseRequestBuilder }

// Request returns OMASettingStringXMLRequest
func (b *OMASettingStringXMLRequestBuilder) Request() *OMASettingStringXMLRequest {
	return &OMASettingStringXMLRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OMASettingStringXMLRequest is request for OMASettingStringXML
type OMASettingStringXMLRequest struct{ BaseRequest }

// Get performs GET request for OMASettingStringXML
func (r *OMASettingStringXMLRequest) Get(ctx context.Context) (resObj *OMASettingStringXML, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OMASettingStringXML
func (r *OMASettingStringXMLRequest) Update(ctx context.Context, reqObj *OMASettingStringXML) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OMASettingStringXML
func (r *OMASettingStringXMLRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}