package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))
	for stringIndex, stringValue := range strings {
		floatValue, err := strconv.ParseFloat(stringValue, 64)
		if err != nil {
			err = errors.New("failed to convert data")
			return nil, err
		}
		floats[stringIndex] = floatValue
	}
	return floats, nil
}
