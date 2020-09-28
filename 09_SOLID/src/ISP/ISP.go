package ISP

import (
	"fmt"
)

type VehicleService interface {
	repairWheels()
	repairLights()
	repairAirbags()
	// repairAC()
}

type WheelsService interface {
	repairWheels()
}
type LightsService interface {
	repairLights()
}
type AirbagsService interface {
	repairAirbags()
}

func Run() {
	fmt.Println("Interface Segregation principle (ISP)")

	fmt.Println(".....................................")
}
