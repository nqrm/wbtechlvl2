package pattern

import (
	"fmt"
	"math"
)

/*
Реализовать паттерн «Посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

Применимость:
	- Когда необходимо добавить операцию над всем элементами сложной структуры объектов
	- Когда над объектами сложной структуры объектов надо выполнять некоторые не связанные между собой операции


Плюсы:
	-  Упрощает добавление операций, работающих со сложными структурами объектов.
	-  Объединяет родственные операции в одной структуре.

Минусы:
	-  Паттерн не оправдан, если иерархия элементов часто меняется.
	-  Может привести к нарушению инкапсуляции элементов.
*/

type Square struct {
	side float64
}

func (s *Square) Accept(v IVisitor) {
	v.forSquare(s)
}

type Circle struct {
	radius float64
}

func (c *Circle) Accept(v IVisitor) {
	v.forCircle(c)
}

type Rectangle struct {
	length float64
	width  float64
}

func (r *Rectangle) Accept(v IVisitor) {
	v.forRectangle(r)
}

type IShape interface {
	Accept(IVisitor)
}

type IVisitor interface {
	forSquare(*Square)
	forRectangle(*Rectangle)
	forCircle(*Circle)
}

type AreaVisitor struct {
	area float64
}

func (a *AreaVisitor) forSquare(s *Square) {
	a.area = s.side * s.side
}

func (a *AreaVisitor) forRectangle(r *Rectangle) {
	a.area = r.length * r.width
}

func (a *AreaVisitor) forCircle(c *Circle) {
	a.area = math.Pi * c.radius * c.radius
}

func VisitorPattern() {

	s := &Square{5}
	c := &Circle{3}

	areaVisitor := &AreaVisitor{}

	s.Accept(areaVisitor)
	fmt.Println(areaVisitor.area)
	c.Accept(areaVisitor)
	fmt.Println(areaVisitor.area)

}
