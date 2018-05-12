package nmea

const (
	// PrefixIIDBT prefix
	PrefixIIDBT = "IIDBT"
)

// IIDBT represents water depth below transducer.
type IIDBT struct {
	BaseSentence
	DepthFeet         float64
	DepthMeters       float64
	DepthFathoms      float64
}

// $IIDBT,038.94,f,011.87,M,006.41,F*2B
// same reading in feet, meters, fathoms
func newIIDBT(s BaseSentence) (IIDBT, error) {
	p := newParser(s, PrefixIIDBT)
	return IIDBT{
		BaseSentence:  s,
		DepthFeet:     p.Float64(0, "depth (feet)"),
		DepthMeters:   p.Float64(2, "depth (m)"),
		DepthFathoms:  p.Float64(4, "depth (fathoms)"),
	}, p.Err()
}

