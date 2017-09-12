package main
import "fmt"

func main()  {
    i := 1
    for i<=10 { // or
    // for i := 1; i<=10; i++ {
        fmt.Println(i)
        if i % 2 == 0 {
            fmt.Println("Odd")
        } else {
            fmt.Println("Even")
        }
        
        switch i {
            case 1: fmt.Println("One")
            case 3: fmt.Println("Three")
            case 7: fmt.Println("Seven")
            case 10: fmt.Println("Ten")
            default: fmt.Println("Unknown number")
        }

        fmt.Println("----next")
        i++
    }
}
