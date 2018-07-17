package nmea

const (
	// PrefixPSILCD1 prefix
	PrefixPSILCD1 = "PSILCD1"
)

// PSILCD1 represents Custom Data.
// Proprietary sentence for Garmin/Nexus/Silva NX2 sailboat instruments.
// CAD is often used to indicate most favourable AVA.
// CFD is often used as a performance index in percent.
type PSILCD1 struct {
	BaseSentence
	CustomAngularData  float64
	CustomFixpointData float64
}

// $PSILCD1,130.50,99.80*2D
func newPSILCD1(s BaseSentence) (PSILCD1, error) {
	p := newParser(s, PrefixPSILCD1)
	return PSILCD1{
		BaseSentence:       s,
		CustomAngularData:  p.Float64(0, "custom angular data"),
		CustomFixpointData: p.Float64(1, "custom fixpoint data"),
	}, p.Err()
}
