package add

import (
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		name  string
		arg1  int
		arg2  int
		want  int
		iserr bool
	}{
		{"ケース1", 1, 1, 2, false},
		{"ケース2", 1, 2, 3, false},
		{"ケース3", 2, 3, 5, false},
		{"ケース4", 4, 5, 9, false},
		{"ケース5", 5, 6, 11, false},
		{"ケース6", 6, 7, 13, false},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Add(tt.arg1, tt.arg2)
			if actual != tt.want {
				t.Errorf("expected : %v, but actual %v", tt.want, actual)
			}

			if (err != nil) != tt.iserr {
				t.Errorf("expected : nil, but actual %v", err)
			}
		})
	}
}

func TestMinus(t *testing.T) {
	cases := []struct {
		name  string
		arg1  int
		arg2  int
		want  int
		iserr bool
	}{
		{"ケース1", 1, 1, 0, false},
		{"ケース2", 1, 2, -1, false},
		{"ケース3", 2, 3, -1, false},
		{"ケース4", 4, 5, -1, false},
		{"ケース5", 5, 6, -1, false},
		{"ケース6", 6, 7, -1, false},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Minus(tt.arg1, tt.arg2)
			if actual != tt.want {
				t.Errorf("expected : %v, but actual %v", tt.want, actual)
			}

			if (err != nil) != tt.iserr {
				t.Errorf("expected : nil, but actual %v", err)
			}
		})
	}
}

func TestDouble(t *testing.T) {
	cases := []struct {
		name  string
		arg1  int
		arg2  int
		want  int
		iserr bool
	}{
		{"ケース1", 1, 1, 1, false},
		{"ケース2", 1, 2, 2, false},
		{"ケース3", 2, 3, 6, false},
		{"ケース4", 4, 5, 20, false},
		{"ケース5", 5, 6, 30, false},
		{"ケース6", 6, 7, 42, false},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Double(tt.arg1, tt.arg2)
			if actual != tt.want {
				t.Errorf("expected : %v, but actual %v", tt.want, actual)
			}

			if (err != nil) != tt.iserr {
				t.Errorf("expected : nil, but actual %v", err)
			}
		})
	}
}
