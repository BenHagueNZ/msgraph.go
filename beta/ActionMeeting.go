// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// AttendanceRecords returns request builder for AttendanceRecord collection
func (b *MeetingAttendanceReportRequestBuilder) AttendanceRecords() *MeetingAttendanceReportAttendanceRecordsCollectionRequestBuilder {
	bb := &MeetingAttendanceReportAttendanceRecordsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/attendanceRecords"
	return bb
}

// MeetingAttendanceReportAttendanceRecordsCollectionRequestBuilder is request builder for AttendanceRecord collection rcn
type MeetingAttendanceReportAttendanceRecordsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AttendanceRecord collection
func (b *MeetingAttendanceReportAttendanceRecordsCollectionRequestBuilder) Request() *MeetingAttendanceReportAttendanceRecordsCollectionRequest {
	return &MeetingAttendanceReportAttendanceRecordsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AttendanceRecord item
func (b *MeetingAttendanceReportAttendanceRecordsCollectionRequestBuilder) ID(id string) *AttendanceRecordRequestBuilder {
	bb := &AttendanceRecordRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MeetingAttendanceReportAttendanceRecordsCollectionRequest is request for AttendanceRecord collection
type MeetingAttendanceReportAttendanceRecordsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AttendanceRecord collection
func (r *MeetingAttendanceReportAttendanceRecordsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AttendanceRecord, error) {
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
	var values []AttendanceRecord
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
			value  []AttendanceRecord
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

// GetN performs GET request for AttendanceRecord collection, max N pages
func (r *MeetingAttendanceReportAttendanceRecordsCollectionRequest) GetN(ctx context.Context, n int) ([]AttendanceRecord, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AttendanceRecord collection
func (r *MeetingAttendanceReportAttendanceRecordsCollectionRequest) Get(ctx context.Context) ([]AttendanceRecord, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AttendanceRecord collection
func (r *MeetingAttendanceReportAttendanceRecordsCollectionRequest) Add(ctx context.Context, reqObj *AttendanceRecord) (resObj *AttendanceRecord, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// CustomQuestions returns request builder for MeetingRegistrationQuestion collection
func (b *MeetingRegistrationRequestBuilder) CustomQuestions() *MeetingRegistrationCustomQuestionsCollectionRequestBuilder {
	bb := &MeetingRegistrationCustomQuestionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/customQuestions"
	return bb
}

// MeetingRegistrationCustomQuestionsCollectionRequestBuilder is request builder for MeetingRegistrationQuestion collection rcn
type MeetingRegistrationCustomQuestionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MeetingRegistrationQuestion collection
func (b *MeetingRegistrationCustomQuestionsCollectionRequestBuilder) Request() *MeetingRegistrationCustomQuestionsCollectionRequest {
	return &MeetingRegistrationCustomQuestionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MeetingRegistrationQuestion item
func (b *MeetingRegistrationCustomQuestionsCollectionRequestBuilder) ID(id string) *MeetingRegistrationQuestionRequestBuilder {
	bb := &MeetingRegistrationQuestionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MeetingRegistrationCustomQuestionsCollectionRequest is request for MeetingRegistrationQuestion collection
type MeetingRegistrationCustomQuestionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MeetingRegistrationQuestion collection
func (r *MeetingRegistrationCustomQuestionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MeetingRegistrationQuestion, error) {
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
	var values []MeetingRegistrationQuestion
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
			value  []MeetingRegistrationQuestion
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

// GetN performs GET request for MeetingRegistrationQuestion collection, max N pages
func (r *MeetingRegistrationCustomQuestionsCollectionRequest) GetN(ctx context.Context, n int) ([]MeetingRegistrationQuestion, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MeetingRegistrationQuestion collection
func (r *MeetingRegistrationCustomQuestionsCollectionRequest) Get(ctx context.Context) ([]MeetingRegistrationQuestion, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MeetingRegistrationQuestion collection
func (r *MeetingRegistrationCustomQuestionsCollectionRequest) Add(ctx context.Context, reqObj *MeetingRegistrationQuestion) (resObj *MeetingRegistrationQuestion, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Registrants returns request builder for MeetingRegistrantBase collection
func (b *MeetingRegistrationBaseRequestBuilder) Registrants() *MeetingRegistrationBaseRegistrantsCollectionRequestBuilder {
	bb := &MeetingRegistrationBaseRegistrantsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/registrants"
	return bb
}

// MeetingRegistrationBaseRegistrantsCollectionRequestBuilder is request builder for MeetingRegistrantBase collection rcn
type MeetingRegistrationBaseRegistrantsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MeetingRegistrantBase collection
func (b *MeetingRegistrationBaseRegistrantsCollectionRequestBuilder) Request() *MeetingRegistrationBaseRegistrantsCollectionRequest {
	return &MeetingRegistrationBaseRegistrantsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MeetingRegistrantBase item
func (b *MeetingRegistrationBaseRegistrantsCollectionRequestBuilder) ID(id string) *MeetingRegistrantBaseRequestBuilder {
	bb := &MeetingRegistrantBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MeetingRegistrationBaseRegistrantsCollectionRequest is request for MeetingRegistrantBase collection
type MeetingRegistrationBaseRegistrantsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MeetingRegistrantBase collection
func (r *MeetingRegistrationBaseRegistrantsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MeetingRegistrantBase, error) {
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
	var values []MeetingRegistrantBase
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
			value  []MeetingRegistrantBase
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

// GetN performs GET request for MeetingRegistrantBase collection, max N pages
func (r *MeetingRegistrationBaseRegistrantsCollectionRequest) GetN(ctx context.Context, n int) ([]MeetingRegistrantBase, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MeetingRegistrantBase collection
func (r *MeetingRegistrationBaseRegistrantsCollectionRequest) Get(ctx context.Context) ([]MeetingRegistrantBase, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MeetingRegistrantBase collection
func (r *MeetingRegistrationBaseRegistrantsCollectionRequest) Add(ctx context.Context, reqObj *MeetingRegistrantBase) (resObj *MeetingRegistrantBase, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// MeetingActivityStatistics returns request builder for MeetingActivityStatistics collection
func (b *UserAnalyticsActivityStatisticsCollectionRequestBuilder) MeetingActivityStatistics() *UserAnalyticsActivityStatisticsCollectionMeetingActivityStatisticsCollectionRequestBuilder {
	bb := &UserAnalyticsActivityStatisticsCollectionMeetingActivityStatisticsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// UserAnalyticsActivityStatisticsCollectionMeetingActivityStatisticsCollectionRequestBuilder is request builder for MeetingActivityStatistics collection rcn
type UserAnalyticsActivityStatisticsCollectionMeetingActivityStatisticsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MeetingActivityStatistics collection
func (b *UserAnalyticsActivityStatisticsCollectionMeetingActivityStatisticsCollectionRequestBuilder) Request() *UserAnalyticsActivityStatisticsCollectionMeetingActivityStatisticsCollectionRequest {
	return &UserAnalyticsActivityStatisticsCollectionMeetingActivityStatisticsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MeetingActivityStatistics item
func (b *UserAnalyticsActivityStatisticsCollectionMeetingActivityStatisticsCollectionRequestBuilder) ID(id string) *MeetingActivityStatisticsRequestBuilder {
	bb := &MeetingActivityStatisticsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// UserAnalyticsActivityStatisticsCollectionMeetingActivityStatisticsCollectionRequest is request for MeetingActivityStatistics collection
type UserAnalyticsActivityStatisticsCollectionMeetingActivityStatisticsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MeetingActivityStatistics collection
func (r *UserAnalyticsActivityStatisticsCollectionMeetingActivityStatisticsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MeetingActivityStatistics, error) {
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
	var values []MeetingActivityStatistics
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
			value  []MeetingActivityStatistics
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

// GetN performs GET request for MeetingActivityStatistics collection, max N pages
func (r *UserAnalyticsActivityStatisticsCollectionMeetingActivityStatisticsCollectionRequest) GetN(ctx context.Context, n int) ([]MeetingActivityStatistics, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MeetingActivityStatistics collection
func (r *UserAnalyticsActivityStatisticsCollectionMeetingActivityStatisticsCollectionRequest) Get(ctx context.Context) ([]MeetingActivityStatistics, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MeetingActivityStatistics collection
func (r *UserAnalyticsActivityStatisticsCollectionMeetingActivityStatisticsCollectionRequest) Add(ctx context.Context, reqObj *MeetingActivityStatistics) (resObj *MeetingActivityStatistics, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *MeetingAttendanceReportRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// MeetingRegistrant returns request builder for MeetingRegistrant collection
func (b *MeetingRegistrationBaseRegistrantsCollectionRequestBuilder) MeetingRegistrant() *MeetingRegistrationBaseRegistrantsCollectionMeetingRegistrantCollectionRequestBuilder {
	bb := &MeetingRegistrationBaseRegistrantsCollectionMeetingRegistrantCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// MeetingRegistrationBaseRegistrantsCollectionMeetingRegistrantCollectionRequestBuilder is request builder for MeetingRegistrant collection rcn
type MeetingRegistrationBaseRegistrantsCollectionMeetingRegistrantCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MeetingRegistrant collection
func (b *MeetingRegistrationBaseRegistrantsCollectionMeetingRegistrantCollectionRequestBuilder) Request() *MeetingRegistrationBaseRegistrantsCollectionMeetingRegistrantCollectionRequest {
	return &MeetingRegistrationBaseRegistrantsCollectionMeetingRegistrantCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MeetingRegistrant item
func (b *MeetingRegistrationBaseRegistrantsCollectionMeetingRegistrantCollectionRequestBuilder) ID(id string) *MeetingRegistrantRequestBuilder {
	bb := &MeetingRegistrantRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MeetingRegistrationBaseRegistrantsCollectionMeetingRegistrantCollectionRequest is request for MeetingRegistrant collection
type MeetingRegistrationBaseRegistrantsCollectionMeetingRegistrantCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MeetingRegistrant collection
func (r *MeetingRegistrationBaseRegistrantsCollectionMeetingRegistrantCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MeetingRegistrant, error) {
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
	var values []MeetingRegistrant
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
			value  []MeetingRegistrant
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

// GetN performs GET request for MeetingRegistrant collection, max N pages
func (r *MeetingRegistrationBaseRegistrantsCollectionMeetingRegistrantCollectionRequest) GetN(ctx context.Context, n int) ([]MeetingRegistrant, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MeetingRegistrant collection
func (r *MeetingRegistrationBaseRegistrantsCollectionMeetingRegistrantCollectionRequest) Get(ctx context.Context) ([]MeetingRegistrant, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MeetingRegistrant collection
func (r *MeetingRegistrationBaseRegistrantsCollectionMeetingRegistrantCollectionRequest) Add(ctx context.Context, reqObj *MeetingRegistrant) (resObj *MeetingRegistrant, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *MeetingRegistrationRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *MeetingRegistrationBaseRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *MeetingRegistrationQuestionRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
