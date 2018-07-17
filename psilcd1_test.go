package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var psilcd1tests = []struct {
	name string
	raw  string
	err  string
	msg  PSILCD1
}{
	{
		name: "good sentence",
		raw:  "$PSILCD1,130.50,99.80*0F",
		msg: PSILCD1{
			CustomAngularData:  130.5,
			CustomFixpointData: 99.8,
		},
	},
	{
		name: "invalid custom fixpoint data",
		raw:  "$PSILCD1,130.50,A*68",
		err:  "nmea: PSILCD1 invalid custom fixpoint data: A",
	},
}

func TestPSILCD1(t *testing.T) {
	for _, tt := range psilcd1tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				psiltbs := m.(PSILCD1)
				psiltbs.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, psiltbs)
			}
		})
	}
}
