// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// BookingAppointmentCancelRequestParameter undocumented
type BookingAppointmentCancelRequestParameter struct {
	// CancellationMessage undocumented
	CancellationMessage *string `json:"cancellationMessage,omitempty"`
}

// BookingBusinessPublishRequestParameter undocumented
type BookingBusinessPublishRequestParameter struct {
}

// BookingBusinessUnpublishRequestParameter undocumented
type BookingBusinessUnpublishRequestParameter struct {
}

// Appointments returns request builder for BookingAppointment collection
func (b *BookingBusinessRequestBuilder) Appointments() *BookingBusinessAppointmentsCollectionRequestBuilder {
	bb := &BookingBusinessAppointmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/appointments"
	return bb
}

// BookingBusinessAppointmentsCollectionRequestBuilder is request builder for BookingAppointment collection
type BookingBusinessAppointmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for BookingAppointment collection
func (b *BookingBusinessAppointmentsCollectionRequestBuilder) Request() *BookingBusinessAppointmentsCollectionRequest {
	return &BookingBusinessAppointmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for BookingAppointment item
func (b *BookingBusinessAppointmentsCollectionRequestBuilder) ID(id string) *BookingAppointmentRequestBuilder {
	bb := &BookingAppointmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// BookingBusinessAppointmentsCollectionRequest is request for BookingAppointment collection
type BookingBusinessAppointmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for BookingAppointment collection
func (r *BookingBusinessAppointmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]BookingAppointment, error) {
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
	var values []BookingAppointment
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
			value  []BookingAppointment
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

// GetN performs GET request for BookingAppointment collection, max N pages
func (r *BookingBusinessAppointmentsCollectionRequest) GetN(ctx context.Context, n int) ([]BookingAppointment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for BookingAppointment collection
func (r *BookingBusinessAppointmentsCollectionRequest) Get(ctx context.Context) ([]BookingAppointment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for BookingAppointment collection
func (r *BookingBusinessAppointmentsCollectionRequest) Add(ctx context.Context, reqObj *BookingAppointment) (resObj *BookingAppointment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// CalendarView returns request builder for BookingAppointment collection
func (b *BookingBusinessRequestBuilder) CalendarView() *BookingBusinessCalendarViewCollectionRequestBuilder {
	bb := &BookingBusinessCalendarViewCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/calendarView"
	return bb
}

// BookingBusinessCalendarViewCollectionRequestBuilder is request builder for BookingAppointment collection
type BookingBusinessCalendarViewCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for BookingAppointment collection
func (b *BookingBusinessCalendarViewCollectionRequestBuilder) Request() *BookingBusinessCalendarViewCollectionRequest {
	return &BookingBusinessCalendarViewCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for BookingAppointment item
func (b *BookingBusinessCalendarViewCollectionRequestBuilder) ID(id string) *BookingAppointmentRequestBuilder {
	bb := &BookingAppointmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// BookingBusinessCalendarViewCollectionRequest is request for BookingAppointment collection
type BookingBusinessCalendarViewCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for BookingAppointment collection
func (r *BookingBusinessCalendarViewCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]BookingAppointment, error) {
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
	var values []BookingAppointment
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
			value  []BookingAppointment
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

// GetN performs GET request for BookingAppointment collection, max N pages
func (r *BookingBusinessCalendarViewCollectionRequest) GetN(ctx context.Context, n int) ([]BookingAppointment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for BookingAppointment collection
func (r *BookingBusinessCalendarViewCollectionRequest) Get(ctx context.Context) ([]BookingAppointment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for BookingAppointment collection
func (r *BookingBusinessCalendarViewCollectionRequest) Add(ctx context.Context, reqObj *BookingAppointment) (resObj *BookingAppointment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Customers returns request builder for BookingCustomer collection
func (b *BookingBusinessRequestBuilder) Customers() *BookingBusinessCustomersCollectionRequestBuilder {
	bb := &BookingBusinessCustomersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/customers"
	return bb
}

// BookingBusinessCustomersCollectionRequestBuilder is request builder for BookingCustomer collection
type BookingBusinessCustomersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for BookingCustomer collection
func (b *BookingBusinessCustomersCollectionRequestBuilder) Request() *BookingBusinessCustomersCollectionRequest {
	return &BookingBusinessCustomersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for BookingCustomer item
func (b *BookingBusinessCustomersCollectionRequestBuilder) ID(id string) *BookingCustomerRequestBuilder {
	bb := &BookingCustomerRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// BookingBusinessCustomersCollectionRequest is request for BookingCustomer collection
type BookingBusinessCustomersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for BookingCustomer collection
func (r *BookingBusinessCustomersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]BookingCustomer, error) {
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
	var values []BookingCustomer
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
			value  []BookingCustomer
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

// GetN performs GET request for BookingCustomer collection, max N pages
func (r *BookingBusinessCustomersCollectionRequest) GetN(ctx context.Context, n int) ([]BookingCustomer, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for BookingCustomer collection
func (r *BookingBusinessCustomersCollectionRequest) Get(ctx context.Context) ([]BookingCustomer, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for BookingCustomer collection
func (r *BookingBusinessCustomersCollectionRequest) Add(ctx context.Context, reqObj *BookingCustomer) (resObj *BookingCustomer, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Services returns request builder for BookingService collection
func (b *BookingBusinessRequestBuilder) Services() *BookingBusinessServicesCollectionRequestBuilder {
	bb := &BookingBusinessServicesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/services"
	return bb
}

// BookingBusinessServicesCollectionRequestBuilder is request builder for BookingService collection
type BookingBusinessServicesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for BookingService collection
func (b *BookingBusinessServicesCollectionRequestBuilder) Request() *BookingBusinessServicesCollectionRequest {
	return &BookingBusinessServicesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for BookingService item
func (b *BookingBusinessServicesCollectionRequestBuilder) ID(id string) *BookingServiceRequestBuilder {
	bb := &BookingServiceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// BookingBusinessServicesCollectionRequest is request for BookingService collection
type BookingBusinessServicesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for BookingService collection
func (r *BookingBusinessServicesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]BookingService, error) {
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
	var values []BookingService
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
			value  []BookingService
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

// GetN performs GET request for BookingService collection, max N pages
func (r *BookingBusinessServicesCollectionRequest) GetN(ctx context.Context, n int) ([]BookingService, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for BookingService collection
func (r *BookingBusinessServicesCollectionRequest) Get(ctx context.Context) ([]BookingService, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for BookingService collection
func (r *BookingBusinessServicesCollectionRequest) Add(ctx context.Context, reqObj *BookingService) (resObj *BookingService, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// StaffMembers returns request builder for BookingStaffMember collection
func (b *BookingBusinessRequestBuilder) StaffMembers() *BookingBusinessStaffMembersCollectionRequestBuilder {
	bb := &BookingBusinessStaffMembersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/staffMembers"
	return bb
}

// BookingBusinessStaffMembersCollectionRequestBuilder is request builder for BookingStaffMember collection
type BookingBusinessStaffMembersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for BookingStaffMember collection
func (b *BookingBusinessStaffMembersCollectionRequestBuilder) Request() *BookingBusinessStaffMembersCollectionRequest {
	return &BookingBusinessStaffMembersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for BookingStaffMember item
func (b *BookingBusinessStaffMembersCollectionRequestBuilder) ID(id string) *BookingStaffMemberRequestBuilder {
	bb := &BookingStaffMemberRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// BookingBusinessStaffMembersCollectionRequest is request for BookingStaffMember collection
type BookingBusinessStaffMembersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for BookingStaffMember collection
func (r *BookingBusinessStaffMembersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]BookingStaffMember, error) {
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
	var values []BookingStaffMember
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
			value  []BookingStaffMember
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

// GetN performs GET request for BookingStaffMember collection, max N pages
func (r *BookingBusinessStaffMembersCollectionRequest) GetN(ctx context.Context, n int) ([]BookingStaffMember, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for BookingStaffMember collection
func (r *BookingBusinessStaffMembersCollectionRequest) Get(ctx context.Context) ([]BookingStaffMember, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for BookingStaffMember collection
func (r *BookingBusinessStaffMembersCollectionRequest) Add(ctx context.Context, reqObj *BookingStaffMember) (resObj *BookingStaffMember, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
