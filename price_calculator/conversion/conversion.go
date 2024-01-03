package conversion

import (
	"errors"
	"strconv"
)

func StringToFloats(strings []string) ([]float64, error) {
	var prices []float64
	for _, stringVal := range strings {
		floatPrice, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, errors.New("failed to conver string to float")
		}
		prices = append(prices, floatPrice)
	}
	return prices, nil
}
