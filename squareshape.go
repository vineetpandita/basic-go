package main

import (
	"fmt"
	"sort"
)

type Shape interface {
	area() float32
}


type Square struct {
	side float32
}


type Rectangle struct {
	len float32
	br float32
sort.Interface
}

func (s Square) area() float32 {
	return s.side * s.side
}


func (r Rectangle) area() float32 {
	return r.br * r.len
}

func main() {

	s := Square{10}
	r := Rectangle{10,11}

	info(s)
	info(r)
}

func info(shp Shape)  {

	fmt.Println(shp)
	fmt.Println(shp.area())



}