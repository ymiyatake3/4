package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func getLinks() {
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
        // Read a line
        line := scanner.Text()

        // Split A\tB to [A, B]
        array := strings.Split(line, "\t")

        // Convert each number to int
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


func main() {
    //var matrix [49][49] boolean
    getLinks()
}
