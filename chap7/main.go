package main
import (
    "fmt";
    "math"
)

// structs
type Circle struct {
    x, y, r float64
}
func (c *Circle) area() float64  {
    return math.Pi * c.r*c.r
}

type Rectangle struct {
    x1, y1, x2, y2 float64
}
func (r *Rectangle) distance() float64{
    var a = r.x2 - r.x1
    var b = r.y2 - r.y1
    return math.Sqrt(a*a + b*b)
}
func (r *Rectangle) area() float64 {
    var l = r.distance()
    var w = r.distance()
    return l * w
}


//embedded types
type Person struct {
    Name string
}
func (p *Person) talk(){
    fmt.Println("Person method")
}
type Android struct {
    Person
    Model string
}

// interfaces
type Shape interface {
    area() float64
}


// func
func totalArea(s ...Shape) float64 {
    var total = 0.0
    for _, s := range s{
        total += s.area()
    }
    return total
}
func main()  {
    var c = &Circle{10, 15, 4.16}
    fmt.Printf("%.2f\n", c.area())
    var r = &Rectangle{10,12,7,20}
    fmt.Printf("%.2f\n", r.area())
    fmt.Printf("%.2f\n", totalArea(c, r))
    var p = new(Person)
    p.Name = "Heron"
    p.talk()
    var a = new(Android)
    a.talk()
}
