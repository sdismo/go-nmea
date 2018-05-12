package nmea


const (
	// PrefixIIMWD prefix
	PrefixIIMWD = "IIMWD"
)

// IIMVD represents wind direction and speed.
type IIMWD struct {
	BaseSentence
	TrueWindDirection      float64
	MagneticWindDirection  float64
	WindSpeedKnots         float64
	WindSpeedMPS           float64
}

// The direction from which the wind blows across the earth’s surface, with respect to north,
// and the speed of the wind.
// $IIMWD,276.94,T,270.84,M,06.47,N,03.33,M*45
// wind direction true, wind direction magnetic, wind speed knots, wind speed m/s
func newIIMWD(s BaseSentence) (IIMWD, error) {
	p := newParser(s, PrefixIIMWD)
	return IIMWD{
		BaseSentence:           s,
		TrueWindDirection:      p.Float64(0, "true wind direction"),
		MagneticWindDirection:  p.Float64(2, "magnetic wind direction"),
		WindSpeedKnots:         p.Float64(4, "wind speed (knots)"),
		WindSpeedMPS:           p.Float64(6, "wind speed (m/s)"),
	}, p.Err()
}


