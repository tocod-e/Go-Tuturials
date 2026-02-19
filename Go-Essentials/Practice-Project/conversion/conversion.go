package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(inputStr []string) ([]float64, error) {
	var floats []float64
	for _, strVal := range inputStr {
		strVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return nil, errors.New("Failed to convert string to float.")
		}
		floats = append(floats, strVal)
	}
	return floats, nil
}
