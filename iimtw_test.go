package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var iimtwtests = []struct {
	name string
	raw  string
	err  string
	msg  IIMTW
}{
	{
		name: "good sentence",
		raw:  "$IIMTW,10.5,C*17",
		msg: IIMTW{
			Temperature:   10.5,
		},
	},
	{
		name: "invalid temperature",
		raw:  "$IIMTW,A,C*4C",
		err:  "nmea: IIMTW invalid temperature: A",
	},
}

func TestIIMTW(t *testing.T) {
	for _, tt := range iimtwtests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				iimtw := m.(IIMTW)
				iimtw.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, iimtw)
			}
		})
	}
}
