package models

type NwdaAnalyticsInfoNfLoadResponse struct {
	EventId         EventId                   `json:"eventId" yaml:"eventId" bson:"eventId"`
	AnalysisType    AnalysisType              `json:"analysisType" yaml:"analysisType" bson:"analysisType"`
	TargetPerid     int64                     `json:"targetPerid" yaml:"targetPerid" bson:"targetPerid"`
	OffSet          int64                     `json:"offSet" yaml:"offSet" bson:"offSet"`
	AnaliticsNfLoad []NwdaAnalyticsInfoNfLoad `json:"analiticsNfLoad" yaml:"analiticsNfLoad" bson:"analiticsNfLoad"`
}
