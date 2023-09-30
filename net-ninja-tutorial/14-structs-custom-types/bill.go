package main

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func makeBill(name string) bill {
	return bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
}
