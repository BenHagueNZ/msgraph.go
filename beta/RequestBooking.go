// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// BookingAppointmentRequestBuilder is request builder for BookingAppointment
type BookingAppointmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingAppointmentRequest
func (b *BookingAppointmentRequestBuilder) Request() *BookingAppointmentRequest {
	return &BookingAppointmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingAppointmentRequest is request for BookingAppointment
type BookingAppointmentRequest struct{ BaseRequest }

// Get performs GET request for BookingAppointment
func (r *BookingAppointmentRequest) Get(ctx context.Context) (resObj *BookingAppointment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingAppointment
func (r *BookingAppointmentRequest) Update(ctx context.Context, reqObj *BookingAppointment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingAppointment
func (r *BookingAppointmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingBusinessRequestBuilder is request builder for BookingBusiness
type BookingBusinessRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingBusinessRequest
func (b *BookingBusinessRequestBuilder) Request() *BookingBusinessRequest {
	return &BookingBusinessRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingBusinessRequest is request for BookingBusiness
type BookingBusinessRequest struct{ BaseRequest }

// Get performs GET request for BookingBusiness
func (r *BookingBusinessRequest) Get(ctx context.Context) (resObj *BookingBusiness, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingBusiness
func (r *BookingBusinessRequest) Update(ctx context.Context, reqObj *BookingBusiness) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingBusiness
func (r *BookingBusinessRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingCurrencyRequestBuilder is request builder for BookingCurrency
type BookingCurrencyRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingCurrencyRequest
func (b *BookingCurrencyRequestBuilder) Request() *BookingCurrencyRequest {
	return &BookingCurrencyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingCurrencyRequest is request for BookingCurrency
type BookingCurrencyRequest struct{ BaseRequest }

// Get performs GET request for BookingCurrency
func (r *BookingCurrencyRequest) Get(ctx context.Context) (resObj *BookingCurrency, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingCurrency
func (r *BookingCurrencyRequest) Update(ctx context.Context, reqObj *BookingCurrency) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingCurrency
func (r *BookingCurrencyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingCustomQuestionRequestBuilder is request builder for BookingCustomQuestion
type BookingCustomQuestionRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingCustomQuestionRequest
func (b *BookingCustomQuestionRequestBuilder) Request() *BookingCustomQuestionRequest {
	return &BookingCustomQuestionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingCustomQuestionRequest is request for BookingCustomQuestion
type BookingCustomQuestionRequest struct{ BaseRequest }

// Get performs GET request for BookingCustomQuestion
func (r *BookingCustomQuestionRequest) Get(ctx context.Context) (resObj *BookingCustomQuestion, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingCustomQuestion
func (r *BookingCustomQuestionRequest) Update(ctx context.Context, reqObj *BookingCustomQuestion) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingCustomQuestion
func (r *BookingCustomQuestionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingCustomerRequestBuilder is request builder for BookingCustomer
type BookingCustomerRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingCustomerRequest
func (b *BookingCustomerRequestBuilder) Request() *BookingCustomerRequest {
	return &BookingCustomerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingCustomerRequest is request for BookingCustomer
type BookingCustomerRequest struct{ BaseRequest }

// Get performs GET request for BookingCustomer
func (r *BookingCustomerRequest) Get(ctx context.Context) (resObj *BookingCustomer, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingCustomer
func (r *BookingCustomerRequest) Update(ctx context.Context, reqObj *BookingCustomer) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingCustomer
func (r *BookingCustomerRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingCustomerInformationRequestBuilder is request builder for BookingCustomerInformation
type BookingCustomerInformationRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingCustomerInformationRequest
func (b *BookingCustomerInformationRequestBuilder) Request() *BookingCustomerInformationRequest {
	return &BookingCustomerInformationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingCustomerInformationRequest is request for BookingCustomerInformation
type BookingCustomerInformationRequest struct{ BaseRequest }

// Get performs GET request for BookingCustomerInformation
func (r *BookingCustomerInformationRequest) Get(ctx context.Context) (resObj *BookingCustomerInformation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingCustomerInformation
func (r *BookingCustomerInformationRequest) Update(ctx context.Context, reqObj *BookingCustomerInformation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingCustomerInformation
func (r *BookingCustomerInformationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingCustomerInformationBaseRequestBuilder is request builder for BookingCustomerInformationBase
type BookingCustomerInformationBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingCustomerInformationBaseRequest
func (b *BookingCustomerInformationBaseRequestBuilder) Request() *BookingCustomerInformationBaseRequest {
	return &BookingCustomerInformationBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingCustomerInformationBaseRequest is request for BookingCustomerInformationBase
type BookingCustomerInformationBaseRequest struct{ BaseRequest }

// Get performs GET request for BookingCustomerInformationBase
func (r *BookingCustomerInformationBaseRequest) Get(ctx context.Context) (resObj *BookingCustomerInformationBase, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingCustomerInformationBase
func (r *BookingCustomerInformationBaseRequest) Update(ctx context.Context, reqObj *BookingCustomerInformationBase) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingCustomerInformationBase
func (r *BookingCustomerInformationBaseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingNamedEntityRequestBuilder is request builder for BookingNamedEntity
type BookingNamedEntityRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingNamedEntityRequest
func (b *BookingNamedEntityRequestBuilder) Request() *BookingNamedEntityRequest {
	return &BookingNamedEntityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingNamedEntityRequest is request for BookingNamedEntity
type BookingNamedEntityRequest struct{ BaseRequest }

// Get performs GET request for BookingNamedEntity
func (r *BookingNamedEntityRequest) Get(ctx context.Context) (resObj *BookingNamedEntity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingNamedEntity
func (r *BookingNamedEntityRequest) Update(ctx context.Context, reqObj *BookingNamedEntity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingNamedEntity
func (r *BookingNamedEntityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingPersonRequestBuilder is request builder for BookingPerson
type BookingPersonRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingPersonRequest
func (b *BookingPersonRequestBuilder) Request() *BookingPersonRequest {
	return &BookingPersonRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingPersonRequest is request for BookingPerson
type BookingPersonRequest struct{ BaseRequest }

// Get performs GET request for BookingPerson
func (r *BookingPersonRequest) Get(ctx context.Context) (resObj *BookingPerson, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingPerson
func (r *BookingPersonRequest) Update(ctx context.Context, reqObj *BookingPerson) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingPerson
func (r *BookingPersonRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingQuestionAnswerRequestBuilder is request builder for BookingQuestionAnswer
type BookingQuestionAnswerRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingQuestionAnswerRequest
func (b *BookingQuestionAnswerRequestBuilder) Request() *BookingQuestionAnswerRequest {
	return &BookingQuestionAnswerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingQuestionAnswerRequest is request for BookingQuestionAnswer
type BookingQuestionAnswerRequest struct{ BaseRequest }

// Get performs GET request for BookingQuestionAnswer
func (r *BookingQuestionAnswerRequest) Get(ctx context.Context) (resObj *BookingQuestionAnswer, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingQuestionAnswer
func (r *BookingQuestionAnswerRequest) Update(ctx context.Context, reqObj *BookingQuestionAnswer) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingQuestionAnswer
func (r *BookingQuestionAnswerRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingQuestionAssignmentRequestBuilder is request builder for BookingQuestionAssignment
type BookingQuestionAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingQuestionAssignmentRequest
func (b *BookingQuestionAssignmentRequestBuilder) Request() *BookingQuestionAssignmentRequest {
	return &BookingQuestionAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingQuestionAssignmentRequest is request for BookingQuestionAssignment
type BookingQuestionAssignmentRequest struct{ BaseRequest }

// Get performs GET request for BookingQuestionAssignment
func (r *BookingQuestionAssignmentRequest) Get(ctx context.Context) (resObj *BookingQuestionAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingQuestionAssignment
func (r *BookingQuestionAssignmentRequest) Update(ctx context.Context, reqObj *BookingQuestionAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingQuestionAssignment
func (r *BookingQuestionAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingReminderRequestBuilder is request builder for BookingReminder
type BookingReminderRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingReminderRequest
func (b *BookingReminderRequestBuilder) Request() *BookingReminderRequest {
	return &BookingReminderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingReminderRequest is request for BookingReminder
type BookingReminderRequest struct{ BaseRequest }

// Get performs GET request for BookingReminder
func (r *BookingReminderRequest) Get(ctx context.Context) (resObj *BookingReminder, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingReminder
func (r *BookingReminderRequest) Update(ctx context.Context, reqObj *BookingReminder) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingReminder
func (r *BookingReminderRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingSchedulingPolicyRequestBuilder is request builder for BookingSchedulingPolicy
type BookingSchedulingPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingSchedulingPolicyRequest
func (b *BookingSchedulingPolicyRequestBuilder) Request() *BookingSchedulingPolicyRequest {
	return &BookingSchedulingPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingSchedulingPolicyRequest is request for BookingSchedulingPolicy
type BookingSchedulingPolicyRequest struct{ BaseRequest }

// Get performs GET request for BookingSchedulingPolicy
func (r *BookingSchedulingPolicyRequest) Get(ctx context.Context) (resObj *BookingSchedulingPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingSchedulingPolicy
func (r *BookingSchedulingPolicyRequest) Update(ctx context.Context, reqObj *BookingSchedulingPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingSchedulingPolicy
func (r *BookingSchedulingPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingServiceRequestBuilder is request builder for BookingService
type BookingServiceRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingServiceRequest
func (b *BookingServiceRequestBuilder) Request() *BookingServiceRequest {
	return &BookingServiceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingServiceRequest is request for BookingService
type BookingServiceRequest struct{ BaseRequest }

// Get performs GET request for BookingService
func (r *BookingServiceRequest) Get(ctx context.Context) (resObj *BookingService, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingService
func (r *BookingServiceRequest) Update(ctx context.Context, reqObj *BookingService) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingService
func (r *BookingServiceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingStaffMemberRequestBuilder is request builder for BookingStaffMember
type BookingStaffMemberRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingStaffMemberRequest
func (b *BookingStaffMemberRequestBuilder) Request() *BookingStaffMemberRequest {
	return &BookingStaffMemberRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingStaffMemberRequest is request for BookingStaffMember
type BookingStaffMemberRequest struct{ BaseRequest }

// Get performs GET request for BookingStaffMember
func (r *BookingStaffMemberRequest) Get(ctx context.Context) (resObj *BookingStaffMember, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingStaffMember
func (r *BookingStaffMemberRequest) Update(ctx context.Context, reqObj *BookingStaffMember) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingStaffMember
func (r *BookingStaffMemberRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingWorkHoursRequestBuilder is request builder for BookingWorkHours
type BookingWorkHoursRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingWorkHoursRequest
func (b *BookingWorkHoursRequestBuilder) Request() *BookingWorkHoursRequest {
	return &BookingWorkHoursRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingWorkHoursRequest is request for BookingWorkHours
type BookingWorkHoursRequest struct{ BaseRequest }

// Get performs GET request for BookingWorkHours
func (r *BookingWorkHoursRequest) Get(ctx context.Context) (resObj *BookingWorkHours, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingWorkHours
func (r *BookingWorkHoursRequest) Update(ctx context.Context, reqObj *BookingWorkHours) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingWorkHours
func (r *BookingWorkHoursRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingWorkTimeSlotRequestBuilder is request builder for BookingWorkTimeSlot
type BookingWorkTimeSlotRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingWorkTimeSlotRequest
func (b *BookingWorkTimeSlotRequestBuilder) Request() *BookingWorkTimeSlotRequest {
	return &BookingWorkTimeSlotRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingWorkTimeSlotRequest is request for BookingWorkTimeSlot
type BookingWorkTimeSlotRequest struct{ BaseRequest }

// Get performs GET request for BookingWorkTimeSlot
func (r *BookingWorkTimeSlotRequest) Get(ctx context.Context) (resObj *BookingWorkTimeSlot, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingWorkTimeSlot
func (r *BookingWorkTimeSlotRequest) Update(ctx context.Context, reqObj *BookingWorkTimeSlot) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingWorkTimeSlot
func (r *BookingWorkTimeSlotRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type BookingAppointmentCancelRequestBuilder struct{ BaseRequestBuilder }

// Cancel action undocumentedrav
func (b *BookingAppointmentRequestBuilder) Cancel(reqObj *BookingAppointmentCancelRequestParameter) *BookingAppointmentCancelRequestBuilder {
	bb := &BookingAppointmentCancelRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Cancel"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type BookingAppointmentCancelRequest struct{ BaseRequest }

func (b *BookingAppointmentCancelRequestBuilder) Request() *BookingAppointmentCancelRequest {
	return &BookingAppointmentCancelRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *BookingAppointmentCancelRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type BookingBusinessGetStaffAvailabilityRequestBuilder struct{ BaseRequestBuilder }

// GetStaffAvailability action undocumentedrac
func (b *BookingBusinessRequestBuilder) GetStaffAvailability(reqObj *BookingBusinessGetStaffAvailabilityRequestParameter) *BookingBusinessGetStaffAvailabilityRequestBuilder {
	bb := &BookingBusinessGetStaffAvailabilityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/GetStaffAvailability"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type BookingBusinessGetStaffAvailabilityRequest struct{ BaseRequest }

func (b *BookingBusinessGetStaffAvailabilityRequestBuilder) Request() *BookingBusinessGetStaffAvailabilityRequest {
	return &BookingBusinessGetStaffAvailabilityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *BookingBusinessGetStaffAvailabilityRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]StaffAvailabilityItem, error) {
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
	var values []StaffAvailabilityItem
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
			value  []StaffAvailabilityItem
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

func (r *BookingBusinessGetStaffAvailabilityRequest) PostN(ctx context.Context, n int) ([]StaffAvailabilityItem, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

func (r *BookingBusinessGetStaffAvailabilityRequest) Post(ctx context.Context) ([]StaffAvailabilityItem, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

type BookingBusinessPublishRequestBuilder struct{ BaseRequestBuilder }

// Publish action undocumentedrav
func (b *BookingBusinessRequestBuilder) Publish(reqObj *BookingBusinessPublishRequestParameter) *BookingBusinessPublishRequestBuilder {
	bb := &BookingBusinessPublishRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Publish"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type BookingBusinessPublishRequest struct{ BaseRequest }

func (b *BookingBusinessPublishRequestBuilder) Request() *BookingBusinessPublishRequest {
	return &BookingBusinessPublishRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *BookingBusinessPublishRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type BookingBusinessUnpublishRequestBuilder struct{ BaseRequestBuilder }

// Unpublish action undocumentedrav
func (b *BookingBusinessRequestBuilder) Unpublish(reqObj *BookingBusinessUnpublishRequestParameter) *BookingBusinessUnpublishRequestBuilder {
	bb := &BookingBusinessUnpublishRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Unpublish"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type BookingBusinessUnpublishRequest struct{ BaseRequest }

func (b *BookingBusinessUnpublishRequestBuilder) Request() *BookingBusinessUnpublishRequest {
	return &BookingBusinessUnpublishRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *BookingBusinessUnpublishRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
