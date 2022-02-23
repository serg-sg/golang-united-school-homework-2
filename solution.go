package square

// Define custom int type to hold sides number and update CalcSquare signature
// by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

// Определить пользовательский тип int для хранения номера сторон и обновить опеределение CalcSquare,
// заменив #yourTypeNameHere#

// Определяем константы для представления сторон 0, 3 и 4. В тесте используются мнемосхемы:
// SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// это как:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

// Реализовать функцию для вычисления площади равносторонней фигурки по правилам:
// func CalcSquare (sideLen float64, sideNum intCustomType) float64
// Функция CalcSquare должна возвращать правильную площадь для:
// равносторонний треугольник (3 стороны),
// квадрат (4 стороны)
// окружность (0 сторон) (считайте sideLen как радиус)
// если передан какой-либо другой параметр sideNum, вернуть 0
// встроенная константа Pi должна использоваться для обхода теста

import (
	"math"
)

// Определяем тип для хранения номера сторон
type Side int

// Определяем константы для представления сторон
const (
	SidesCircle = iota
	_
	_
	SidesTriangle
	SidesSquare
)

func CalcSquare(sideLen float64, sidesNum Side) float64 {
	switch sidesNum {
	default:
		{
			return 0
		}
	case 0:
		{
			// Площадь круга
			return math.Pi * math.Pow(sideLen, 2)
		}
	case 3:
		{
			// Площаль равностороннего треугольника
			return (math.Pow(sideLen, 2) * math.Pow(3, 0.5) / 4)
		}
	case 4:
		{
			// Площадь квадрата
			return math.Pow(sideLen, 2)
		}

	}
}
