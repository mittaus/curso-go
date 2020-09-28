package casoVehiculo

import "fmt"

type IVehicle interface {
	getVehicleIdentifier() string
	getWheels() (wheels int)
	getMaxSpeed() (wheels int)
}

type Vehicle struct {
	name string
}

type Car struct {
	Vehicle
}

type Bike struct {
	Vehicle
}

func (v Vehicle) getVehicleIdentifier() string {
	return v.name
}

func vehicleWheels(a []IVehicle) {
	for i := 0; i < len(a); i++ {
		// if a[i].getVehicleIdentifier() == "carro" {
		// 	fmt.Println(a[i].getMaxSpeed())
		// }
		// if a[i].getVehicleIdentifier() == "bicicleta" {
		// 	fmt.Println(a[i].getMaxSpeed())
		// }
		fmt.Println(a[i].getMaxSpeed())
	}
}

func (c Car) getWheels() (wheels int)  { return 4 }
func (b Bike) getWheels() (wheels int) { return 2 }

func (c Car) getMaxSpeed() (wheels int)  { return 100 }
func (b Bike) getMaxSpeed() (wheels int) { return 180 }

func main() {
	carro := Car{Vehicle{"carro"}}
	bici := Bike{Vehicle{"bicicleta"}}

	misVehiculos := []IVehicle{}
	misVehiculos = append(misVehiculos, carro, bici)
	vehicleWheels(misVehiculos)
	fmt.Println(carro, bici)
}
