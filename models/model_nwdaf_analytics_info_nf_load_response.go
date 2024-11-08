package models

type NwdaAnalyticsInfoNfLoadResponse struct {
	EventId         EventId                   `json:"eventId" yaml:"eventId" bson:"eventId"`
	AnaliticsNfLoad []NwdaAnalyticsInfoNfLoad `json:"analiticsNfLoad" yaml:"analiticsNfLoad" bson:"analiticsNfLoad"`
}
