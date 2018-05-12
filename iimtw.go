package nmea

const (
	// PrefixIIMTW prefix
	PrefixIIMTW = "IIMTW"
)

// IIMTW represents water temperature.
type IIMTW struct {
	BaseSentence
	Temperature   float64
}

// $IIMTW,10.5,C*17
// Temp always in degrees Celsius
func newIIMTW(s BaseSentence) (IIMTW, error) {
	p := newParser(s, PrefixIIMTW)
	return IIMTW{
		BaseSentence:  s,
		Temperature:   p.Float64(0, "temperature"),
	}, p.Err()
}

