package nmea

const (
	// PrefixIIMTW prefix
	PrefixIIMTW = "IIMTW"
)

// IIMTW represents water temperature.
type IIMTW struct {
	BaseSentence
	WaterTemperature float64
}

// $IIMTW,10.5,C*17
// Water temperature always in degrees Celsius
func newIIMTW(s BaseSentence) (IIMTW, error) {
	p := newParser(s, PrefixIIMTW)
	return IIMTW{
		BaseSentence:     s,
		WaterTemperature: p.Float64(0, "water temperature (C)"),
	}, p.Err()
}
