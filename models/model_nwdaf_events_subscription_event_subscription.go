package models

// Represents a subscription to a single event.
type NwdafEventsSubscriptionEventSubscription struct {
	// \"false\" represents not applicable for all slices. \"true\" represents applicable for all slices.
	AnySlice bool `json:"anySlice,omitempty" yaml:"anySlice" bson:"anySlice,omitempty"`
	// Identification(s) of application to which the subscription applies.
	AppIds []string `json:"appIds,omitempty" yaml:"appIds" bson:"appIds,omitempty"`
	// Identification(s) of DNN to which the subscription applies.
	Dnns           []string                   `json:"dnns,omitempty" yaml:"dnns" bson:"dnns,omitempty"`
	Dnais          []string                   `json:"dnais,omitempty" yaml:"dnais" bson:"dnais,omitempty"`
	Event          NwdafEvent                 `json:"event" yaml:"event" bson:"event,omitempty"`
	ExtraReportReq *EventReportingRequirement `json:"extraReportReq,omitempty" yaml:"extraReportReq" bson:"extraReportReq,omitempty"`
	// Identification(s) of LADN DNN to indicate the LADN service area as the AOI.
	LadnDnns []string `json:"ladnDnns,omitempty" yaml:"ladnDnns" bson:"ladnDnns,omitempty"`
	// Indicates that the NWDAF shall report the corresponding network slice load level to the NF  service consumer where the load level of the network slice identified by snssais is  reached.
	LoadLevelThreshold int32                                     `json:"loadLevelThreshold,omitempty" yaml:"loadLevelThreshold" bson:"loadLevelThreshold,omitempty"`
	NotificationMethod NwdafEventsSubscriptionNotificationMethod `json:"notificationMethod,omitempty" yaml:"notificationMethod" bson:"notificationMethod,omitempty"`
	MatchingDir        MatchingDirection                         `json:"matchingDir,omitempty" yaml:"matchingDir" bson:"matchingDir,omitempty"`
	// Shall be supplied in order to start reporting when an average load level is reached.
	NfLoadLvlThds []ThresholdLevel        `json:"nfLoadLvlThds,omitempty" yaml:"nfLoadLvlThds" bson:"nfLoadLvlThds,omitempty"`
	NfInstanceIds []string                `json:"nfInstanceIds,omitempty" yaml:"nfInstanceIds" bson:"nfInstanceIds,omitempty"`
	NfSetIds      []string                `json:"nfSetIds,omitempty" yaml:"nfSetIds" bson:"nfSetIds,omitempty"`
	NfTypes       []NrfNfManagementNfType `json:"nfTypes,omitempty" yaml:"nfTypes" bson:"nfTypes,omitempty"`
	NetworkArea   *NetworkAreaInfo        `json:"networkArea,omitempty" yaml:"networkArea" bson:"networkArea,omitempty"`
	VisitedAreas  []NetworkAreaInfo       `json:"visitedAreas,omitempty" yaml:"visitedAreas" bson:"visitedAreas,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxTopAppUlNbr int32 `json:"maxTopAppUlNbr,omitempty" yaml:"maxTopAppUlNbr" bson:"maxTopAppUlNbr,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxTopAppDlNbr int32                    `json:"maxTopAppDlNbr,omitempty" yaml:"maxTopAppDlNbr" bson:"maxTopAppDlNbr,omitempty"`
	NsiIdInfos     []NsiIdInfo              `json:"nsiIdInfos,omitempty" yaml:"nsiIdInfos" bson:"nsiIdInfos,omitempty"`
	NsiLevelThrds  []int32                  `json:"nsiLevelThrds,omitempty" yaml:"nsiLevelThrds" bson:"nsiLevelThrds,omitempty"`
	QosRequ        *QosRequirement          `json:"qosRequ,omitempty" yaml:"qosRequ" bson:"qosRequ,omitempty"`
	QosFlowRetThds []RetainabilityThreshold `json:"qosFlowRetThds,omitempty" yaml:"qosFlowRetThds" bson:"qosFlowRetThds,omitempty"`
	RanUeThrouThds []string                 `json:"ranUeThrouThds,omitempty" yaml:"ranUeThrouThds" bson:"ranUeThrouThds,omitempty"`
	// indicating a time in seconds.
	RepetitionPeriod int32 `json:"repetitionPeriod,omitempty" yaml:"repetitionPeriod" bson:"repetitionPeriod,omitempty"`
	// Identification(s) of network slice to which the subscription applies. It corresponds to  snssais in the data model definition of 3GPP TS 29.520.
	Snssaia          []Snssai                      `json:"snssaia,omitempty" yaml:"snssaia" bson:"snssaia,omitempty"`
	TgtUe            *TargetUeInformation          `json:"tgtUe,omitempty" yaml:"tgtUe" bson:"tgtUe,omitempty"`
	CongThresholds   []ThresholdLevel              `json:"congThresholds,omitempty" yaml:"congThresholds" bson:"congThresholds,omitempty"`
	NwPerfRequs      []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty" yaml:"nwPerfRequs" bson:"nwPerfRequs,omitempty"`
	BwRequs          []BwRequirement               `json:"bwRequs,omitempty" yaml:"bwRequs" bson:"bwRequs,omitempty"`
	ExcepRequs       []Exception                   `json:"excepRequs,omitempty" yaml:"excepRequs" bson:"excepRequs,omitempty"`
	ExptAnaType      ExpectedAnalyticsType         `json:"exptAnaType,omitempty" yaml:"exptAnaType" bson:"exptAnaType,omitempty"`
	ExptUeBehav      *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty" yaml:"exptUeBehav" bson:"exptUeBehav,omitempty"`
	RatFreqs         []RatFreqInformation          `json:"ratFreqs,omitempty" yaml:"ratFreqs" bson:"ratFreqs,omitempty"`
	ListOfAnaSubsets []AnalyticsSubset             `json:"listOfAnaSubsets,omitempty" yaml:"listOfAnaSubsets" bson:"listOfAnaSubsets,omitempty"`
	DisperReqs       []DispersionRequirement       `json:"disperReqs,omitempty" yaml:"disperReqs" bson:"disperReqs,omitempty"`
	RedTransReqs     []RedundantTransmissionExpReq `json:"redTransReqs,omitempty" yaml:"redTransReqs" bson:"redTransReqs,omitempty"`
	WlanReqs         []WlanPerformanceReq          `json:"wlanReqs,omitempty" yaml:"wlanReqs" bson:"wlanReqs,omitempty"`
	UpfInfo          *UpfInformation               `json:"upfInfo,omitempty" yaml:"upfInfo" bson:"upfInfo,omitempty"`
	AppServerAddrs   []AddrFqdn                    `json:"appServerAddrs,omitempty" yaml:"appServerAddrs" bson:"appServerAddrs,omitempty"`
	DnPerfReqs       []DnPerformanceReq            `json:"dnPerfReqs,omitempty" yaml:"dnPerfReqs" bson:"dnPerfReqs,omitempty"`
}