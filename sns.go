package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func makeNamesMap() map[int]string{
    var fp *os.File
    var err error

    fp, err = os.Open("nicknames.txt")
    if err != nil {
        panic(err)
    }
    defer fp.Close()

    scanner := bufio.NewScanner(fp)

    // number : name
    names := map[int]string{}

    for scanner.Scan() {
        // Read a line
        line := scanner.Text()

        // Split A\tB to [A, B]
        array := strings.Split(line, "\t")

        num, _ := strconv.Atoi(array[0])
        names[num] = array[1]
    }

    return names
}

func makeLinksArray() [][]int{
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

    return links
}


func main() {
    var matrix [49][49] bool
    links := makeLinksArray()

    // Put link datas into adjacency matrix
    for i := 0; i < len(links); i++ {
        from := links[i][0]
        to := links[i][1]
        matrix[from][to] = true
    }

    names := makeNamesMap()
    myName := "alex"

    
}
