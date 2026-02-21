package main
//import "fmt"

//func main() {
	//p := MakePoint(10.5, 20.3, "Hello World")
	//fmt.Println(p)
//}


type Point struct{
	X float32
	Y float32
	Text string
}

func MakePoint(x, y float32, text string) Point {
	return Point{
		X: x,
		Y: y,
		Text: text,
	}
}