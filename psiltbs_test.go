package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var psiltbstests = []struct {
	name string
	raw  string
	err  string
	msg  PSILTBS
}{
	{
		name: "good sentence",
		raw:  "$PSILTBS,8.6,N*2D",
		msg: PSILTBS{
			TargetBoatSpeed: 8.6,
		},
	},
	{
		name: "invalid target boat speed",
		raw:  "$PSILTBS,A,N*4C",
		err:  "nmea: PSILTBS invalid target boat speed (knots): A",
	},
}

func TestPSILTBS(t *testing.T) {
	for _, tt := range psiltbstests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				psiltbs := m.(PSILTBS)
				psiltbs.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, psiltbs)
			}
		})
	}
}
