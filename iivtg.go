package nmea

// Note: xxVTG exists both as GPVTG and IIVTG depending on kind of talker eqm.
// IEC 61162-1 (2nd ed) VTG have one extra 'mode indicator' field at the end not
// listed by http://aprs.gids.nl/nmea/#vtg

const (
	// PrefixIIVTG prefix
	PrefixIIVTG = "IIVTG"

	// A = Autonomous mode
	AutonomousMode = "A"
	//D = Differential mode
	DifferentialMode = "D"
	//E = Estimated (dead reckoning) mode
	EstimatedMode = "E"
	//M = Manual input mode
	ManualMode = "M"
	//S = Simulator mode
	SimulatorMode = "S"
	//N = Data not valid
	DataNotValid = "N"
)

// IIVTG represents track & speed data.
type IIVTG struct {
	BaseSentence
	TrueTrack        float64
	MagneticTrack    float64
	GroundSpeedKnots float64
	GroundSpeedKPH   float64
	ModeIndicator    string
}

// $IIVTG,018.99,T,012.89,M,7.54,N,,,A*6C
// $IIVTG,170.91,T,164.81,M,7.37,N,,,D*63
// $IIVTG,,,,,,N,,,N*69
func newIIVTG(s BaseSentence) (IIVTG, error) {
	p := newParser(s, PrefixIIVTG)
	return IIVTG{
		BaseSentence:     s,
		TrueTrack:        p.Float64(0, "true track"),
		MagneticTrack:    p.Float64(2, "magnetic track"),
		GroundSpeedKnots: p.Float64(4, "ground speed (knots)"),
		GroundSpeedKPH:   p.Float64(6, "ground speed (km/h)"),
		ModeIndicator:    p.EnumString(8, "mode indicator", AutonomousMode, DifferentialMode, EstimatedMode, ManualMode, SimulatorMode, DataNotValid),
	}, p.Err()
}
