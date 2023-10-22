package fizzbuzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	cases := []struct  {
		name string
		want string
		arg  int
		iserr  bool
	}{
		{"ケース1","0",0,false},
		{"ケース2","1",1,false},
		{"ケース3","2",2,false},
		{"ケース4","fizz",3,false},
		{"ケース5","4",4,false},
		{"ケース6","buzz",5,false},
	}

	for _, tt := range cases {
		t.Run(tt.name,func(t *testing.T) {
			actual, err := FizzBuzz(tt.arg)

			if tt.want != actual {
				t.Errorf("expected : %v, but actual %v", tt.want,actual)
			}

			if (err != nil) != tt.iserr {
				t.Errorf("expected : nil, but actual %v",err)
			}
		})
	}
}