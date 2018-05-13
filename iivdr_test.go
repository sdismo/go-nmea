package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var iivdrtests = []struct {
	name string
	raw  string
	err  string
	msg  IIVDR
}{
	{
		name: "good sentence",
		raw:  "$IIVDR,179.22,T,173.83,M,0.15,N*0C",
		msg: IIVDR{
			CurrentTrueDirection:      179.22,
			CurrentMagneticDirection:  173.83,
			CurrentSpeedKnots:         0.15,
		},
	},
	{
		name: "invalid magnetic direction",
		raw:  "$IIVDR,179.22,T,A,M,0.15,N*5D",
		err:  "nmea: IIVDR invalid current magnetic direction: A",
	},
}

func TestIIVDR(t *testing.T) {
	for _, tt := range iivdrtests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				iivdr := m.(IIVDR)
				iivdr.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, iivdr)
			}
		})
	}
}
