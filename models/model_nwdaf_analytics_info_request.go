package models

type NwdafAnalyticsInfoRequest struct {
	EventId       EventId  `json:"eventId" yaml:"eventId" bson:"eventId"`
	NfInstanceIds []string `json:"nfInstanceIds" yaml:"nfInstanceIds" bson:"nfInstanceIds"`
	NfTypes       []NfType `json:"nfTypes" yaml:"nfTypes" bson:"nfTypes"`
}
