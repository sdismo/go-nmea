package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var iivhwtests = []struct {
	name string
	raw  string
	err  string
	msg  IIVHW
}{
	{
		name: "good sentence",
		raw:  "$IIVHW,140,T,134,M,05.23,N,09.68,K*55",
		msg: IIVHW{
			TrueHeading:      140,
			MagneticHeading:  134,
			SpeedKnots:       5.23,
			SpeedKMH:         9.68,
		},
	},
	{
		name: "invalid magnetic heading",
		raw:  "$IIVHW,140,T,A,M,05.23,N,09.68,K*22",
		err:  "nmea: IIVHW invalid magnetic heading: A",
	},
}

func TestIIVHW(t *testing.T) {
	for _, tt := range iivhwtests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				iivhw := m.(IIVHW)
				iivhw.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, iivhw)
			}
		})
	}
}
