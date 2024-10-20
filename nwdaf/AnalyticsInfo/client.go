package AnalyticsInfo

// APIClient manages communication with the Nnwdaf_AnalyticsInfo API v1.2.2
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	NWDAFAnalyticsDocumentApi *NWDAFAnalyticsDocumentApiService
	NWDAFContextDocumentApi   *NWDAFContextDocumentApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.NWDAFAnalyticsDocumentApi = (*NWDAFAnalyticsDocumentApiService)(&c.common)
	c.NWDAFContextDocumentApi = (*NWDAFContextDocumentApiService)(&c.common)

	return c
}
