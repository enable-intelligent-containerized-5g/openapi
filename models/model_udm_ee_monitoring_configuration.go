package models

type UdmEeMonitoringConfiguration struct {
	EventType                      UdmEeEventType                       `json:"eventType" yaml:"eventType" bson:"eventType,omitempty"`
	ImmediateFlag                  bool                                 `json:"immediateFlag,omitempty" yaml:"immediateFlag" bson:"immediateFlag,omitempty"`
	LocationReportingConfiguration *UdmEeLocationReportingConfiguration `json:"locationReportingConfiguration,omitempty" yaml:"locationReportingConfiguration" bson:"locationReportingConfiguration,omitempty"`
	AssociationType                UdmEeAssociationType                 `json:"associationType,omitempty" yaml:"associationType" bson:"associationType,omitempty"`
	DatalinkReportCfg              *DatalinkReportingConfiguration      `json:"datalinkReportCfg,omitempty" yaml:"datalinkReportCfg" bson:"datalinkReportCfg,omitempty"`
	LossConnectivityCfg            *LossConnectivityCfg                 `json:"lossConnectivityCfg,omitempty" yaml:"lossConnectivityCfg" bson:"lossConnectivityCfg,omitempty"`
	// indicating a time in seconds.
	MaximumLatency int32 `json:"maximumLatency,omitempty" yaml:"maximumLatency" bson:"maximumLatency,omitempty"`
	// indicating a time in seconds.
	MaximumResponseTime  int32 `json:"maximumResponseTime,omitempty" yaml:"maximumResponseTime" bson:"maximumResponseTime,omitempty"`
	SuggestedPacketNumDl int32 `json:"suggestedPacketNumDl,omitempty" yaml:"suggestedPacketNumDl" bson:"suggestedPacketNumDl,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn                   string                          `json:"dnn,omitempty" yaml:"dnn" bson:"dnn,omitempty"`
	SingleNssai           *Snssai                         `json:"singleNssai,omitempty" yaml:"singleNssai" bson:"singleNssai,omitempty"`
	PduSessionStatusCfg   *UdmEePduSessionStatusCfg       `json:"pduSessionStatusCfg,omitempty" yaml:"pduSessionStatusCfg" bson:"pduSessionStatusCfg,omitempty"`
	ReachabilityForSmsCfg ReachabilityForSmsConfiguration `json:"reachabilityForSmsCfg,omitempty" yaml:"reachabilityForSmsCfg" bson:"reachabilityForSmsCfg,omitempty"`
	// String uniquely identifying MTC provider information.
	MtcProviderInformation string                                 `json:"mtcProviderInformation,omitempty" yaml:"mtcProviderInformation" bson:"mtcProviderInformation,omitempty"`
	AfId                   string                                 `json:"afId,omitempty" yaml:"afId" bson:"afId,omitempty"`
	ReachabilityForDataCfg *UdmEeReachabilityForDataConfiguration `json:"reachabilityForDataCfg,omitempty" yaml:"reachabilityForDataCfg" bson:"reachabilityForDataCfg,omitempty"`
	IdleStatusInd          bool                                   `json:"idleStatusInd,omitempty" yaml:"idleStatusInd" bson:"idleStatusInd,omitempty"`
}
