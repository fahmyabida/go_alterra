package main

func hitungLuasSegitga(alas, tinggi float64) func() float64 {
	return func() float64 {
		return 0.5 * alas * tinggi
	}
}
