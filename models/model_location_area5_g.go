package models

// Represents a user location area when the UE is attached to 5G.
type LocationArea5G struct {
	// Identifies a list of geographic area of the user where the UE is located.
	GeographicAreas []GeographicArea `json:"geographicAreas,omitempty" yaml:"geographicAreas" bson:"geographicAreas,omitempty"`
	// Identifies a list of civic addresses of the user where the UE is located.
	CivicAddresses []CivicAddress   `json:"civicAddresses,omitempty" yaml:"civicAddresses" bson:"civicAddresses,omitempty"`
	NwAreaInfo     *NetworkAreaInfo `json:"nwAreaInfo,omitempty" yaml:"nwAreaInfo" bson:"nwAreaInfo,omitempty"`
}
