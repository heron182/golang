package main
import "fmt"

const name string = "Heron" //global scope

//converts from farenheint to celsius
func f_to_c(){
    fmt.Print("Enter the temperature in F: ")
    var input float64
    fmt.Scanf("%f", &input)
    fmt.Printf("Temperature in C is %.2f\n",
        (input-32) * 5/9 )
}

func main()  {
    var x string = `Hello World`
    fmt.Println(x)
    y := 45 //shorter inference declaration
    y += 10
    fmt.Println(y)
    fmt.Println(name)

    // declaring multiple variables
    var (
        a = 10
        b = "Johnson"
        c = true
    )
    fmt.Println(a,b,c)
    f_to_c()
}
