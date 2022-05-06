package purchase

import (
	"fmt"
	"strings"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {

	if strings.EqualFold("truck", kind) || strings.EqualFold("car", kind) {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var vehicle string
	if option1 < option2 {
		vehicle = option1
	} else {
		vehicle = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", vehicle)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * .80
	} else if age >= 10 {
		return originalPrice * .5
	}
	return originalPrice * .7
}
