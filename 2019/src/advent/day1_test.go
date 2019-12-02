package main

import "testing"

func Test_calculateFuelValue(t *testing.T) {
	fuelValue := calculateFuelValue(12)
	if fuelValue != 2 {
		t.Error("Fuel Value not correct. Expected ", 2, " and got ", fuelValue)
	}

	fuelValue = calculateFuelValue(1969)
	if fuelValue != 966 {
		t.Error("Fuel Value not correct. Expected ", 966, " and got ", fuelValue)
	}
}
