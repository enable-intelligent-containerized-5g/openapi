package models

type UdmEeEeSubscription struct {
	// String providing an URI formatted according to RFC 3986.
	CallbackReference string `json:"callbackReference" yaml:"callbackReference" bson:"callbackReference,omitempty"`
	// A map (list of key-value pairs where ReferenceId serves as key) of MonitoringConfigurations
	MonitoringConfigurations map[string]UdmEeMonitoringConfiguration `json:"monitoringConfigurations" yaml:"monitoringConfigurations" bson:"monitoringConfigurations,omitempty"`
	ReportingOptions         *UdmEeReportingOptions                  `json:"reportingOptions,omitempty" yaml:"reportingOptions" bson:"reportingOptions,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string       `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
	SubscriptionId    string       `json:"subscriptionId,omitempty" yaml:"subscriptionId" bson:"subscriptionId,omitempty"`
	ContextInfo       *ContextInfo `json:"contextInfo,omitempty" yaml:"contextInfo" bson:"contextInfo,omitempty"`
	EpcAppliedInd     bool         `json:"epcAppliedInd,omitempty" yaml:"epcAppliedInd" bson:"epcAppliedInd,omitempty"`
	// Fully Qualified Domain Name
	ScefDiamHost string `json:"scefDiamHost,omitempty" yaml:"scefDiamHost" bson:"scefDiamHost,omitempty"`
	// Fully Qualified Domain Name
	ScefDiamRealm       string `json:"scefDiamRealm,omitempty" yaml:"scefDiamRealm" bson:"scefDiamRealm,omitempty"`
	NotifyCorrelationId string `json:"notifyCorrelationId,omitempty" yaml:"notifyCorrelationId" bson:"notifyCorrelationId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	SecondCallbackRef string `json:"secondCallbackRef,omitempty" yaml:"secondCallbackRef" bson:"secondCallbackRef,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi            string   `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi,omitempty"`
	ExcludeGpsiList []string `json:"excludeGpsiList,omitempty" yaml:"excludeGpsiList" bson:"excludeGpsiList,omitempty"`
	IncludeGpsiList []string `json:"includeGpsiList,omitempty" yaml:"includeGpsiList" bson:"includeGpsiList,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	DataRestorationCallbackUri string `json:"dataRestorationCallbackUri,omitempty" yaml:"dataRestorationCallbackUri" bson:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              bool   `json:"udrRestartInd,omitempty" yaml:"udrRestartInd" bson:"udrRestartInd,omitempty"`
}
