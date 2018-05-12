package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var iivtgtests = []struct {
	name string
	raw  string
	err  string
	msg  IIVTG
}{
	{
		name: "good sentence autonomous",
		raw:  "$IIVTG,018.99,T,012.89,M,7.54,N,,,A*6C",
		msg: IIVTG{
			TrueTrack:        18.99,
			MagneticTrack:    12.89,
			GroundSpeedKnots: 7.54,
			GroundSpeedKPH:   0,
			ModeIndicator:    AutonomousMode,
		},
	},
	{
		name: "good sentence differential",
		raw:  "$IIVTG,170.91,T,164.81,M,7.37,N,,,D*63",
		msg: IIVTG{
			TrueTrack:        170.91,
			MagneticTrack:    164.81,
			GroundSpeedKnots: 7.37,
			GroundSpeedKPH:   0,
			ModeIndicator:    DifferentialMode,
		},
	},
	{
		name: "good sentence no valid data",
		raw:  "$IIVTG,,,,,,N,,,N*69",
		msg: IIVTG{
			TrueTrack:        0,
			MagneticTrack:    0,
			GroundSpeedKnots: 0,
			GroundSpeedKPH:   0,
			ModeIndicator:    DataNotValid,
		},
	},
	{
		name: "bad true track",
		raw:  "$IIVTG,T,018.99,T,012.89,M,7.54,N,,,A*14",
		err:  "nmea: IIVTG invalid true track: T",
	},
}

func TestIIVTG(t *testing.T) {
	for _, tt := range iivtgtests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				iivtg := m.(IIVTG)
				iivtg.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, iivtg)
			}
		})
	}
}
