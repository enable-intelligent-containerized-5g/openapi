package models

import (
	"time"
)

// Represents the description of analytics with parameters as relevant for the requesting NF  service consumer.
type NwdafAnalyticsInfoAnalyticsData struct {
	// string with format 'date-time' as defined in OpenAPI.
	Start *time.Time `json:"start,omitempty" yaml:"start" bson:"start,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty" yaml:"expiry" bson:"expiry,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStampGen *time.Time             `json:"timeStampGen,omitempty" yaml:"timeStampGen" bson:"timeStampGen,omitempty"`
	AnaMetaInfo  *AnalyticsMetadataInfo `json:"anaMetaInfo,omitempty" yaml:"anaMetaInfo" bson:"anaMetaInfo,omitempty"`
	// The slices and their load level information.
	SliceLoadLevelInfos []SliceLoadLevelInformation                    `json:"sliceLoadLevelInfos,omitempty" yaml:"sliceLoadLevelInfos" bson:"sliceLoadLevelInfos,omitempty"`
	NsiLoadLevelInfos   []NsiLoadLevelInfo                             `json:"nsiLoadLevelInfos,omitempty" yaml:"nsiLoadLevelInfos" bson:"nsiLoadLevelInfos,omitempty"`
	NfLoadLevelInfos    []NfLoadLevelInformation                       `json:"nfLoadLevelInfos,omitempty" yaml:"nfLoadLevelInfos" bson:"nfLoadLevelInfos,omitempty"`
	NwPerfs             []NetworkPerfInfo                              `json:"nwPerfs,omitempty" yaml:"nwPerfs" bson:"nwPerfs,omitempty"`
	SvcExps             []NwdafEventsSubscriptionServiceExperienceInfo `json:"svcExps,omitempty" yaml:"svcExps" bson:"svcExps,omitempty"`
	QosSustainInfos     []QosSustainabilityInfo                        `json:"qosSustainInfos,omitempty" yaml:"qosSustainInfos" bson:"qosSustainInfos,omitempty"`
	UeMobs              []UeMobility                                   `json:"ueMobs,omitempty" yaml:"ueMobs" bson:"ueMobs,omitempty"`
	UeComms             []UeCommunication                              `json:"ueComms,omitempty" yaml:"ueComms" bson:"ueComms,omitempty"`
	UserDataCongInfos   []UserDataCongestionInfo                       `json:"userDataCongInfos,omitempty" yaml:"userDataCongInfos" bson:"userDataCongInfos,omitempty"`
	AbnorBehavrs        []AbnormalBehaviour                            `json:"abnorBehavrs,omitempty" yaml:"abnorBehavrs" bson:"abnorBehavrs,omitempty"`
	SmccExps            []SmcceInfo                                    `json:"smccExps,omitempty" yaml:"smccExps" bson:"smccExps,omitempty"`
	DisperInfos         []DispersionInfo                               `json:"disperInfos,omitempty" yaml:"disperInfos" bson:"disperInfos,omitempty"`
	RedTransInfos       []RedundantTransmissionExpInfo                 `json:"redTransInfos,omitempty" yaml:"redTransInfos" bson:"redTransInfos,omitempty"`
	WlanInfos           []WlanPerformanceInfo                          `json:"wlanInfos,omitempty" yaml:"wlanInfos" bson:"wlanInfos,omitempty"`
	DnPerfInfos         []DnPerfInfo                                   `json:"dnPerfInfos,omitempty" yaml:"dnPerfInfos" bson:"dnPerfInfos,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat string `json:"suppFeat,omitempty" yaml:"suppFeat" bson:"suppFeat,omitempty"`
}
