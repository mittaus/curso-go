package Shapes

type ShapeInterface interface {
	CalcArea() float64
	GetName() string
	SetWidth(float64)
	SetHeight(float64)
	SetRadius(float64)
}

// type Vehicle struct {
// }

// type Car struct {
// 	Vehicle
// }

// type Bike struct {
// 	Vehicle
// }

// func (c Car) getWheels() (wheels int)  { return 100 }
// func (b Bike) getWheels() (wheels int) { return 180 }

// func (v Vehicle) getVehicleIdentifier() string                {}
// func (v Vehicle) saveVehicle(vIn Vehicle)                     {}
// func (v Vehicle) getVehicle(identifier string) (vOut Vehicle) {}

// func (v Vehicle) vehicleWheels(a []Vehicle) {
// 	for i := 0; i <= len(a); i++ {
// 		if a[i].getVehicleIdentifier() == "car" {
// 			fmt.Println("4")
// 		}
// 		if a[i].getVehicleIdentifier() == "bike" {
// 			fmt.Println("2")
// 		}
// 	}
// }
