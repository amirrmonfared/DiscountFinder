package crawler

import (
	"testing"
)

type testCase struct {
	htmlTag string
	want    float64
}

func TestGetPrice(t *testing.T) {

	cases := []testCase{
		{"4.879,42 TL4.896,96 TL4.908,82 TL4.896,96 TL4.934,30 TL", 4.879},
		{"4.079,42 TL4.896,96 TL4.908,82 TL4.896,96 TL4.934,30 TL", 4.079},
		{"4.870,42 TL4.896,96 TL4.908,82 TL4.896,96 TL4.934,30 TL", 4.870},
		{"35.879,42 TL4.896,96 TL4.908,82 TL4.896,96 TL4.934,30 TL", 35.879},
		{"35.079,42 TL4.896,96 TL4.908,82 TL4.896,96 TL4.934,30 TL", 35.079},
		{"52,26 TL4.896,96 TL4.908,82 TL4.896,96 TL4.934,30 TL", 52},
		{"50,1 TL4.896,96 TL4.908,82 TL4.896,96 TL4.934,30 TL", 50},
	}

	for _, tc := range cases {
		got, err := GetPrice(tc.htmlTag)
		if tc.want != got {
			t.Errorf("Expected %f, but got %f err %v:", tc.want, got, err)
		}

	}

}
