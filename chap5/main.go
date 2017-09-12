package main
import "fmt"

func main()  {
    // var grades [5]float64
    // grades[0] = 54.2
    // grades[1] = 10
    // grades[2] = 22.7
    // grades[3] = 87.1
    // grades[4] = 56
    grades := [5]float64{18.4, 10, 55.7,
                         100, 71.55}

    var total float64 = 0
    // for i := 0; i < len(grades); i++ {
    //     total += grades[i]
    // }

    // _ tells the compiler we dont use that variable
    // index, value
    for _, value := range grades {
        total += value
    }
    fmt.Printf("%.2f\n", total / float64(len(grades)))

    // slice1 := make([]float64, 6)
    var arr = [5]int{18,20,30,40,50}
    var slice1 = arr[0:3]
    fmt.Println("Array", arr)
    fmt.Println("slice1 [0:3]", slice1)
    // 2 elements are added to the slice
    // the slice and its undelying arr are modified
    slice1 = append(slice1, 1,2)
    fmt.Println("Array after append 2 itens", arr)
    fmt.Println("slice1 after append 2 itens", slice1)


    // 4 elements are added to the slice
    // 4 + len(slice) = 7
    // > len(arr) so a new slice is created
    // and returned by append
    var slice2 = append(slice1, 1,2,3,4)
    fmt.Println(
        `New slice created after appending 4 itens to arr`,
        slice2)

    copy(slice2, slice1)
    fmt.Println(
        "slice1 after trying to copy slice2 into it.",
        slice1)


    // maps
    fmt.Println("---maps---")
    var elements = make(map[string]string)
    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["Li"] = "Lithium"
    elements["Be"] = "Beryllium"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["N"] = "Nitrogen"
    elements["O"] = "Oxygen"
    elements["F"] = "Fluorine"
    elements["Ne"] = "Neon"
    // or var elements = map[string]string{"key": "value"}

    fmt.Println(elements["Li"])
    // if short syntax
    // if key is found then do something
    if name, found := elements["B"]; found {
        fmt.Println(name)
    }
}
