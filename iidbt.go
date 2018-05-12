package nmea

const (
	// PrefixIIDBT prefix
	PrefixIIDBT = "IIDBT"
)

// IIDBT represents depth below transducer.
type IIDBT struct {
	BaseSentence
	Depth         float64
}

// $IIDBT,038.94,f,011.87,M,006.41,F*2B
// same reading in feet, meters, fathoms
func newIIDBT(s BaseSentence) (IIDBT, error) {
	p := newParser(s, PrefixIIDBT)
	return IIDBT{
		BaseSentence:  s,
		Depth:         p.Float64(2, "depth"),
	}, p.Err()
}

