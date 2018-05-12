package nmea

const (
	// PrefixIIVHW prefix
	PrefixIIVHW = "IIVHW"
)

// IIVHW represents speed through water and vessel heading.
type IIVHW struct {
	BaseSentence
	TrueHeading      float64
	MagneticHeading  float64
	SpeedKnots       float64
	SpeedKPH         float64
}

// $IIVHW,140,T,134,M,05.23,N,09.68,K*55
// true heading, magnetic heading, speed in knots, speed in km/h
func newIIVHW(s BaseSentence) (IIVHW, error) {
	p := newParser(s, PrefixIIVHW)
	return IIVHW{
		BaseSentence:     s,
		TrueHeading:      p.Float64(0, "true heading"),
		MagneticHeading:  p.Float64(2, "magnetic heading"),
		SpeedKnots:       p.Float64(4, "speed (knots)"),
		SpeedKPH:         p.Float64(6, "speed (km/h)"),
	}, p.Err()
}

