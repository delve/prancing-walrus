package helpers

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

var powers = map[string]float64{
	"o": math.Pow10(0),
	"k": math.Pow10(3),
	"m": math.Pow10(6),
	"b": math.Pow10(9),
	"t": math.Pow10(12),
	"p": math.Pow10(15),
	"e": math.Pow10(18),
	"z": math.Pow10(21),
	"y": math.Pow10(24),
	"r": math.Pow10(27),
	"q": math.Pow10(30),
}

var numberWMagnitude, _ = regexp.Compile("[0-9]+[" + strings.Join(maps.Keys(powers), "") + "]")
var number, _ = regexp.Compile("[0-9]+")

func DebreviateNumber(abbrv string) (value float64, err error) {
	abbrv = strings.TrimSpace(strings.ToLower(abbrv))
	abbrvSegments := [2]string{"",""}

	switch {
	case numberWMagnitude.Match([]byte(abbrv)):
		abbrvSegments[0] = abbrv[:len(abbrv)-1]
		abbrvSegments[1] = string(abbrv[len(abbrv)-1])
	case number.Match([]byte(abbrv)):
		abbrvSegments[0] = abbrv
		abbrvSegments[1] = "o"
	}


	power, ok := powers[abbrvSegments[1]]
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
