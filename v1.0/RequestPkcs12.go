// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// Pkcs12CertificateRequestBuilder is request builder for Pkcs12Certificate
type Pkcs12CertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns Pkcs12CertificateRequest
func (b *Pkcs12CertificateRequestBuilder) Request() *Pkcs12CertificateRequest {
	return &Pkcs12CertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Pkcs12CertificateRequest is request for Pkcs12Certificate
type Pkcs12CertificateRequest struct{ BaseRequest }

// Get performs GET request for Pkcs12Certificate
func (r *Pkcs12CertificateRequest) Get(ctx context.Context) (resObj *Pkcs12Certificate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Pkcs12Certificate
func (r *Pkcs12CertificateRequest) Update(ctx context.Context, reqObj *Pkcs12Certificate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Pkcs12Certificate
func (r *Pkcs12CertificateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Pkcs12CertificateInformationRequestBuilder is request builder for Pkcs12CertificateInformation
type Pkcs12CertificateInformationRequestBuilder struct{ BaseRequestBuilder }

// Request returns Pkcs12CertificateInformationRequest
func (b *Pkcs12CertificateInformationRequestBuilder) Request() *Pkcs12CertificateInformationRequest {
	return &Pkcs12CertificateInformationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Pkcs12CertificateInformationRequest is request for Pkcs12CertificateInformation
type Pkcs12CertificateInformationRequest struct{ BaseRequest }

// Get performs GET request for Pkcs12CertificateInformation
func (r *Pkcs12CertificateInformationRequest) Get(ctx context.Context) (resObj *Pkcs12CertificateInformation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Pkcs12CertificateInformation
func (r *Pkcs12CertificateInformationRequest) Update(ctx context.Context, reqObj *Pkcs12CertificateInformation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Pkcs12CertificateInformation
func (r *Pkcs12CertificateInformationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}