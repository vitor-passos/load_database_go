package common

import "strconv"

func StringToBoolean(value string) bool {
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	return boolValue
}

func ParseToFloatWithDefault(value string, defaultValue float64) float64 {
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil || floatValue <= 0 {
		return defaultValue
	}
	return floatValue
}
