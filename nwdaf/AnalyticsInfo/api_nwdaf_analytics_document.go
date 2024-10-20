package AnalyticsInfo

import (
	"context"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/enable-intelligent-containerized-5g/openapi"
	"github.com/enable-intelligent-containerized-5g/openapi/models"
)

// Linger please
var (
	_ context.Context
)

type NWDAFAnalyticsDocumentApiService service

/*
NWDAFAnalyticsDocumentApiService Read a NWDAF Analytics
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param EventId - Identify the analytics.
 * @param AnaReq - Identifies the analytics reporting requirement information.
 * @param EventFilter - Identify the analytics.
 * @param SupportedFeatures - To filter irrelevant responses related to unsupported features.
 * @param TgtUe - Identify the target UE information.

@return GetNWDAFAnalyticsResponse
*/

// GetNWDAFAnalyticsRequest
type GetNWDAFAnalyticsRequest struct {
	EventId           *models.EventId
	AnaReq            *models.EventReportingRequirement
	EventFilter       *models.NwdafAnalyticsInfoEventFilter
	SupportedFeatures *string
	TgtUe             *models.TargetUeInformation
}

func (r *GetNWDAFAnalyticsRequest) SetEventId(EventId models.EventId) {
	r.EventId = &EventId
}

func (r *GetNWDAFAnalyticsRequest) SetAnaReq(AnaReq models.EventReportingRequirement) {
	r.AnaReq = &AnaReq
}

func (r *GetNWDAFAnalyticsRequest) SetEventFilter(EventFilter models.NwdafAnalyticsInfoEventFilter) {
	r.EventFilter = &EventFilter
}

func (r *GetNWDAFAnalyticsRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}

func (r *GetNWDAFAnalyticsRequest) SetTgtUe(TgtUe models.TargetUeInformation) {
	r.TgtUe = &TgtUe
}

type GetNWDAFAnalyticsResponse struct {
	NwdafAnalyticsInfoAnalyticsData models.NwdafAnalyticsInfoAnalyticsData
}

type GetNWDAFAnalyticsError struct {
	ProblemDetails                     models.ProblemDetails
	ProblemDetailsAnalyticsInfoRequest models.ProblemDetailsAnalyticsInfoRequest
}

func (a *NWDAFAnalyticsDocumentApiService) GetNWDAFAnalytics(ctx context.Context, request *GetNWDAFAnalyticsRequest) (*GetNWDAFAnalyticsResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetNWDAFAnalyticsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/analytics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.EventId == nil {
		return nil, openapi.ReportError("EventId must be non nil")
	} else {
		localVarQueryParams.Add("event-id", openapi.ParameterToString(request.EventId, "multi"))
	}
	if request.AnaReq != nil {
		localVarQueryParams.Add("ana-req", openapi.ParameterToString(request.AnaReq, "application/json"))
	}
	if request.EventFilter != nil {
		localVarQueryParams.Add("event-filter", openapi.ParameterToString(request.EventFilter, "application/json"))
	}
	if request.SupportedFeatures != nil {
		localVarQueryParams.Add("supported-features", openapi.ParameterToString(request.SupportedFeatures, "multi"))
	}
	if request.TgtUe != nil {
		localVarQueryParams.Add("tgt-ue", openapi.ParameterToString(request.TgtUe, "application/json"))
	}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	err = localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.StatusCode,
	}

	switch localVarHTTPResponse.StatusCode {
	case 200:
		err = openapi.Deserialize(&localVarReturnValue.NwdafAnalyticsInfoAnalyticsData, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 204:
		return &localVarReturnValue, nil
	case 400:
		var v GetNWDAFAnalyticsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v GetNWDAFAnalyticsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v GetNWDAFAnalyticsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v GetNWDAFAnalyticsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 406:
		return &localVarReturnValue, nil
	case 414:
		var v GetNWDAFAnalyticsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v GetNWDAFAnalyticsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v GetNWDAFAnalyticsError
		err = openapi.Deserialize(&v.ProblemDetailsAnalyticsInfoRequest, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v GetNWDAFAnalyticsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	default:
		return nil, apiError
	}
}
