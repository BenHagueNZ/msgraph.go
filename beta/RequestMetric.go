// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// MetricTimeSeriesDataPointRequestBuilder is request builder for MetricTimeSeriesDataPoint
type MetricTimeSeriesDataPointRequestBuilder struct{ BaseRequestBuilder }

// Request returns MetricTimeSeriesDataPointRequest
func (b *MetricTimeSeriesDataPointRequestBuilder) Request() *MetricTimeSeriesDataPointRequest {
	return &MetricTimeSeriesDataPointRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MetricTimeSeriesDataPointRequest is request for MetricTimeSeriesDataPoint
type MetricTimeSeriesDataPointRequest struct{ BaseRequest }

// Get performs GET request for MetricTimeSeriesDataPoint
func (r *MetricTimeSeriesDataPointRequest) Get(ctx context.Context) (resObj *MetricTimeSeriesDataPoint, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MetricTimeSeriesDataPoint
func (r *MetricTimeSeriesDataPointRequest) Update(ctx context.Context, reqObj *MetricTimeSeriesDataPoint) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MetricTimeSeriesDataPoint
func (r *MetricTimeSeriesDataPointRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
