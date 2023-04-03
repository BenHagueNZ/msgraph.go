// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// SimulationAutomations returns request builder for SimulationAutomation collection rcn
func (b *AttackSimulationRootRequestBuilder) SimulationAutomations() *AttackSimulationRootSimulationAutomationsCollectionRequestBuilder {
	bb := &AttackSimulationRootSimulationAutomationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/simulationAutomations"
	return bb
}

// AttackSimulationRootSimulationAutomationsCollectionRequestBuilder is request builder for SimulationAutomation collection
type AttackSimulationRootSimulationAutomationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SimulationAutomation collection
func (b *AttackSimulationRootSimulationAutomationsCollectionRequestBuilder) Request() *AttackSimulationRootSimulationAutomationsCollectionRequest {
	return &AttackSimulationRootSimulationAutomationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SimulationAutomation item
func (b *AttackSimulationRootSimulationAutomationsCollectionRequestBuilder) ID(id string) *SimulationAutomationRequestBuilder {
	bb := &SimulationAutomationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AttackSimulationRootSimulationAutomationsCollectionRequest is request for SimulationAutomation collection
type AttackSimulationRootSimulationAutomationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SimulationAutomation collection
func (r *AttackSimulationRootSimulationAutomationsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SimulationAutomation, error) {
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
	var values []SimulationAutomation
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
			value  []SimulationAutomation
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

// GetN performs GET request for SimulationAutomation collection, max N pages
func (r *AttackSimulationRootSimulationAutomationsCollectionRequest) GetN(ctx context.Context, n int) ([]SimulationAutomation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SimulationAutomation collection
func (r *AttackSimulationRootSimulationAutomationsCollectionRequest) Get(ctx context.Context) ([]SimulationAutomation, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SimulationAutomation collection
func (r *AttackSimulationRootSimulationAutomationsCollectionRequest) Add(ctx context.Context, reqObj *SimulationAutomation) (resObj *SimulationAutomation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Simulations returns request builder for Simulation collection rcn
func (b *AttackSimulationRootRequestBuilder) Simulations() *AttackSimulationRootSimulationsCollectionRequestBuilder {
	bb := &AttackSimulationRootSimulationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/simulations"
	return bb
}

// AttackSimulationRootSimulationsCollectionRequestBuilder is request builder for Simulation collection
type AttackSimulationRootSimulationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Simulation collection
func (b *AttackSimulationRootSimulationsCollectionRequestBuilder) Request() *AttackSimulationRootSimulationsCollectionRequest {
	return &AttackSimulationRootSimulationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Simulation item
func (b *AttackSimulationRootSimulationsCollectionRequestBuilder) ID(id string) *SimulationRequestBuilder {
	bb := &SimulationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AttackSimulationRootSimulationsCollectionRequest is request for Simulation collection
type AttackSimulationRootSimulationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Simulation collection
func (r *AttackSimulationRootSimulationsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Simulation, error) {
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
	var values []Simulation
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
			value  []Simulation
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

// GetN performs GET request for Simulation collection, max N pages
func (r *AttackSimulationRootSimulationsCollectionRequest) GetN(ctx context.Context, n int) ([]Simulation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Simulation collection
func (r *AttackSimulationRootSimulationsCollectionRequest) Get(ctx context.Context) ([]Simulation, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Simulation collection
func (r *AttackSimulationRootSimulationsCollectionRequest) Add(ctx context.Context, reqObj *Simulation) (resObj *Simulation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *AttackSimulationRootRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
