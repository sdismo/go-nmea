package nmea

const (
	// PrefixPSILTBS prefix
	PrefixPSILTBS = "PSILTBS"
)

// PSILTBS represents Target Boat Speed.
// Proprietary sentence for Garmin/Nexus/Silva NX2 sailboat instruments.
// Indicates expected sailboat performance given heading, wind
// conditions and stored polar performance tables.
type PSILTBS struct {
	BaseSentence
	TargetBoatSpeed float64
}

// $PSILTBS,8.6,N*2D
func newPSILTBS(s BaseSentence) (PSILTBS, error) {
	p := newParser(s, PrefixPSILTBS)
	return PSILTBS{
		BaseSentence:    s,
		TargetBoatSpeed: p.Float64(0, "target boat speed (knots)"),
	}, p.Err()
}
