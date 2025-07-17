package main

func main() {
	smallVehicle := NewVehicle("ABC123", SMALL)
	mediumVehicle := NewVehicle("XYZ789", MEDIUM)
	largeVehicle := NewVehicle("M1234", LARGE)

	parkingLot := GetParkingLotInstance()
	parkingLot.AddLevel(NewLevel(1, 10))

	parkingLot.ParkVehicle(smallVehicle)
	parkingLot.ParkVehicle(mediumVehicle)
	parkingLot.ParkVehicle(largeVehicle)

	parkingLot.ParkVehicle(largeVehicle)

	parkingLot.DisplayAvailability()

	parkingLot.UnparkVehicle(mediumVehicle)

	parkingLot.DisplayAvailability()
}
