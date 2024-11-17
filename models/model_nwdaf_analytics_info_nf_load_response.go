package models

type NwdafAnalyticsInfoNfLoadResponse struct {
	EventId         EventId                    `json:"eventId" yaml:"eventId" bson:"eventId"`
	AnalysisType    AnalysisType               `json:"analysisType" yaml:"analysisType" bson:"analysisType"`
	TargetPeriod    int64                      `json:"targetPeriod" yaml:"targetPeriod" bson:"targetPeriod"`
	OffSet          int64                      `json:"offSet" yaml:"offSet" bson:"offSet"`
	AnaliticsNfLoad []NwdafAnalyticsInfoNfLoad `json:"analiticsNfLoad" yaml:"analiticsNfLoad" bson:"analiticsNfLoad"`
}
