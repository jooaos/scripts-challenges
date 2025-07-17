package main

type Level struct {
	floor        int
	parkingSpots []*ParkingSpot
}

func NewLevel(floor int, numSpots int) *Level {
	level := &Level{floor: floor}
	smallSpots := int(float64(numSpots) * 0.50)
	mediumSpots := int(float64(numSpots) * 0.40)
	largeSpots := int(float64(numSpots) * 0.10)

	for i := 1; i <= smallSpots; i++ {
		level.parkingSpots = append(level.parkingSpots, NewParkingSpot(i, SMALL))
	}

	for i := mediumSpots; i <= mediumSpots; i++ {
		level.parkingSpots = append(level.parkingSpots, NewParkingSpot(i, MEDIUM))
	}

	for i := 1; i <= largeSpots; i++ {
		level.parkingSpots = append(level.parkingSpots, NewParkingSpot(i, LARGE))
	}
	return level
}

func (l *Level) ParkVehicle(vehicle *Vehicle) bool {
	for _, spot := range l.parkingSpots {
		if spot.IsAvailable() && spot.GetVehicleType() == vehicle.GetType() {
			spot.ParkVehicle(vehicle)
			return true
		}
	}
	return false
}

func (l *Level) UnparkVehicle(vehicle *Vehicle) bool {
	for _, spot := range l.parkingSpots {
		if !spot.IsAvailable() && spot.GetParkedVehicle() == vehicle {
			spot.UnparkVehicle()
			return true
		}
	}
	return false
}

func (l *Level) DisplayAvailability() {
	for _, spot := range l.parkingSpots {
		status := "Available"
		if !spot.IsAvailable() {
			status = "Occupied"
		}
		println("Level:", l.floor, "Type:", spot.GetVehicleType(), "Spot:", spot.GetSpotNumber(), "Status:", status)
	}
}
