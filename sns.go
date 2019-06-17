package main

import (
    "fmt"
    "os"
)


func main() {
    f, err := os.Open("nicknames.txt")
    if err != nil {
        fmt.Println("error")
    }
    defer f.Close()
}
