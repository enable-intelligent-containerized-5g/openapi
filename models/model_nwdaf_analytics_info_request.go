package models

type NwdafAnalyticsInfoRequest struct {
	EventId  EventId                 `json:"eventId" yaml:"eventId" bson:"eventId"`
	NfSetIds []string                `json:"nfSetIds" yaml:"nfSetIds" bson:"nfSetIds"`
	NfTypes  []NrfNfManagementNfType `json:"nfTypes" yaml:"nfTypes" bson:"nfTypes"`
}
