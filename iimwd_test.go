package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var iimwdtests = []struct {
	name string
	raw  string
	err  string
	msg  IIMWD
}{
	{
		name: "good sentence",
		raw:  "$IIMWD,276.94,T,270.84,M,06.47,N,03.33,M*45",
		msg: IIMWD{
			TrueWindDirection:     276.94,
			MagneticWindDirection: 270.84,
			WindSpeedKnots:        6.47,
			WindSpeedMPS:          3.33,
		},
	},
	{
		name: "invalid wind speed knots",
		raw:  "$IIMWD,276.94,T,270.84,M,A,N,03.33,M*2F",
		err:  "nmea: IIMWD invalid wind speed (knots): A",
	},
}

func TestIIMWD(t *testing.T) {
	for _, tt := range iimwdtests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				iimwd := m.(IIMWD)
				iimwd.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, iimwd)
			}
		})
	}
}
