package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)


func main() {
    //var matrix [49][49] int

    var fp *os.File
    var err error

    fp, err = os.Open("links.txt")
    if err != nil {
        panic(err)
    }
    defer fp.Close()

    scanner := bufio.NewScanner(fp)
    for scanner.Scan() {
        line := scanner.Text()
        s := strings.Split(line, "\t")
        fmt.Println(s)
        //matrix = append(matrix, line)
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }
    //fmt.Println(matrix)
}
