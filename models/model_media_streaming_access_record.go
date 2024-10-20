package models

import (
	"time"
)

type MediaStreamingAccessRecord struct {
	// string with format 'date-time' as defined in OpenAPI.
	Timestamp                         *time.Time             `json:"timestamp" yaml:"timestamp" bson:"timestamp,omitempty"`
	MediaStreamHandlerEndpointAddress *EndpointAddress       `json:"mediaStreamHandlerEndpointAddress" yaml:"mediaStreamHandlerEndpointAddress" bson:"mediaStreamHandlerEndpointAddress,omitempty"`
	ApplicationServerEndpointAddress  *EndpointAddress       `json:"applicationServerEndpointAddress" yaml:"applicationServerEndpointAddress" bson:"applicationServerEndpointAddress,omitempty"`
	SessionIdentifier                 string                 `json:"sessionIdentifier,omitempty" yaml:"sessionIdentifier" bson:"sessionIdentifier,omitempty"`
	RequestMessage                    map[string]interface{} `json:"requestMessage" yaml:"requestMessage" bson:"requestMessage,omitempty"`
	CacheStatus                       CacheStatus            `json:"cacheStatus,omitempty" yaml:"cacheStatus" bson:"cacheStatus,omitempty"`
	ResponseMessage                   map[string]interface{} `json:"responseMessage" yaml:"responseMessage" bson:"responseMessage,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	ProcessingLatency float32                `json:"processingLatency" yaml:"processingLatency" bson:"processingLatency,omitempty"`
	ConnectionMetrics map[string]interface{} `json:"connectionMetrics,omitempty" yaml:"connectionMetrics" bson:"connectionMetrics,omitempty"`
}
