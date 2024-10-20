package models

// Represents the network slice and optionally the associated network slice instance and the  load level information.
type NsiLoadLevelInfo struct {
	// Load level information of the network slice and the optionally associated network slice  instance.
	LoadLevelInformation int32   `json:"loadLevelInformation" yaml:"loadLevelInformation" bson:"loadLevelInformation,omitempty"`
	Snssai               *Snssai `json:"snssai" yaml:"snssai" bson:"snssai,omitempty"`
	// Contains the Identifier of the selected Network Slice instance
	NsiId    string         `json:"nsiId,omitempty" yaml:"nsiId" bson:"nsiId,omitempty"`
	ResUsage *ResourceUsage `json:"resUsage,omitempty" yaml:"resUsage" bson:"resUsage,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	NumOfExceedLoadLevelThr int32            `json:"numOfExceedLoadLevelThr,omitempty" yaml:"numOfExceedLoadLevelThr" bson:"numOfExceedLoadLevelThr,omitempty"`
	ExceedLoadLevelThrInd   bool             `json:"exceedLoadLevelThrInd,omitempty" yaml:"exceedLoadLevelThrInd" bson:"exceedLoadLevelThrInd,omitempty"`
	NetworkArea             *NetworkAreaInfo `json:"networkArea,omitempty" yaml:"networkArea" bson:"networkArea,omitempty"`
	TimePeriod              *TimeWindow      `json:"timePeriod,omitempty" yaml:"timePeriod" bson:"timePeriod,omitempty"`
	// Each element indicates the time elapsed between times each threshold is met or exceeded or crossed. The start time and end time are the exact time stamps of the resource usage threshold is reached or exceeded. May be present if the \"listOfAnaSubsets\" attribute is  provided and the maximum number of instances shall not exceed the value provided in the  \"numOfExceedLoadLevelThr\" attribute.
	ResUsgThrCrossTimePeriod []TimeWindow   `json:"resUsgThrCrossTimePeriod,omitempty" yaml:"resUsgThrCrossTimePeriod" bson:"resUsgThrCrossTimePeriod,omitempty"`
	NumOfUes                 *NumberAverage `json:"numOfUes,omitempty" yaml:"numOfUes" bson:"numOfUes,omitempty"`
	NumOfPduSess             *NumberAverage `json:"numOfPduSess,omitempty" yaml:"numOfPduSess" bson:"numOfPduSess,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty" yaml:"confidence" bson:"confidence,omitempty"`
}