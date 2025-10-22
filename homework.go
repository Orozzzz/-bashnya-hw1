package main

import "fmt"

func main() {
	
    maxNumber := 12307
    var num int
    fmt.Print("Enter a any number: ")
    fmt.Scanln(&num)

    if num > maxNumber {
    fmt.Printf("Error! The number %d great maxnumber %d\n", num, maxNumber)
    return
    }

    for num < maxNumber {
        if num < 0 {
            num = num * -1
        } else if num % 7 == 0 {
            num *= 39
        } else if num % 9 == 0 {
            num = num * 13 + 1
            continue
        } else {
            num = (num + 2) * 3
    }

    if num % 13 == 0 && num % 9 == 0 {
        fmt.Println("service error")
        return
    } else {
        num += 1
    }
    }

    fmt.Printf("Result: %d\n", num)

}