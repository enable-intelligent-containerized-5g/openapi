package models

type UdmPpLocationArea struct {
	// Identifies a list of geographic area of the user where the UE is located.
	GeographicAreas []GeographicArea `json:"geographicAreas,omitempty" yaml:"geographicAreas" bson:"geographicAreas,omitempty"`
	// Identifies a list of civic addresses of the user where the UE is located.
	CivicAddresses []CivicAddress   `json:"civicAddresses,omitempty" yaml:"civicAddresses" bson:"civicAddresses,omitempty"`
	NwAreaInfo     *NetworkAreaInfo `json:"nwAreaInfo,omitempty" yaml:"nwAreaInfo" bson:"nwAreaInfo,omitempty"`
	UmtTime        *UmtTime         `json:"umtTime,omitempty" yaml:"umtTime" bson:"umtTime,omitempty"`
}
