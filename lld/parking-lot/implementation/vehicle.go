package main

type VehicleType int

const (
	SMALL VehicleType = iota
	MEDIUM
	LARGE
)

type Vehicle struct {
	License string
	Type    VehicleType
}

func NewVehicle(license string, typeOf VehicleType) *Vehicle {
	return &Vehicle{
		License: license,
		Type:    typeOf,
	}
}

func (v *Vehicle) GetLicense() string {
	return v.License
}

func (v *Vehicle) GetType() VehicleType {
	return v.Type
}
