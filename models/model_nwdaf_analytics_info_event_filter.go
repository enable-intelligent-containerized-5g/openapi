package models

// Represents the event filters used to identify the requested analytics.
type NwdafAnalyticsInfoEventFilter struct {
	// \"false\" represents not applicable for all slices. \"true\" represents applicable for all slices.
	AnySlice bool `json:"anySlice,omitempty" yaml:"anySlice" bson:"anySlice,omitempty"`
	// Identification(s) of network slice.
	Snssais []Snssai `json:"snssais,omitempty" yaml:"snssais" bson:"snssais,omitempty"`
	AppIds  []string `json:"appIds,omitempty" yaml:"appIds" bson:"appIds,omitempty"`
	Dnns    []string `json:"dnns,omitempty" yaml:"dnns" bson:"dnns,omitempty"`
	Dnais   []string `json:"dnais,omitempty" yaml:"dnais" bson:"dnais,omitempty"`
	// Identification(s) of LADN DNN to indicate the LADN service area as the AOI.
	LadnDnns     []string          `json:"ladnDnns,omitempty" yaml:"ladnDnns" bson:"ladnDnns,omitempty"`
	NetworkArea  *NetworkAreaInfo  `json:"networkArea,omitempty" yaml:"networkArea" bson:"networkArea,omitempty"`
	VisitedAreas []NetworkAreaInfo `json:"visitedAreas,omitempty" yaml:"visitedAreas" bson:"visitedAreas,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxTopAppUlNbr int32 `json:"maxTopAppUlNbr,omitempty" yaml:"maxTopAppUlNbr" bson:"maxTopAppUlNbr,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxTopAppDlNbr   int32                         `json:"maxTopAppDlNbr,omitempty" yaml:"maxTopAppDlNbr" bson:"maxTopAppDlNbr,omitempty"`
	NfInstanceIds    []string                      `json:"nfInstanceIds,omitempty" yaml:"nfInstanceIds" bson:"nfInstanceIds,omitempty"`
	NfSetIds         []string                      `json:"nfSetIds,omitempty" yaml:"nfSetIds" bson:"nfSetIds,omitempty"`
	NfTypes          []NrfNfManagementNfType       `json:"nfTypes,omitempty" yaml:"nfTypes" bson:"nfTypes,omitempty"`
	NsiIdInfos       []NsiIdInfo                   `json:"nsiIdInfos,omitempty" yaml:"nsiIdInfos" bson:"nsiIdInfos,omitempty"`
	QosRequ          *QosRequirement               `json:"qosRequ,omitempty" yaml:"qosRequ" bson:"qosRequ,omitempty"`
	NwPerfTypes      []NetworkPerfType             `json:"nwPerfTypes,omitempty" yaml:"nwPerfTypes" bson:"nwPerfTypes,omitempty"`
	BwRequs          []BwRequirement               `json:"bwRequs,omitempty" yaml:"bwRequs" bson:"bwRequs,omitempty"`
	ExcepIds         []ExceptionId                 `json:"excepIds,omitempty" yaml:"excepIds" bson:"excepIds,omitempty"`
	ExptAnaType      ExpectedAnalyticsType         `json:"exptAnaType,omitempty" yaml:"exptAnaType" bson:"exptAnaType,omitempty"`
	ExptUeBehav      *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty" yaml:"exptUeBehav" bson:"exptUeBehav,omitempty"`
	RatFreqs         []RatFreqInformation          `json:"ratFreqs,omitempty" yaml:"ratFreqs" bson:"ratFreqs,omitempty"`
	DisperReqs       []DispersionRequirement       `json:"disperReqs,omitempty" yaml:"disperReqs" bson:"disperReqs,omitempty"`
	RedTransReqs     []RedundantTransmissionExpReq `json:"redTransReqs,omitempty" yaml:"redTransReqs" bson:"redTransReqs,omitempty"`
	WlanReqs         []WlanPerformanceReq          `json:"wlanReqs,omitempty" yaml:"wlanReqs" bson:"wlanReqs,omitempty"`
	ListOfAnaSubsets []AnalyticsSubset             `json:"listOfAnaSubsets,omitempty" yaml:"listOfAnaSubsets" bson:"listOfAnaSubsets,omitempty"`
	UpfInfo          *UpfInformation               `json:"upfInfo,omitempty" yaml:"upfInfo" bson:"upfInfo,omitempty"`
	AppServerAddrs   []AddrFqdn                    `json:"appServerAddrs,omitempty" yaml:"appServerAddrs" bson:"appServerAddrs,omitempty"`
	DnPerfReqs       []DnPerformanceReq            `json:"dnPerfReqs,omitempty" yaml:"dnPerfReqs" bson:"dnPerfReqs,omitempty"`
}
