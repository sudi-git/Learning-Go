package main
//import "fmt"

//func main() {
//	c := GetCircle(5)

//	fmt.Println(c.Radius)
//	fmt.Println(c.Diameter)
//	fmt.Println(c.Area)
//	fmt.Println(c.Perimeter)
//}

type Circle struct {
	Radius    float32
	Diameter  float32
	Area      float32
	Perimeter float32
}

func GetCircle(r float32) Circle {
	pi := float32(3.14)

	return Circle{
		Radius:    r,
		Diameter:  2 * r,
		Area:      pi * r * r,
		Perimeter: 2 * pi * r,
	}
}

