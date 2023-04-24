// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// BusinessFlowRecordDecisionsRequestParameter undocumented
type BusinessFlowRecordDecisionsRequestParameter struct {
	// ReviewResult undocumented
	ReviewResult *string `json:"reviewResult,omitempty"`
	// Justification undocumented
	Justification *string `json:"justification,omitempty"`
}

// BusinessScenarioPlannerGetPlanRequestParameter undocumented
type BusinessScenarioPlannerGetPlanRequestParameter struct {
	// Target undocumented
	Target *BusinessScenarioTaskTargetBase `json:"target,omitempty"`
}

// Planner is navigation property rn
func (b *BusinessScenarioRequestBuilder) Planner() *BusinessScenarioPlannerRequestBuilder {
	bb := &BusinessScenarioPlannerRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/planner"
	return bb
}

// PlanConfiguration is navigation property rn
func (b *BusinessScenarioPlannerRequestBuilder) PlanConfiguration() *PlannerPlanConfigurationRequestBuilder {
	bb := &PlannerPlanConfigurationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/planConfiguration"
	return bb
}

// TaskConfiguration is navigation property rn
func (b *BusinessScenarioPlannerRequestBuilder) TaskConfiguration() *PlannerTaskConfigurationRequestBuilder {
	bb := &PlannerTaskConfigurationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/taskConfiguration"
	return bb
}

// Tasks returns request builder for BusinessScenarioTask collection
func (b *BusinessScenarioPlannerRequestBuilder) Tasks() *BusinessScenarioPlannerTasksCollectionRequestBuilder {
	bb := &BusinessScenarioPlannerTasksCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tasks"
	return bb
}

// BusinessScenarioPlannerTasksCollectionRequestBuilder is request builder for BusinessScenarioTask collection rcn
type BusinessScenarioPlannerTasksCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for BusinessScenarioTask collection
func (b *BusinessScenarioPlannerTasksCollectionRequestBuilder) Request() *BusinessScenarioPlannerTasksCollectionRequest {
	return &BusinessScenarioPlannerTasksCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for BusinessScenarioTask item
func (b *BusinessScenarioPlannerTasksCollectionRequestBuilder) ID(id string) *BusinessScenarioTaskRequestBuilder {
	bb := &BusinessScenarioTaskRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// BusinessScenarioPlannerTasksCollectionRequest is request for BusinessScenarioTask collection
type BusinessScenarioPlannerTasksCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for BusinessScenarioTask collection
func (r *BusinessScenarioPlannerTasksCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]BusinessScenarioTask, error) {
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
	var values []BusinessScenarioTask
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
			value  []BusinessScenarioTask
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

// GetN performs GET request for BusinessScenarioTask collection, max N pages
func (r *BusinessScenarioPlannerTasksCollectionRequest) GetN(ctx context.Context, n int) ([]BusinessScenarioTask, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for BusinessScenarioTask collection
func (r *BusinessScenarioPlannerTasksCollectionRequest) Get(ctx context.Context) ([]BusinessScenarioTask, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for BusinessScenarioTask collection
func (r *BusinessScenarioPlannerTasksCollectionRequest) Add(ctx context.Context, reqObj *BusinessScenarioTask) (resObj *BusinessScenarioTask, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *BusinessFlowRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *BusinessFlowTemplateRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *BusinessScenarioRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *BusinessScenarioPlanReferenceRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *BusinessScenarioPlannerRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// BusinessScenarioTask returns request builder for BusinessScenarioTask collection
func (b *PlannerTasksCollectionRequestBuilder) BusinessScenarioTask() *PlannerTasksCollectionBusinessScenarioTaskCollectionRequestBuilder {
	bb := &PlannerTasksCollectionBusinessScenarioTaskCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// PlannerTasksCollectionBusinessScenarioTaskCollectionRequestBuilder is request builder for BusinessScenarioTask collection rcn
type PlannerTasksCollectionBusinessScenarioTaskCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for BusinessScenarioTask collection
func (b *PlannerTasksCollectionBusinessScenarioTaskCollectionRequestBuilder) Request() *PlannerTasksCollectionBusinessScenarioTaskCollectionRequest {
	return &PlannerTasksCollectionBusinessScenarioTaskCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for BusinessScenarioTask item
func (b *PlannerTasksCollectionBusinessScenarioTaskCollectionRequestBuilder) ID(id string) *BusinessScenarioTaskRequestBuilder {
	bb := &BusinessScenarioTaskRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerTasksCollectionBusinessScenarioTaskCollectionRequest is request for BusinessScenarioTask collection
type PlannerTasksCollectionBusinessScenarioTaskCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for BusinessScenarioTask collection
func (r *PlannerTasksCollectionBusinessScenarioTaskCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]BusinessScenarioTask, error) {
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
	var values []BusinessScenarioTask
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
			value  []BusinessScenarioTask
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

// GetN performs GET request for BusinessScenarioTask collection, max N pages
func (r *PlannerTasksCollectionBusinessScenarioTaskCollectionRequest) GetN(ctx context.Context, n int) ([]BusinessScenarioTask, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for BusinessScenarioTask collection
func (r *PlannerTasksCollectionBusinessScenarioTaskCollectionRequest) Get(ctx context.Context) ([]BusinessScenarioTask, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for BusinessScenarioTask collection
func (r *PlannerTasksCollectionBusinessScenarioTaskCollectionRequest) Add(ctx context.Context, reqObj *BusinessScenarioTask) (resObj *BusinessScenarioTask, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}