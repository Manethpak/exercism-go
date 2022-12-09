package purchase

import (
	"fmt"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if (kind == "car" || kind == "truck") {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	res := option1
	if (option1 > option2) {
		res = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", res)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var coeff float64 = 0.5
	if (age < 3) {
		coeff = 0.8
	} else if (age < 10) {
		coeff = 0.7
	}

	return coeff * originalPrice
}
