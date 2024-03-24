package helpers

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func DebreviateNumber(abbrv string) (value float64, err error) {
	powers := map[byte]float64{
		'k': math.Pow10(3),
		'm': math.Pow10(6),
		'b': math.Pow10(9),
		't': math.Pow10(12),
		'p': math.Pow10(15),
		'e': math.Pow10(18),
		'z': math.Pow10(21),
		'y': math.Pow10(24),
		'r': math.Pow10(27),
		'q': math.Pow10(30),
	}

	abbrv = strings.TrimSpace(strings.ToLower(abbrv))
	abbrvSegments := [2]string{
		abbrv[:len(abbrv)-1],
		string(abbrv[len(abbrv)-1]),
	}

	power, ok := powers[abbrvSegments[1][0]]
	if !ok {
		err = fmt.Errorf("unknown magnitude abbreviation %s", abbrvSegments[1])
		return 0, err
	}

	base, err := strconv.ParseFloat(abbrvSegments[0], 64)
	if err != nil {
		err = fmt.Errorf("could not parse number %s", abbrvSegments[0])
		return 0, err
	}

	return base * power, nil
}
