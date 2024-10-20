package models

// Represents the RAT type and/or Frequency information.
type RatFreqInformation struct {
	// Set to \"true\" to indicate to handle all the frequencies the NWDAF received, otherwise  set to \"false\" or omit. The \"allFreq\" attribute and the \"freq\" attribute are mutually  exclusive.
	AllFreq bool `json:"allFreq,omitempty" yaml:"allFreq" bson:"allFreq,omitempty"`
	// Set to \"true\" to indicate to handle all the RAT Types the NWDAF received, otherwise  set to \"false\" or omit. The \"allRat\" attribute and the \"ratType\" attribute are mutually  exclusive.
	AllRat bool `json:"allRat,omitempty" yaml:"allRat" bson:"allRat,omitempty"`
	// Integer value indicating the ARFCN applicable for a downlink, uplink or bi-directional (TDD) NR global frequency raster, as definition of \"ARFCN-ValueNR\" IE in clause 6.3.2 of 3GPP TS 38.331.
	Freq            int32             `json:"freq,omitempty" yaml:"freq" bson:"freq,omitempty"`
	RatType         RatType           `json:"ratType,omitempty" yaml:"ratType" bson:"ratType,omitempty"`
	SvcExpThreshold *ThresholdLevel   `json:"svcExpThreshold,omitempty" yaml:"svcExpThreshold" bson:"svcExpThreshold,omitempty"`
	MatchingDir     MatchingDirection `json:"matchingDir,omitempty" yaml:"matchingDir" bson:"matchingDir,omitempty"`
}
