package main

 
type Point struct{
	X float32
	Y float32
	Text string
}

func PointDiff(p1, p2 Point) Point {
//checking if p1's X &Y coordinate is greater than or equal to p2's X&Y
	if p1.X>=p2.X && p1.Y>=p2.Y{
	return p1
	}else{
	return p2
	}
}