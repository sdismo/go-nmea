package nmea

const (
	// PrefixIIVDR prefix
	PrefixIIVDR = "IIVDR"
)

// IIVHW represents water current set and drift.
// The direction towards which a current flows (set) and speed (drift) of a current.
type IIVDR struct {
	BaseSentence
	CurrentTrueDirection     float64
	CurrentMagneticDirection float64
	CurrentSpeedKnots        float64
}

// $IIVDR,179.22,T,173.83,M,0.15,N*0C
func newIIVDR(s BaseSentence) (IIVDR, error) {
	p := newParser(s, PrefixIIVDR)
	return IIVDR{
		BaseSentence:             s,
		CurrentTrueDirection:     p.Float64(0, "current true direction"),
		CurrentMagneticDirection: p.Float64(2, "current magnetic direction"),
		CurrentSpeedKnots:        p.Float64(4, "current speed (knots)"),
	}, p.Err()
}
