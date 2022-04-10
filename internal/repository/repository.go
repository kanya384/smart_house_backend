package repository

type Repository struct {
	Users           UsersRepository
	Controllers     ControllersRepository
	ControllerTypes ControllerTypes
	HouseParts      HouseParts
	Houses          Houses
	DeviceTypes     DeviceTypes
	Devices         Devices
	Pins            Pins
}
