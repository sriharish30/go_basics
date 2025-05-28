package conversion

import (
	"errors"
	"strconv"
)

func StringtoFloat(strings []string) ([]float64, error) {
	var floats []float64

	for _, stringval := range strings {
		floatval, err := strconv.ParseFloat(stringval, 64)

		if err != nil {
			return nil, errors.New("Failed to convert string to float")
		}
		floats = append(floats, floatval)
	}
	return floats, nil
}
