package common

import "testing"

func TestValidateCpf(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{
			name:  "With 041.091.641-25 string value should return true",
			value: "041.091.641-25",
			want:  true,
		},
		{
			name:  "With 123456789 string value should return false",
			value: "1234567",
			want:  false,
		},
		{
			name:  "With NULL string value should return false",
			value: "NULL",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatingCPF(tt.value); got != tt.want {
				t.Errorf("ValidatingCPF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateCnpj(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{
			name:  "With 79.379.491/0001-83 string value should return true",
			value: "79.379.491/0001-83",
			want:  true,
		},
		{
			name:  "With 123456789 string value should return false",
			value: "123456789",
			want:  false,
		},
		{
			name:  "With NULL string value should return false",
			value: "NULL",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatingCNPJ(tt.value); got != tt.want {
				t.Errorf("ValidatingCNPJ() = %v, want %v", got, tt.want)
			}
		})
	}
}
