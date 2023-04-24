// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// OpticalCharacterRecognitionConfigurationRequestBuilder is request builder for OpticalCharacterRecognitionConfiguration
type OpticalCharacterRecognitionConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns OpticalCharacterRecognitionConfigurationRequest
func (b *OpticalCharacterRecognitionConfigurationRequestBuilder) Request() *OpticalCharacterRecognitionConfigurationRequest {
	return &OpticalCharacterRecognitionConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OpticalCharacterRecognitionConfigurationRequest is request for OpticalCharacterRecognitionConfiguration
type OpticalCharacterRecognitionConfigurationRequest struct{ BaseRequest }

// Get performs GET request for OpticalCharacterRecognitionConfiguration
func (r *OpticalCharacterRecognitionConfigurationRequest) Get(ctx context.Context) (resObj *OpticalCharacterRecognitionConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OpticalCharacterRecognitionConfiguration
func (r *OpticalCharacterRecognitionConfigurationRequest) Update(ctx context.Context, reqObj *OpticalCharacterRecognitionConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OpticalCharacterRecognitionConfiguration
func (r *OpticalCharacterRecognitionConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
