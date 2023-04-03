// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// Runs returns request builder for SimulationAutomationRun collection rcn
func (b *SimulationAutomationRequestBuilder) Runs() *SimulationAutomationRunsCollectionRequestBuilder {
	bb := &SimulationAutomationRunsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/runs"
	return bb
}

// SimulationAutomationRunsCollectionRequestBuilder is request builder for SimulationAutomationRun collection
type SimulationAutomationRunsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SimulationAutomationRun collection
func (b *SimulationAutomationRunsCollectionRequestBuilder) Request() *SimulationAutomationRunsCollectionRequest {
	return &SimulationAutomationRunsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SimulationAutomationRun item
func (b *SimulationAutomationRunsCollectionRequestBuilder) ID(id string) *SimulationAutomationRunRequestBuilder {
	bb := &SimulationAutomationRunRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SimulationAutomationRunsCollectionRequest is request for SimulationAutomationRun collection
type SimulationAutomationRunsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SimulationAutomationRun collection
func (r *SimulationAutomationRunsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SimulationAutomationRun, error) {
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
	var values []SimulationAutomationRun
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
			value  []SimulationAutomationRun
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

// GetN performs GET request for SimulationAutomationRun collection, max N pages
func (r *SimulationAutomationRunsCollectionRequest) GetN(ctx context.Context, n int) ([]SimulationAutomationRun, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SimulationAutomationRun collection
func (r *SimulationAutomationRunsCollectionRequest) Get(ctx context.Context) ([]SimulationAutomationRun, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SimulationAutomationRun collection
func (r *SimulationAutomationRunsCollectionRequest) Add(ctx context.Context, reqObj *SimulationAutomationRun) (resObj *SimulationAutomationRun, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Simulation is navigation property rn
func (b *SimulationRequestBuilder) Simulation() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// SimulationAutomation is navigation property rn
func (b *SimulationAutomationRequestBuilder) SimulationAutomation() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// SimulationAutomationRun is navigation property rn
func (b *SimulationAutomationRunRequestBuilder) SimulationAutomationRun() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
