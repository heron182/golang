package main
import "fmt"

func compute_avg(grades []float64) float64  {
    var total = 0.0
    for _, grade := range grades{
        total += grade
    }
    return total / float64(len(grades))
}

func add(a ...float64) (total float64) {
    total = 0.0
    for _, value := range a{
        total += value
    }
    return
}

//closure support
func userPool() func() int  {
    var pool = 3
    return func() int {
        if pool > 0 {
            pool--
            return pool
        } else {
            return 0
        }
    }
}

// defer example
func fn1()  {
    fmt.Println("This will print first")
}

func fn2()  {
    fmt.Println("This will print last by using defer")
}

// panic example
func recover_func()  {
    var str = recover()
    fmt.Println("Recovering...", str)
}

// pointers and references
// aka * and & operators
func swap_value(x *int, y *int){
    var swap = new(int) // creates a pointer
    swap = y
    y = x
    x = swap
    fmt.Println("Inside swap_value x,y", *x, *y)
}

func main()  {
    defer recover_func()
    defer fn2()
    fn1()
    var grades = []float64{55.5,10,41.8,15.3,89.4}
    var average = compute_avg(grades)
    fmt.Println(average)
    fmt.Println(add(10,20,30))
    fmt.Println(add(grades...))
    var popUser = userPool()
    fmt.Println(popUser()) // 2
    fmt.Println(popUser()) // 1
    fmt.Println(popUser()) // 0
    var x = 10
    var y = 5
    swap_value(&x, &y)
    fmt.Println("After swap_value x,y", x, y)
    panic("Unexpect error")
}
