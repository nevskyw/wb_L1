package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func CreatePoint(x, y float64) *Point { //конструктор
	var point Point
	point.x = x
	point.y = y
	return &point
}

func GetDistance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {

	p := CreatePoint(123, 43)
	p2 := CreatePoint(345.3, 24.556)
	fmt.Println(GetDistance(p, p2))

	//fmt.Println(p.x) //не можем обратиться к координатам точек

}

// type Point struct {
// 	x float64
// 	y float64
// }

// //конструктор под разное количество аргументов
// // функция NewPoint
// func NewPoint(i ...float64) *Point {

// 	if len(i) == 2 {
// 		return &Point{
// 			x: i[0],
// 			y: i[1],
// 		}
// 	}
// 	return &Point{}

// }

// //подсчет расстояния по формуле
// func (p1 *Point) DistanceTo(p2 *Point) float64 {
// 	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
// }

// func main() {
// 	p1 := NewPoint(1, 2, 3, 4)
// 	p2 := NewPoint()
// 	fmt.Println(p1.DistanceTo(p2))
// }

// Адреса golang
// https://golangify.com/pointers
// конструктор
// https://golang-blog.blogspot.com/2020/04/constructors-in-golang.html
