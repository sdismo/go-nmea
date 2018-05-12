package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var iidbttests = []struct {
	name string
	raw  string
	err  string
	msg  IIDBT
}{
	{
		name: "good sentence",
		raw:  "$IIDBT,038.94,f,011.87,M,006.41,F*2B",
		msg: IIDBT{
			Depth:   11.87,
		},
	},
	{
		name: "invalid metric depth",
		raw:  "$IIDBT,038.94,f,A,M,006.41,F*7B",
		err:  "nmea: IIDBT invalid depth: A",
	},
}

func TestIIDBT(t *testing.T) {
	for _, tt := range iidbttests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				iidbt := m.(IIDBT)
				iidbt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, iidbt)
			}
		})
	}
}
