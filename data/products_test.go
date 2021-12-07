package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Chocolate",
		Price: 7.25,
		SKU:   "abc-cdf-efg",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
