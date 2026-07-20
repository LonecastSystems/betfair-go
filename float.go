package betfair

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Float is a float64 that also accepts Betfair's documented string forms
// for undefined values: "NaN", "Infinity", and "-Infinity".
type Float float64

func (f *Float) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if len(data) == 0 || bytes.Equal(data, []byte("null")) {
		*f = 0
		return nil
	}
	if data[0] == '"' {
		var text string
		if err := json.Unmarshal(data, &text); err != nil {
			return err
		}
		switch strings.ToLower(strings.TrimSpace(text)) {
		case "", "null":
			*f = 0
		case "nan":
			*f = Float(math.NaN())
		case "inf", "+inf", "infinity", "+infinity":
			*f = Float(math.Inf(1))
		case "-inf", "-infinity":
			*f = Float(math.Inf(-1))
		default:
			value, err := strconv.ParseFloat(text, 64)
			if err != nil {
				return fmt.Errorf("unsupported Betfair float %q", text)
			}
			*f = Float(value)
		}
		return nil
	}
	var value float64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = Float(value)
	return nil
}
