package tests

import (
	"load_database_go/common"
	"testing"
)

func TestStringToBoolean(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{
			name:  "With true string value should return true",
			value: "true",
			want:  true,
		},
		{
			name:  "With false string value should return false",
			value: "false",
			want:  false,
		},
		{
			name:  "With 1 string value should return true",
			value: "1",
			want:  true,
		},
		{
			name:  "With 0 string value should return false",
			value: "0",
			want:  false,
		},
		{
			name:  "With t string value should return true",
			value: "t",
			want:  true,
		},
		{
			name:  "With f string value should return false",
			value: "f",
			want:  false,
		},
		{
			name:  "With T string value should return true",
			value: "T",
			want:  true,
		},
		{
			name:  "With F string value should return false",
			value: "F",
			want:  false,
		},
		{
			name:  "With TRUE string value should return true",
			value: "TRUE",
			want:  true,
		},
		{
			name:  "With FALSE string value should return false",
			value: "FALSE",
			want:  false,
		},
		{
			name:  "With True string value should return true",
			value: "True",
			want:  true,
		},
		{
			name:  "With False string value should return false",
			value: "False",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := common.StringToBoolean(tt.value); got != tt.want {
				t.Errorf("StringToBoolean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseToFloatWithDefault(t *testing.T) {
	tests := []struct {
		name         string
		value        string
		defaultValue float64
		want         float64
	}{
		{
			name:         "With NULL string value should return 0.0",
			value:        "NULL",
			defaultValue: 0.0,
			want:         0.0,
		},
		{
			name:         "With -1 string value should return 0.0",
			value:        "-1",
			defaultValue: 0.0,
			want:         0.0,
		},
		{
			name:         "With -909.75 string value should return 0.0",
			value:        "-909.75",
			defaultValue: 0.0,
			want:         0.0,
		},
		{
			name:         "With 89.75 string value should return 89.75",
			value:        "89.75",
			defaultValue: 0.0,
			want:         89.75,
		},
		{
			name:         "With 26 string value should return 26",
			value:        "26",
			defaultValue: 0.0,
			want:         26.00,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := common.ParseToFloatWithDefault(tt.value, tt.defaultValue); got != tt.want {
				t.Errorf("ParseToFloatWithDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
