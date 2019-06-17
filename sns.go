package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)


func main() {
    //var matrix [49][49] boolean

    var fp *os.File
    var err error

    fp, err = os.Open("links.txt")
    if err != nil {
        panic(err)
    }
    defer fp.Close()

    scanner := bufio.NewScanner(fp)
    links := [][]int{}
    cnt := 0
    for scanner.Scan() {
        line := scanner.Text()
        array := strings.Split(line, "\t")
        from, _ := strconv.Atoi(array[0])
        to, _ := strconv.Atoi(array[1])
        link := []int{from, to}
        links = append(links, link)

        cnt++
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }
    fmt.Println(links)
}
