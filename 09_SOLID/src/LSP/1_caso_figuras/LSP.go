package LSP

import (
	"fmt"

	"example.com/username/solid/src/LSP/1_caso_figuras/AreaCalculator"
	"example.com/username/solid/src/LSP/1_caso_figuras/Shapes"
)

func Run() {
	fmt.Println("Liskov Substitution principle (LSP)")

	var circle Shapes.Circle
	var square Shapes.Square
	var rec Shapes.Rectangle

	var calcOutputter AreaCalculator.CalculatorOutputter
	var areaCalc AreaCalculator.AreaCalculator

	circle.Make("circle")
	square.Make("square")
	rec.Make("rectangle")

	var _circle Shapes.ShapeInterface = &circle
	var _square Shapes.ShapeInterface = &square
	var _rec Shapes.ShapeInterface = &rec

	_circle.SetRadius(12.6)
	_square.SetHeight(4.3)
	_rec.SetHeight(6.5)
	_rec.SetWidth(4.8)

	areaCalc.AddShape(_circle)
	areaCalc.AddShape(_square)
	areaCalc.AddShape(_rec)

	areaResult := areaCalc.GetShapesArea()
	fmt.Println(string(calcOutputter.JSON(areaResult)))
	fmt.Println("........................................................")
}
